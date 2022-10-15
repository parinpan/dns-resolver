package resolver

import (
	"context"
	"encoding/json"
	"github.com/parinpan/dns-resolver/pkg/dns"
	"net/http"
)

type resolverService interface {
	Resolve(ctx context.Context, question dns.Question) (DataGroup, error)
}

type Request struct {
	HostName   string `json:"hostname"`
	RecordType string `json:"record_type"`
}

type Response struct {
	ErrorMessage *string     `json:"error_message,omitempty"`
	StatusCode   int         `json:"status_code"`
	Data         interface{} `json:"data"`
}

func Handler(service resolverService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeResponse(w, Response{
				ErrorMessage: stringToRef("could not parse the request"),
				StatusCode:   http.StatusInternalServerError,
				Data:         nil,
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
				Data:         nil,
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
		w.Write([]byte(`{"error_message": "unexpected error happened", "status_code": 500, "data": []}`))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(bytes)
}

func stringToRef(s string) *string {
	return &s
}
