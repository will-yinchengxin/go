package dto

import (
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type ListSplitter struct {
	LimitSize int
	Message   []*primitive.Message
	CurrIndex int
}

func NewListSplitter(msg []*primitive.Message) *ListSplitter {
	return &ListSplitter{
		LimitSize: 10 * 1, // 这里动态调整即可
		CurrIndex: 0,
		Message:   msg,
	}
}

func (l *ListSplitter) ListSplitter(messageList []*primitive.Message) {
	l.Message = messageList
}

func (l *ListSplitter) HasNext() bool {
	return l.CurrIndex < len(l.Message)
}

func (l *ListSplitter) Next() []*primitive.Message {
	nextIndex := l.CurrIndex
	totalSize := 0
	for nextIndex < len(l.Message) {
		msg := l.Message[nextIndex]

		/*
			tmpSize:是一条消息的空间大小
			tmpSize=topic的长度和Body的长度
		*/
		tmpSize := len(msg.Topic) + len(msg.Body)
		//properties有标签等信息
		properties := msg.GetProperties()
		for key, val := range properties {
			tmpSize += len(key) + len(val)
		}
		//日志头信息
		tmpSize = tmpSize + 20 //for log overhead

		//如果一条消息就超过4MB，记录下来，
		if tmpSize > l.LimitSize {
			if nextIndex-l.CurrIndex == 0 {
				//用于截取列表长度
				nextIndex++
			}
			break
		}
		//超过1MB，就发送
		if tmpSize+totalSize > l.LimitSize {
			break
		} else {
			//累加totalSize
			totalSize += tmpSize
		}

		nextIndex++
	}

	//截取消息列表，返回
	SubList := l.Message[l.CurrIndex:nextIndex]
	l.CurrIndex = nextIndex
	return SubList
}
