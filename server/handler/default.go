package handler

import (
	"local.packages/gen/models"
	"local.packages/gen/restapi/serialtocsv/common"

	"github.com/go-openapi/swag"
)

func CreateDefaultError(msg string) *common.GetDataDefault {
	// NOTE: all default behavior is common
	payload := &models.ErrorV1{
		Code:    swag.String("400"),
		Message: swag.String(msg),
	}
	return common.NewGetDataDefault(400).WithPayload(payload)
}
