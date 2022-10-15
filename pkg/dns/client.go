package dns

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"time"
)

const (
	defaultServer = "8.8.4.4:53"
	udpBlockSize  = 1024 * 10
)

type dnsResolverClient interface {
	ExchangeContext(ctx context.Context, m *dns.Msg, a string) (r *dns.Msg, rtt time.Duration, err error)
}

type Question struct {
	HostName string
	Type     string
}

type Data struct {
	ID    string
	Type  string
	Key   string
	Value interface{}
}

type ResolverClient struct {
	Client dnsResolverClient
}

func (r *ResolverClient) ExchangeContext(ctx context.Context, question Question) ([]Data, error) {
	query := dns.Msg{
		MsgHdr: dns.MsgHdr{
			Id:                 dns.Id(),
			RecursionAvailable: true,
			RecursionDesired:   true,
			Opcode:             dns.OpcodeQuery,
		},
		Question: questionsToClientQuestions(question),
	}

	query.SetEdns0(udpBlockSize, true)

	message, _, err := r.Client.ExchangeContext(ctx, &query, defaultServer)
	if err != nil {
		return nil, fmt.Errorf("could not get a response from client, with an error: %w", err)
	}

	return messageToResponses(message, question), nil
}

func messageToResponses(message *dns.Msg, question Question) (data []Data) {
	for _, answer := range message.Answer {
		header := answer.Header()

		if !(dns.StringToType[question.Type] == header.Rrtype || dns.StringToType[question.Type] == dns.TypeANY) {
			continue
		}

		if mapper, ok := mappers[header.Rrtype]; ok {
			mapper(answer, &data)
		} else {
			mappers[dns.TypeANY](answer, &data)
		}
	}
	return data
}

func questionsToClientQuestions(questions ...Question) (qs []dns.Question) {
	for _, question := range questions {
		qs = append(qs, dns.Question{
			Name:   dns.Fqdn(question.HostName),
			Qtype:  dns.StringToType[question.Type],
			Qclass: dns.ClassINET,
		})
	}
	return qs
}
