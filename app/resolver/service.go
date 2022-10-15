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
	RType string `json:"-"`
}

type DataGroup map[string][][]Data

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
	group := make(map[string][]Data)

	for _, item := range data {
		group[item.ID] = append(group[item.ID], Data{
			Key:   item.Key,
			Value: fmt.Sprint(item.Value),
			RType: item.Type,
		})
	}

	for _, items := range group {
		if len(items) > 0 {
			rtype := items[0].RType
			dg[rtype] = append(dg[rtype], items)
		}
	}

	return dg
}
