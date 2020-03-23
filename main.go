package main

import (
	"fmt"
	"github.com/ivansukach/speed-control-grpc/protocol"
	"github.com/ivansukach/speed-control-http/config"
	"github.com/ivansukach/speed-control-http/handlers"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client started")
	cfg := config.Load()
	opts := grpc.WithInsecure() //WithInsecure returns a DialOption which disables transport security for this ClientConn.
	//// Note that transport security is required unless WithInsecure is set.
	clientConnInterface, err := grpc.Dial(cfg.SpeedControlGRPCEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnInterface.Close()
	client := protocol.NewSpeedControlServiceClient(clientConnInterface)
	sls := handlers.NewHandler(client)
	e := echo.New()
	e.POST("/create", sls.Add)
	//e.POST("/update", bs.Update)
	//e.POST("/delete", bs.Delete)
	//e.POST("/getById", bs.GetById)
	//e.POST("/listing", bs.Listing)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
