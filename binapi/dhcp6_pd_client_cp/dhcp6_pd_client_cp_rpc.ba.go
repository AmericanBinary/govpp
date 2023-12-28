// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package dhcp6_pd_client_cp

import (
	"context"

	api "github.com/americanbinary/govpp/api"
)

// RPCService defines RPC service dhcp6_pd_client_cp.
type RPCService interface {
	DHCP6PdClientEnableDisable(ctx context.Context, in *DHCP6PdClientEnableDisable) (*DHCP6PdClientEnableDisableReply, error)
	IP6AddDelAddressUsingPrefix(ctx context.Context, in *IP6AddDelAddressUsingPrefix) (*IP6AddDelAddressUsingPrefixReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) DHCP6PdClientEnableDisable(ctx context.Context, in *DHCP6PdClientEnableDisable) (*DHCP6PdClientEnableDisableReply, error) {
	out := new(DHCP6PdClientEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IP6AddDelAddressUsingPrefix(ctx context.Context, in *IP6AddDelAddressUsingPrefix) (*IP6AddDelAddressUsingPrefixReply, error) {
	out := new(IP6AddDelAddressUsingPrefixReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
