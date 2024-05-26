package providers

import (
	"net"
	"testing"
)

func TestAkamaiIps(t *testing.T) {
	t.Parallel()
	ips, err := AkamaiIps()
	if err != nil {
		t.Fatal(err)
	}

	for _, prefix := range ips {
		ipv4Addr, _, err := net.ParseCIDR(prefix)
		if err != nil {
			t.Errorf("Akamai IP Address is invalid %s", ipv4Addr)
		}
	}
}
