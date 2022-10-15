package dns

import (
	"fmt"
	"github.com/miekg/dns"
)

type converter func(answer dns.RR, data *[]Data)

var mappers = map[uint16]converter{
	dns.TypeA:      typeA,
	dns.TypeAAAA:   typeAAAA,
	dns.TypeCAA:    typeCAA,
	dns.TypeCNAME:  typeCNAME,
	dns.TypeDNSKEY: typeDNSKEY,
	dns.TypeDS:     typeDS,
	dns.TypeMX:     typeMX,
	dns.TypeNS:     typeNS,
	dns.TypePTR:    typePTR,
	dns.TypeSOA:    typeSOA,
	dns.TypeSRV:    typeSRV,
	dns.TypeTLSA:   typeTLSA,
	dns.TypeTXT:    typeTXT,
}

func typeA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.A)
	appendData(data, Data{Type: "A", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "A", Key: "DATA", Value: ans.A.String()})
}

func typeAAAA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.AAAA)
	appendData(data, Data{Type: "AAAA", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "AAAA", Key: "DATA", Value: ans.AAAA.String()})
}

func typeCAA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.CAA)
	value := fmt.Sprint(ans.Flag, " ", ans.Tag, fmt.Sprintf(" \"%s\"", ans.Value))
	appendData(data, Data{Type: "CAA", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "CAA", Key: "DATA", Value: value})
}

func typeCNAME(answer dns.RR, data *[]Data) {

}

func typeDNSKEY(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.DNSKEY)
	appendData(data, Data{Type: "DNSKEY", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "DNSKEY", Key: "FLAGS", Value: ans.Flags})
	appendData(data, Data{Type: "DNSKEY", Key: "ALGORITHM", Value: ans.Algorithm})
	appendData(data, Data{Type: "DNSKEY", Key: "PROTOCOL", Value: ans.Protocol})
	appendData(data, Data{Type: "DNSKEY", Key: "KEY", Value: ans.PublicKey})
}

func typeDS(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.DS)
	appendData(data, Data{Type: "DS", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "DS", Key: "KEYTAG", Value: ans.KeyTag})
	appendData(data, Data{Type: "DS", Key: "ALGORITHM", Value: ans.Algorithm})
	appendData(data, Data{Type: "DS", Key: "DIGEST TYPE", Value: ans.DigestType})
	appendData(data, Data{Type: "DS", Key: "DIGEST", Value: ans.Digest})
}

func typeMX(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.MX)
	appendData(data, Data{Type: "DS", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "DS", Key: "EXCHANGE", Value: ans.Mx})
	appendData(data, Data{Type: "DS", Key: "PREFERENCE", Value: ans.Preference})

}

func typeNS(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.NS)
	appendData(data, Data{Type: "DS", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "DS", Key: "TARGET", Value: ans.Ns})
}

func typePTR(answer dns.RR, data *[]Data) {

}

func typeSOA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.SOA)
	appendData(data, Data{Type: "DS", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "DS", Key: "MNAME", Value: ans.Ns})
	appendData(data, Data{Type: "DS", Key: "RNAME", Value: ans.Mbox})
	appendData(data, Data{Type: "DS", Key: "SERIAL", Value: ans.Serial})
	appendData(data, Data{Type: "DS", Key: "REFRESH", Value: ans.Refresh})
	appendData(data, Data{Type: "DS", Key: "RETRY", Value: ans.Retry})
	appendData(data, Data{Type: "DS", Key: "EXPIRE", Value: ans.Expire})
}

func typeSRV(answer dns.RR, data *[]Data) {

}

func typeTLSA(answer dns.RR, data *[]Data) {

}

func typeTXT(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.TXT)
	appendData(data, Data{Type: "DS", Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{Type: "DS", Key: "VALUE", Value: ans.Txt})
}

func appendData(data *[]Data, new Data) {
	*data = append(*data, new)
}
