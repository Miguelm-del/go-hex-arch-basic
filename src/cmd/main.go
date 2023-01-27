package main

import (
	"github.com/Miguelm-del/go-hex-arch-basic/src/internal/adapters/app/api"
	"github.com/Miguelm-del/go-hex-arch-basic/src/internal/adapters/core/arithmetic"
	gRPC "github.com/Miguelm-del/go-hex-arch-basic/src/internal/adapters/framework/left/grpc"
	"github.com/Miguelm-del/go-hex-arch-basic/src/internal/adapters/framework/right/db"
	"github.com/Miguelm-del/go-hex-arch-basic/src/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	var dbaseAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection %v", err)
	}
	defer dbaseAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()
}
