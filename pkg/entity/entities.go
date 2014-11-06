package entity

import (
	"errors"
	"fmt"

	"github.com/racker/perigee"
	"github.com/rackspace/gophercloud"
)

type Entity struct {
	ID          int               `json:"id"`
	Label       string            `json:"label"`
	AgentID     string            `json:"agent_id,omitempty"`
	IPAddresses map[string]string `json:"ip_addresses,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type EntityList struct {
	Values   []Entity          `json:"values"`
	Metadata map[string]string `json:"metadata"`
}

func GetEntity(client gophercloud.ServiceClient, entityId int) (Entity, error) {
	var entity Entity
	reqUrl := fmt.Sprintf(
		"%s/entities/%s/",
		client.ResourceBase,
		entityId,
	)

	_, err := perigee.Request("GET", reqUrl, perigee.Options{
		MoreHeaders: client.ProviderClient.AuthenticatedHeaders(),
		Results:     entity,
		OkCodes:     []int{200},
	})

	if err != nil {
		return Entity{}, errors.New(
			fmt.Sprintf("Error during http request: %v", err),
		)
	}

	return entity, nil
}
