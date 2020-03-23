package handlers

import (
	"context"
	"github.com/ivansukach/speed-control-grpc/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (slc *SpeedLimitControl) Add(c echo.Context) error {
	log.Info("Add")
	record := new(RecordModel)

	if err := c.Bind(record); err != nil {
		log.Errorf("echo.Context Error Create %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	r := protocol.Record{
		Date:   record.Date,
		Number: record.Number,
		Speed:  record.Speed,
	}

	_, err := slc.client.Add(context.Background(),
		&protocol.AddRequest{Record: &r})
	if err != nil {
		log.Errorf("gRPC Error Add %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
