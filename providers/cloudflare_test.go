package providers

import (
	"net"
	"testing"
)

func TestCloudflareIps(t *testing.T) {
	t.Parallel()
	ips, err := CloudflareIps()
	if err != nil {
		t.Fatal(err)
	}

	for _, prefix := range ips {
		ipv4Addr, _, err := net.ParseCIDR(prefix)
		if err != nil {
			t.Errorf("Cloudflare IP Address is invalid %s", ipv4Addr)
		}
	}
}
