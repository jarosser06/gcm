package cloudmonitoring

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
)

const RaxEndpoint = "https://monitoring.api.rackspacecloud.com/v1.0/"

type Client struct {
	ProviderClient *gophercloud.ProviderClient
	Endpoint       string
	TenantID       string
}

func NewClient(authOpts gophercloud.AuthOptions) Client {
	provider, err := rackspace.AuthenticatedClient(authOpts)

	if err != nil {
		panic(err)
	}

	client := Client{
		ProviderClient: provider,
		Endpoint:       RaxEndpoint,
	}

	return client
}
