package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// 获取md5字符串
func GetMD5(str string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(hex.EncodeToString(cipherStr))
}

// 登录密码加密
func EncryptPassword(in string) {
	h1 := md5.New()
	h1.Write([]byte(in))
	f1 := hex.EncodeToString(h1.Sum(nil))
	h2 := md5.New()
	h2.Write([]byte(f1))
	//TODO 建议将盐值存在用户表中，创建的时候随机生成盐值
	h2.Write([]byte("a6b85a82044f39d2ec12db39834be19868f654a0"))
	f2 := hex.EncodeToString(h2.Sum(nil))
	fmt.Println( f2)
}

// 隐藏手机号
func HideAccount(account string) string {
	if len(account) == 11 {
		account = account[:3] + "******" + account[9:]
	}
	return account
}

