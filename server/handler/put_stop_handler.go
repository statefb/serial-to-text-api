package handler

import (
	"app/server/gen/restapi/serialtocsv/common"
	"app/server/states"
	"github.com/go-openapi/runtime/middleware"
)

type PutStopHandler struct{}

func (h *PutStopHandler) Handle(params common.PutStopParams) middleware.Responder {
	if states.GetState() != states.Collecting {
		return CreateDefaultError("state must be collecting.")
	}

	// create and send via specified protocol
	states.SetState(states.Sending)

	// TODO

	// back to waiting
	states.SetState(states.Waiting)
	return common.NewPutStopOK()
}
