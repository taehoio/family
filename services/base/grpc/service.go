package grpc

import "google.golang.org/grpc"

type IFace interface {
	RegisterService(*grpc.Server)
}
