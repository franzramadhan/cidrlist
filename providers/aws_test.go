package providers

import (
	"net"
	"strings"
	"testing"
)

func TestAwsIps(t *testing.T) {
	t.Parallel()
	aws := AwsInputs{}
	ips, err := aws.AwsIps()
	if err != nil {
		t.Fatal(err)
	}

	for _, prefix := range strings.Split(ips, ",") {
		ipv4Addr, _, err := net.ParseCIDR(prefix)
		if err != nil {
			t.Errorf("AWS IP Address is invalid %s", ipv4Addr)
		}
	}
}
