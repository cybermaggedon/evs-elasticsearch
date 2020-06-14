//
// ElasticSearch loader.  Maps events to an ElasticSearch schema.
//

package main

import (
	evs "github.com/cybermaggedon/evs-golang-api"
	"log"
	"os"
	"strconv"
)

const ()

type ElasticSearch struct {

	// Embed EventAnalytic framework
	evs.EventAnalytic

	loader *Loader
}

// Initialisation
func (e *ElasticSearch) Init(binding string) error {

	lc := NewLoader()

	if val, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok {
		lc = lc.Url(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_READ_ALIAS"); ok {
		lc = lc.ReadAlias(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_WRITE_ALIAS"); ok {
		lc = lc.WriteAlias(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_TEMPLATE"); ok {
		lc = lc.Template(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_SHARDS"); ok {
		shards, _ := strconv.Atoi(val)
		lc = lc.Shards(shards)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_BOX_TYPE"); ok {
		lc = lc.BoxType(val)
	}

	var err error
	e.loader, err = lc.Build()
	if err != nil {
		return err
	}

	e.EventAnalytic.Init(binding, []string{}, e)
	return nil
}

// Event handler for new events.
func (e *ElasticSearch) Event(ev *evs.Event, p map[string]string) error {

	log.Print("event")

	obs := Convert(ev)

	err := e.loader.Load(obs)
	if (err != nil) {
		return err
	}

	return nil
}

func main() {

	e := &ElasticSearch{}

	binding, ok := os.LookupEnv("INPUT")
	if !ok {
		binding = "ioc"
	}

	e.Init(binding)

	log.Print("Initialisation complete.")

	e.Run()

}
