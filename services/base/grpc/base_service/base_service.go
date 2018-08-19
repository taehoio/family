package base_service

import "google.golang.org/grpc"

type IFace interface {
	RegisterService(*grpc.Server)
}
