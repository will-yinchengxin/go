package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func TestFile() {
	//file, err := os.Open("/d/Project/test/Dingtalk_20210820143551.jpg")
	//defer file.Close()
	//if err != nil {
	//	fmt.Println("open file failed!, err:", err)
	//	return
	//}

	content, err := ioutil.ReadFile("D:\\Project\\test\\file\\Dingtalk_20210820143551.jpg")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func UploadAndSave(c *gin.Context)  {
	// 文件进行本地存储，方便后续点播云上传
	file, err := c.FormFile("file")
	if err != nil {
		//codeType := utils.GetFormFileError
		//utils.Error(c, codeType)
		return
	}
	if path.Ext(file.Filename) == ".mp4" {
		dir := filepath.Dir("E:\\upload\\") // Windows: "E:\\upload\\ Linux: /studio/
		// 创建目录
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)

			os.Chmod(dir, os.ModePerm)
		}
		dst := path.Join(dir, file.Filename) // 文件地址
		err = c.SaveUploadedFile(file, dst)
		/*
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			out, err := os.Create(dst)
			if err != nil {
				return err
			}
			defer out.Close()

			_, err = io.Copy(out, src)
			return err
		*/
		if err != nil {
			//codeType := utils.SaveFormFileError
			//utils.Error(c, codeType)
			return
		}
	}
}