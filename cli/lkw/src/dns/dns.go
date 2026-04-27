package dns

import (
	"github.com/medvesek/infrastructure/lkw/lib/cloudflare"
	"github.com/medvesek/infrastructure/lkw/src/baseconfig"
)

func SetupDomain(domain string) error {
	cloudflareToken := baseconfig.CloudflareToken()
	ip := baseconfig.ServerIp()

	cloudflareClient := cloudflare.New(cloudflareToken)

	_, err := cloudflareClient.EnsureARecord(domain, ip)

	return err
}

func SetupDomains(domains []string) error {
	for _, domain := range domains {
		if err := SetupDomain(domain); err != nil {
			return err
		}
	}
	return nil
}
