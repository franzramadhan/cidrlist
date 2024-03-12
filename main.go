package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/franzramadhan/cidrlist/providers"
)

var (
	app          = kingpin.New("cidrlist", "A command-line tool to fetch IP address ranges from well known providers.")
	getIp        = app.Command("get", "Get IP from provider (cloudflare, akamai, aws, gcp, and google).")
	providerName = getIp.Arg("provider", "The name of the provider.").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case getIp.FullCommand():
		switch name := *providerName; name {
		case "akamai":
			ips, err := providers.AkamaiIps()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s:\n%s", strings.ToUpper(name), ips)
		case "aws":
			aws := providers.AwsInputs{}
			ips, err := aws.AwsIps()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s:\n%s", strings.ToUpper(name), ips)
		case "cloudflare":
			ips, err := providers.CloudflareIps()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s:\n%s", strings.ToUpper(name), ips)
		case "gcp":
			ips, err := providers.GcpIps()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s:\n%s", strings.ToUpper(name), ips)
		case "google":
			ips, err := providers.GoogleIps()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s:\n%s", strings.ToUpper(name), ips)

		case "":
			log.Fatal("invalid provider")
		}
	}
}
