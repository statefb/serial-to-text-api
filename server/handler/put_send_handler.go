package handler

import (
	"local.packages/gen/restapi/serialtocsv/common"
	"local.packages/sender"
	"local.packages/states"

	"github.com/go-openapi/runtime/middleware"
)

type PutSendHandler struct{}

func (h *PutSendHandler) Handle(params common.PutSendParams) middleware.Responder {
	if states.GetState() != states.Waiting {
		return CreateDefaultError("state must be waiting.")
	}
	sender, err := sender.GetSender()
	if err != nil {
		panic(err)
	}

	// TODO: data to send must be request body
	err = sender.Send(params.Body)
	if err != nil {
		panic(err)
	}

	return common.NewPutSendOK()
}
