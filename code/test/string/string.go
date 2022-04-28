package string

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int64  `json:"status"`
}

type Employee struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Title      string `gorm:"column:title" json:"title" form:"title"`
	Content    string `gorm:"column:content" json:"content" form:"content"`
	Status     int64  `gorm:"column:status" json:"status" form:"status"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeleteTime int64  `gorm:"column:delete_time" json:"delete_time" form:"delete_time"`
}

/*
0.2.1.101  => 02110100 去掉点号，向后补0凑齐8位，=> 去掉前面的0，最终结果是 2110100
1.1.1.001=> 11100100  去掉点号，向后补0凑齐8位 =>最终结果是 11100100

对比版本大小：11100100 > 2110100
*/
func getChar(str string) string {
	str = strings.Replace(str, ".", "", -1)
	if len(str) > 8 {
		return ""
	}
	// 进行补零
	if len(str) < 8 {
		zeroLen := 8 - len(str)
		var tmpStr bytes.Buffer
		tmpStr.WriteString(str)
		for i := 0; i < zeroLen; i++ {
			tmpStr.WriteString("0")
		}
		str = tmpStr.String()
	}
	for _, ch := range str {
		if string(ch) == "0" {
			str = strings.TrimLeft(str,"0")
		}
	}
	return str
}

func stringSprintf() {
	k := time.Now()
	var s string
	for i := 0; i < 10000; i++ {
		s = fmt.Sprintf("%s%s", s, strconv.Itoa(i))
	}
	fmt.Println("stringSprintf", time.Now().Sub(k))
}

func stringBuilder() {
	//var str bytes.Buffer
	//str.WriteString("will ")
	//str.WriteString("test ")
	//str.WriteString("home")
	//fmt.Println(str.String())



	//stringSprintf()
	//byteBuffer()
	//stringsBuilder()
	//dsn := "root:123456@tcp(127.0.0.1:13306)/will?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	return
	//}

	//user := User{Title: "test", Content: "this is content", Status: 1}
	//employee := Employee{}
	//copier.Copy(&employee, &user)
	//fmt.Printf("%#v\n", employee)
	//db.Table("test").Create(&employee)

	// 进行更新
	//userAno := User{
	//	Title: "this1",
	//}
	//employeeAno := Employee{}
	//copier.Copy(&employeeAno, &userAno)
	//// 加入select 后 填入什么字段更新什么字段
	//// 当不使用select的时候，默认更新含有内容的字段
	//fmt.Println(employeeAno)
	//db.Table("test").Select("title", "content", "status").Where("id = ?", 56).Updates(employeeAno)
	//fmt.Println(employeeAno.Content)
}

func byteBuffer() {
	k := time.Now()
	var s bytes.Buffer
	for i := 0; i < 10000; i++ {
		s.WriteString(strconv.Itoa(i))
	}
	fmt.Println("byteBuffer", time.Now().Sub(k))
}

func stringsBuilder() {
	k := time.Now()
	var s strings.Builder
	for i := 0; i < 10000; i++ {
		s.WriteString(strconv.Itoa(i))
	}
	fmt.Println("stringsBuilder", time.Now().Sub(k))
}

func replaceSpace(s string) string {
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] == ' ' {
			n := len(ss)
			ss = append(ss, ss[n-2:]...)
			fmt.Println(string(ss))
			for k := n - 1; k > i+2; k-- {
				ss[k] = ss[k-2]
			}
			ss[i] = '%'
			ss[i+1] = '2'
			ss[i+2] = '0'
		}
	}
	return string(ss)
}