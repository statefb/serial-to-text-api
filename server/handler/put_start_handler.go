package handler

import (
	"app/server/data"
	"app/server/gen/restapi/serialtocsv/common"
	"app/server/states"

	"github.com/go-openapi/runtime/middleware"
)

type PutStartHandler struct{}

func (h *PutStartHandler) Handle(params common.PutStartParams) middleware.Responder {
	if states.GetState() != states.Waiting {
		return CreateDefaultError("state must be waiting.")
	}
	// memorize lot info
	data.SetKeyData(*params.Body)
	// open serial
	err := data.Open()
	if err != nil {
		panic(err)
	}
	// start collecting data
	data.Start()

	states.SetState(states.Collecting)
	return common.NewPutStartOK()
}
