package gjson

import (
	_ "embed"
	"fmt"
	"github.com/tidwall/gjson"
)

var (
  //go:embed vrpm.json
  JsonData string
  LocalData string = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
)

/*
gjson 是 tidwall 下的经典开源库，和 sjson 常常一起出现。前者代表【get json】，后者代表【set json】，分别侧重于读写 json 结构。
其实从 json 结构中拿到某个属性值，并不是一件很难的事。定义一个结构体，json.Unmarshal 就能搞定。

使用 gjson 原因：
  简单：反序列化的解法下，毕竟你还得定义一个结构体，而很多时候，读取的这个属性近在眼前，我们只是需要一个轻量级的解决方案；
  快速：对整个 json 进行反序列化是很耗性能的，明明我们只需要某个固定层级关系的属性，何必把整个 json 都解析出来呢？我们需要一个高性能的方案，以远低于 json.Unmarshal 的成本来读取某个值。
*/
func main() {
  // 读取 value
  value := gjson.Get(JsonData, "servicePort")
  fmt.Println(value)
  
  
  // 校验 value 是否存在
  value := gjson.Get(LocalData, "name.last")
  if !value.Exists() {
    println("no last name")
  } else {
    println(value.String())
  }
  
  // 迭代器
  result := gjson.Get(json, "programmers")
  result.ForEach(func(key, value gjson.Result) bool {
    println(value.String()) 
    return true // keep iterating
  })
}
