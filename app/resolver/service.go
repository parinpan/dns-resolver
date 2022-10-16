package resolver

import (
	"context"
	"fmt"
	"github.com/parinpan/dns-resolver/pkg/dns"
	"log"
)

type DataGroup map[string][][]Data
type DataGroupWithNS map[string]DataGroup

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	RType string `json:"-"`
}

type dnsResolverClient interface {
	ExchangeContext(ctx context.Context, question dns.Question) ([]dns.Data, error)
}

type Resolver struct {
	Client dnsResolverClient
}

func (r *Resolver) ResolveNS(ctx context.Context, question dns.Question) (DataGroupWithNS, error) {
	dg := make(DataGroupWithNS)

	nameservers, err := r.fetchNameservers(ctx, question.HostName)
	if err != nil {
		return nil, err
	}

	type nsAnswer struct {
		ns  string
		dg  DataGroup
		err error
	}

	nsLen := len(nameservers)
	nsChan := make(chan nsAnswer, nsLen)

	for _, ns := range nameservers {
		go func(ns string) {
			result, err := r.Resolve(ctx, dns.Question{
				HostName: ns,
				Type:     question.Type,
			})

			nsChan <- nsAnswer{
				ns:  ns,
				dg:  result,
				err: err,
			}
		}(ns)
	}

	for i := 0; i < nsLen; i++ {
		answer := <-nsChan
		dg[answer.ns] = answer.dg
	}

	return dg, nil
}

func (r *Resolver) Resolve(ctx context.Context, question dns.Question) (DataGroup, error) {
	data, err := r.Client.ExchangeContext(ctx, question)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return dataToDataGroup(data), nil
}

func (r *Resolver) fetchNameservers(ctx context.Context, hostname string) ([]string, error) {
	data, err := r.Client.ExchangeContext(ctx, dns.Question{
		HostName: hostname,
		Type:     "NS",
	})

	if err != nil {
		return nil, err
	}

	return groupNS(data), nil
}

func groupNS(data []dns.Data) []string {
	var ns []string

	for _, item := range data {
		if item.Key == "TARGET" {
			ns = append(ns, fmt.Sprint(item.Value))
		}
	}

	return ns
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
