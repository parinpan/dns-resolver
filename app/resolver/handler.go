package resolver

import (
	"context"
	"encoding/json"
	"github.com/parinpan/dns-resolver/pkg/dns"
	"net/http"
)

type resolverService interface {
	ResolveNS(ctx context.Context, question dns.Question) (DataGroupWithNS, error)
	Resolve(ctx context.Context, question dns.Question) (DataGroup, error)
}

type Request struct {
	HostName   string `json:"hostname"`
	RecordType string `json:"record_type"`
}

type Response struct {
	ErrorMessage *string     `json:"error_message"`
	StatusCode   int         `json:"status_code"`
	Data         interface{} `json:"data"`
}

func HandlerNS(service resolverService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeResponse(w, Response{
				ErrorMessage: stringToRef("could not parse the request"),
				StatusCode:   http.StatusInternalServerError,
			})
			return
		}

		if req.HostName == "" || req.RecordType == "" {
			writeResponse(w, Response{
				ErrorMessage: stringToRef("both hostname and record type must be specified in the request"),
				StatusCode:   http.StatusBadRequest,
			})
			return
		}

		data, err := service.ResolveNS(r.Context(), dns.Question{
			HostName: req.HostName,
			Type:     req.RecordType,
		})

		if err != nil {
			writeResponse(w, Response{
				ErrorMessage: stringToRef("could not resolve the dns, error: " + err.Error()),
				StatusCode:   http.StatusInternalServerError,
			})
			return
		}

		writeResponse(w, Response{
			StatusCode: http.StatusOK,
			Data:       data,
		})
	}
}

func Handler(service resolverService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeResponse(w, Response{
				ErrorMessage: stringToRef("could not parse the request"),
				StatusCode:   http.StatusInternalServerError,
			})
			return
		}

		if req.HostName == "" || req.RecordType == "" {
			writeResponse(w, Response{
				ErrorMessage: stringToRef("both hostname and record type must be specified in the request"),
				StatusCode:   http.StatusBadRequest,
			})
			return
		}

		data, err := service.Resolve(r.Context(), dns.Question{
			HostName: req.HostName,
			Type:     req.RecordType,
		})

		if err != nil {
			writeResponse(w, Response{
				ErrorMessage: stringToRef("could not resolve the dns, error: " + err.Error()),
				StatusCode:   http.StatusInternalServerError,
			})
			return
		}

		writeResponse(w, Response{
			StatusCode: http.StatusOK,
			Data:       data,
		})
	}
}

func writeResponse(w http.ResponseWriter, resp Response) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(`{"error_message": "unexpected error happened", "status_code": 500, "data": null}`))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(bytes)
}

func stringToRef(s string) *string {
	return &s
}
