// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package tls_openssl

import (
	"context"

	api "github.com/americanbinary/govpp/api"
)

// RPCService defines RPC service tls_openssl.
type RPCService interface {
	TLSOpensslSetEngine(ctx context.Context, in *TLSOpensslSetEngine) (*TLSOpensslSetEngineReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) TLSOpensslSetEngine(ctx context.Context, in *TLSOpensslSetEngine) (*TLSOpensslSetEngineReply, error) {
	out := new(TLSOpensslSetEngineReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
