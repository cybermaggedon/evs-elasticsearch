//
// ElasticSearch loader.  Maps events to an ElasticSearch schema.
//

package main

import (
	evs "github.com/cybermaggedon/evs-golang-api"
	"log"
	"fmt"
	"os"
	"encoding/json"
)

const (
)

type ElasticSearch struct {

	// Embed EventAnalytic framework
	evs.EventAnalytic

}

// Initialisation
func (e *ElasticSearch) Init(binding string) error {
	e.EventAnalytic.Init(binding, []string{}, e)
	return nil
}

// Event handler for new events.
func (e *ElasticSearch) Event(ev *evs.Event, properties map[string]string) error {

	log.Print("event")

	obs := Convert(ev)

	b, err := json.Marshal(obs)
	if err != nil {
		log.Printf("JSON encode: %v", err)
		return err
	}

	fmt.Println(string(b))

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
