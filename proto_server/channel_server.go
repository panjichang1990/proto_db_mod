package proto_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"time"
)

type ChannelServer interface {
	CheckToken(ctx context.Context, in *OpenPre, opts ...client.CallOption) (*OperationRep, error)
	PhoneRegOrLogin(ctx context.Context, in *PhoneRegOrLoginReq, opts ...client.CallOption) (*RegOrLoginRep, error)
	AccountRegOrLogin(ctx context.Context, in *AccountRegOrLoginReq, opts ...client.CallOption) (*RegOrLoginRep, error)
	SendCode(ctx context.Context, in *SendCodeReq, opts ...client.CallOption) (*OperationRep, error)
	GetAccounts(ctx context.Context, in *GetAccountsReq, opts ...client.CallOption) (*GetAccountRep, error)
	BindPhone(ctx context.Context, in *BindPhoneReq, opts ...client.CallOption) (*OperationRep, error)
	RealAuth(ctx context.Context, in *RealAuthReq, opts ...client.CallOption) (*OperationRep, error)
}

type ChannelServerHandler interface {
	CheckToken(context.Context, *OpenPre, *OperationRep) error
	PhoneRegOrLogin(context.Context, *PhoneRegOrLoginReq, *RegOrLoginRep) error
	AccountRegOrLogin(context.Context, *AccountRegOrLoginReq, *RegOrLoginRep) error
	SendCode(context.Context, *SendCodeReq, *OperationRep) error
	GetAccounts(context.Context, *GetAccountsReq, *GetAccountRep) error
	BindPhone(context.Context, *BindPhoneReq, *OperationRep) error
	RealAuth(context.Context, *RealAuthReq, *OperationRep) error
}

type channelHandler struct {
	ChannelServerHandler
}

func (h *channelHandler) CheckToken(ctx context.Context, in *OpenPre, out *OperationRep) error {
	return h.ChannelServerHandler.CheckToken(ctx, in, out)
}

func (h *channelHandler) PhoneRegOrLogin(ctx context.Context, in *PhoneRegOrLoginReq, out *RegOrLoginRep) error {
	return h.ChannelServerHandler.PhoneRegOrLogin(ctx, in, out)
}

func (h *channelHandler) AccountRegOrLogin(ctx context.Context, in *AccountRegOrLoginReq, out *RegOrLoginRep) error {
	return h.ChannelServerHandler.AccountRegOrLogin(ctx, in, out)
}

func (h *channelHandler) SendCode(ctx context.Context, in *SendCodeReq, out *OperationRep) error {
	return h.ChannelServerHandler.SendCode(ctx, in, out)
}

func (h *channelHandler) GetAccounts(ctx context.Context, in *GetAccountsReq, out *GetAccountRep) error {
	return h.ChannelServerHandler.GetAccounts(ctx, in, out)
}

func (h *channelHandler) BindPhone(ctx context.Context, in *BindPhoneReq, out *OperationRep) error {
	return h.ChannelServerHandler.BindPhone(ctx, in, out)
}

func (h *channelHandler) RealAuth(ctx context.Context, in *RealAuthReq, out *OperationRep) error {
	return h.ChannelServerHandler.RealAuth(ctx, in, out)
}

func RegisterChannelHandler(s server.Server, hdlr ChannelServerHandler, opts ...server.HandlerOption) {
	type channelServer interface {
		CheckToken(context.Context, *OpenPre, *OperationRep) error
		PhoneRegOrLogin(context.Context, *PhoneRegOrLoginReq, *RegOrLoginRep) error
		AccountRegOrLogin(context.Context, *AccountRegOrLoginReq, *RegOrLoginRep) error
		SendCode(context.Context, *SendCodeReq, *OperationRep) error
		GetAccounts(context.Context, *GetAccountsReq, *GetAccountRep) error
		BindPhone(context.Context, *BindPhoneReq, *OperationRep) error
		RealAuth(context.Context, *RealAuthReq, *OperationRep) error
	}
	type ChannelServer struct {
		channelServer
	}
	h := &channelHandler{hdlr}
	err := s.Handle(s.NewHandler(&ChannelServer{h}, opts...))
	if err != nil {
		fmt.Println(err)
	}
}

type channelService struct {
	c           client.Client
	serviceName string
}

func NewChannelService(serviceName string) ChannelServer {
	reg := consul.NewRegistry()
	cc := client.NewClient(client.Registry(reg), client.DialTimeout(1*time.Second))
	return &channelService{
		c:           cc,
		serviceName: serviceName,
	}
}

func (c *channelService) CheckToken(ctx context.Context, in *OpenPre, opts ...client.CallOption) (*OperationRep, error) {
	req := c.c.NewRequest(c.serviceName, "ChannelServer.CheckToken", in)
	out := new(OperationRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelService) PhoneRegOrLogin(ctx context.Context, in *PhoneRegOrLoginReq, opts ...client.CallOption) (*RegOrLoginRep, error) {
	req := c.c.NewRequest(c.serviceName, "ChannelServer.PhoneRegOrLogin", in)
	out := new(RegOrLoginRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelService) AccountRegOrLogin(ctx context.Context, in *AccountRegOrLoginReq, opts ...client.CallOption) (*RegOrLoginRep, error) {
	req := c.c.NewRequest(c.serviceName, "ChannelServer.AccountRegOrLogin", in)
	out := new(RegOrLoginRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelService) SendCode(ctx context.Context, in *SendCodeReq, opts ...client.CallOption) (*OperationRep, error) {
	req := c.c.NewRequest(c.serviceName, "ChannelServer.SendCode", in)
	out := new(OperationRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelService) GetAccounts(ctx context.Context, in *GetAccountsReq, opts ...client.CallOption) (*GetAccountRep, error) {
	req := c.c.NewRequest(c.serviceName, "ChannelServer.GetAccounts", in)
	out := new(GetAccountRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelService) BindPhone(ctx context.Context, in *BindPhoneReq, opts ...client.CallOption) (*OperationRep, error) {
	req := c.c.NewRequest(c.serviceName, "ChannelServer.BindPhone", in)
	out := new(OperationRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelService) RealAuth(ctx context.Context, in *RealAuthReq, opts ...client.CallOption) (*OperationRep, error) {
	req := c.c.NewRequest(c.serviceName, "ChannelServer.RealAuth", in)
	out := new(OperationRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
