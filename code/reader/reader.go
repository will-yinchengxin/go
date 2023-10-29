package reader

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var NoBody struct {
	io.Reader
	io.Closer
}

// IOReader 从头再读取 io.Reader
func IOReader() error {
	r, err := GetDomainList()
	if err != nil {
		log.Fatal(err)
		return err
	}
	r1 := new(http.Response)
	*r1 = *r
	if r1.ContentLength == 0 && r1.Body != nil {
		// Is it actually 0 length? Or just unknown?
		var buf [1]byte
		n, err := r1.Body.Read(buf[:])
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			// Reset it to a known zero reader, in case underlying one
			// is unhappy being read repeatedly.
			r1.Body = NoBody
		} else {
			r1.ContentLength = -1
			r1.Body = struct {
				io.Reader
				io.Closer
			}{
				io.MultiReader(bytes.NewReader(buf[:1]), r.Body),
				r.Body,
			}
		}
	}

	return nil
}

func GetDomainList() (*http.Response, error) {
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/v1/kv/global/service/delivery", "http://127.0.0.1:8455"), nil)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Status Not OK")
	}
	return resp, nil
}
