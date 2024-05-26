package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/franzramadhan/cidrlist"
)

var (
	app          = kingpin.New("cidrlist", "A command-line tool to fetch IP address ranges from well known providers.")
	getIp        = app.Command("get", fmt.Sprintf("Get IP from provider (%s).", strings.Join(cidrlist.GetProviderNames(), ", ")))
	providerName = getIp.Arg("provider", "The name of the provider.").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case getIp.FullCommand():
		target := strings.ToLower(*providerName)
		provider := cidrlist.Provider(target)
		ips, err := cidrlist.Get(provider)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s:\n%s", strings.ToUpper(target), ips)
	}
}
