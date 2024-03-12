package providers

import (
	"encoding/json"
	"net/http"
	"strings"
)

/*

Get IPs used on the Cloudflare/JD Cloud network, see https://www.cloudflare.com/ips for Cloudflare IPs or https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD Cloud IPs.

Source: https://developers.cloudflare.com/api/operations/cloudflare-i-ps-cloudflare-ip-details
*/

const (
	URL_CLOUDFLARE = "https://api.cloudflare.com/client/v4/ips"
)

type Cloudflare struct {
	Result struct {
		Ipv4Cidrs []string `json:"ipv4_cidrs"`
		Ipv6Cidrs []string `json:"ipv6_cidrs"`
		Etag      string   `json:"etag"`
	} `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

func CloudflareIps() (string, error) {
	var err error
	var client = &http.Client{}
	var data *Cloudflare

	request, err := http.NewRequest(http.MethodGet, URL_CLOUDFLARE, nil)
	if err != nil {
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	results := strings.Join(append(data.Result.Ipv4Cidrs, data.Result.Ipv6Cidrs...), ",")

	return results, nil
}
