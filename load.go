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
	"strconv"
	"time"
)

type Loader struct {
	EsConfig
	client        *elastic.Client
	bps           *elastic.BulkProcessor
	event_latency prometheus.Summary
	flushed       prometheus.Summary
	committed     prometheus.Summary
	indexed       prometheus.Summary
	succeeded     prometheus.Summary
	failed        prometheus.Summary
}

const (
	FIELDS_LIMIT    = "index.mapping.total_fields.limit"
	SHARDS          = "number_of_shards"
	REPLICAS        = "number_of_replicas"
	SHARDS_PER_NODE = "routing.allocation.total_shards_per_node"
	BOX_TYPE        = "routing.allocation.include.box_type"
)

func (l *Loader) GenerateTemplate() Mapping {

	number_of_shards := strconv.Itoa(l.shards)

	template := Mapping{
		"template": l.template + "*",
		"aliases": Mapping{
			l.read_alias: Mapping{},
		},
		"settings": Mapping{
			FIELDS_LIMIT:    5000,
			SHARDS:          number_of_shards,
			REPLICAS:        1,
			SHARDS_PER_NODE: number_of_shards,
		},
		"mappings": mapping,
	}

	if l.box_type != "" {
		template["settings"].(Mapping)[BOX_TYPE] = l.box_type
	}

	return template

}

func (l *Loader) InitMetrics() {

	l.event_latency = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "elasticsearch_event_latency",
			Help: "Latency from cyberprobe to store",
		})
	prometheus.MustRegister(l.event_latency)

	l.flushed = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "elasticsearch_flushed",
			Help: "Number of flush interval invocations",
		})
	prometheus.MustRegister(l.flushed)

	l.committed = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "elasticsearch_committed",
			Help: "Numer of bulk request commits",
		})
	prometheus.MustRegister(l.committed)

	l.indexed = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "elasticsearch_indexed",
			Help: "Number of requests indexes",
		})
	prometheus.MustRegister(l.indexed)

	l.succeeded = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "elasticsearch_succeeded",
			Help: "Number of successful requests",
		})
	prometheus.MustRegister(l.succeeded)

	l.failed = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name: "elasticsearch_failed",
			Help: "Number of failed request",
		})
	prometheus.MustRegister(l.failed)

}

func (l *Loader) InitClient() {

	// Open ElasticSearch connection.

	for {

		var err error
		l.client, err = elastic.NewClient(elastic.SetURL(l.url))
		if err != nil {
			log.Printf("Elasticsearch connection: %v", err)
			time.Sleep(1 * time.Second * 5)
			continue
		}

		break

	}

}

func (l *Loader) InitTemplate() {

	for {

		tmplExists, err := l.client.IndexTemplateExists(l.template).
			Do(context.Background())
		if err != nil {
			log.Printf("Elasticsearch: %s", err.Error())
			time.Sleep(time.Second * 5)
			continue
		}

		template := l.GenerateTemplate()

		template_json, err := json.Marshal(&template)
		if err != nil {
			log.Printf("Couldn't encode template to JSON: %v",
				err)
		}

		ipt, err := l.client.IndexPutTemplate(l.template).
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
			break
		}

		time.Sleep(time.Second * 1)

		ci, err := l.client.CreateIndex(l.write_alias + "-000001").
			Do(context.Background())

		if err != nil {
			log.Printf("(CreateEmptyIndex) (ignored): %s",
				err.Error())
		} else {
			if !ci.Acknowledged {
				log.Print("Create index not acknowledged.")
			} else {
				log.Print("Index created.")
			}
		}

		time.Sleep(time.Second * 1)

		ara, err := l.client.Alias().
			Add(l.write_alias+"-000001", l.write_alias).
			Do(context.Background())

		if err != nil {
			log.Printf("(AddWriteAlias) (ignored): %v", err)
		} else {
			if !ara.Acknowledged {
				log.Print("Add write alias not acknowledged.")
			} else {
				log.Print("Write Alias Added")
			}
		}

	}
}

func (l *Loader) InitBulkProcessor() error {

	var err error
	l.bps, err = l.client.BulkProcessor().
		Name("Worker-1").
		Workers(5).
		BulkActions(25).
		BulkSize(5 * 1024 * 1024).
		FlushInterval(1 * time.Second).
		Stats(true).
		Do(context.Background())
	if err != nil {
		log.Printf("BulkProcess: %v", err)
		return err
	}

	return nil

}

func (l *Loader) StatsObserver() {
	
	prev := l.bps.Stats()

	for {

		stats := l.bps.Stats()

		// Stats is a running count, so observe the deltas
		l.flushed.Observe(float64(stats.Flushed - prev.Flushed))
		l.committed.Observe(float64(stats.Committed - prev.Committed))
		l.indexed.Observe(float64(stats.Indexed - prev.Indexed))
		l.succeeded.Observe(float64(stats.Succeeded - prev.Succeeded))
		l.failed.Observe(float64(stats.Failed - prev.Failed))

		prev = stats

		time.Sleep(time.Second)

	}

}

func (l *Loader) Init() error {

	l.InitMetrics()
	l.InitClient()
	l.InitTemplate()
	err := l.InitBulkProcessor()
	if err != nil {
		return err
	}

	// Stats manager
	go l.StatsObserver()

	return nil

}

func (l *Loader) Load(ob *Observation) error {

	bir := elastic.NewBulkIndexRequest().
		Doc(ob).
		Id(ob.Id).
		Index(l.write_alias).
		Type(l.object)

	ts := time.Now()
	go l.recordLatency(ts, ob)

	l.bps.Add(bir)

	return nil

}

func (l *Loader) recordLatency(ts time.Time, ob *Observation) {
	obsTime, err := time.Parse(time.RFC3339, ob.Time)
	if err != nil {
		log.Printf("Date Parse Error: %s", err.Error())
	}
	latency := ts.Sub(obsTime)
	l.event_latency.Observe(float64(latency))
}

func (lc EsConfig) NewLoader() (*Loader, error) {
	l := &Loader{
		EsConfig: lc,
	}
	err := l.Init()
	if err != nil {
		return nil, err
	}
	return l, nil
}
