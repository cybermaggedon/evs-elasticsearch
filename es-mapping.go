package main

type Mapping map[string]interface{}

var mapping = Mapping{"cyberprobe": obMapping}

var obMapping = Mapping{
	"properties": Mapping{
		"id": Mapping{
			"type": "keyword",
		},
		"time": Mapping{
			"type": "date",
		},
		"url": Mapping{
			"type": "keyword",
		},
		"action": Mapping{
			"type": "keyword",
		},
		"device": Mapping{
			"type": "keyword",
		},
		"network": Mapping{
			"type": "keyword",
		},
		"origin": Mapping{
			"type": "keyword",
		},
		"risk": Mapping{
			"type": "float",
		},
		"operations": Mapping{
			"properties": Mapping{
				"unknown": Mapping{
					"type": "keyword",
				},
			},
		},
		"dns": Mapping{
			"properties": Mapping{
				"type": Mapping{
					"type": "keyword",
				},
				"query": Mapping{
					"properties": Mapping{
						"name": Mapping{
							"type": "keyword",
						},
						"type": Mapping{
							"type": "keyword",
						},
						"class": Mapping{
							"type": "keyword",
						},
					},
				},
				"answer": Mapping{
					"properties": Mapping{
						"name": Mapping{
							"type": "keyword",
						},
						"type": Mapping{
							"type": "keyword",
						},
						"class": Mapping{
							"type": "keyword",
						},
						"address": Mapping{
							"type": "keyword",
						},
					},
				},
			},
		},
		"http": Mapping{
			"properties": Mapping{
				"method": Mapping{
					"type": "keyword",
				},
				"status": Mapping{
					"type": "keyword",
				},
				"code": Mapping{
					"type": "integer",
				},
				"header": Mapping{
					"properties": Mapping{
						"User-Agent": Mapping{
							"type": "keyword",
						},
						"Host": Mapping{
							"type": "keyword",
						},
						"Content-Type": Mapping{
							"type": "keyword",
						},
						"Server": Mapping{
							"type": "keyword",
						},
						"Connection": Mapping{
							"type": "keyword",
						},
					},
				},
			},
		},
		"ftp": Mapping{
			"properties": Mapping{
				"command": Mapping{
					"type": "keyword",
				},
				"status": Mapping{
					"type": "integer",
				},
				"text": Mapping{
					"type": "text",
				},
			},
		},
		"icmp": Mapping{
			"properties": Mapping{
				"type": Mapping{
					"type": "integer",
				},
				"code": Mapping{
					"type": "integer",
				},
			},
		},
		"sip": Mapping{
			"properties": Mapping{
				"method": Mapping{
					"type": "keyword",
				},
				"from": Mapping{
					"type": "keyword",
				},
				"to": Mapping{
					"type": "keyword",
				},
				"status": Mapping{
					"type": "keyword",
				},
				"code": Mapping{
					"type": "integer",
				},
			},
		},
		"smtp": Mapping{
			"properties": Mapping{
				"command": Mapping{
					"type": "keyword",
				},
				"from": Mapping{
					"type": "keyword",
				},
				"to": Mapping{
					"type": "keyword",
				},
				"status": Mapping{
					"type": "keyword",
				},
				"text": Mapping{
					"type": "text",
				},
				"code": Mapping{
					"type": "integer",
				},
			},
		},
		"ntp": Mapping{
			"properties": Mapping{
				"version": Mapping{
					"type": "integer",
				},
				"mode": Mapping{
					"type": "integer",
				},
			},
		},
		"unrecognised_payload": Mapping{
			"properties": Mapping{
				"sha1": Mapping{
					"type": "keyword",
				},
				"length": Mapping{
					"type": "integer",
				},
			},
		},
		"src": Mapping{
			"properties": Mapping{
				"ipv4": Mapping{
					"type": "ip",
				},
				"ipv6": Mapping{
					"type": "ip",
				},
				"tcp": Mapping{
					"type": "integer",
				},
				"udp": Mapping{
					"type": "integer",
				},
			},
		},
		"dest": Mapping{
			"properties": Mapping{
				"ipv4": Mapping{
					"type": "ip",
				},
				"ipv6": Mapping{
					"type": "ip",
				},
				"tcp": Mapping{
					"type": "integer",
				},
				"udp": Mapping{
					"type": "integer",
				},
			},
		},
		"location": Mapping{
			"properties": Mapping{
				"src": Mapping{
					"properties": Mapping{
						"city": Mapping{
							"type": "keyword",
						},
						"iso": Mapping{
							"type": "keyword",
						},
						"country": Mapping{
							"type": "keyword",
						},
						"asnum": Mapping{
							"type": "integer",
						},
						"asorg": Mapping{
							"type": "keyword",
						},
						"position": Mapping{
							"type": "geo_point",
						},
						"accuracy": Mapping{
							"type": "integer",
						},
						"postcode": Mapping{
							"type": "keyword",
						},
					},
				},
				"dest": Mapping{
					"properties": Mapping{
						"city": Mapping{
							"type": "keyword",
						},
						"iso": Mapping{
							"type": "keyword",
						},
						"country": Mapping{
							"type": "keyword",
						},
						"asnum": Mapping{
							"type": "integer",
						},
						"asorg": Mapping{
							"type": "keyword",
						},
						"position": Mapping{
							"type": "geo_point",
						},
						"accuracy": Mapping{
							"type": "integer",
						},
						"postcode": Mapping{
							"type": "keyword",
						},
					},
				},
			},
		},
		"indicators": Mapping{
			"properties": Mapping{
				"id": Mapping{
					"type": "keyword",
				},
				"type": Mapping{
					"type": "keyword",
				},
				"value": Mapping{
					"type": "keyword",
				},
				"description": Mapping{
					"type": "keyword",
				},
				"category": Mapping{
					"type": "keyword",
				},
				"author": Mapping{
					"type": "keyword",
				},
				"source": Mapping{
					"type": "keyword",
				},
				"probability": Mapping{
					"type": "float",
				},
			},
		},
	},
}
