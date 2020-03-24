package handlers

import (
	"context"
	"fmt"
	"github.com/ivansukach/speed-control-grpc/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (slc *SpeedLimitControl) Listing(c echo.Context) error {
	log.Info("Listing")
	record := new(RecordModel)
	if err := c.Bind(record); err != nil {
		log.Errorf("echo.Context Error GetMinMax %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	resp, err := slc.client.Listing(context.Background(), &protocol.ListingRequest{Date: record.Date,
		SpeedLimit: record.Speed})
	if err != nil {
		log.Errorf("gRPC Error Listing %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	for _, value := range resp.Record {
		fmt.Println(value)
	}
	return c.JSON(http.StatusOK, resp.Record)
}
