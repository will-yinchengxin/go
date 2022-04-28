package inArray

import "reflect"

// import ref "reflect"

// 判断 数组、切片、map中是否存在某个数值
func InArray(obj interface{}, target interface{}) bool {

	/*
	var data = map[string]interface{}{"name": "will", "age": 18}

	dataVal := ref.ValueOf(data)
	fmt.Println(dataVal) // map[age:18 name:will]
	MapIndex := dataVal.MapIndex(ref.ValueOf("name")).IsValid()
	fmt.Println(MapIndex) // true
	*/

	targetValue := reflect.ValueOf(obj)
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == target {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(target)).IsValid() {
			return true
		}
	}
	return false
}
