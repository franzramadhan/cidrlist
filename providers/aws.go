package providers

import (
	"encoding/json"
	"net/http"
	"slices"
)

/*

AWS publishes its current IP address ranges in JSON format. With this information, you can identify traffic from AWS. You can also use this information to allow or deny traffic to or from some AWS services.

Reference: https://docs.aws.amazon.com/vpc/latest/userguide/aws-ip-ranges.html

*/

const (
	URL_AWS = "https://ip-ranges.amazonaws.com/ip-ranges.json"
)

type Aws struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix           string `json:"ip_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"prefixes"`
	Ipv6Prefixes []struct {
		Ipv6Prefix         string `json:"ipv6_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"ipv6_prefixes"`
}

type AwsInputs struct {
	IpType string
	// Region  string
	// Service string
}

func (a *AwsInputs) AwsIps() ([]string, error) {
	var err error
	var client = &http.Client{}
	var data *Aws
	var prefixes []string

	request, err := http.NewRequest(http.MethodGet, URL_AWS, nil)
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

	switch ipType := a.IpType; ipType {
	case "":
		for _, prefix := range data.Prefixes {
			prefixes = append(prefixes, prefix.IPPrefix)
		}

		for _, prefix := range data.Ipv6Prefixes {
			prefixes = append(prefixes, prefix.Ipv6Prefix)
		}
	case "ipv4":
		for _, prefix := range data.Prefixes {
			prefixes = append(prefixes, prefix.IPPrefix)
		}
	case "ipv6":
		for _, prefix := range data.Ipv6Prefixes {
			prefixes = append(prefixes, prefix.Ipv6Prefix)
		}
	}

	slices.Sort(prefixes)
	prefixes = slices.Compact(prefixes)

	return prefixes, nil
}
