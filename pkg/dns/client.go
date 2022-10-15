package dns

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"time"
)

const (
	defaultServer = "1.1.1.1:53"
)

type dnsResolverClient interface {
	ExchangeContext(ctx context.Context, m *dns.Msg, a string) (r *dns.Msg, rtt time.Duration, err error)
}

type Question struct {
	HostName string
	Type     string
}

type Data struct {
	Type  string
	Key   string
	Value interface{}
}

type ResolverClient struct {
	Client dnsResolverClient
}

func (r *ResolverClient) ExchangeContext(ctx context.Context, questions ...Question) ([]Data, error) {
	query := dns.Msg{
		MsgHdr: dns.MsgHdr{
			Id:               dns.Id(),
			RecursionDesired: true,
		},
		Question: questionsToClientQuestions(questions...),
	}

	message, _, err := r.Client.ExchangeContext(ctx, &query, defaultServer)
	if err != nil {
		return nil, fmt.Errorf("could not get a response from client, with an error: %w", err)
	}

	return messageToResponses(message), nil
}

func messageToResponses(message *dns.Msg) (data []Data) {
	for _, answer := range message.Answer {
		header := answer.Header()
		mappers[header.Rrtype](answer, &data)
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
