package cloudflare

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zones"
)

type CloudflareClient struct {
	client *cloudflare.Client
}

func New(token string) *CloudflareClient {
	return &CloudflareClient{
		client: cloudflare.NewClient(option.WithAPIToken(token)),
	}
}

func (c *CloudflareClient) EnsureARecord(domain string, ip string) (dns.RecordResponse, error) {
	baseDomain := getBaseDomain(domain)
	zone, err := c.findZone(baseDomain)

	if err != nil {
		return dns.RecordResponse{}, err
	}

	record, err := c.findARecord(zone.ID, domain)

	if err != nil {
		return dns.RecordResponse{}, err
	}

	if record.ID == "" {
		newRecord, err := c.createARecord(zone.ID, domain, ip)

		if err != nil {
			return dns.RecordResponse{}, err
		}

		record = *newRecord
	}

	if record.Content != ip {
		_, err = c.updateARecord(record.ID, zone.ID, domain, ip)

		if err != nil {
			return dns.RecordResponse{}, err
		}
	}

	return record, nil
}

func (c *CloudflareClient) CreateARecord(domain string, ip string) (*dns.RecordResponse, error) {
	baseDomain := getBaseDomain(domain)
	zone, err := c.findZone(baseDomain)

	if err != nil {
		return nil, err
	}

	return c.createARecord(zone.ID, domain, ip)
}

func (c *CloudflareClient) RemoveARecord(domain string) (*dns.RecordDeleteResponse, error) {
	baseDomain := getBaseDomain(domain)
	zone, err := c.findZone(baseDomain)

	if err != nil {
		return nil, err
	}

	record, err := c.findARecord(zone.ID, domain)

	if err != nil {
		return nil, err
	}

	return c.removeARecord(zone.ID, record.ID)
}

func (c *CloudflareClient) FindARecord(domain string) (dns.RecordResponse, error) {
	baseDomain := getBaseDomain(domain)
	zone, err := c.findZone(baseDomain)

	if err != nil {
		return dns.RecordResponse{}, err
	}

	return c.findARecord(zone.ID, domain)
}

func (c *CloudflareClient) findZone(domain string) (zones.Zone, error) {
	response, err := c.client.Zones.List(context.TODO(), zones.ZoneListParams{
		Name: cloudflare.String(domain),
	})

	if err != nil {
		return zones.Zone{}, err
	}

	if len(response.Result) == 0 {
		return zones.Zone{}, fmt.Errorf("No results for zone %s", domain)
	}

	zone := response.Result[0]

	return zone, err
}

func (c *CloudflareClient) findARecord(zoneId string, domain string) (dns.RecordResponse, error) {
	response, err := c.client.DNS.Records.List(context.TODO(), dns.RecordListParams{
		ZoneID: cloudflare.String(zoneId),
		Name: cloudflare.F(dns.RecordListParamsName{
			Exact: cloudflare.String(domain),
		}),
	})

	if err != nil {
		return dns.RecordResponse{}, err
	}

	if len(response.Result) == 0 {
		return dns.RecordResponse{}, nil
	}

	record := response.Result[0]

	return record, err
}

func (c *CloudflareClient) removeARecord(zoneId string, recordId string) (*dns.RecordDeleteResponse, error) {
	return c.client.DNS.Records.Delete(context.TODO(), recordId, dns.RecordDeleteParams{
		ZoneID: cloudflare.String(zoneId),
	})
}

func (c *CloudflareClient) createARecord(zoneId string, domain string, ip string) (*dns.RecordResponse, error) {
	return c.client.DNS.Records.New(context.TODO(), dns.RecordNewParams{
		ZoneID: cloudflare.String(zoneId),
		Body: dns.ARecordParam{
			Type:    cloudflare.F(dns.ARecordTypeA),
			Name:    cloudflare.String(domain),
			Content: cloudflare.String(ip),
		},
	})
}

func (c *CloudflareClient) updateARecord(dnsRecordId string, zoneId string, domain string, ip string) (*dns.RecordResponse, error) {
	return c.client.DNS.Records.Update(context.TODO(), dnsRecordId, dns.RecordUpdateParams{
		ZoneID: cloudflare.String(zoneId),
		Body: dns.ARecordParam{
			Type:    cloudflare.F(dns.ARecordTypeA),
			Name:    cloudflare.String(domain),
			Content: cloudflare.String(ip),
		},
	})
}

func getBaseDomain(full string) string {
	parts := strings.Split(full, ".")

	if len(parts) < 3 {
		return strings.Join(parts, ".")
	}

	return strings.Join(parts[len(parts)-2:], ".")
}
