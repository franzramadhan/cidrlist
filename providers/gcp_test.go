package providers

import (
	"net"
	"strings"
	"testing"
)

func TestGcpIps(t *testing.T) {
	t.Parallel()
	ips, err := GcpIps()
	if err != nil {
		t.Fatal(err)
	}

	for _, prefix := range strings.Split(ips, ",") {
		ipv4Addr, _, err := net.ParseCIDR(prefix)
		if err != nil {
			t.Errorf("GCP IP Address is invalid %s", ipv4Addr)
		}
	}
}
