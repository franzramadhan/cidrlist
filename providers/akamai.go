package providers

import (
	"strings"
)

/*
Origin IP Access Control List (Origin IP ACL) offers some protection for your origin server by restricting traffic to ​Akamai​-controlled IP addresses.
​Akamai​ maintains a small and stable list of IP addresses that you use in policy rules in your origin server's firewall.
These IP addresses are represented in a list using classless inter-domain routing (CIDR).
CIDR is an IP addressing scheme that improves the allocation of IP addresses by using a single IP address with a prefix at the end to designate many, unique IP addresses. With Origin IP ACL, requests from edge servers to your origin will always be sourced from an address in one of these prefixes.

Reference: https://techdocs.akamai.com/origin-ip-acl/docs/welcome
*/
var AkamaiCidrs = []string{
	"23.32.0.0/11",
	"23.192.0.0/11",
	"2.16.0.0/13",
	"104.64.0.0/10",
	"184.24.0.0/13",
	"23.0.0.0/12",
	"95.100.0.0/15",
	"92.122.0.0/15",
	"184.50.0.0/15",
	"88.221.0.0/16",
	"23.64.0.0/14",
	"72.246.0.0/15",
	"96.16.0.0/15",
	"96.6.0.0/15",
	"69.192.0.0/16",
	"23.72.0.0/13",
	"173.222.0.0/15",
	"118.214.0.0/16",
	"184.84.0.0/14",
	"2a02:26f0::/32",
	"2600:1400::/24",
	"2405:9600::/32",
}

func AkamaiIps() (string, error) {
	results := strings.Join(AkamaiCidrs, ",")
	return results, nil
}
