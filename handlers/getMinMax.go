package handlers

import (
	"context"
	"github.com/ivansukach/speed-control-grpc/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (slc *SpeedLimitControl) GetMinMax(c echo.Context) error {
	log.Info("GetMinMax")
	record := new(RecordModel)
	if err := c.Bind(record); err != nil {
		log.Errorf("echo.Context Error GetMinMax %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	resp, err := slc.client.GetMinMax(context.Background(), &protocol.GetMinMaxRequest{Date: record.Date})
	if err != nil {
		log.Errorf("gRPC Error GetMinMax %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	log.Println(resp.Min)
	log.Println(resp.Max)
	obj := make([]*protocol.Record, 0)
	obj = append(obj, resp.Min)
	obj = append(obj, resp.Max)
	return c.JSON(http.StatusOK, obj)
}
