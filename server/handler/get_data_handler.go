package handler

import (
	"fmt"

	"local.packages/data"
	"local.packages/gen/models"
	"local.packages/gen/restapi/serialtocsv/common"
	"local.packages/states"

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
	// payload := []*models.CollectedData{}
	// payload := make([]*models.CollectedData, config.NewConf().Serial.MaxRecordSize)
	d := data.GetData()
	payload := make([]*models.CollectedData, len(d))
	for i := 0; i < len(d); i++ {
		payload[i] = &d[i]
	}
	// for i, x := range data.GetData() {
	// 	fmt.Printf("x: %+v\n", x)
	// 	fmt.Printf("&x: %+v\n", &x)
	// 	// payload = append(payload, &x)
	// 	payload[i] = &x
	// }
	fmt.Printf("payload: %+v\n", payload)
	return common.NewGetDataOK().WithPayload(payload)
}
