# `evs-elasticsearch`

Eventstream analytic for Cyberprobe event streams.  Subscribes to Pulsar
for Cyberprobe events and produces events which are stored in an index in
ElasticSearch.

## Getting Started

The target deployment product is a container engine.  The analytic expects
a Pulsar service to be running, along with an ElasticSearch service.

```
  docker run -d \
      -e PULSAR_BROKER=pulsar://<PULSAR-HOST>:6650 \
      -e ELASTICSEARCH_URL=http://<ES-HOST>:9200 \
      -p 8088:8088 \
      docker.io/cybermaggedon/evs-elasticsearch:<VERSION>
```
      
### Prerequisites

You need to have a container deployment system e.g. Podman, Docker, Moby.

You need to have an ElasticSearch service running.  To run a single-node,
non-production service, try...

```
  docker run -d -p 9200:9200 -e discovery.type=single-node \
       elasticsearch:7.7.1
```

You also need a Pulsar exchange, being fed by events from Cyberprobe.

### Installing

The easiest way is to use the containers we publish to Docker hub.
See https://hub.docker.com/r/cybermaggedon/evs-elasticsearch

```
  docker pull docker.io/cybermaggedon/evs-elasticsearch:<VERSION>
```

If you want to build this yourself, you can just clone the Github repo,
and type `make`.

## Deployment configuration

The following environment variables are used to configure:

| Variable | Purpose | Default |
|----------|---------|---------|
| `INPUT` | Specifies the Pulsar topic to subscribe to.  This is just the topic part of the URL e.g. `cyberprobe`. | `ioc` |
| `METRICS_PORT` | Specifies the port number to serve Prometheus metrics on.  If not set, metrics will not be served. The container has a default setting of 8088. | `8088` |
| `ELASTICSEARCH_URL` | Specifies the URL of the ElasticSearch service. | `http://localhost:9200` |
| `ELASTICSEARCH_READ_ALIAS` | Specifies the alias used by clients to access the indexes. | `cyberprobe` |
| `ELASTICSEARCH_WRITE_ALIAS` | Specifies the alias used by this loader to write to the indexes. | `active-cyberprobe` |
| `ELASTICSEARCH_TEMPLATE` | Specifies the template name used to clone new index partitions. | `active-cyberprobe` |
| `ELASTICSEARCH_SHARDS` | Number of replica copies of data.  Use a higher number for higher data resilience. | `1` |


