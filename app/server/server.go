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

	r := &resolver.Resolver{
		Client: &dnsPkg.ResolverClient{
			Client: &dns.Client{},
		},
	}

	staticDir := http.Dir("web/build")
	http.Handle("/", http.FileServer(staticDir))

	http.HandleFunc("/resolve", allowCorsMiddleware(resolver.Handler(r)))
	http.HandleFunc("/resolve/ns", allowCorsMiddleware(resolver.HandlerNS(r)))

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

func allowCorsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		isPreflight := r.Method == http.MethodOptions &&
			r.Header.Get("Access-Control-Request-Method") != ""

		if isPreflight {
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.WriteHeader(http.StatusNoContent)
		} else {
			next(w, r)
		}
	}
}
