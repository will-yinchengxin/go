package reflect

import (
	"fmt"
	"reflect"
	"strings"
)

/*
反射三大定律:
	- 反射可以将接口类型变量 转换为“反射类型对象”；
	- 反射可以将 “反射类型对象”转换为 接口类型变量；
	- 如果要修改 “反射类型对象” 其类型必须是 可写的；

reflect.TypeOf()  获取接口的类型
reflect.ValueOf() 获得接口值的值
两种方法均返回对象,也就是我们提及的反射类型对象

第一定律举例:
	var age interface{} = 25
	t := reflect.TypeOf(age)
	v := reflect.ValueOf(age)
	fmt.Println(t, v) // int 25
	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t)  // *reflect.rtype
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v) // reflect.Value

第二定律举例:
	reflect.Value 结构体会接受interface方法,返回一个interface类型变量
	只有 Value 才能逆向转换，而 Type 则不行
	-----------------------------------------------------------
	var age interface{} = 25
	v := reflect.ValueOf(age)

	i := v.Interface()
	fmt.Printf("新对象的类型 %T, 值为 %v\n", i, i) // 新对象的类型 int, 值为 25
	ii := v.Interface().(int)
	fmt.Println(ii) // 25

	type  ====>  reflect.TypeOf()  =====>  reflect.Type object

		 ||   ====>   reflect.ValueOf()   ====> ||
		data 								    reflect.Value object
 		 ||   <====   Interface()         <===  ||

第三定律举例:
	因此在反射的规则里
		不是接收变量指针创建的反射对象，是不具备『可写性』的
		是否具备『可写性』，可使用 CanSet() 来获取得知
		对不具备『可写性』的对象进行修改，是没有意义的，也认为是不合法的，因此会报错。

	var name string = "test"
    v := reflect.ValueOf(name)
    fmt.Println("可写性为:", v.CanSet()) // false

	要让反射对象具备可写性，需要注意两点
		创建反射对象时传入变量的指针
		使用 Elem()函数返回指针指向的数据

	var name string = "test"
    v1 := reflect.ValueOf(&name)
    fmt.Println("v1 可写性为:", v1.CanSet())  // v1 可写性为: false
    v2 := v1.Elem()
    fmt.Println("v2 可写性为:", v2.CanSet())  // v2 可写性为: true

	demo:
		var name string = "test reflect"
		fmt.Println("初始值为", name)

		v1 := reflect.ValueOf(&name)
		v2 := v1.Elem()

		v2.SetSting("give another value")
		fmt.Println("now the value is", name)


类型 reflect.Value 有一个方法 Type()，它会返回一个 reflect.Type 类型的对象。Type和 Value都有一个名为 Kind 的方法，
它会返回一个常量，表示底层数据的类型，常见值有：Uint、Float64、Slice等
*/

type rawVal struct {
	DecimalSeparator  string `jpath:"userContext.conversationCredentials.sessionToken"`
	GroupingSeparator string `jpath:"userContext.valid"`
	GroupPattern      string `jpath:"userContext.cobrandId"json:"group_pattern"`
}

var TestMap = map[string]interface{}{
	"userContext": map[string]interface{}{
		"userContext": map[string]interface{}{"sessionToken": "06142010_1:b8d011fefbab8bf1753391b074ffedf9578612d676ed2b7f073b5785b"},
		"valid": "true",
		"cobrandId": "1.0000004e+07"}}

func GetStruct() {
	var rawVal rawVal
	//rawVal := rawVal{}
	mapToStruct(TestMap, &rawVal)
}

func mapToStruct(m map[string]interface{}, rawVal interface{}) (bool, error) {
	decoded := false

	var val reflect.Value
	reflectRawValue := reflect.ValueOf(rawVal)
	kind := reflectRawValue.Kind()
	//fmt.Println(reflectRawValue, kind) // &{  } ptr

	switch kind {
	case reflect.Ptr:
		// 获取值类型
		val = reflectRawValue.Elem() // fmt.Println(val) // {}
		if val.Kind() != reflect.Struct {
			return decoded, fmt.Errorf("Incompatible Type : %v : Looking For Struct", kind)
		}
	case reflect.Struct:
		var ok bool
		val, ok = rawVal.(reflect.Value)
		if ok == false {
			return decoded, fmt.Errorf("Incompatible Type : %v : Looking For reflect.Value", kind)
		}
	default:
		return decoded, fmt.Errorf("Incompatible Type : %v", kind)
	}

	// 遍历结构体字段，以 GroupPattern 为例
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		fmt.Println(valueField)
		//---------------------------------------------------
		indexName := val.Type().Field(i).Name
		fmt.Println(indexName) // GroupPattern
		indexIndex := val.Type().Field(i).Index
		fmt.Println(indexIndex) // [2]
		typeField := val.Type().Field(i)
		fmt.Println(typeField) // {GroupPattern  string jpath:"userContext.cobrandId"json:"group_pattern" 32 [2] false}
		tag := typeField.Tag
		fmt.Println(tag) // jpath:"userContext.cobrandId"json:"group_pattern"
		tagValue := tag.Get("jpath")
		fmt.Println(tagValue) // userContext.cobrandId
		//-----------------------------------------------------

		keys := strings.Split(tagValue, ".")
		data := findData(m, keys)
		if data != nil {
			// 更改反转状态
			decoded = true
			err := Decode("", data, valueField)
			if err != nil {
				return false, err
			}
		}
	}
	return decoded, nil
}

// 多维map查找指定键值
func findData(m map[string]interface{}, keys []string) interface{} {
	if len(keys) == 1 {
		if value, ok := m[keys[0]]; ok == true {
			return value
		}
		return nil
	}

	if value, ok := m[keys[0]]; ok == true {
		if m, ok := value.(map[string]interface{}); ok == true {
			return findData(m, keys[1:])
		}
	}

	return nil
}

func Decode(name string, data interface{}, val reflect.Value) error {
	if data == nil {
		return nil
	}
	dataVal := reflect.ValueOf(data)
	// IsValid 校验的是flag
	if !dataVal.IsValid() {
		// 如果value是为经过校验的，那么返回对应类型的 0 值
		val.Set(reflect.Zero(val.Type()))
		return nil
	}

	return nil
}

func ArrayColumn(array interface{}, key string) (list []interface{}) {
	valueOf := reflect.ValueOf(array)
	for i := 0; i < valueOf.Len(); i++ {
		one := valueOf.Index(i)
		one = one.Elem() // 指针不能操作
		value := one.FieldByName(key)
		kind := value.Kind()
		switch kind {
		case reflect.Int:
			list = append(list, value.Int())
		case reflect.String:
			list = append(list, value.String())
		default:
			list = append(list, value)
		}
	}
	return list
}
