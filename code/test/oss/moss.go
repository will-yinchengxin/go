package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"path/filepath"
	"strconv"
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
	BucketName		string
	Endpoint		string
	DwnLoadPartNum  int64
}

func (o *MOSS) InitConfig() *MOSS {
	return &MOSS{
		AccessKeyId: AppConfig.Oss.AccessKeyId,
		AccessKeySecret: AppConfig.Oss.AccessKeySecret,
		Host: AppConfig.Oss.Host,
		UploadDir: AppConfig.Oss.UploadDir,
		ExpireTime: AppConfig.Oss.ExpireTime,
		BucketName: AppConfig.Oss.BucketName,
		Endpoint: AppConfig.Oss.Endpoint,
		DwnLoadPartNum: AppConfig.Oss.DwnLoadPartNum,
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

// 断点续传下载
// FileName 文件下载地址 例如/a/b/c.txt
// DownloadPath 文件下载存储路径 例如 /root/c.txt
func (o *MOSS) DownLoad(param AliDownLoad) error {
	client, err := oss.New(o.Endpoint, o.AccessKeyId, o.AccessKeySecret)
	if err != nil {
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return err
	}

	meta, err := bucket.GetObjectDetailedMeta(param.FileName)
	if err != nil {
		return err
	}

	// 计算分片大小
	objectSize, err := strconv.ParseInt(meta.Get("Content-Length"), 10, 64)
	partSize := int64(100 * 1024)
	if objectSize > o.DwnLoadPartNum {
		partSize = objectSize / o.DwnLoadPartNum
	}

	routines := o.DwnLoadPartNum / 4
	if routines == 0 {
		routines = 1
	}

	// 判断文件夹是否存在，不存在先创建文件夹
	dir := filepath.Dir(param.DownloadPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
		// os.MkdirAll(dir, os.ModePerm) // 创建多级目录
		os.Chmod(dir, os.ModePerm)
	}

	err = bucket.DownloadFile(param.FileName, param.DownloadPath, partSize, oss.Routines(int(routines)), oss.Checkpoint(true, ""))
	if err != nil {
		return err
	}
	return nil
}

// 删除文件名称 例如 /root/c.txt
func (o *MOSS) DeleteByFileName(fileName string) error {
	client, err := oss.New(o.Endpoint, o.AccessKeyId, o.AccessKeySecret)
	if err != nil {
		return err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return err
	}

	// 返回删除成功的文件。
	err = bucket.DeleteObject(fileName)
	if err != nil {
		return err
	}
	return nil

}

func GetGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
