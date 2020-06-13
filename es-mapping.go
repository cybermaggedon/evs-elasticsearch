package main

var mapping = TypeMapping{"cyberprobe": obMapping}

var obMapping = TypeMapping{
	"properties": TypeMapping{
		"id": TypeMapping{
			"type": "keyword",
		},
		"time": TypeMapping{
			"type": "date",
		},
		"url": TypeMapping{
			"type": "keyword",
		},
		"action": TypeMapping{
			"type": "keyword",
		},
		"device": TypeMapping{
			"type": "keyword",
		},
		"network": TypeMapping{
			"type": "keyword",
		},
		"origin": TypeMapping{
			"type": "keyword",
		},
		"risk": TypeMapping{
			"type": "float",
		},
		"operations": TypeMapping{
			"properties": TypeMapping{
				"unknown": TypeMapping{
					"type": "keyword",
				},
			},
		},
		"dns": TypeMapping{
			"properties": TypeMapping{
				"type": TypeMapping{
					"type": "keyword",
				},
				"query": TypeMapping{
					"properties": TypeMapping{
						"name": TypeMapping{
							"type": "keyword",
						},
						"type": TypeMapping{
							"type": "keyword",
						},
						"class": TypeMapping{
							"type": "keyword",
						},
					},
				},
				"answer": TypeMapping{
					"properties": TypeMapping{
						"name": TypeMapping{
							"type": "keyword",
						},
						"type": TypeMapping{
							"type": "keyword",
						},
						"class": TypeMapping{
							"type": "keyword",
						},
						"address": TypeMapping{
							"type": "keyword",
						},
					},
				},
			},
		},
		"http": TypeMapping{
			"properties": TypeMapping{
				"method": TypeMapping{
					"type": "keyword",
				},
				"status": TypeMapping{
					"type": "keyword",
				},
				"code": TypeMapping{
					"type": "integer",
				},
				"header": TypeMapping{
					"properties": TypeMapping{
						"User-Agent": TypeMapping{
							"type": "keyword",
						},
						"Host": TypeMapping{
							"type": "keyword",
						},
						"Content-Type": TypeMapping{
							"type": "keyword",
						},
						"Server": TypeMapping{
							"type": "keyword",
						},
						"Connection": TypeMapping{
							"type": "keyword",
						},
					},
				},
			},
		},
		"ftp": TypeMapping{
			"properties": TypeMapping{
				"command": TypeMapping{
					"type": "keyword",
				},
				"status": TypeMapping{
					"type": "integer",
				},
				"text": TypeMapping{
					"type": "text",
				},
			},
		},
		"icmp": TypeMapping{
			"properties": TypeMapping{
				"type": TypeMapping{
					"type": "integer",
				},
				"code": TypeMapping{
					"type": "integer",
				},
			},
		},
		"sip": TypeMapping{
			"properties": TypeMapping{
				"method": TypeMapping{
					"type": "keyword",
				},
				"from": TypeMapping{
					"type": "keyword",
				},
				"to": TypeMapping{
					"type": "keyword",
				},
				"status": TypeMapping{
					"type": "keyword",
				},
				"code": TypeMapping{
					"type": "integer",
				},
			},
		},
		"smtp": TypeMapping{
			"properties": TypeMapping{
				"command": TypeMapping{
					"type": "keyword",
				},
				"from": TypeMapping{
					"type": "keyword",
				},
				"to": TypeMapping{
					"type": "keyword",
				},
				"status": TypeMapping{
					"type": "keyword",
				},
				"text": TypeMapping{
					"type": "text",
				},
				"code": TypeMapping{
					"type": "integer",
				},
			},
		},
		"ntp": TypeMapping{
			"properties": TypeMapping{
				"version": TypeMapping{
					"type": "integer",
				},
				"mode": TypeMapping{
					"type": "integer",
				},
			},
		},
		"unrecognised_payload": TypeMapping{
			"properties": TypeMapping{
				"sha1": TypeMapping{
					"type": "keyword",
				},
				"length": TypeMapping{
					"type": "integer",
				},
			},
		},
		"src": TypeMapping{
			"properties": TypeMapping{
				"ipv4": TypeMapping{
					"type": "ip",
				},
				"ipv6": TypeMapping{
					"type": "ip",
				},
				"tcp": TypeMapping{
					"type": "integer",
				},
				"udp": TypeMapping{
					"type": "integer",
				},
			},
		},
		"dest": TypeMapping{
			"properties": TypeMapping{
				"ipv4": TypeMapping{
					"type": "ip",
				},
				"ipv6": TypeMapping{
					"type": "ip",
				},
				"tcp": TypeMapping{
					"type": "integer",
				},
				"udp": TypeMapping{
					"type": "integer",
				},
			},
		},
		"location": TypeMapping{
			"properties": TypeMapping{
				"src": TypeMapping{
					"properties": TypeMapping{
						"city": TypeMapping{
							"type": "keyword",
						},
						"iso": TypeMapping{
							"type": "keyword",
						},
						"country": TypeMapping{
							"type": "keyword",
						},
						"asnum": TypeMapping{
							"type": "integer",
						},
						"asorg": TypeMapping{
							"type": "keyword",
						},
						"position": TypeMapping{
							"type": "geo_point",
						},
						"accuracy": TypeMapping{
							"type": "integer",
						},
						"postcode": TypeMapping{
							"type": "keyword",
						},
					},
				},
				"dest": TypeMapping{
					"properties": TypeMapping{
						"city": TypeMapping{
							"type": "keyword",
						},
						"iso": TypeMapping{
							"type": "keyword",
						},
						"country": TypeMapping{
							"type": "keyword",
						},
						"asnum": TypeMapping{
							"type": "integer",
						},
						"asorg": TypeMapping{
							"type": "keyword",
						},
						"position": TypeMapping{
							"type": "geo_point",
						},
						"accuracy": TypeMapping{
							"type": "integer",
						},
						"postcode": TypeMapping{
							"type": "keyword",
						},
					},
				},
			},
		},
		"indicators": TypeMapping{
			"properties": TypeMapping{
				"id": TypeMapping{
					"type": "keyword",
				},
				"type": TypeMapping{
					"type": "keyword",
				},
				"value": TypeMapping{
					"type": "keyword",
				},
				"description": TypeMapping{
					"type": "keyword",
				},
				"category": TypeMapping{
					"type": "keyword",
				},
				"author": TypeMapping{
					"type": "keyword",
				},
				"source": TypeMapping{
					"type": "keyword",
				},
				"probability": TypeMapping{
					"type": "float",
				},
			},
		},
	},
}
