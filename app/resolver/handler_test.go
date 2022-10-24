package resolver

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/parinpan/dns-resolver/pkg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerNS(t *testing.T) {
	t.Run("returns 400 when request could not be parsed", func(t *testing.T) {
		mockService := newResolverServiceMock(t)
		mux := http.NewServeMux()
		mux.HandleFunc("/resolve/ns", HandlerNS(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve/ns", bytes.NewReader([]byte(`!!!!!`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns 400 when request is not complete", func(t *testing.T) {
		mockService := newResolverServiceMock(t)
		mux := http.NewServeMux()
		mux.HandleFunc("/resolve/ns", HandlerNS(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve/ns", bytes.NewReader([]byte(`{}`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns 500 when service returned error", func(t *testing.T) {
		mockService := newResolverServiceMock(t)

		// mock functions
		mockService.On("ResolveNS", mock.Anything, dns.Question{
			HostName: "fachr.in",
			Type:     "ANY",
		}).Return(nil, errors.New("error"))

		mux := http.NewServeMux()
		mux.HandleFunc("/resolve/ns", HandlerNS(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve/ns", bytes.NewReader([]byte(`{
			"hostname": "fachr.in",
			"record_type": "ANY"
		}`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("returns 200 when service returned a response", func(t *testing.T) {
		mockService := newResolverServiceMock(t)

		expectedData := DataGroupWithNS{
			"ns1": {
				"TXT": {
					{
						Data{
							Key:   "test1",
							Value: "test1",
							RType: "TEXT",
						},
					},
				},
			},
		}

		// mock functions
		mockService.On("ResolveNS", mock.Anything, dns.Question{
			HostName: "fachr.in",
			Type:     "ANY",
		}).Return(expectedData, nil)

		mux := http.NewServeMux()
		mux.HandleFunc("/resolve/ns", HandlerNS(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve/ns", bytes.NewReader([]byte(`{
			"hostname": "fachr.in",
			"record_type": "ANY"
		}`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		expectedResponse := Response{
			ErrorMessage: nil,
			StatusCode:   http.StatusOK,
			Data:         expectedData,
		}

		expectedBytes, err := json.Marshal(expectedResponse)
		actualBytes, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expectedBytes, actualBytes)
	})
}

func TestHandler(t *testing.T) {
	t.Run("returns 400 when request could not be parsed", func(t *testing.T) {
		mockService := newResolverServiceMock(t)
		mux := http.NewServeMux()
		mux.HandleFunc("/resolve", Handler(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve", bytes.NewReader([]byte(`!!!!!`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns 400 when request is not complete", func(t *testing.T) {
		mockService := newResolverServiceMock(t)
		mux := http.NewServeMux()
		mux.HandleFunc("/resolve", Handler(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve", bytes.NewReader([]byte(`{}`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("returns 500 when service returned error", func(t *testing.T) {
		mockService := newResolverServiceMock(t)

		// mock functions
		mockService.On("Resolve", mock.Anything, dns.Question{
			HostName: "fachr.in",
			Type:     "ANY",
		}).Return(nil, errors.New("error"))

		mux := http.NewServeMux()
		mux.HandleFunc("/resolve", Handler(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve", bytes.NewReader([]byte(`{
			"hostname": "fachr.in",
			"record_type": "ANY"
		}`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("returns 200 when service returned a response", func(t *testing.T) {
		mockService := newResolverServiceMock(t)

		expectedData := DataGroup{
			"TXT": {
				{
					Data{
						Key:   "test1",
						Value: "test1",
						RType: "TEXT",
					},
				},
			},
		}

		// mock functions
		mockService.On("Resolve", mock.Anything, dns.Question{
			HostName: "fachr.in",
			Type:     "ANY",
		}).Return(expectedData, nil)

		mux := http.NewServeMux()
		mux.HandleFunc("/resolve", Handler(mockService))

		srv := httptest.NewServer(mux)
		defer srv.Close()

		req, err := http.NewRequest("POST", srv.URL+"/resolve", bytes.NewReader([]byte(`{
			"hostname": "fachr.in",
			"record_type": "ANY"
		}`)))
		assert.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer resp.Body.Close()

		expectedResponse := Response{
			ErrorMessage: nil,
			StatusCode:   http.StatusOK,
			Data:         expectedData,
		}

		expectedBytes, err := json.Marshal(expectedResponse)
		actualBytes, err := io.ReadAll(resp.Body)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expectedBytes, actualBytes)
	})
}
