package upload

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	Uuid "github.com/satori/go.uuid"
	"net/http"
)

//gin.SetMode("release")
//route := gin.Default()
//// 注册中间件
//route.POST("/upload", upload.DealWithUpload)
//fmt.Println("start server 8080")
//route.Run(":8080")


// 127.0.0.1:8080/upload
// post form-data file 格式类型文件
func DealWithUpload(c *gin.Context) {
	//验证文件
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    9000,
			"message": "缺少文件",
			"data":    "",
		})
		return
	}
	//先判断文件大小
	if file.Size > 5*1024*1024 { //5m大小
		c.JSON(http.StatusOK, gin.H{
			"code":    9001,
			"message": "超出大小",
			"data":    "",
		})
		return
	}

	//判断文件类型
	rFile, _ := file.Open()

	// defer uFile.Close()
	// 即刻关闭
	rFile.Close()

	allowed := []string{"image/png", "image/jpeg", "image/jpg"}
	mime, err := mimetype.DetectReader(rFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    9002,
			"message": "读取信息错误",
			"data":    "",
		})
		return
	}

	if !mimetype.EqualsAny(mime.String(), allowed...) {
		c.JSON(http.StatusOK, gin.H{
			"code":    9003,
			"message": "上传类型不属于所属范围",
			"data":    "",
		})
		return
	}

	key := "dist/profile/picture/" + "omsPreview"
	fileName := Uuid.NewV1().String() + mime.Extension() //
	data := make(map[string]interface{})
	data["key"] = key
	data["fileName"] = fileName
	c.JSON(http.StatusOK, gin.H{
		"code":    9004,
		"message": "that is ok",
		"data": data,
	})
}

/*
package v1

import (
	"account-uniform/app/do/request"
	"account-uniform/app/do/response"
	"account-uniform/app/utils"
	"account-uniform/app/utils/http"
	"account-uniform/consts"
	"account-uniform/core"
	logs "account-uniform/nvwa/yclogs"
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"path"
	"strconv"
)

type Upload struct {
}

func (u *Upload) Avatar(req request.UploadFile, res *response.UploadFile, ctx context.Context) (codeType *utils.CodeType) {

	if req.UploadFile == nil {
		return utils.ImageZeroError
	}
	fileSize := req.UploadFile.Size
	fmt.Println(path.Ext(req.UploadFile.Filename)) // .jpg
	fmt.Println(path.Base(req.UploadFile.Filename)) // Dingtalk_20210820143551.jpg
	match, _ :=path.Match(".jpg", path.Ext(req.UploadFile.Filename))
	fmt.Println(match) // true
	fileExt := path.Ext(req.UploadFile.Filename) //获取文件后缀
	if ok := u.isAllowed(req.UploadFile, []string{"image/png", "image/jpeg"}); !ok {
		return utils.ImageFormatError
	}
	if ok := u.isSizeOk(int(fileSize), 4*1024*1024); !ok {
		return utils.ImageSizeError

	}
	filePath := "avatar"
	fileName := req.UnionId + fileExt
	err, url := u.uploadCommon(req.UploadFile, filePath, fileName)
	if err != nil {
		return utils.ImageUploadError
	}
	res.Url = url
	return
}

func (u Upload) isAllowed(file *multipart.FileHeader, allowed []string) bool {
	rFile, _ := file.Open()
	defer rFile.Close()                       //
	mime, err := mimetype.DetectReader(rFile) //获取文件mimetype
	fmt.Println(mime.String())  // image/jpeg
	fmt.Println(mime.Extension()) // .jpg
	fmt.Println(mime.Parent()) // application/octet-stream
	if err != nil {
		return false
	}
	if !mimetype.EqualsAny(mime.String(), allowed...) {
		return false
	}
	return true
}

func (u Upload) isSizeOk(size int, limit int) bool {
	if size > limit {
		return false
	}
	return true
}

func (u *Upload) uploadCommon(file *multipart.FileHeader, path, fileName string) (err error, urlPath string) {
	url := core.AppConfig.OBSConfig.Url + core.AppConfig.OBSConfig.CommonUpload

	headers := make(map[string]string)
	headers["app_id"] = strconv.Itoa(consts.APP_ID)
	headers["bucket_id"] = core.AppConfig.OBSConfig.BucketID
	headers["sk"] = core.AppConfig.OBSConfig.SK
	urlPath, err, _ = http.PostFile(url, file, headers, path, fileName)
	if err != nil {
		_ = core.Log.Error(logs.TraceFormatter{
			Trace: logrus.Fields{
				"err":  err.Error(),
				"func": "UploadCommon",
			},
		})
	}
	return err, urlPath
}

*/