package oss

import (
	"fmt"
)

type Oss struct {
	Moss *MOSS
}

func (e *Oss) SetConfig() *Oss {
	SetConfig()
	return &Oss{}
}

func (e *Oss) Policy()  {
	res, err := e.Moss.InitConfig().GetPolicyToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
