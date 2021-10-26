package snowflake

import (
	"time"
	sf "github.com/bwmarrin/snowflake"
)

/*
	const (
		StartTime = "2021-10-13"
		MachineID = 1
	)
	snowflake.Init(consts.StartTime, consts.MachineID)
	snowflake.GenID()
*/

var Node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	Node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return Node.Generate().Int64()
}