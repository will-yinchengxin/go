package uesr

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"hash/crc32"
	"strconv"
)

func GetUid() string{
	//生成uuid
	uV4 := uuid.NewV4()
	uuidHash := strconv.Itoa(int(crc32.ChecksumIEEE([]byte(uV4.String())))) // 2047730008
	fmt.Println(uuidHash)
	return uuidHash
}

func EncodePassword(in string) string {
	//密码加密
	h1 := md5.New()
	h1.Write([]byte(in))
	f1 := hex.EncodeToString(h1.Sum(nil))
	h2 := md5.New()
	h2.Write([]byte(f1))
	//TODO 建议将盐值存在用户表中，创建的时候随机生成盐值
	h2.Write([]byte("willTest")) // consts.Salt
	f2 := hex.EncodeToString(h2.Sum(nil))
	return f2
}
