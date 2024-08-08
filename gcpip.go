package gcpip

import (
	"net/netip"
)

//go:generate go run update_ips.go

// IsGcpIP returns true if the ip address falls within one of the known
// GCP ip ranges.
func IsGcpIP(ip netip.Addr) bool {
	r := Range(ip)
	return r != nil
}

// Range returns the ip range and metadata an address falls within.
// If the IP is not an GCP IP address it returns nil
func Range(ip netip.Addr) *IPRange {
	_, r, ok := cidrTbl.Lookup(ip)
	if ok {
		return &r
	}
	return nil
}

type IPRange struct {
	Prefix  netip.Prefix
	Service string
	Scope   string
}
