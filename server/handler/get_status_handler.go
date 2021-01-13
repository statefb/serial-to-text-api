package handler

import (
	"app/server/gen/restapi/serialtocsv/common"
	"app/server/states"
	"github.com/go-openapi/runtime/middleware"
)

type GetStatusHandler struct{}

func (h *GetStatusHandler) Handle(params common.GetStatusParams) middleware.Responder {
	return common.NewGetStatusOK().WithPayload(states.GetState().String())
}
