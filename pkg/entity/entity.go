package entity

import (
	"errors"

	cloudmonitoring "github.com/jarosser06/gcm/pkg"
	"github.com/racker/perigee"
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

func GetEntity(client cloudmonitoring.Client, entityId int) (Entity, error) {
	if client.EmptyTenant {
		return _, errors.New("Tenant cannot be empty")
	}

	var res CreateResult

	reqUrl := fmt.sprintf(
		"%sentities/%s/",
		cloudmonitoring.monitoringEndpoint,
		entityId,
	)

	_, res.Err = perigee.Request("GET", reqUrl, perigee.Options{
		MoreHeaders: client.AuthenticatedHeaders(),
		Results:     res.Body,
		OkCodes:     []int{200},
	})
}
