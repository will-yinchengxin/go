package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"os"
	"path/filepath"
)

func GetPolicyToken() () {
	// Signature = URL-Encode( Base64( HMAC-SHA1( YourSecretAccessKeyID, UTF-8-Encoding-Of( StringToSign ) ) ) )
	/*
			AAE5DG77IBL4U7SXAH9F
			uLZIA0rQS5qgtdSQ8joFx4G3d4Yl9AAFDZQiY9Hd
			obs.cn-east-3.myhuaweicloud.com
			1646822097

			GET


			3293557925
			/test

			obs.cn-east-3.myhuaweicloud.com
			?AccessKeyId=AAE5DG77IBL4U7SXAH9F
			&Expires=3293557925
			&Signature=J13uBGIcIPWDRcF46A4fYZocpSM%3D

		https://obs-community.obs.cn-north-1.myhuaweicloud.com/sign/query_signature.html
	*/

	param := "GET" + "\n" +
		"" + "\n" +
		"" + "\n" +
		"3293557925" + "\n" +
		"" + "/test"
	fmt.Println(param)
	mac := hmac.New(sha1.New, []byte("uLZIA0rQS5qgtdSQ8joFx4G3d4Yl9AAFDZQiY9Hd"))
	_, err := mac.Write([]byte(param))
	if err != nil {
		fmt.Println(err)
		return
	}
	str2 := obs.UrlEncode(base64.StdEncoding.EncodeToString(mac.Sum(nil)), false)
	fmt.Println(base64.StdEncoding.EncodeToString(mac.Sum(nil)), str2)
	return
}
// ------------------------- 华为 obs 存储 ---------------------------
type Obs struct {
	BucketName      string
	ObjectKey       string
	AK              string
	SK              string
	EndPoint        string
	ExpireTime      int32
	Host            string
	Dir             string
	DownLoadPartNum int64
	DomainName      string
	UserName        string
	UserPassword    string
}

func (o *Obs) Init() *Obs {
	o.BucketName = AppConfig.Obs.BucketName
	o.ObjectKey = AppConfig.Obs.ObjectKey
	o.AK = AppConfig.Obs.AK
	o.SK = AppConfig.Obs.SK
	o.EndPoint = AppConfig.Obs.EndPoint
	o.ExpireTime = AppConfig.Obs.ExpireTime
	o.Host = AppConfig.Obs.Host
	o.Dir = AppConfig.Obs.Dir
	o.DownLoadPartNum = AppConfig.Obs.DownLoadPartNum
	o.DomainName = AppConfig.Obs.DomainName
	o.UserName = AppConfig.Obs.UserName
	o.UserPassword = AppConfig.Obs.UserPassword
	return o
}
/*
test:
	var err error
	defer func() {
		utils.ErrorLog(err)
	}()
	//param := AliDownLoad {
	//	FileName: "100072/ota/a..zip",
	//	DownloadPath: "E:\\pic\\",
	//}
	//err = o.Moss.
	//	Init().
	//	DownLoad(param)

	err = o.Moss.
		Init().
		DeleteByFileName("100072/ota/a..zip")

	if err != nil {
		codeType = utils.GlobalError
		return
	}
	return
*/
func (o *Obs) DownLoad(param AliDownLoad) error {
	input := &obs.DownloadFileInput{}
	input.Bucket = o.BucketName
	input.Key = param.FileName

	// 进行分段计算
	input.EnableCheckpoint = true      // 开启断点续传模式
	input.PartSize = 100 * 1024 * 1024 // 指定分段大小为100MB
	input.TaskNum = 5                  // 指定分段下载时的最大并发数

	dir := filepath.Dir(param.DownloadPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
		os.Chmod(dir, os.ModePerm)
	}
	input.DownloadFile = param.DownloadPath

	obsClient, err := obs.New(o.AK, o.SK, o.EndPoint)
	if err != nil {
		return err
	}
	_, err = obsClient.DownloadFile(input)
	if err != nil {
		return err
	}
	return nil
}

// 删除文件名称
func (o *Obs) DeleteByFileName(fileName string) error {
	input := &obs.DeleteObjectInput{}
	input.Bucket = o.BucketName
	input.Key = fileName

	obsClient, err := obs.New(o.AK, o.SK, o.EndPoint)
	if err != nil {
		return err
	}
	_, err = obsClient.DeleteObject(input)
	if err != nil {
		return err
	}
	return nil
}
