// Code generated by Kitex v0.9.1. DO NOT EDIT.

package followservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	follow "tiktok/kitex_gen/follow"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Action": kitex.NewMethodInfo(
		actionHandler,
		newFollowServiceActionArgs,
		newFollowServiceActionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListFollowing": kitex.NewMethodInfo(
		listFollowingHandler,
		newFollowServiceListFollowingArgs,
		newFollowServiceListFollowingResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListFollower": kitex.NewMethodInfo(
		listFollowerHandler,
		newFollowServiceListFollowerArgs,
		newFollowServiceListFollowerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListFriend": kitex.NewMethodInfo(
		listFriendHandler,
		newFollowServiceListFriendArgs,
		newFollowServiceListFriendResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	followServiceServiceInfo                = NewServiceInfo()
	followServiceServiceInfoForClient       = NewServiceInfoForClient()
	followServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return followServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return followServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return followServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "FollowService"
	handlerType := (*follow.FollowService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "follow",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func actionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceActionArgs)
	realResult := result.(*follow.FollowServiceActionResult)
	success, err := handler.(follow.FollowService).Action(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceActionArgs() interface{} {
	return follow.NewFollowServiceActionArgs()
}

func newFollowServiceActionResult() interface{} {
	return follow.NewFollowServiceActionResult()
}

func listFollowingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceListFollowingArgs)
	realResult := result.(*follow.FollowServiceListFollowingResult)
	success, err := handler.(follow.FollowService).ListFollowing(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceListFollowingArgs() interface{} {
	return follow.NewFollowServiceListFollowingArgs()
}

func newFollowServiceListFollowingResult() interface{} {
	return follow.NewFollowServiceListFollowingResult()
}

func listFollowerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceListFollowerArgs)
	realResult := result.(*follow.FollowServiceListFollowerResult)
	success, err := handler.(follow.FollowService).ListFollower(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceListFollowerArgs() interface{} {
	return follow.NewFollowServiceListFollowerArgs()
}

func newFollowServiceListFollowerResult() interface{} {
	return follow.NewFollowServiceListFollowerResult()
}

func listFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*follow.FollowServiceListFriendArgs)
	realResult := result.(*follow.FollowServiceListFriendResult)
	success, err := handler.(follow.FollowService).ListFriend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFollowServiceListFriendArgs() interface{} {
	return follow.NewFollowServiceListFriendArgs()
}

func newFollowServiceListFriendResult() interface{} {
	return follow.NewFollowServiceListFriendResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Action(ctx context.Context, req *follow.ActionReq) (r *follow.ActionResp, err error) {
	var _args follow.FollowServiceActionArgs
	_args.Req = req
	var _result follow.FollowServiceActionResult
	if err = p.c.Call(ctx, "Action", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFollowing(ctx context.Context, req *follow.ListFollowingReq) (r *follow.ListFollowingResp, err error) {
	var _args follow.FollowServiceListFollowingArgs
	_args.Req = req
	var _result follow.FollowServiceListFollowingResult
	if err = p.c.Call(ctx, "ListFollowing", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFollower(ctx context.Context, req *follow.ListFollowerReq) (r *follow.ListFollowerResp, err error) {
	var _args follow.FollowServiceListFollowerArgs
	_args.Req = req
	var _result follow.FollowServiceListFollowerResult
	if err = p.c.Call(ctx, "ListFollower", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFriend(ctx context.Context, req *follow.ListFriendReq) (r *follow.ListFriendResp, err error) {
	var _args follow.FollowServiceListFriendArgs
	_args.Req = req
	var _result follow.FollowServiceListFriendResult
	if err = p.c.Call(ctx, "ListFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
