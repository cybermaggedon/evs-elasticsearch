
// ES model - presenting an event protobuf as an ElasticSearch event

package main

import (
	evs "github.com/cybermaggedon/evs-golang-api"
	"github.com/golang/protobuf/ptypes"
	"net"
	"encoding/binary"
	"strconv"
)

const mappings = `
{
		    "cyberprobe": {
		      "properties": {
		        "id": {
		          "type": "keyword"
		        },
		        "time": {
		          "type": "date"
		        },
		        "url": {
		          "type": "keyword"
		        },
		        "action": {
		          "type": "keyword"
		        },
		        "device": {
		          "type": "keyword"
		        },
		        "network": {
		          "type": "keyword"
		        },
		        "origin": {
		          "type": "keyword"
		        },
		        "risk": {
		          "type": "float"
		        },
				"operations" : {
					"properties": {
						"unknown": {
							"type": "keyword"
						}
					}
				},
		        "dns": {
		          "properties": {
		            "type": {
		              "type": "keyword"
		            },
		            "query": {
		              "properties": {
		                "name": {
		                  "type": "keyword"
		                },
		                "type": {
		                  "type": "keyword"
		                },
		                "class": {
		                  "type": "keyword"
		                }
		              }
		            },
		            "answer": {
		              "properties": {
		                "name": {
		                  "type": "keyword"
		                },
		                "type": {
		                  "type": "keyword"
		                },
		                "class": {
		                  "type": "keyword"
		                },
		                "address": {
		                  "type": "keyword"
		                }
		              }
		            }
		          }
		        },
		        "http": {
		          "properties": {
		            "method": {
		              "type": "keyword"
		            },
		            "status": {
		              "type": "keyword"
		            },
		            "code": {
		              "type": "integer"
		            },
		            "header": {
		              "properties": {
		                "User-Agent": {
		                  "type": "keyword"
		                },
		                "Host": {
		                  "type": "keyword"
		                },
		                "Content-Type": {
		                  "type": "keyword"
		                },
		                "Server": {
		                  "type": "keyword"
		                },
		                "Connection": {
		                  "type": "keyword"
		                }
		              }
		            }
		          }
		        },
		        "ftp": {
		          "properties": {
		            "command": {
		              "type": "keyword"
		            },
		            "status": {
		              "type": "integer"
		            },
		            "text": {
		              "type": "text"
		            }
		          }
		        },
		        "icmp": {
		          "properties": {
		            "type": {
		              "type": "integer"
		            },
		            "code": {
		              "type": "integer"
		            }
		          }
		        },
		        "sip": {
		          "properties": {
		            "method": {
		              "type": "keyword"
		            },
		            "from": {
		              "type": "keyword"
		            },
		            "to": {
		              "type": "keyword"
		            },
		            "status": {
		              "type": "keyword"
		            },
		            "code": {
		              "type": "integer"
		            }
		          }
		        },
		        "smtp": {
		          "properties": {
		            "command": {
		              "type": "keyword"
		            },
		            "from": {
		              "type": "keyword"
		            },
		            "to": {
		              "type": "keyword"
		            },
		            "status": {
		              "type": "keyword"
		            },
		            "text": {
		              "type": "text"
		            },
		            "code": {
		              "type": "integer"
		            }
		          }
		        },
		        "ntp": {
		          "properties": {
		            "version": {
		              "type": "integer"
		            },
		            "mode": {
		              "type": "integer"
		            }
		          }
		        },
		        "unrecognised_payload": {
		          "properties": {
		            "sha1": {
		              "type": "keyword"
		            },
		            "length": {
		              "type": "integer"
		            }
		          }
		        },
		        "src": {
		          "properties": {
		            "ipv4": {
		              "type": "ip"
		            },
		            "ipv6": {
		              "type": "ip"
		            },
		            "tcp": {
		              "type": "integer"
		            },
		            "udp": {
		              "type": "integer"
		            }
		          }
		        },
		        "dest": {
		          "properties": {
		            "ipv4": {
		              "type": "ip"
		            },
		            "ipv6": {
		              "type": "ip"
		            },
		            "tcp": {
		              "type": "integer"
		            },
		            "udp": {
		              "type": "integer"
		            }
		          }
		        },
		        "location": {
		          "properties": {
		            "src": {
		              "properties": {
		                "city": {
		                  "type": "keyword"
		                },
		                "iso": {
		                  "type": "keyword"
		                },
		                "country": {
		                  "type": "keyword"
		                },
		                "asnum": {
		                  "type": "integer"
		                },
		                "asorg": {
		                  "type": "keyword"
		                },
		                "position": {
		                  "type": "geo_point"
		                },
		                "accuracy": {
		                  "type": "integer"
		                },
		                "postcode": {
		                  "type": "keyword"
		                }
		              }
		            },
		            "dest": {
		              "properties": {
		                "city": {
		                  "type": "keyword"
		                },
		                "iso": {
		                  "type": "keyword"
		                },
		                "country": {
		                  "type": "keyword"
		                },
		                "asnum": {
		                  "type": "integer"
		                },
		                "asorg": {
		                  "type": "keyword"
		                },
		                "position": {
		                  "type": "geo_point"
		                },
		                "accuracy": {
		                  "type": "integer"
		                },
		                "postcode": {
		                  "type": "keyword"
		                }
		              }
		            }
		          }
		        },
		        "indicators": {
		          "properties": {
		            "id": {
		              "type": "keyword"
		            },
		            "type": {
		              "type": "keyword"
		            },
		            "value": {
		              "type": "keyword"
		            },
		            "description": {
		              "type": "keyword"
		            },
		            "category": {
		              "type": "keyword"
		            },
		            "author": {
		              "type": "keyword"
		            },
		            "source": {
		              "type": "keyword"
		            },
		            "probability": {
		              "type": "float"
		            }
		          }
		        }
		      }
		    }
		  }
		}`


type ObQuery struct {
	Name  []string `json:"name,omitempty"`
	Class []string `json:"class,omitempty"`
	Type  []string `json:"type,omitempty"`
}

type ObAnswer struct {
	Name    []string `json:"name,omitempty"`
	Class   []string `json:"class,omitempty"`
	Type    []string `json:"type,omitempty"`
	Address []string `json:"address,omitempty"`
}

type ObIndicators struct {
	Id          []string `json:"id,omitempty"`
	Type        []string `json:"type,omitempty"`
	Value       []string `json:"value,omitempty"`
	Description []string `json:"description,omitempty"`
	Category    []string `json:"category,omitempty"`
	Author      []string `json:"author,omitempty"`
	Source      []string `json:"source,omitempty"`
	Probability []float32 `json:"probability,omitempty"`
}

type ObHttp struct {
	Method string            `json:"method,omitempty"`
	Header map[string]string `json:"header,omitempty"`
	Status string            `json:"status,omitempty"`
	Code   int32               `json:"code,omitempty"`
}

type ObFtp struct {
	Command string   `json:"command,omitempty"`
	Status  int32      `json:"status,omitempty"`
	Text    []string `json:"text,omitempty"`
}

type ObIcmp struct {
	Type int32 `json:"type"`
	Code int32 `json:"code"`
}

type ObDns struct {
	Query  *ObQuery  `json:"query,omitempty"`
	Answer *ObAnswer `json:"answer,omitempty"`
	Type   string    `json:"type,omitempty"`
}

type ObSip struct {
	Method string `json:"method,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Status string `json:"status,omitempty"`
	Code   int32    `json:"code,omitempty"`
}

type ObSmtp struct {
	Command string   `json:"command,omitempty"`
	From    string   `json:"from,omitempty"`
	To      []string `json:"to,omitempty"`
	Status  int32      `json:"status,omitempty"`
	Text    []string `json:"text,omitempty"`
}

type ObNtp struct {
	Version int32 `json:"version"`
	Mode    int32 `json:"mode"`
}

type ObUnrecog struct {
	Sha1   string `json:"sha1,omitempty"`
	Length int64    `json:"length,omitempty"`
}

type Observation struct {
	Id      string `json:"id,omitempty"`
	Action  string `json:"action,omitempty"`
	Device  string `json:"device,omitempty"`
	Network string `json:"network,omitempty"`
	Origin  string `json:"origin,omitempty"`
	Time    string `json:"time,omitempty"`

	Src  map[string][]string `json:"src,omitempty"`
	Dest map[string][]string `json:"dest,omitempty"`

	Url string `json:"url,omitempty"`

	Http *ObHttp `json:"http,omitempty"`

	Icmp *ObIcmp `json:"icmp,omitempty"`

	Dns *ObDns `json:"dns,omitempty"`

	Ftp *ObFtp `json:"ftp,omitempty"`

	Sip *ObSip `json:"sip,omitempty"`

	Smtp *ObSmtp `json:"smtp,omitempty"`

	Ntp *ObNtp `json:"ntp,omitempty"`

	Unrecognised *ObUnrecog `json:"unrecognised_payload,omitempty"`

	Location *evs.Locations `json:"location,omitempty"`

	Indicators *ObIndicators `json:"indicators,omitempty"`

	Risk float64 `json:"risk"`

	Operations map[string]string `json:"operations,omitempty"`
}

// Converts a 32-bit int to an IP address
// FIXME: Copied from detector, put in a library
func int32ToIp(ipLong uint32) net.IP {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	return net.IP(ipByte)
}

// Converts a byte array to an IP address. This is for IPv6 addresses.
func bytesToIp(b []byte) net.IP {
	return net.IP(b)
}

func Convert(ev *evs.Event) *Observation {

	ob := &Observation{}
	
	ob.Id = ev.Id
	ob.Action = ev.Action.String()
	ob.Device = ev.Device
	ob.Network = ev.Network
	ob.Origin = ev.Origin.String()
	tm, _ := ptypes.Timestamp(ev.Time)
	ob.Time = tm.String()
	ob.Url = ev.Url
	ob.Location = ev.Location
//	ob.Risk = ev.Risk

	switch d := ev.Detail.(type) {
		case *evs.Event_DnsMessage:

		msg := d.DnsMessage

		if len(msg.Query) > 0 {
			ob.Dns = &ObDns{}
			query := &ObQuery{}
			for _, val := range msg.Query {
				query.Name = append(query.Name, val.Name)
				query.Type = append(query.Type, val.Type)
				query.Class = append(query.Class, val.Class)
			}
			ob.Dns.Query = query
		}

		if len(msg.Answer) > 0 {
			answer := &ObAnswer{}
			for _, val := range msg.Answer {
				answer.Name = append(answer.Name, val.Name)
				answer.Type = append(answer.Type, val.Type)
				answer.Class = append(answer.Class, val.Class)
				// FIXME:
				//				answer.Address = append(answer.Address,
				//					val.Address)

			}
			ob.Dns.Answer = answer
		}
		break

	case *evs.Event_HttpRequest:
		ob.Http = &ObHttp{
			Method: d.HttpRequest.Method,
			Header: d.HttpRequest.Header,
		}
		break

	case *evs.Event_HttpResponse:
		ob.Http = &ObHttp{
			Status: d.HttpResponse.Status,
			Code:   d.HttpResponse.Code,
			Header: d.HttpResponse.Header,
		}
		break
		
	case *evs.Event_FtpCommand:
		ob.Ftp = &ObFtp{
			Command: d.FtpCommand.Command,
		}
		break
		
	case *evs.Event_FtpResponse:
		ob.Ftp = &ObFtp{
			Status: d.FtpResponse.Status,
			Text:   d.FtpResponse.Text,
		}
		break
		
	case *evs.Event_Icmp:
		ob.Icmp = &ObIcmp{
			Type: d.Icmp.Type,
			Code: d.Icmp.Code,
		}
		break
		
	case *evs.Event_SipRequest:
		ob.Sip = &ObSip{
			Method: d.SipRequest.Method,
			From:   d.SipRequest.From,
			To:     d.SipRequest.To,
		}
		break
		
	case *evs.Event_SipResponse:
		ob.Sip = &ObSip{
			Code:   d.SipResponse.Code,
			Status: d.SipResponse.Status,
			From:   d.SipResponse.From,
			To:     d.SipResponse.To,
		}
		break
		
	case *evs.Event_SmtpCommand:
		ob.Smtp = &ObSmtp{
			Command: d.SmtpCommand.Command,
		}
		break
		
	case *evs.Event_SmtpResponse:
		ob.Smtp = &ObSmtp{
			Status: d.SmtpResponse.Status,
			Text:   d.SmtpResponse.Text,
		}
		break
		
	case *evs.Event_SmtpData:
		ob.Smtp = &ObSmtp{
			From: d.SmtpData.From,
			To:   d.SmtpData.To,
		}
		break
		
	case *evs.Event_NtpTimestamp:
		ob.Ntp = &ObNtp{
			Version: d.NtpTimestamp.Version,
			Mode:    d.NtpTimestamp.Mode,
		}
		break
		
	case *evs.Event_NtpControl:
		ob.Ntp = &ObNtp{
			Version: d.NtpControl.Version,
			Mode:    d.NtpControl.Mode,
		}
		break
		
	case *evs.Event_NtpPrivate:
		ob.Ntp = &ObNtp{
			Version: d.NtpPrivate.Version,
			Mode:    d.NtpPrivate.Mode,
		}
		break
		
	case *evs.Event_UnrecognisedStream:
		ob.Unrecognised = &ObUnrecog{
//			Sha1:   d.UnrecognisedStream.PayloadHash,
//			Length: d.UnrecognisedStream.PayloadLength,
		}
		break
		
	case *evs.Event_UnrecognisedDatagram:
		ob.Unrecognised = &ObUnrecog{
//			Sha1:   d.UnrecognisedDatagram.PayloadHash,
//			Length: d.UnrecognisedDatagram.PayloadLength,
		}
		break

	}

	// Indicators.
	if len(ev.Indicators) > 0 {
		ob.Indicators = &ObIndicators{}
		for _, val := range ev.Indicators {
			ob.Indicators.Id = append(ob.Indicators.Id, val.Id)
			ob.Indicators.Type =
				append(ob.Indicators.Type, val.Type)
			ob.Indicators.Value =
				append(ob.Indicators.Value, val.Value)
			ob.Indicators.Category =
				append(ob.Indicators.Category, val.Category)
			ob.Indicators.Description =
				append(ob.Indicators.Description,
					val.Description)
			ob.Indicators.Author =
				append(ob.Indicators.Author, val.Author)
			ob.Indicators.Source =
				append(ob.Indicators.Source, val.Source)
			// FIXME:
//			ob.Indicators.Probability =
//				append(ob.Indicators.Probability,
//				val.Probability)
		}
	}

	ob.Src = make(map[string][]string, 0)
	ob.Dest = make(map[string][]string, 0)

	for _, val := range ev.Src {
		var cls, addr string
		switch val.Protocol {
		case evs.Protocol_ipv4:
			cls = "ipv4"
			addr = int32ToIp(val.Address.GetIpv4()).String()
			break
		case evs.Protocol_ipv6:
			cls = "ipv6"
			addr = bytesToIp(val.Address.GetIpv6()).String()
			break
		case evs.Protocol_tcp:
			cls = "tcp"
			addr = strconv.Itoa(int(val.Address.GetPort()))
			break
		case evs.Protocol_udp:
			cls = "udp"
			addr = strconv.Itoa(int(val.Address.GetPort()))
		default:
			continue
		}
		if _, ok := ob.Src[cls]; !ok {
			ob.Src[cls] = []string{}
		}
		ob.Src[cls] = append(ob.Src[cls], addr)
	}

	for _, val := range ev.Dest {
		var cls, addr string
		switch val.Protocol {
		case evs.Protocol_ipv4:
			cls = "ipv4"
			addr = int32ToIp(val.Address.GetIpv4()).String()
			break
		case evs.Protocol_ipv6:
			cls = "ipv6"
			addr = bytesToIp(val.Address.GetIpv6()).String()
			break
		case evs.Protocol_tcp:
			cls = "tcp"
			addr = strconv.Itoa(int(val.Address.GetPort()))
			break
		case evs.Protocol_udp:
			cls = "udp"
			addr = strconv.Itoa(int(val.Address.GetPort()))
		default:
			continue
		}
		if _, ok := ob.Dest[cls]; !ok {
			ob.Dest[cls] = []string{}
		}
		ob.Dest[cls] = append(ob.Dest[cls], addr)
	}

	return ob

}
