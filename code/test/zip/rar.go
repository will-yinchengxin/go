package zip

import (
	"encoding/xml"
	"github.com/mholt/archiver/v3"
	"io/ioutil"
	"strings"
)

func Rar() error{
	z := archiver.NewRar()
	if err := z.Walk("E:\\ota\\whole-ota(2).rar", func(f archiver.File) error {
		defer f.Close()
		Split := strings.Split(f.Name(), "/")
		if len(Split) == 2 && Split[1] == ".rar" {
			data, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}
			Manifest := Manifest{}
			err = xml.Unmarshal(data, &Manifest)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}