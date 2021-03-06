package handler

import (
	"local.packages/data"
	"local.packages/gen/restapi/serialtocsv/common"
	"local.packages/states"

	"github.com/go-openapi/runtime/middleware"
)

type PutResetHandler struct{}

func (h *PutResetHandler) Handle(params common.PutResetParams) middleware.Responder {
	data.ResetKeyData()
	err := data.Reset()
	if err != nil {
		panic(err)
	}
	states.SetState(states.Waiting)
	return common.NewPutResetOK()
}
