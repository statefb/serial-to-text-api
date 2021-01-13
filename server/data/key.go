package data

import "app/server/gen/models"

var keyData *models.KeyData

func SetKeyData(kd *models.KeyData) {
	keyData = kd
}

func GetKeyData() *models.KeyData {
	return keyData
}

func ResetKeyData() {
	keyData = nil
}
