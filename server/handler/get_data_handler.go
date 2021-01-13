package handler

import (
	"github.com/go-openapi/runtime/middleware"
	// "github.com/go-openapi/swag"
	// "github.com/go-openapi/strfmt"
	// "github.com/go-openapi/strfmt/conv"
	"app/server/gen/restapi/serialtocsv/common"
	// "app/server/gen/models"
	"app/server/data"
	"app/server/states"
	// "time"
)

type GetDataHandler struct{}

func (h *GetDataHandler) Handle(params common.GetDataParams) middleware.Responder {
	if states.GetState() != states.Waiting {
		return CreateDefaultError("state must be waiting.")
	}
	payload := data.GetData()
	return common.NewGetDataOK().WithPayload(payload)
}
