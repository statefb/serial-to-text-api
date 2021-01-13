package converter

import (
	"app/server/gen/models"
	"encoding/json"
)

type outputData struct {
	key  *models.KeyData         `json: "key"`
	data []*models.CollectedData `json: "data"`
}

type JsonConverter struct {
	key  *models.KeyData
	data []*models.CollectedData
}

func NewJsonConverter(key *models.KeyData, data []*models.CollectedData) *JsonConverter {
	return &JsonConverter{key, data}
}

func (jc *JsonConverter) Convert() ([]byte, error) {
	o := &outputData{
		key:  jc.key,
		data: jc.data,
	}
	content, err := json.Marshal(o)
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
