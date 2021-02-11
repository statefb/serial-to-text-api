package sender

import (
	"math/rand"
	"time"

	"local.packages/gen/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randName(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func convertCollectData(cd []*models.CollectedData) []models.CollectedData {
	cdList := []models.CollectedData{}
	for _, d := range cd {
		cdList = append(cdList, *d)
	}
	return cdList
}
