package zip

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)
//-------------------------- zip to xml ------------------------------
type Manifest struct {
	XMLName       xml.Name `xml:"manifest"`
	Text          string   `xml:",chardata"`
	Upgraded      string   `xml:"upgraded,attr"`
	UpgradeNumber string   `xml:"upgradeNumber,attr"`
	Version       string   `xml:"version,attr"`
	VersionNumber string   `xml:"versionNumber,attr"`
	Device        string   `xml:"device,attr"`
}

func UnZip() {
	fr, err := zip.OpenReader("E:\\ota\\package_info.zip")
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	for _, f := range fr.Reader.File {
		if f.Name == "package_info.xml" {
			rc, err := f.Open()
			defer rc.Close()
			if err != nil {
				log.Fatal(err)
			}
			data, err := ioutil.ReadAll(rc)
			Manifest := Manifest{}
			err = xml.Unmarshal(data, &Manifest)
			fmt.Println(Manifest.VersionNumber, Manifest.Version, Manifest.Device)
		}
	}
}
// -------------------------
func IsZip(src string) (bool, error) {
	f, err := os.Stat(src)
	if err != nil {
		return false, nil
	}

	if f.IsDir() {
		return false, nil
	}

	return ExtZip == filepath.Ext(f.Name()), nil

}

func IsDir(src string) (bool, error) {
	f, err := os.Stat(src)
	if err != nil {
		return false, err
	}

	return f.IsDir(), nil
}

func IsFileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

const (
	ExtZip = ".zip"
	// maxFileCount 压缩文件最大文件数
	maxFileCount = 1500
	// singleFileMaxSize 单个文件解压和总文件解压不能超过10G
	singleFileMaxSize uint64 = 10737418240
	fileMaxSize       int64  = 10737418240
)

/*
func main() {
	err := UnCompress("/tmp/ginProject-one.zip")
	if err != nil {
		fmt.Println(err)
	}
}
*/
func UnCompress(src, dst string) error {
	isZip, err := IsZip(src)
	if err != nil {
		return err
	}
	if !isZip {
		return errors.New("src is not a zip file")
	}

	isDir, err := IsDir(dst)
	if err != nil {
		return err
	}
	if !isDir {
		return errors.New("dst is not a directory")
	}

	srcReader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer srcReader.Close()

	if len(srcReader.File) > maxFileCount {
		return errors.New("zip file have too many files")
	}
	for _, f := range srcReader.File {
		// safety requirements
		if strings.Contains(f.Name, "./") ||
			strings.Contains(f.Name, ".\\") ||
			strings.Contains(f.Name, "..") {
		}
		if f.UncompressedSize64 > singleFileMaxSize {
			return errors.New("single file exceeds the maximum size")
		}
	}
	err = writeReader(srcReader, dst)
	if err != nil {
		return err
	}

	return nil
}

func writeReader(srcReader *zip.ReadCloser, dst string) error {
	var totalSize int64
	for _, f := range srcReader.File {
		if totalSize > fileMaxSize {
			return errors.New("the total size has exceeds the upper limit")
		}
		fileName := f.Name
		targetFilePath := filepath.Join(dst, fileName)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(targetFilePath, f.Mode())
			if err != nil {
				return err
			}
			continue
		}

		isExist, err := IsFileExist(targetFilePath)
		if err != nil {
			return err
		}
		if isExist {
			return errors.New("the target path has exist")
		}

		if err := os.MkdirAll(path.Dir(targetFilePath), 0700); err != nil {
			return err
		}

		writeSize, err := writeFile(f, targetFilePath)
		if err != nil {
			return err
		}
		totalSize += writeSize
	}

	return nil
}

func writeFile(f *zip.File, targetFilePath string) (int64, error) {
	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		return 0, err
	}
	defer targetFile.Close()
	file, err := f.Open()
	if err != nil {
		return 0, err
	}
	defer file.Close()

	writeSize, err := io.Copy(targetFile, file)
	if err != nil {
		return 0, err
	}

	return writeSize, nil
}

//func GetSrcZips(src string) ([]string, error) {
//	files, err := ioutil.ReadDir(src)
//	if err != nil {
//		return nil, err
//	}
//	fileNames := []string{}
//	for _, f := range files {
//		ext := filepath.Ext(f.Name())
//		if ext != ExtZip {
//			continue
//		}
//
//		fileNames = append(fileNames, strings.TrimSuffix(f.Name(), ext))
//	}
//
//	return fileNames, nil
//}

//func Compress(srcDir, dstZipPath string) error {
//	if filepath.Ext(filepath.Base(dstZipPath)) != ExtZip {
//		return errors.New("not a zip file")
//	}
//
//	dstDir := filepath.Dir(dstZipPath)
//	isDstDir, err := IsDir(dstDir)
//	if err != nil {
//		return err
//	}
//	if !isDstDir {
//		return errors.New("dstDir is not a directory")
//	}
//
//	isSrcDir, err := IsDir(srcDir)
//	if err != nil {
//		return err
//	}
//	if !isSrcDir {
//		return errors.New("srcDirs is not a directory")
//	}
//
//	f, err := os.Create(dstZipPath)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	zw := zip.NewWriter(f)
//	defer zw.Close()
//
//	files, err := ioutil.ReadDir(srcDir)
//	if err != nil {
//		return err
//	}
//
//	for _, fi := range files {
//		if err := compress(fi, srcDir, "", zw); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func compress(fi os.FileInfo, fileDir string, subName string, zw *zip.Writer) error {
//	if fi.IsDir() {
//		fileDir := filepath.Join(subName, fi.Name())
//		if subName != "" {
//			subName = filepath.Join(subName, fi.Name())
//		} else {
//			subName = fi.Name()
//		}
//
//		header, err := zip.FileInfoHeader(fi)
//		if err != nil {
//			return err
//		}
//		header.Name = subName + "/"
//
//		_, err = zw.CreateHeader(header)
//		if err != nil {
//			return err
//		}
//
//		files, err := ioutil.ReadDir(fileDir)
//		if err != nil {
//			return err
//		}
//		for _, fi := range files {
//			if err := compress(fi, fileDir, subName, zw); err != nil {
//				return err
//			}
//		}
//	} else {
//		filePath := filepath.Join(fileDir, fi.Name())
//		f, err := os.Open(filePath)
//		if err != nil {
//			return err
//		}
//		defer f.Close()
//
//		header, err := zip.FileInfoHeader(fi)
//		if err != nil {
//			return err
//		}
//		if subName != "" {
//			header.Name = filepath.Join(subName, fi.Name())
//		}
//		writer, err := zw.CreateHeader(header)
//		if err != nil {
//			return err
//		}
//		_, err = io.Copy(writer, f)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
