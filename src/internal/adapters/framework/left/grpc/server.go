package rpc

import (
	"github.com/Miguelm-del/go-hex-arch-basic/src/internal/adapters/framework/left/grpc/pb"
	"github.com/Miguelm-del/go-hex-arch-basic/src/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}
