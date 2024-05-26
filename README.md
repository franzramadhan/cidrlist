# cidrlist

Simple cli to get CIDR list from well known providers

## Available Providers

| Provider name | Constant name | Source of IP addresses |
| :- | :- | :- |
| `akamai` | `cidrlist.ProviderAkamai` | [Akamai Origin IP ACL](https://techdocs.akamai.com/origin-ip-acl/docs/welcome) |
| `aws` | `cidrlist.ProviderAWS` | [AWS IP Address Ranges](https://docs.aws.amazon.com/vpc/latest/userguide/aws-ip-ranges.html) |
| `cloudflare` | `cidrlist.ProviderCloudflare` | [Cloudflare IP Ranges](https://www.cloudflare.com/ips/) |
| `gcp` | `cidrlist.ProviderGCP` | [Google Cloud Global & Regional IP Address Ranges](https://support.google.com/a/answer/10026322?hl=en) |
| `google` | `cidrlist.ProviderGoogle` | [Google IP Address Ranges](https://support.google.com/a/answer/10026322?hl=en) |

## Usage

Install the library using the following command:

```go
go get -u github.com/franzramadhan/cidrlist
```

Implementation example:

```go
package main

import (
	"fmt"

	"github.com/franzramadhan/cidrlist"
)

func main() {
	ips, err := cidrlist.Get(cidrlist.ProviderAWS)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("IP of provider %s:\n", cidrlist.ProviderAWS)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
```

## CLI Usage

Install cidrlist CLI using the following command:

```go
go install github.com/franzramadhan/cidrlist/cmd/cidrlist@latest
```

Usage:

```go
cidrlist get <provider-name>
cidrlist get cloudflare
cidrlist get gcp
```
