package resolver

import (
	"net"

	dto "github.com/prometheus/client_model/go"
)

type dnsResolver struct {
}

// NewDNSResolver creates a resolver that uses DNS to find hosts
func NewDNSResolver() Resolver {
	return dnsResolver{}
}

func (r dnsResolver) Resolve(target string) ([]Result, error) {
	addrs, err := net.LookupHost(target)
	if err != nil {
		return nil, err
	}

	var results = make([]Result, 0, len(addrs))
	for _, addr := range addrs {
		results = append(results, Result{IP: addr, Label: createLabelPair("ip", addr)})
	}

	return results, nil
}

func createLabelPair(name string, value string) *dto.LabelPair {
	return &dto.LabelPair{
		Name:  &name,
		Value: &value,
	}
}