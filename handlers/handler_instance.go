package handlers

import (
	"github.com/rs/zerolog"
	"github.com/zakisk/socio-server/data"
	"github.com/zakisk/socio-server/models"
)

type Handler struct {
	log       zerolog.Logger
	dbHandler *data.DBHandler
}

func NewHandlerInstance(log zerolog.Logger, dbHandler *data.DBHandler) models.HandlerInterface {
	return &Handler{log: log, dbHandler: dbHandler}
}
