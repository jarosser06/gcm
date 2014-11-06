package monitoring

import "github.com/rackspace/gophercloud"

func NewMonitoringV1(client *gophercloud.ProviderClient) (*gophercloud.ServiceClient, error) {
	endpoint := "https://monitoring.api.rackspacecloud.com/"

	return &gophercloud.ServiceClient{
		ProviderClient: client,
		Endpoint:       endpoint,
		ResourceBase:   endpoint + "v1.0",
	}, nil
}
