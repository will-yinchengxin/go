package regexp

import (
	"fmt"
	"regexp"
	"strings"
)

func Regexp() {
	var url = "https://developer-fat.shadowcreator.com/file#房贷首付"
	r, _ := regexp.Compile("#(.*)")
	b := r.MatchString(url)
	if !b {
		fmt.Println("链接未能匹配出标题内容")
	}
	fmt.Println(strings.TrimLeft(r.FindString(url), "#"))
}
/*

func VersionCheck(f validator.FieldLevel) bool {
	val := f.Field().String()
	if ok, _ := regexp.MatchString(`^[0-9]{1,3}\.[0-9]{1,2}\.[0-9]{1,2}$`, val); ok {
		return true
	}
	return false
}

// CheckUrl url checker
func CheckUrl(f validator.FieldLevel) bool {
	val := f.Field().String()
	urlPartten := "^(http|https|ftp)\\://([a-zA-Z0-9\\.\\-]+(\\:[a-zA-Z0-9\\.&amp;%\\$\\-]+)*@)*((25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])|localhost|([a-zA-Z0-9\\-]+\\.)*[a-zA-Z0-9\\-]+\\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(\\:[0-9]+)*(/($|[a-zA-Z0-9\\.\\,\\?\\'\\\\\\+&amp;%\\$#\\=~_\\-]+))*$"
	if ok, _ := regexp.MatchString(urlPartten, val); ok {
		return true
	}
	return false
}
*/

func CheckEmail(email string) bool {
	if pass, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); pass {
		return true
	}
	return false
}