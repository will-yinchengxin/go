package sign

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type signalHandler func(s os.Signal, arg interface{})

type signalSet struct {
	m map[os.Signal]signalHandler
}

func NewSignalSet() (*signalSet) {
	ss := signalSet{}
	ss.m = make(map[os.Signal]signalHandler)
	return &ss

	// ss := new(signalSet)
	// return ss
}

func (s *signalSet) register(os os.Signal, handler signalHandler) {
	if _, ok := s.m[os]; !ok {
		s.m[os] = handler
	}
}

func (s *signalSet) handle(sign os.Signal, arg interface{}) error {
	if _, ok := s.m[sign]; ok {
		s.m[sign](sign, arg)
		return nil
	} else {
		return fmt.Errorf("no handler function for sign %v \n", sign)
	}
}

// --------------------------------test demo-----------------------------
func TestDemo() {
	// -------------------------reg----------------------------
	ss := NewSignalSet()
	handler := func(os os.Signal, arg interface{}) {
		fmt.Printf("handle signal: %v \n", os)
	}

	ss.register(syscall.SIGINT, handler)
	ss.register(syscall.SIGUSR1, handler)
	ss.register(syscall.SIGUSR2, handler)
	// --------------------------------------------------------
	for {
		c := make(chan os.Signal)

		// 拓展操作
		//var sigs []os.Signal
		//for sig := range ss.m {
		//	sigs = append(sigs, sig)
		//}

		signal.Notify(c)
		sig := <- c

		err := ss.handle(sig, nil)
		if err != nil {
			fmt.Printf("unknow signal received %v \n", sig)
			os.Exit(1)
		}
	}
	/*
	# ps -ef | grep main
	root       927   921  0 08:08 pts/5    00:00:00 go run main.go
	root      1059   927  0 08:08 pts/5    00:00:00 /tmp/go-build2884475841/b001/exe/main
	root      1077   911  0 08:11 pts/4    00:00:00 grep main

	# kill -2 927
	# kill -10 1059
	# kill -12 1059

	----------------------------------------------------------------------------------------
	# go run main.go
	handle signal: interrupt             // kill -2 927
	handle signal: user defined signal 1 // kill -10 1059
	handle signal: user defined signal 2 // kill -12 1059
	unknow signal received urgent I/O condition
	exit status 1
	*/
}