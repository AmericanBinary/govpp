// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package rdma

import (
	"context"

	api "github.com/americanbinary/govpp/api"
)

// RPCService defines RPC service rdma.
type RPCService interface {
	RdmaCreate(ctx context.Context, in *RdmaCreate) (*RdmaCreateReply, error)
	RdmaCreateV2(ctx context.Context, in *RdmaCreateV2) (*RdmaCreateV2Reply, error)
	RdmaCreateV3(ctx context.Context, in *RdmaCreateV3) (*RdmaCreateV3Reply, error)
	RdmaDelete(ctx context.Context, in *RdmaDelete) (*RdmaDeleteReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) RdmaCreate(ctx context.Context, in *RdmaCreate) (*RdmaCreateReply, error) {
	out := new(RdmaCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) RdmaCreateV2(ctx context.Context, in *RdmaCreateV2) (*RdmaCreateV2Reply, error) {
	out := new(RdmaCreateV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) RdmaCreateV3(ctx context.Context, in *RdmaCreateV3) (*RdmaCreateV3Reply, error) {
	out := new(RdmaCreateV3Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) RdmaDelete(ctx context.Context, in *RdmaDelete) (*RdmaDeleteReply, error) {
	out := new(RdmaDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
