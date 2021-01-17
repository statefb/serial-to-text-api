package converter

import (
	"app/server/gen/models"
	"encoding/json"
)

type outputData struct {
	Key  *models.KeyData         `json:"key"`
	Data *[]models.CollectedData `json:"data"`
}

// type outputData struct {
// 	Value int `json: "value"`
// }

type JsonConverter struct {
	key  *models.KeyData
	data *[]models.CollectedData
}

func NewJsonConverter(key *models.KeyData, data *[]models.CollectedData) *JsonConverter {
	return &JsonConverter{key, data}
}

func (jc *JsonConverter) Convert() ([]byte, error) {
	o := outputData{
		Key:  jc.key,
		Data: jc.data,
	}
	// o := outputData{Value: 4}
	content, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return nil, err
	}
	return content, nil
}

// func transform(key *models.KeyData, data *models.CollectedData) *JsonData {
// 	val := strconv.FormatFloat(*jc.collectedData.Value, 'f', 4, 64)
// 	return &JsonData{
// 		lotNo:     *kd.LotID,
// 		bagNo:     *kd.BagID,
// 		timestamp: cd.Timestamp.String(),
// 		value:     val,
// 	}
// }
