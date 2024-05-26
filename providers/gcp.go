package providers

import (
	"encoding/json"
	"net/http"
	"slices"
)

/*
As an administrator, you can use these lists when you need a range of IP addresses for Google APIs and services' default domains:
IP ranges that Google makes available to users on the internet
Global and regional external IP address ranges for customers' Google Cloud resources
The default domains' IP address ranges for Google APIs and services fit within the list of ranges between these 2 sources. (Subtract the usable ranges from the complete list.)

Source: https://support.google.com/a/answer/10026322?hl=en
*/
const (
	URL_GCP = "https://www.gstatic.com/ipranges/cloud.json"
)

type Gcp struct {
	SyncToken    string `json:"syncToken"`
	CreationTime string `json:"creationTime"`
	Prefixes     []struct {
		Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
		Service    string `json:"service"`
		Scope      string `json:"scope"`
		Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	} `json:"prefixes"`
}

func GcpIps() ([]string, error) {
	var err error
	var client = &http.Client{}
	var data *Gcp
	var prefixes []string

	request, err := http.NewRequest(http.MethodGet, URL_GCP, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	for _, prefix := range data.Prefixes {
		if prefix.Ipv4Prefix != "" {
			prefixes = append(prefixes, prefix.Ipv4Prefix)
		}
		if prefix.Ipv6Prefix != "" {
			prefixes = append(prefixes, prefix.Ipv6Prefix)
		}
	}

	slices.Sort(prefixes)
	prefixes = slices.Compact(prefixes)

	return prefixes, nil
}
