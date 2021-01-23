package handler

import (
	"local.packages/gen/restapi/serialtocsv/common"
	"local.packages/states"

	"github.com/go-openapi/runtime/middleware"
)

type GetStatusHandler struct{}

func (h *GetStatusHandler) Handle(params common.GetStatusParams) middleware.Responder {
	return common.NewGetStatusOK().WithPayload(states.GetState().String())
}
