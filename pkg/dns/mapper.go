package dns

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/miekg/dns"
)

type converter func(answer dns.RR, data *[]Data)

var mappers = map[uint16]converter{
	dns.TypeANY:    typeANY,
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
	dns.TypeTSIG:   typeTSIG,
	dns.TypeRRSIG:  typeRSIG,
}

func typeA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.A)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "DATA", Value: ans.A.String()})
}

func typeAAAA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.AAAA)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "DATA", Value: ans.AAAA.String()})
}

func typeCAA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.CAA)
	id := uuid.New()
	value := fmt.Sprint(ans.Flag, " ", ans.Tag, fmt.Sprintf(" \"%s\"", ans.Value))
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "DATA", Value: value})
}

func typeCNAME(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.CNAME)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TARGET", Value: ans.Target})
}

func typeDNSKEY(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.DNSKEY)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "FLAGS", Value: ans.Flags})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "ALGORITHM", Value: ans.Algorithm})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "PROTOCOL", Value: ans.Protocol})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "KEY", Value: ans.PublicKey})
}

func typeDS(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.DS)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "KEYTAG", Value: ans.KeyTag})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "ALGORITHM", Value: ans.Algorithm})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "DIGEST TYPE", Value: ans.DigestType})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "DIGEST", Value: ans.Digest})
}

func typeMX(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.MX)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "EXCHANGE", Value: ans.Mx})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "PREFERENCE", Value: ans.Preference})
}

func typeNS(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.NS)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TARGET", Value: ans.Ns})
}

func typePTR(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.PTR)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "PTR", Value: ans.Ptr})
}

func typeSOA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.SOA)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "MNAME", Value: ans.Ns})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "RNAME", Value: ans.Mbox})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "SERIAL", Value: ans.Serial})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "REFRESH", Value: ans.Refresh})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "RETRY", Value: ans.Retry})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "EXPIRE", Value: ans.Expire})
}

func typeSRV(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.SRV)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TARGET", Value: ans.Target})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "PORT", Value: ans.Port})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "PRIORITY", Value: ans.Priority})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "WEIGHT", Value: ans.Weight})
}

func typeTLSA(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.TLSA)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "MATCHING TYPE", Value: ans.MatchingType})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "USAGE", Value: ans.Usage})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "SELECTOR", Value: ans.Selector})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "CERTIFICATE", Value: ans.Certificate})
}

func typeTXT(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.TXT)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "VALUE", Value: ans.Txt})
}

func typeANY(answer dns.RR, data *[]Data) {
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: answer.Header().Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "DATA", Value: answer.Header().String()})
}

func typeTSIG(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.TSIG)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "ALGORITHM", Value: ans.Algorithm})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "MAC", Value: ans.MAC})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "DATA", Value: ans.OtherData})
}

func typeRSIG(answer dns.RR, data *[]Data) {
	ans := answer.(*dns.RRSIG)
	id := uuid.New()
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "TTL", Value: ans.Hdr.Ttl})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "MAC", Value: dns.Type(ans.TypeCovered).String()})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "ALGORITHM", Value: ans.Algorithm})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "LABELS", Value: ans.Labels})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "EXPIRATION", Value: dns.TimeToString(ans.Expiration)})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "INCEPTION", Value: dns.TimeToString(ans.Inception)})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "KEY TAG", Value: dns.TimeToString(ans.Inception)})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "SIGNER", Value: ans.SignerName})
	appendData(data, Data{ID: id.String(), Type: answerRType(answer), Key: "SIGNATURE", Value: ans.Signature})
}

func answerRType(answer dns.RR) string {
	return dns.TypeToString[answer.Header().Rrtype]
}

func appendData(data *[]Data, new Data) {
	*data = append(*data, new)
}
