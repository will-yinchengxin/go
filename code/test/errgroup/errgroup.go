package errgroup

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func ErrGroup()  {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.baidu.com/",
		"http://www.noexist11111111.com/",
	}
	g := new(errgroup.Group)
	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Printf("get [%s] success: [%d] \n", url, resp.StatusCode)
			return resp.Body.Close()
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("all success!")
	}
}
