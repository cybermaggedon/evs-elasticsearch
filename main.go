//
// ElasticSearch loader.  Maps events to an ElasticSearch schema.
//

package main

import (
	evs "github.com/cybermaggedon/evs-golang-api"
	pb "github.com/cybermaggedon/evs-golang-api/protos"
	"log"
)

const ()

type ElasticSearch struct {
	*EsConfig

	// Embed event analytic framework
	*evs.EventSubscriber
	*evs.EventProducer
	evs.Interruptible

	loader *Loader
}

// Initialisation
func NewElasticSearch(ec *EsConfig) *ElasticSearch {

	e := &ElasticSearch{
		EsConfig: ec,
	}

	var err error
	e.EventSubscriber, err = evs.NewEventSubscriber(e, e)
	if err != nil {
		log.Fatal(err)
	}

	e.EventProducer, err = evs.NewEventProducer(e)
	if err != nil {
		log.Fatal(err)
	}

	e.RegisterStop(e)

	e.loader, err = e.NewLoader()
	if err != nil {
		log.Fatal(err)
	}

	return e
}

// Event handler for new events.
func (e *ElasticSearch) Event(ev *pb.Event, p map[string]string) error {

	obs := Convert(ev)

	err := e.loader.Load(obs)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	ec := NewEsConfig()
	e := NewElasticSearch(ec)
	log.Print("Initialisation complete")
	e.Run()
	log.Print("Shutdown.")

}
