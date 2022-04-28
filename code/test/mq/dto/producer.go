package dto

import (
	"strconv"
)

type Data struct {
	ProjectId string `json:"projectId" form:"projectId"`
	MapId     int    `json:"mapId" form:"mapId"`
	MakeUrl   string `json:"makeUrl" form:"makeUrl"`
}

func GetData() []Data {
	ProducerData := []Data{}
	for i := 0; i < 500000; i++ {
		data := Data{
			ProjectId: strconv.Itoa(i),
			MapId:     i,
			MakeUrl:   "test" + strconv.Itoa(i),
		}
		ProducerData = append(ProducerData, data)
	}
	return ProducerData
}
