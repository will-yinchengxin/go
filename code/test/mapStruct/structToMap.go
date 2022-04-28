package mapStruct

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type MuchInfo struct {
	Class    int64    `json:"class"`
	UserInfo UserInfo `json:"userInfo"`
}

/*
这种方式, 对于直接输出而不操作的情况,没有任何问题,但是如果需要对输出的结果进行操作,就要注意,它 json.Marshal 会将其 int64 转换为 float64
*/
func NormalStructToMap() {
	u1 := UserInfo{Name: "will", Age: 3}

	b, _ := json.Marshal(&u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
	/*
		结果:
			key:name value:will value type:string
			key:age value:3 value type:float64
	*/
}

func TestReflectStructToMap() (map[string]interface{}, error) {
	u1 := &UserInfo{
		Name: "will",
		Age:  3,
	}
	tagName := "json"
	out := make(map[string]interface{})

	v := reflect.ValueOf(u1)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("only accept struct type")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i) // {Name  string json:"name" 0 [0] false}
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

func ReflectStructToMap() (map[string]interface{}, error) {
	u1 := &UserInfo{
		Name: "will",
		Age:  3,
	}
	tagName := "json"

	out := make(map[string]interface{})

	v := reflect.ValueOf(u1)
	if v.Kind() == reflect.Ptr {
		fmt.Println(v.Elem()) // 如果结构体取了指针地址, 需要使用 Elem() 获取元素对象, {will 3}
		v = v.Elem()
	}

	fmt.Println(v.Kind())
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("StructToMap only accepet struct pointer")
	}

	fmt.Println(v.Type()) // mapStruct.UserInfo
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(t.Field(i)) // {Name  string json:"name" 0 [0] false}
		fi := t.Field(i)
		fmt.Println(fi.Name) // Name
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	// 这里类型就是正确的 int64 了
	for k, v := range out {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	}
	return out, nil
}

// 多层 struct 嵌套转单层 map, 需要注意的是, 避免嵌套的结构体有重名的问题
func MultiStructToMap() (map[string]interface{}, error) {
	in := MuchInfo{Class: 3, UserInfo: UserInfo{Name: "will", Age: 3}}
	tag := "json"

	// 当前函数只接收struct类型
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr { // 结构体指针
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	out := make(map[string]interface{}, 8)
	queue := make([]interface{}, 0, 2)
	queue = append(queue, in)
	fmt.Println(len(queue), cap(queue), len(out)) // 1, 2, 0

	for len(queue) > 0 {
		v := reflect.ValueOf(queue[0])
		if v.Kind() == reflect.Ptr { // 结构体指针
			v = v.Elem()
		}
		queue = queue[1:]
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			vi := v.Field(i)
			if vi.Kind() == reflect.Ptr { // 内嵌指针
				vi = vi.Elem()
				if vi.Kind() == reflect.Struct { // 结构体
					queue = append(queue, vi.Interface())
				} else {
					ti := t.Field(i)
					if tagValue := ti.Tag.Get(tag); tagValue != "" {
						// 存入map
						out[tagValue] = vi.Interface()
					}
				}
				break
			}
			if vi.Kind() == reflect.Struct { // 内嵌结构体
				queue = append(queue, vi.Interface())
				break
			}
			// 一般字段
			ti := t.Field(i)
			if tagValue := ti.Tag.Get(tag); tagValue != "" {
				// 存入map
				out[tagValue] = vi.Interface()
			}
		}
	}
	return out, nil
}
