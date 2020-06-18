// ES model - presenting an event protobuf as an ElasticSearch event

package main

import (
	evs "github.com/cybermaggedon/evs-golang-api"
	pb "github.com/cybermaggedon/evs-golang-api/protos"
	"github.com/golang/protobuf/ptypes"
)

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
	Id          []string  `json:"id,omitempty"`
	Type        []string  `json:"type,omitempty"`
	Value       []string  `json:"value,omitempty"`
	Description []string  `json:"description,omitempty"`
	Category    []string  `json:"category,omitempty"`
	Author      []string  `json:"author,omitempty"`
	Source      []string  `json:"source,omitempty"`
	Probability []float32 `json:"probability,omitempty"`
}

type ObHttp struct {
	Method string            `json:"method,omitempty"`
	Header map[string]string `json:"header,omitempty"`
	Status string            `json:"status,omitempty"`
	Code   int32             `json:"code,omitempty"`
}

type ObFtp struct {
	Command string   `json:"command,omitempty"`
	Status  int32    `json:"status,omitempty"`
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
	Code   int32  `json:"code,omitempty"`
}

type ObSmtp struct {
	Command string   `json:"command,omitempty"`
	From    string   `json:"from,omitempty"`
	To      []string `json:"to,omitempty"`
	Status  int32    `json:"status,omitempty"`
	Text    []string `json:"text,omitempty"`
}

type ObNtp struct {
	Version int32 `json:"version"`
	Mode    int32 `json:"mode"`
}

type ObUnrecog struct {
	Sha1   string `json:"sha1,omitempty"`
	Length int64  `json:"length,omitempty"`
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

	Location *pb.Locations `json:"location,omitempty"`

	Indicators *ObIndicators `json:"indicators,omitempty"`

	Risk float64 `json:"risk"`

	Operations map[string]string `json:"operations,omitempty"`
}

func Convert(ev *pb.Event) *Observation {

	ob := &Observation{}

	ob.Id = ev.Id
	ob.Action = ev.Action.String()
	ob.Device = ev.Device
	ob.Network = ev.Network
	ob.Origin = ev.Origin.String()
	tm, _ := ptypes.Timestamp(ev.Time)
	ob.Time = tm.Format("2006-01-02T15:04:05.999Z")
	ob.Url = ev.Url
	ob.Location = ev.Location
	//	ob.Risk = ev.Risk

	switch d := ev.Detail.(type) {
	case *pb.Event_DnsMessage:

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
				if answer.Address != nil {
					ad := evs.AddressToString(val.Address)
					answer.Address = append(answer.Address,
						ad)
				} else {
					answer.Address = append(answer.Address,
						"")
				}

			}
			ob.Dns.Answer = answer
		}
		break

	case *pb.Event_HttpRequest:
		ob.Http = &ObHttp{
			Method: d.HttpRequest.Method,
			Header: d.HttpRequest.Header,
		}
		break

	case *pb.Event_HttpResponse:
		ob.Http = &ObHttp{
			Status: d.HttpResponse.Status,
			Code:   d.HttpResponse.Code,
			Header: d.HttpResponse.Header,
		}
		break

	case *pb.Event_FtpCommand:
		ob.Ftp = &ObFtp{
			Command: d.FtpCommand.Command,
		}
		break

	case *pb.Event_FtpResponse:
		ob.Ftp = &ObFtp{
			Status: d.FtpResponse.Status,
			Text:   d.FtpResponse.Text,
		}
		break

	case *pb.Event_Icmp:
		ob.Icmp = &ObIcmp{
			Type: d.Icmp.Type,
			Code: d.Icmp.Code,
		}
		break

	case *pb.Event_SipRequest:
		ob.Sip = &ObSip{
			Method: d.SipRequest.Method,
			From:   d.SipRequest.From,
			To:     d.SipRequest.To,
		}
		break

	case *pb.Event_SipResponse:
		ob.Sip = &ObSip{
			Code:   d.SipResponse.Code,
			Status: d.SipResponse.Status,
			From:   d.SipResponse.From,
			To:     d.SipResponse.To,
		}
		break

	case *pb.Event_SmtpCommand:
		ob.Smtp = &ObSmtp{
			Command: d.SmtpCommand.Command,
		}
		break

	case *pb.Event_SmtpResponse:
		ob.Smtp = &ObSmtp{
			Status: d.SmtpResponse.Status,
			Text:   d.SmtpResponse.Text,
		}
		break

	case *pb.Event_SmtpData:
		ob.Smtp = &ObSmtp{
			From: d.SmtpData.From,
			To:   d.SmtpData.To,
		}
		break

	case *pb.Event_NtpTimestamp:
		ob.Ntp = &ObNtp{
			Version: d.NtpTimestamp.Version,
			Mode:    d.NtpTimestamp.Mode,
		}
		break

	case *pb.Event_NtpControl:
		ob.Ntp = &ObNtp{
			Version: d.NtpControl.Version,
			Mode:    d.NtpControl.Mode,
		}
		break

	case *pb.Event_NtpPrivate:
		ob.Ntp = &ObNtp{
			Version: d.NtpPrivate.Version,
			Mode:    d.NtpPrivate.Mode,
		}
		break

	case *pb.Event_UnrecognisedStream:
		ob.Unrecognised = &ObUnrecog{
			//			Sha1:   d.UnrecognisedStream.PayloadHash,
			//			Length: d.UnrecognisedStream.PayloadLength,
		}
		break

	case *pb.Event_UnrecognisedDatagram:
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
		cls = val.Protocol.String()
		if val.Address != nil {
			addr = evs.AddressToString(val.Address)
		}
		if _, ok := ob.Src[cls]; !ok {
			ob.Src[cls] = []string{}
		}
		ob.Src[cls] = append(ob.Src[cls], addr)
	}

	for _, val := range ev.Dest {
		var cls, addr string
		cls = val.Protocol.String()
		if val.Address != nil {
			addr = evs.AddressToString(val.Address)
		}
		if _, ok := ob.Dest[cls]; !ok {
			ob.Dest[cls] = []string{}
		}
		ob.Dest[cls] = append(ob.Dest[cls], addr)
	}

	return ob

}
