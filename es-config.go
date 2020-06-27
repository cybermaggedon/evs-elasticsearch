package main

import (
	"github.com/cybermaggedon/evs-golang-api"
	"os"
	"strconv"
)

type EsConfig struct {
	*evs.Config
	url         string
	template    string
	write_alias string
	read_alias  string
	shards      int
	object      string
	box_type    string
}

func NewEsConfig() *EsConfig {

	c := &EsConfig{
		Config:      evs.NewConfig("evs-elasticsearch", "ioc", nil),
		url:         "http://localhost:9200",
		read_alias:  "cyberprobe",
		write_alias: "active-cyberprobe",
		template:    "active-cyberprobe",
		shards:      1,
	}

	if val, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok {
		c.Url(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_READ_ALIAS"); ok {
		c.ReadAlias(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_WRITE_ALIAS"); ok {
		c.WriteAlias(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_TEMPLATE"); ok {
		c.Template(val)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_SHARDS"); ok {
		shards, _ := strconv.Atoi(val)
		c.Shards(shards)
	}
	if val, ok := os.LookupEnv("ELASTICSEARCH_BOX_TYPE"); ok {
		c.BoxType(val)
	}

	return c

}

func (lc *EsConfig) Url(url string) {
	lc.url = url
}

func (lc *EsConfig) ReadAlias(val string) {
	lc.read_alias = val
}

func (lc *EsConfig) WriteAlias(val string) {
	lc.write_alias = val
}

func (lc *EsConfig) Template(val string) {
	lc.template = val
}

func (lc *EsConfig) Shards(val int) {
	lc.shards = val
}

func (lc *EsConfig) BoxType(val string) {
	lc.box_type = val
}
