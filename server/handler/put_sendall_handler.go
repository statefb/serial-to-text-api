package handler

import (
	"local.packages/gen/restapi/serialtocsv/common"
	"local.packages/sender"
	"local.packages/states"

	"github.com/go-openapi/runtime/middleware"
)

type PutSendallHandler struct{}

func (h *PutSendallHandler) Handle(params common.PutSendallParams) middleware.Responder {
	if states.GetState() != states.Waiting {
		return CreateDefaultError("state must be waiting.")
	}
	sender, err := sender.GetSender()
	if err != nil {
		panic(err)
	}

	err = sender.SendAll()
	if err != nil {
		panic(err)
	}

	return common.NewPutSendallOK()
}
