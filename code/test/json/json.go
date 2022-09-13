package json

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/vmihailenco/msgpack"
)

// 这里使用msgpack实现消息的序列化。messagepack是一个高效的二进制序列化协议。
// 相比json编码后的数据的体积更小，编解码的速度更快。redis script也支持messagepack。
func MsgpackMarshal() {
	type Item struct {
		Foo string
	}

	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
	// Output: bar
}

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
