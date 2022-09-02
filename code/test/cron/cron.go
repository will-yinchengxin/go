package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"math/rand"
	"os"
	"time"
)

func CronJob() {
	c := cron.New()
	defer c.Stop()
	err := c.AddFunc("0/5 * * * * ? ", func() {
		os.Create("./filetest/" + randCreator(5) + "will.txt")
	})

	if err != nil {
		fmt.Println("start cron failed", err.Error())
		return
	}

	c.Start()
	fmt.Println("start cron task success")
	for {
		time.Sleep(time.Second)
	}
}

func randCreator(l int) string {
	str := "0123456789abcdefghigklmnopqrstuvwxyz"
	strList := []byte(str)

	result := []byte{}
	i := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i = i + 1
	}
	return string(result)
}
