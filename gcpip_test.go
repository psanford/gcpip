package gcpip

import (
	"fmt"
	"net/netip"
	"testing"
)

func TestGcpip(t *testing.T) {
	gcpIPs := []netip.Addr{
		netip.MustParseAddr("34.1.208.21"),
		netip.MustParseAddr("2600:1900:4280::1"),
	}

	for _, addr := range gcpIPs {
		if !IsGcpIP(addr) {
			t.Errorf("Expected %s to match gcp ip but did not", addr)
		}
	}

	nonGcpIPs := []netip.Addr{
		netip.MustParseAddr("127.0.0.12"),
		netip.MustParseAddr("10.48.20.96"),
		netip.MustParseAddr("8.8.8.8"),
		netip.MustParseAddr("2a05:d03a:8000::1"),
	}
	for _, addr := range nonGcpIPs {
		if IsGcpIP(addr) {
			t.Errorf("%s is not an GCP ip address, but it matched", addr)
		}
	}
}

func BenchmarkLookup(b *testing.B) {
	tests := []struct {
		ip    string
		isGCP bool
	}{
		{"2a05:d07f:e0ff::ffff", true},
		{"54.74.0.27", true},
		{"2a05:d03a:8000::1", true},
		{"100.23.255.254", true},
		{"57.180.0.0", true},
		{"127.0.0.12", false},
		{"10.48.20.96", false},
		{"8.8.8.8", false},
		{"2606:4700:4700::1111", false},
	}

	for _, tc := range tests {
		ip := netip.MustParseAddr(tc.ip)
		b.Run(tc.ip, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isGcp := IsGcpIP(ip)

				if isGcp != tc.isGCP {
					b.Fatalf("%s got isGcp=%t expected=%t", tc.ip, isGcp, tc.isGCP)
				}
			}
		})
	}
}

func ExampleRange() {
	ip := netip.MustParseAddr("104.155.192.1")
	r := Range(ip)
	fmt.Println(r.Prefix)
	fmt.Println(r.Service)
	fmt.Println(r.Scope)
	// Output:
	// 104.155.192.0/19
	// Google Cloud
	// asia-east1
}

func ExampleIsGcpIP() {
	ips := []netip.Addr{
		netip.MustParseAddr("104.155.192.1"),
		netip.MustParseAddr("54.74.0.27"),
		netip.MustParseAddr("127.0.0.1"),
	}
	for _, ip := range ips {
		if IsGcpIP(ip) {
			fmt.Printf("%s is GCP\n", ip)
		} else {
			fmt.Printf("%s is NOT GCP\n", ip)
		}
	}
	// Output:
	// 104.155.192.1 is GCP
	// 54.74.0.27 is NOT GCP
	// 127.0.0.1 is NOT GCP
}
