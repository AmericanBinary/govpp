// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package pg

import (
	"context"

	api "github.com/americanbinary/govpp/api"
)

// RPCService defines RPC service pg.
type RPCService interface {
	PgCapture(ctx context.Context, in *PgCapture) (*PgCaptureReply, error)
	PgCreateInterface(ctx context.Context, in *PgCreateInterface) (*PgCreateInterfaceReply, error)
	PgCreateInterfaceV2(ctx context.Context, in *PgCreateInterfaceV2) (*PgCreateInterfaceV2Reply, error)
	PgEnableDisable(ctx context.Context, in *PgEnableDisable) (*PgEnableDisableReply, error)
	PgInterfaceEnableDisableCoalesce(ctx context.Context, in *PgInterfaceEnableDisableCoalesce) (*PgInterfaceEnableDisableCoalesceReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) PgCapture(ctx context.Context, in *PgCapture) (*PgCaptureReply, error) {
	out := new(PgCaptureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PgCreateInterface(ctx context.Context, in *PgCreateInterface) (*PgCreateInterfaceReply, error) {
	out := new(PgCreateInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PgCreateInterfaceV2(ctx context.Context, in *PgCreateInterfaceV2) (*PgCreateInterfaceV2Reply, error) {
	out := new(PgCreateInterfaceV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PgEnableDisable(ctx context.Context, in *PgEnableDisable) (*PgEnableDisableReply, error) {
	out := new(PgEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PgInterfaceEnableDisableCoalesce(ctx context.Context, in *PgInterfaceEnableDisableCoalesce) (*PgInterfaceEnableDisableCoalesceReply, error) {
	out := new(PgInterfaceEnableDisableCoalesceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
