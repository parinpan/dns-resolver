package server

import (
	"context"
	"github.com/miekg/dns"
	"github.com/parinpan/dns-resolver/app/resolver"
	dnsPkg "github.com/parinpan/dns-resolver/pkg/dns"
	"log"
	"net/http"
)

func Start(ctx context.Context, address string) error {
	errChan := make(chan error)

	http.HandleFunc("/resolve", resolver.Handler(&resolver.Resolver{
		Client: &dnsPkg.ResolverClient{
			Client: &dns.Client{},
		},
	}))

	server := &http.Server{
		Addr:    address,
		Handler: http.DefaultServeMux,
	}

	go func() {
		defer func() {
			if err := server.Shutdown(ctx); err != nil {
				log.Fatalln(err.Error())
			}
		}()

		errChan <- server.ListenAndServe()
	}()

	return <-errChan
}
