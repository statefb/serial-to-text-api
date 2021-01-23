package data

import (
	"log"

	"local.packages/gen/models"
)

var keyData models.KeyData

func SetKeyData(kd models.KeyData) {
	keyData = kd
}

func GetKeyData() models.KeyData {
	return keyData
}

func ResetKeyData() {
	log.Printf("reset key data.")
	keyData = models.KeyData{}
}
