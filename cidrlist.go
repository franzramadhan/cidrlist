package cidrlist

import (
	"fmt"

	"github.com/franzramadhan/cidrlist/providers"
)

type Provider string

const (
	ProviderAkamai     Provider = "akamai"
	ProviderAWS        Provider = "aws"
	ProviderCloudflare Provider = "cloudflare"
	ProviderGCP        Provider = "gcp"
	ProviderGoogle     Provider = "google"
)

func GetProviderNames() []string {
	return []string{
		string(ProviderAkamai),
		string(ProviderAWS),
		string(ProviderCloudflare),
		string(ProviderGCP),
		string(ProviderGoogle),
	}
}

func Get(provider Provider) ([]string, error) {
	fnByProviders := map[Provider]func() ([]string, error){
		ProviderAkamai:     providers.AkamaiIps,
		ProviderAWS:        new(providers.AwsInputs).AwsIps,
		ProviderCloudflare: providers.CloudflareIps,
		ProviderGCP:        providers.GcpIps,
		ProviderGoogle:     providers.GoogleIps,
	}

	if fn, ok := fnByProviders[provider]; ok {
		return fn()
	}

	return nil, fmt.Errorf("invalid provider %s", provider)
}
