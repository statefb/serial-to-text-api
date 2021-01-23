package handler

import (
	"local.packages/data"
	"local.packages/gen/restapi/serialtocsv/common"
	"local.packages/states"

	"github.com/go-openapi/runtime/middleware"
)

type PutStartHandler struct{}

func (h *PutStartHandler) Handle(params common.PutStartParams) middleware.Responder {
	if states.GetState() != states.Waiting {
		return CreateDefaultError("state must be waiting.")
	}
	// memorize lot info
	data.SetKeyData(*params.Body)
	// start collecting data
	data.Start()

	states.SetState(states.Collecting)
	return common.NewPutStartOK()
}
