package handler

import (
	"app/server/data"
	"app/server/gen/models"
	"app/server/gen/restapi/serialtocsv/common"
	"app/server/states"

	"github.com/go-openapi/runtime/middleware"
)

type GetDataHandler struct{}

func (h *GetDataHandler) Handle(params common.GetDataParams) middleware.Responder {
	ss := []states.State{
		states.Waiting,
		states.Collecting,
	}
	found := false
	for _, s := range ss {
		if states.GetState() == s {
			found = true
		}
	}
	if !found {
		return CreateDefaultError("state must be waiting.")
	}
	payload := []*models.CollectedData{}
	for _, x := range data.GetData() {
		payload = append(payload, &x)
	}
	return common.NewGetDataOK().WithPayload(payload)
}
