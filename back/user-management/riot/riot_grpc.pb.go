// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: riot.proto

package riot

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RiotAPIService_GetSummonerByName_FullMethodName                 = "/riot.RiotAPIService/GetSummonerByName"
	RiotAPIService_GetChampionMasteriesBySummoner_FullMethodName    = "/riot.RiotAPIService/GetChampionMasteriesBySummoner"
	RiotAPIService_GetLeagueEntriesBySummoner_FullMethodName        = "/riot.RiotAPIService/GetLeagueEntriesBySummoner"
	RiotAPIService_UpdateSummonerByName_FullMethodName              = "/riot.RiotAPIService/UpdateSummonerByName"
	RiotAPIService_UpdateChampionMasteriesBySummoner_FullMethodName = "/riot.RiotAPIService/UpdateChampionMasteriesBySummoner"
	RiotAPIService_UpdateLeagueEntriesBySummoner_FullMethodName     = "/riot.RiotAPIService/UpdateLeagueEntriesBySummoner"
	RiotAPIService_GetChampionBySummonerAndLane_FullMethodName      = "/riot.RiotAPIService/GetChampionBySummonerAndLane"
	RiotAPIService_GetChampionsByTeams_FullMethodName               = "/riot.RiotAPIService/GetChampionsByTeams"
	RiotAPIService_GetTeams_FullMethodName                          = "/riot.RiotAPIService/GetTeams"
)

// RiotAPIServiceClient is the client API for RiotAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RiotAPIServiceClient interface {
	GetSummonerByName(ctx context.Context, in *SummonerByNameRequest, opts ...grpc.CallOption) (*SummonerResponse, error)
	GetChampionMasteriesBySummoner(ctx context.Context, in *ChampionMasteriesRequest, opts ...grpc.CallOption) (*ChampionMasteriesResponse, error)
	GetLeagueEntriesBySummoner(ctx context.Context, in *LeagueEntriesRequest, opts ...grpc.CallOption) (*LeagueEntriesResponse, error)
	UpdateSummonerByName(ctx context.Context, in *SummonerByNameRequest, opts ...grpc.CallOption) (*SummonerResponse, error)
	UpdateChampionMasteriesBySummoner(ctx context.Context, in *ChampionMasteriesRequest, opts ...grpc.CallOption) (*ChampionMasteriesResponse, error)
	UpdateLeagueEntriesBySummoner(ctx context.Context, in *LeagueEntriesRequest, opts ...grpc.CallOption) (*LeagueEntriesResponse, error)
	GetChampionBySummonerAndLane(ctx context.Context, in *ChampionBySummonerAndLaneRequest, opts ...grpc.CallOption) (*ChampionBySummonerAndLaneResponse, error)
	GetChampionsByTeams(ctx context.Context, in *GetChampionsByTeamsRequest, opts ...grpc.CallOption) (*GetChampionsByTeamsResponse, error)
	GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error)
}

type riotAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRiotAPIServiceClient(cc grpc.ClientConnInterface) RiotAPIServiceClient {
	return &riotAPIServiceClient{cc}
}

func (c *riotAPIServiceClient) GetSummonerByName(ctx context.Context, in *SummonerByNameRequest, opts ...grpc.CallOption) (*SummonerResponse, error) {
	out := new(SummonerResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_GetSummonerByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) GetChampionMasteriesBySummoner(ctx context.Context, in *ChampionMasteriesRequest, opts ...grpc.CallOption) (*ChampionMasteriesResponse, error) {
	out := new(ChampionMasteriesResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_GetChampionMasteriesBySummoner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) GetLeagueEntriesBySummoner(ctx context.Context, in *LeagueEntriesRequest, opts ...grpc.CallOption) (*LeagueEntriesResponse, error) {
	out := new(LeagueEntriesResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_GetLeagueEntriesBySummoner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) UpdateSummonerByName(ctx context.Context, in *SummonerByNameRequest, opts ...grpc.CallOption) (*SummonerResponse, error) {
	out := new(SummonerResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_UpdateSummonerByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) UpdateChampionMasteriesBySummoner(ctx context.Context, in *ChampionMasteriesRequest, opts ...grpc.CallOption) (*ChampionMasteriesResponse, error) {
	out := new(ChampionMasteriesResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_UpdateChampionMasteriesBySummoner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) UpdateLeagueEntriesBySummoner(ctx context.Context, in *LeagueEntriesRequest, opts ...grpc.CallOption) (*LeagueEntriesResponse, error) {
	out := new(LeagueEntriesResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_UpdateLeagueEntriesBySummoner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) GetChampionBySummonerAndLane(ctx context.Context, in *ChampionBySummonerAndLaneRequest, opts ...grpc.CallOption) (*ChampionBySummonerAndLaneResponse, error) {
	out := new(ChampionBySummonerAndLaneResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_GetChampionBySummonerAndLane_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) GetChampionsByTeams(ctx context.Context, in *GetChampionsByTeamsRequest, opts ...grpc.CallOption) (*GetChampionsByTeamsResponse, error) {
	out := new(GetChampionsByTeamsResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_GetChampionsByTeams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riotAPIServiceClient) GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error) {
	out := new(GetTeamsResponse)
	err := c.cc.Invoke(ctx, RiotAPIService_GetTeams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RiotAPIServiceServer is the server API for RiotAPIService service.
// All implementations must embed UnimplementedRiotAPIServiceServer
// for forward compatibility
type RiotAPIServiceServer interface {
	GetSummonerByName(context.Context, *SummonerByNameRequest) (*SummonerResponse, error)
	GetChampionMasteriesBySummoner(context.Context, *ChampionMasteriesRequest) (*ChampionMasteriesResponse, error)
	GetLeagueEntriesBySummoner(context.Context, *LeagueEntriesRequest) (*LeagueEntriesResponse, error)
	UpdateSummonerByName(context.Context, *SummonerByNameRequest) (*SummonerResponse, error)
	UpdateChampionMasteriesBySummoner(context.Context, *ChampionMasteriesRequest) (*ChampionMasteriesResponse, error)
	UpdateLeagueEntriesBySummoner(context.Context, *LeagueEntriesRequest) (*LeagueEntriesResponse, error)
	GetChampionBySummonerAndLane(context.Context, *ChampionBySummonerAndLaneRequest) (*ChampionBySummonerAndLaneResponse, error)
	GetChampionsByTeams(context.Context, *GetChampionsByTeamsRequest) (*GetChampionsByTeamsResponse, error)
	GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsResponse, error)
	mustEmbedUnimplementedRiotAPIServiceServer()
}

// UnimplementedRiotAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRiotAPIServiceServer struct {
}

func (UnimplementedRiotAPIServiceServer) GetSummonerByName(context.Context, *SummonerByNameRequest) (*SummonerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSummonerByName not implemented")
}
func (UnimplementedRiotAPIServiceServer) GetChampionMasteriesBySummoner(context.Context, *ChampionMasteriesRequest) (*ChampionMasteriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionMasteriesBySummoner not implemented")
}
func (UnimplementedRiotAPIServiceServer) GetLeagueEntriesBySummoner(context.Context, *LeagueEntriesRequest) (*LeagueEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeagueEntriesBySummoner not implemented")
}
func (UnimplementedRiotAPIServiceServer) UpdateSummonerByName(context.Context, *SummonerByNameRequest) (*SummonerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSummonerByName not implemented")
}
func (UnimplementedRiotAPIServiceServer) UpdateChampionMasteriesBySummoner(context.Context, *ChampionMasteriesRequest) (*ChampionMasteriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChampionMasteriesBySummoner not implemented")
}
func (UnimplementedRiotAPIServiceServer) UpdateLeagueEntriesBySummoner(context.Context, *LeagueEntriesRequest) (*LeagueEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLeagueEntriesBySummoner not implemented")
}
func (UnimplementedRiotAPIServiceServer) GetChampionBySummonerAndLane(context.Context, *ChampionBySummonerAndLaneRequest) (*ChampionBySummonerAndLaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionBySummonerAndLane not implemented")
}
func (UnimplementedRiotAPIServiceServer) GetChampionsByTeams(context.Context, *GetChampionsByTeamsRequest) (*GetChampionsByTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChampionsByTeams not implemented")
}
func (UnimplementedRiotAPIServiceServer) GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeams not implemented")
}
func (UnimplementedRiotAPIServiceServer) mustEmbedUnimplementedRiotAPIServiceServer() {}

// UnsafeRiotAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RiotAPIServiceServer will
// result in compilation errors.
type UnsafeRiotAPIServiceServer interface {
	mustEmbedUnimplementedRiotAPIServiceServer()
}

func RegisterRiotAPIServiceServer(s grpc.ServiceRegistrar, srv RiotAPIServiceServer) {
	s.RegisterService(&RiotAPIService_ServiceDesc, srv)
}

func _RiotAPIService_GetSummonerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummonerByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).GetSummonerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_GetSummonerByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).GetSummonerByName(ctx, req.(*SummonerByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_GetChampionMasteriesBySummoner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChampionMasteriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).GetChampionMasteriesBySummoner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_GetChampionMasteriesBySummoner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).GetChampionMasteriesBySummoner(ctx, req.(*ChampionMasteriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_GetLeagueEntriesBySummoner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeagueEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).GetLeagueEntriesBySummoner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_GetLeagueEntriesBySummoner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).GetLeagueEntriesBySummoner(ctx, req.(*LeagueEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_UpdateSummonerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummonerByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).UpdateSummonerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_UpdateSummonerByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).UpdateSummonerByName(ctx, req.(*SummonerByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_UpdateChampionMasteriesBySummoner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChampionMasteriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).UpdateChampionMasteriesBySummoner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_UpdateChampionMasteriesBySummoner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).UpdateChampionMasteriesBySummoner(ctx, req.(*ChampionMasteriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_UpdateLeagueEntriesBySummoner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeagueEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).UpdateLeagueEntriesBySummoner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_UpdateLeagueEntriesBySummoner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).UpdateLeagueEntriesBySummoner(ctx, req.(*LeagueEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_GetChampionBySummonerAndLane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChampionBySummonerAndLaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).GetChampionBySummonerAndLane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_GetChampionBySummonerAndLane_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).GetChampionBySummonerAndLane(ctx, req.(*ChampionBySummonerAndLaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_GetChampionsByTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChampionsByTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).GetChampionsByTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_GetChampionsByTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).GetChampionsByTeams(ctx, req.(*GetChampionsByTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiotAPIService_GetTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiotAPIServiceServer).GetTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RiotAPIService_GetTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiotAPIServiceServer).GetTeams(ctx, req.(*GetTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RiotAPIService_ServiceDesc is the grpc.ServiceDesc for RiotAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RiotAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "riot.RiotAPIService",
	HandlerType: (*RiotAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSummonerByName",
			Handler:    _RiotAPIService_GetSummonerByName_Handler,
		},
		{
			MethodName: "GetChampionMasteriesBySummoner",
			Handler:    _RiotAPIService_GetChampionMasteriesBySummoner_Handler,
		},
		{
			MethodName: "GetLeagueEntriesBySummoner",
			Handler:    _RiotAPIService_GetLeagueEntriesBySummoner_Handler,
		},
		{
			MethodName: "UpdateSummonerByName",
			Handler:    _RiotAPIService_UpdateSummonerByName_Handler,
		},
		{
			MethodName: "UpdateChampionMasteriesBySummoner",
			Handler:    _RiotAPIService_UpdateChampionMasteriesBySummoner_Handler,
		},
		{
			MethodName: "UpdateLeagueEntriesBySummoner",
			Handler:    _RiotAPIService_UpdateLeagueEntriesBySummoner_Handler,
		},
		{
			MethodName: "GetChampionBySummonerAndLane",
			Handler:    _RiotAPIService_GetChampionBySummonerAndLane_Handler,
		},
		{
			MethodName: "GetChampionsByTeams",
			Handler:    _RiotAPIService_GetChampionsByTeams_Handler,
		},
		{
			MethodName: "GetTeams",
			Handler:    _RiotAPIService_GetTeams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "riot.proto",
}
