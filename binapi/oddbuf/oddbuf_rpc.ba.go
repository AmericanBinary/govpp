// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package oddbuf

import (
	"context"

	api "github.com/americanbinary/govpp/api"
)

// RPCService defines RPC service oddbuf.
type RPCService interface {
	OddbufEnableDisable(ctx context.Context, in *OddbufEnableDisable) (*OddbufEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) OddbufEnableDisable(ctx context.Context, in *OddbufEnableDisable) (*OddbufEnableDisableReply, error) {
	out := new(OddbufEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
