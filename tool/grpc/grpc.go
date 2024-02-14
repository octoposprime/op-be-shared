package tgrpc

import (
	"context"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return handler(ctx, req)
	}
	userIds := md.Get(string(smodel.QueryKeyUid))
	if len(userIds) > 0 {
		ctx = context.WithValue(ctx, smodel.QueryKeyUid, userIds[0])
	}
	return handler(ctx, req)
}
