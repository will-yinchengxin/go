package json

import (
	"encoding/json"
	"fmt"
	"time"
)

func JsonUnmarshal() {
	type FruitBasket struct {
		Name    string    `json:"name"`
		Fruit   []string  `json:"fruit"`
		Id      int64     `json:"id"`
		Created time.Time `json:"created"`
	}

	jsonData := []byte(`
    {
        "name": "Standard",
        "fruit": [
             "Apple",
            "Banana",
            "Orange"
        ],
        "id": 999,
        "created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id)
	fmt.Println(basket.Created)
}
