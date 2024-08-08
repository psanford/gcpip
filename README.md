# gcpip: a Go package to check if an IP belongs to GCP

gcpip is a Go package that allows you to determine if an IP address belongs to GCP.

A cli tool is also included in `cmd/gcpip` for easily checking the status of an ip address.

## Example:

```

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
```

CLI:
```
$ $ gcpip 104.155.192.1
{
  "Prefix": "104.155.192.0/19",
  "Service": "Google Cloud",
  "Scope": "asia-east1"
}
```

## Updating the ip ranges

To update the ip ranges run: `go generate`. This will fetch from https://www.gstatic.com/ipranges/cloud.json and regenerate the `ips.gen.go` file.

## License

MIT
