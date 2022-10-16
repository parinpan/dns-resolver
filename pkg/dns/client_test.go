package dns

import (
	"context"
	"errors"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestResolverClient_ExchangeContext(t *testing.T) {
	t.Run("returns error when client returned an error", func(t *testing.T) {
		mockClient := newDnsResolverClientMock(t)

		// mock function
		mockClient.On("ExchangeContext", mock.Anything, mock.Anything, defaultServer).
			Return(nil, time.Duration(0), errors.New("error"))

		client := &ResolverClient{Client: mockClient}
		data, err := client.ExchangeContext(context.Background(), Question{})

		assert.Nil(t, data)
		assert.Error(t, err)
	})

	t.Run("returns data when client returned proper response", func(t *testing.T) {
		mockHeader := dns.RR_Header{
			Name: "dns.mock",
			Ttl:  100,
		}

		injectHeader := func(rType uint16, header dns.RR_Header) dns.RR_Header {
			header.Rrtype = rType
			return header
		}

		mockResponse := &dns.Msg{
			Answer: []dns.RR{
				&dns.ANY{Hdr: injectHeader(dns.TypeANY, mockHeader)},
				&dns.A{Hdr: injectHeader(dns.TypeA, mockHeader)},
				&dns.AAAA{Hdr: injectHeader(dns.TypeAAAA, mockHeader)},
				&dns.CAA{Hdr: injectHeader(dns.TypeCAA, mockHeader)},
				&dns.CNAME{Hdr: injectHeader(dns.TypeCNAME, mockHeader)},
				&dns.DNSKEY{Hdr: injectHeader(dns.TypeDNSKEY, mockHeader)},
				&dns.DS{Hdr: injectHeader(dns.TypeDS, mockHeader)},
				&dns.MX{Hdr: injectHeader(dns.TypeMX, mockHeader)},
				&dns.NS{Hdr: injectHeader(dns.TypeNS, mockHeader)},
				&dns.PTR{Hdr: injectHeader(dns.TypePTR, mockHeader)},
				&dns.SOA{Hdr: injectHeader(dns.TypeSOA, mockHeader)},
				&dns.SRV{Hdr: injectHeader(dns.TypeSRV, mockHeader)},
				&dns.TLSA{Hdr: injectHeader(dns.TypeTLSA, mockHeader)},
				&dns.TXT{Hdr: injectHeader(dns.TypeTXT, mockHeader)},
				&dns.TSIG{Hdr: injectHeader(dns.TypeTSIG, mockHeader)},
				&dns.RRSIG{Hdr: injectHeader(dns.TypeRRSIG, mockHeader)},
				&dns.APL{Hdr: injectHeader(dns.TypeAPL, mockHeader)},
			},
		}

		for _, answer := range mockResponse.Answer {
			mockClient := newDnsResolverClientMock(t)

			// mock function
			mockClient.On("ExchangeContext", mock.Anything, mock.Anything, defaultServer).
				Return(&dns.Msg{Answer: []dns.RR{answer}}, time.Duration(10), nil)

			challenge := make([]Data, 0)
			client := &ResolverClient{Client: mockClient}

			data, err := client.ExchangeContext(context.Background(), Question{
				HostName: "test",
				Type:     dns.TypeToString[answer.Header().Rrtype],
			})

			if mapper, ok := mappers[answer.Header().Rrtype]; ok {
				mapper(answer, &challenge)
			} else {
				mappers[dns.TypeANY](answer, &challenge)
			}

			// remove id from data and challenge
			for i := 0; i < len(data); i++ {
				challenge[i].ID = ""
				data[i].ID = ""
			}

			assert.Equal(t, len(data), len(challenge))
			assert.EqualValues(t, data, challenge)
			assert.NoError(t, err)
		}
	})
}
