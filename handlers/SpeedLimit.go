package handlers

import "github.com/ivansukach/speed-control-grpc/protocol"

type SpeedLimitControl struct {
	client protocol.SpeedControlServiceClient
}

func NewHandler(client protocol.SpeedControlServiceClient) *SpeedLimitControl {
	return &SpeedLimitControl{client: client}
}
