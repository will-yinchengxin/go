package ip

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"github.com/ppmoon/gbt2260"
)

// 根据ip判断素数地区信息
func IP()  {
	// 这里默认会切换只根目录
	region, err := ip2region.New("ip/ip2region.db")
	defer region.Close()
	if err != nil {
		return
	}
	ip, err := region.MemorySearch("114.92.87.234")
	if err != nil {
		return
	}
	fmt.Println(ip.City)
}

func GetCode(city string) string {
	NewGetGbt2260Table := gbt2260.GetGbt2260Table()
	var code string
	Loop:
	for _, strings := range NewGetGbt2260Table {
		if strings[1] == city {
			code = strings[0]
			break Loop
		}
	}
	return code
}

func GetRegion(code string) []string {
	region := gbt2260.NewGBT2260()
	localCode := region.SearchGBT2260(code)
	return localCode
}
