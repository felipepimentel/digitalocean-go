package api

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/felipepimentel/digitalocean-go/internal/config"
)

type Client struct {
	*godo.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		Client: godo.NewFromToken(cfg.DOToken),
	}
}

func (c *Client) ListDroplets(ctx context.Context) ([]godo.Droplet, error) {
	list := []godo.Droplet{}
	opt := &godo.ListOptions{Page: 1, PerPage: 100}

	for {
		droplets, resp, err := c.Droplets.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		list = append(list, droplets...)
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		opt.Page++
	}

	return list, nil
}

func (c *Client) CreateDroplet(ctx context.Context, name, region, size, image string) (*godo.Droplet, error) {
	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   size,
		Image:  godo.DropletCreateImage{Slug: image},
	}

	droplet, _, err := c.Droplets.Create(ctx, createRequest)
	return droplet, err
}

func (c *Client) DeleteDroplet(ctx context.Context, id int) error {
	_, err := c.Droplets.Delete(ctx, id)
	return err
}

func (c *Client) ListVPCs(ctx context.Context) ([]godo.VPC, error) {
	list := []godo.VPC{}
	opt := &godo.ListOptions{Page: 1, PerPage: 100}

	for {
		vpcs, resp, err := c.VPCs.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		for _, vpc := range vpcs {
			list = append(list, *vpc)
		}
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		opt.Page++
	}

	return list, nil
}

func (c *Client) CreateVPC(ctx context.Context, name, region, ipRange string) (*godo.VPC, error) {
	createRequest := &godo.VPCCreateRequest{
		Name:        name,
		RegionSlug:  region,
		IPRange:     ipRange,
		Description: "Created via DigitalOcean CLI",
	}

	vpc, _, err := c.VPCs.Create(ctx, createRequest)
	return vpc, err
}

func (c *Client) DeleteVPC(ctx context.Context, id string) error {
	_, err := c.VPCs.Delete(ctx, id)
	return err
}

func (c *Client) ListKubernetesClusters(ctx context.Context) ([]*godo.KubernetesCluster, error) {
	list := []*godo.KubernetesCluster{}
	opt := &godo.ListOptions{Page: 1, PerPage: 100}

	for {
		clusters, resp, err := c.Kubernetes.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		list = append(list, clusters...)
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		opt.Page++
	}

	return list, nil
}

func (c *Client) CreateKubernetesCluster(ctx context.Context, name, region, version string, numNodes int) (*godo.KubernetesCluster, error) {
	createRequest := &godo.KubernetesClusterCreateRequest{
		Name:        name,
		RegionSlug:  region,
		VersionSlug: version,
		NodePools: []*godo.KubernetesNodePoolCreateRequest{
			{
				Size:  "s-2vcpu-2gb",
				Name:  "worker-pool",
				Count: numNodes,
			},
		},
	}

	cluster, _, err := c.Kubernetes.Create(ctx, createRequest)
	return cluster, err
}

func (c *Client) DeleteKubernetesCluster(ctx context.Context, id string) error {
	_, err := c.Kubernetes.Delete(ctx, id)
	return err
}

// Change the return type from []*godo.Database to []godo.Database
func (c *Client) ListDatabases(ctx context.Context) ([]godo.Database, error) {
	list := []godo.Database{}
	opt := &godo.ListOptions{Page: 1, PerPage: 100}

	for {
		databases, resp, err := c.Databases.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		// Append without converting to pointer
		list = append(list, databases...)
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		opt.Page++
	}

	return list, nil
}

func (c *Client) CreateDatabase(ctx context.Context, name, engine, version, size, region string) (*godo.Database, error) {
	createRequest := &godo.DatabaseCreateRequest{
		Name:       name,
		EngineSlug: engine,
		Version:    version,
		SizeSlug:   size,
		Region:     region,
	}

	database, _, err := c.Databases.Create(ctx, createRequest)
	return database, err
}

func (c *Client) DeleteDatabase(ctx context.Context, id string) error {
	_, err := c.Databases.Delete(ctx, id)
	return err
}

// Change the GetBillingInfo function to use the correct method
func (c *Client) GetBillingInfo(ctx context.Context) (*godo.Balance, error) {
	balance, _, err := c.Balance.Get(ctx)
	return balance, err
}

func (c *Client) ListDomains(ctx context.Context) ([]godo.Domain, error) {
	list := []godo.Domain{}
	opt := &godo.ListOptions{Page: 1, PerPage: 100}

	for {
		domains, resp, err := c.Domains.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		list = append(list, domains...)
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		opt.Page++
	}

	return list, nil
}

func (c *Client) CreateDomain(ctx context.Context, name string) (*godo.Domain, error) {
	createRequest := &godo.DomainCreateRequest{
		Name: name,
	}

	domain, _, err := c.Domains.Create(ctx, createRequest)
	return domain, err
}

func (c *Client) DeleteDomain(ctx context.Context, name string) error {
	_, err := c.Domains.Delete(ctx, name)
	return err
}

func (c *Client) ListDomainRecords(ctx context.Context, domain string) ([]godo.DomainRecord, error) {
	list := []godo.DomainRecord{}
	opt := &godo.ListOptions{Page: 1, PerPage: 100}

	for {
		records, resp, err := c.Domains.Records(ctx, domain, opt)
		if err != nil {
			return nil, err
		}
		list = append(list, records...)
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		opt.Page++
	}

	return list, nil
}

func (c *Client) CreateDomainRecord(ctx context.Context, domain, recordType, name, data string, priority int) (*godo.DomainRecord, error) {
	createRequest := &godo.DomainRecordEditRequest{
		Type:     recordType,
		Name:     name,
		Data:     data,
		Priority: priority,
	}

	record, _, err := c.Domains.CreateRecord(ctx, domain, createRequest)
	return record, err
}

func (c *Client) DeleteDomainRecord(ctx context.Context, domain string, recordID int) error {
	_, err := c.Domains.DeleteRecord(ctx, domain, recordID)
	return err
}