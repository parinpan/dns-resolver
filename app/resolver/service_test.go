package resolver

import (
	"context"
	"errors"
	"github.com/parinpan/dns-resolver/pkg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestResolver_ResolveNS(t *testing.T) {
	t.Run("returns error when could not fetch nameservers", func(t *testing.T) {
		mockClient := newDnsResolverClientMock(t)

		// mock functions
		mockClient.On("ExchangeContext", mock.Anything, dns.Question{
			HostName: "mock.test",
			Type:     "NS",
		}).Return(nil, errors.New("error"))

		r := &Resolver{Client: mockClient}
		dg, err := r.ResolveNS(context.Background(), dns.Question{
			HostName: "mock.test",
			Type:     "ANT",
		})

		assert.Nil(t, dg)
		assert.Error(t, err)
	})

	t.Run("returns no data when could not resolve records", func(t *testing.T) {
		mockClient := newDnsResolverClientMock(t)

		mockDataNS := []dns.Data{
			{
				Key:   "TARGET",
				Value: "ns1",
			},
			{
				Key:   "TARGET",
				Value: "ns2",
			},
			{
				Key:   "TARGET",
				Value: "ns3",
			},
		}

		// mock functions
		mockClient.On("ExchangeContext", mock.Anything, dns.Question{
			HostName: "mock.test",
			Type:     "NS",
		}).Return(mockDataNS, nil)

		mockClient.On("ExchangeContext", mock.Anything, mock.Anything).
			Return(nil, errors.New("error")).Times(3)

		r := &Resolver{Client: mockClient}
		dg, err := r.ResolveNS(context.Background(), dns.Question{
			HostName: "mock.test",
			Type:     "TXT",
		})

		assert.Len(t, dg, 0)
		assert.NoError(t, err)
	})

	t.Run("returns data when could fetch nameservers", func(t *testing.T) {
		mockClient := newDnsResolverClientMock(t)

		mockDataNS := []dns.Data{
			{
				Key:   "TARGET",
				Value: "ns1",
			},
			{
				Key:   "TARGET",
				Value: "ns2",
			},
			{
				Key:   "TARGET",
				Value: "ns3",
			},
		}

		expectedData := DataGroupWithNS{
			"ns1": {
				"TXT": [][]Data{
					{
						{
							Key:   "test1",
							Value: "test1",
							RType: "TXT",
						},
					},
				},
			},
			"ns2": {
				"TXT": [][]Data{
					{
						{
							Key:   "test2",
							Value: "test2",
							RType: "TXT",
						},
					},
				},
			},
			"ns3": {
				"TXT": [][]Data{
					{
						{
							Key:   "test3",
							Value: "test3",
							RType: "TXT",
						},
					},
				},
			},
		}

		// mock functions
		mockClient.On("ExchangeContext", mock.Anything, dns.Question{
			HostName: "mock.test",
			Type:     "NS",
		}).Return(mockDataNS, nil)

		mockClient.On("ExchangeContext", mock.Anything, dns.Question{
			HostName: "ns1",
			Type:     "TXT",
		}).Return([]dns.Data{{Key: "test1", Value: "test1", Type: "TXT"}}, nil)

		mockClient.On("ExchangeContext", mock.Anything, dns.Question{
			HostName: "ns2",
			Type:     "TXT",
		}).Return([]dns.Data{{Key: "test2", Value: "test2", Type: "TXT"}}, nil)

		mockClient.On("ExchangeContext", mock.Anything, dns.Question{
			HostName: "ns3",
			Type:     "TXT",
		}).Return([]dns.Data{{Key: "test3", Value: "test3", Type: "TXT"}}, nil)

		r := &Resolver{Client: mockClient}
		dg, err := r.ResolveNS(context.Background(), dns.Question{
			HostName: "mock.test",
			Type:     "TXT",
		})

		assert.EqualValues(t, expectedData, dg)
		assert.NoError(t, err)
	})
}
