//
// ElasticSearch loader
//

package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

type Loader struct {
	client         *elastic.Client
	bps            *elastic.BulkProcessor
	url            string
	es_template    string
	es_write_alias string
	es_read_alias  string
	es_shards      int
	es_object      string
	eventLatency   *prometheus.SummaryVec
	recvLabels     prometheus.Labels
}

// Can use this to keep track of ES failures.
// FIXME: Should integrate stats.
var afters, failures int64
var afterRequests int64

func afterFn(executionId int64, requests []elastic.BulkableRequest,
	response *elastic.BulkResponse, err error) {

	curAfters := atomic.AddInt64(&afters, 1)

	if err != nil {
		atomic.AddInt64(&failures, 1)
	}
	curFailures := atomic.LoadInt64(&failures)

	curReqs := atomic.AddInt64(&afterRequests, int64(len(requests)))

	if curReqs%100000 == 0 {
		log.Printf("Stats: batches=%d failed=%d objects=%d", curAfters,
			curFailures, curReqs)
	}

}

func (s *Loader) elasticInit() error {

	// Open ElasticSearch connection.

	for {

		var err error
		s.client, err = elastic.NewClient(elastic.SetURL(s.url))
		if err != nil {
			log.Printf("Elasticsearch connection: %s", err.Error())
			continue
		}

		break

		time.Sleep(1 * time.Second * 5)

	}

	for {

		tmplExists, err := s.client.IndexTemplateExists(s.es_template).
			Do(context.Background())
		if err != nil {
			log.Printf("Elasticsearch: %s", err.Error())
			continue
		}

		number_of_shards := strconv.Itoa(s.es_shards)

		template := Mapping{
			"template": s.es_template + "*",
			"aliases": Mapping{
				s.es_read_alias: Mapping{},
			},
			"settings": Mapping{
				"index.mapping.total_fields.limit":         5000,
				"number_of_shards":                         number_of_shards,
				"number_of_replicas":                       1,
				"routing.allocation.total_shards_per_node": number_of_shards,
				"routing.allocation.include.box_type":      "hot",
			},
			"mappings": mapping,
		}

		template_json, err := json.Marshal(&template)
		if err != nil {
			log.Printf("Couldn't encode template to JSON: %v",
				err)
		}

		ipt, err := s.client.IndexPutTemplate(s.es_template).
			BodyString(string(template_json)).
			Do(context.Background())

		if err != nil {
			log.Printf("(PutTemplateFromJson) (ignored): %s",
				err.Error())
		} else {
			if !ipt.Acknowledged {
				log.Print("Create template not acknowledged.")
			} else {
				log.Print("Template created.")
			}
		}

		if tmplExists {
			log.Print("Index Template Update Complete")
			time.Sleep(time.Second * 5)
			break
		}

		time.Sleep(time.Second * 1)

		ci, err := s.client.CreateIndex(s.es_write_alias + "-000001").
			Do(context.Background())

		if err != nil {
			log.Printf("(CreateEmptyIndex) (ignored): %s", err.Error())
		} else {
			if !ci.Acknowledged {
				log.Print("Create index not acknowledged.")
			} else {
				log.Print("Index created.")
			}
		}

		time.Sleep(time.Second * 1)

		ara, err := s.client.Alias().Add(s.es_write_alias+"-000001", s.es_write_alias).
			Do(context.Background())

		if err != nil {
			log.Printf("(AddWriteAlias) (ignored): %s", err.Error())
		} else {
			if !ara.Acknowledged {
				log.Print("Add write alias not acknowledged.")
			} else {
				log.Print("Write Alias Added")
			}
		}

	}

	time.Sleep(time.Second * 1)

	var err error
	s.bps, err = s.client.BulkProcessor().
		Name("Worker-1").
		Workers(5).
		BulkActions(25).
		BulkSize(5 * 1024 * 1024).
		FlushInterval(1 * time.Second).
		After(afterFn).
		Do(context.Background())
	if err != nil {
		log.Printf("BulkProcess: %v", err)

		// FIXME: Need to retry that one.
		os.Exit(1)
	}

	return nil

}

func getenv(env string, def string) string {
	if val, ok := os.LookupEnv(env); ok {
		return val
	}
	return def
}

func (s *Loader) init() error {

	// Get configuration values from environment variables.
	s.url = getenv("ELASTICSEARCH_URL", "http://localhost:9200")
	s.es_read_alias = getenv("ELASTICSEARCH_READ_ALIAS", "cyberprobe")
	s.es_write_alias = getenv("ELASTICSEARCH_WRITE_ALIAS", "active-cyberprobe")
	s.es_template = getenv("ELASTICSEARCH_TEMPLATE", "active-cyberprobe")
	s.es_shards, _ = strconv.Atoi(getenv("ELASTICSEARCH_SHARDS", "3"))
	s.es_object = getenv("ELASTICSEARCH_OBJECT", "observation")

	//configuration specific to prometheus stats
	/*
		s.recvLabels = prometheus.Labels{}
		s.eventLatency = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Name: "event_latency",
				Help: "Latency from cyberprobe to store",
			},
			[]string{"store"},
		)

		prometheus.MustRegister(s.eventLatency)
	*/

	return s.elasticInit()

}

func (h *Loader) Output(ob Observation, id string) error {

	bir := elastic.NewBulkIndexRequest().
		Doc(ob).
		Id(id).
		Index(h.es_write_alias).
		Type(h.es_object)

	ts := time.Now().UnixNano()
	go h.recordLatency(ts, ob)

	h.bps.Add(bir)

	return nil

}

func (h *Loader) recordLatency(ts int64, ob Observation) {
	obsTime, err := time.Parse(time.RFC3339, ob.Time)
	if err != nil {
		log.Printf("Date Parse Error: %s", err.Error())
	}
	latency := ts - obsTime.UnixNano()
	h.eventLatency.With(h.recvLabels).Observe(float64(latency))
}
