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
	e.POST("/add", sls.Add)
	e.POST("/getMinMax", sls.GetMinMax)
	e.POST("/listing", sls.Listing)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
