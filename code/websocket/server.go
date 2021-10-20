package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

var (
	// 新用户进来进行登记
	enteringChannel = make(chan *User)
	// 用户离开进行等级
	leavingChannel = make(chan *User)
	// 广播专用的用户普通消息channel
	// 缓冲是尽可能避免出现异常状况堵塞,这里简单给了8
	messageChannel = make(chan *Message, 8)
)

func main() {
	// 这里未指定 ip 所以绑定在本机的 ip 上
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	// 用于记录聊天室用户,并进行消息广播
	go broadcaster()

	fmt.Println("start tcp server 2020")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster 用于记录聊天室用户,并进行消息广播
// 1.新用户进来 2.用户普通消息 3.用户离开
func broadcaster() {
	users := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			// 新用户加入
			users[user] = struct{}{}
		case user := <-leavingChannel:
			// 用户离开
			delete(users, user)
			close(user.MessageChannel)
		case msg := <-messageChannel:
			// 给所有在线用户发送消息
			for user := range users {
				// 过滤掉自己的消息
				if user.ID == msg.OwnerID {
					continue
				}
				user.MessageChannel <- msg.Content
			}
		}
	}
}

// 处理连接请求
func handleConn(conn net.Conn) {
	defer conn.Close()
	// 1.新用户进来,构建用户实例
	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	// 2.给当前用户发送欢迎消息,并告知所有用户新用户的加入
	go sendMessage(conn, user.MessageChannel)
	user.MessageChannel <- "Welcome, " + user.String()

	// 3.记录至全局用户表中,避免用锁
	enteringChannel <- user

	// 4.剔除不活跃的用户(5分钟时效)
	var userActive = make(chan struct{})
	go func() {
		d := 5 * time.Minute
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				conn.Close()
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	// 5.读取客户端的输入内容
	input := bufio.NewScanner(conn)
	var msg Message
	for input.Scan() {
		msg.Content = strconv.Itoa(user.ID) + ": " + input.Text()
		msg.OwnerID = user.ID
		messageChannel <- &msg

		// 用户活跃
		userActive <- struct{}{}
	}

	if err := input.Err(); err != nil {
		log.Println("读取数据错误: ", err)
	}

	// 用户离开
	leavingChannel <- user
	msg.Content = "user " + strconv.Itoa(user.ID) + " has left"
	messageChannel <- &msg
}

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

func (u *User) String() string {
	return strconv.Itoa(u.ID)
}

type Message struct {
	OwnerID int
	Content string
}

// 想客户端发送消息
func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

// 生成 userId 工具类
func GenUserID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100000)
}
