// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package dns

import (
	"context"

	api "github.com/americanbinary/govpp/api"
)

// RPCService defines RPC service dns.
type RPCService interface {
	DNSEnableDisable(ctx context.Context, in *DNSEnableDisable) (*DNSEnableDisableReply, error)
	DNSNameServerAddDel(ctx context.Context, in *DNSNameServerAddDel) (*DNSNameServerAddDelReply, error)
	DNSResolveIP(ctx context.Context, in *DNSResolveIP) (*DNSResolveIPReply, error)
	DNSResolveName(ctx context.Context, in *DNSResolveName) (*DNSResolveNameReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) DNSEnableDisable(ctx context.Context, in *DNSEnableDisable) (*DNSEnableDisableReply, error) {
	out := new(DNSEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DNSNameServerAddDel(ctx context.Context, in *DNSNameServerAddDel) (*DNSNameServerAddDelReply, error) {
	out := new(DNSNameServerAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DNSResolveIP(ctx context.Context, in *DNSResolveIP) (*DNSResolveIPReply, error) {
	out := new(DNSResolveIPReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DNSResolveName(ctx context.Context, in *DNSResolveName) (*DNSResolveNameReply, error) {
	out := new(DNSResolveNameReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
