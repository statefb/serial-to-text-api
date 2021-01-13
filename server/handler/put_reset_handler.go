package handler

import (
	"app/server/data"
	"app/server/gen/restapi/serialtocsv/common"
	"app/server/states"

	"github.com/go-openapi/runtime/middleware"
)

type PutResetHandler struct{}

func (h *PutResetHandler) Handle(params common.PutResetParams) middleware.Responder {
	data.ResetKeyData()
	states.SetState(states.Waiting)
	return common.NewPutResetOK()
}
