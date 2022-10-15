package resolver

import (
	"context"
	"fmt"
	"github.com/parinpan/dns-resolver/pkg/dns"
)

type dnsResolverClient interface {
	ExchangeContext(ctx context.Context, question dns.Question) ([]dns.Data, error)
}

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DataGroup map[string][]Data

type Resolver struct {
	Client dnsResolverClient
}

func (r *Resolver) Resolve(ctx context.Context, question dns.Question) (DataGroup, error) {
	data, err := r.Client.ExchangeContext(ctx, question)
	if err != nil {
		return nil, err
	}

	return dataToDataGroup(data), nil
}

func dataToDataGroup(data []dns.Data) DataGroup {
	dg := make(DataGroup)

	for _, item := range data {
		if _, ok := dg[item.Key]; !ok {
			dg[item.Key] = []Data{}
		}

		dg[item.Key] = append(dg[item.Type], Data{
			Key:   item.Key,
			Value: fmt.Sprint(item.Value),
		})
	}

	return dg
}
