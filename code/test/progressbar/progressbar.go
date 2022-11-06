package progressbar

import (
	"github.com/schollz/progressbar"
	"time"
)

func progressbarTest() {
	bar := progressbar.New(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
  
  /*
  $ go run main.go
  100% |████████████████████████████████████████| [4s:0s] 
  */
}
