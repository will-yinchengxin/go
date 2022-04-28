package atomic

import (
	"errors"
	"sync/atomic"
)

// 并发安全的整数数组接口
type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
}

type MyArray struct {
	val atomic.Value
	len int32
}

func (array *MyArray) CheckValue() (err error) {
	if array.val.Load() == nil {
		return errors.New("array is empty")
	}
	return nil
}

func (array *MyArray) CheckIndex(index int32) (err error) {
	if array.len <= index {
		return errors.New("array out of range")
	}
	return nil
}

func (array *MyArray) Set(index int32, elem int) (err error) {
	if err := array.CheckValue(); err != nil {
		return err
	}

	if err := array.CheckIndex(index); err != nil {
		return err
	}

	newArray := make([]int, array.len)
	copy(newArray, array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return nil
}

func (array *MyArray) Get(index int32) (elem int, err error) {
	if err := array.CheckValue();err != nil{
		return 0,err
	}

	if err = array.CheckIndex(index);err!=nil{
		return 0,err
	}

	num := array.val.Load().([]int)[index]
	return num, err
}
