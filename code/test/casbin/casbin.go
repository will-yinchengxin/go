package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func Casbin() {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "r.sub == g.sub && r.obj == p.obj && r.act == p.act")

	a := fileadapter.NewAdapter("D:\\Project\\test\\casbin\\policy.csv")
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "will", "data1", "read")
	check(e, "will", "data2", "write")
	check(e, "yin", "data1", "write")
	check(e, "yin", "data2", "read")
}