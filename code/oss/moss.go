package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io"
	"time"
)

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type MOSS struct {
	AccessKeyId     string
	AccessKeySecret string
	Host            string
	UploadDir       string
	ExpireTime      int64
}

func (o *MOSS) InitConfig() *MOSS {
	return &MOSS{
		AccessKeyId: AppConfig.Oss.AccessKeyId,
		AccessKeySecret: AppConfig.Oss.AccessKeySecret,
		Host: AppConfig.Oss.Host,
		UploadDir: AppConfig.Oss.UploadDir,
		ExpireTime: AppConfig.Oss.ExpireTime,
	}

	//var a = new(MOSS)
	//a.AccessKeyId = AppConfig.Oss.AccessKeyId
	//a.AccessKeySecret = AppConfig.Oss.AccessKeySecret
	//a.Host = AppConfig.Oss.Host
	//a.UploadDir = AppConfig.Oss.UploadDir
	//a.ExpireTime = AppConfig.Oss.ExpireTime
	//return a
}

func (o *MOSS) GetPolicyToken() (policyToken PolicyToken, err error) {
	now := time.Now().Unix()
	expire_end := now + o.ExpireTime
	var tokenExpire = GetGmtIso8601(expire_end)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, o.UploadDir)
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, _ := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(sha1.New, []byte(o.AccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	policyToken.AccessKeyId = o.AccessKeyId
	policyToken.Host = o.Host
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = o.UploadDir
	policyToken.Policy = string(debyte)

	return
}

func GetGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
