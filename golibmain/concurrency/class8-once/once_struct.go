package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// OnceStruct 是一个扩展的sync.Once类型，提供了一个Done方法
type OnceStruct struct {
	sync.Once
}

// Done 返回此Once是否执行过
// 如果执行过则返回true
// 如果没有执行过或者正在执行，返回false
func (o *OnceStruct) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.Once))) == 1
}

func main() {
	var flag OnceStruct
	fmt.Println("flag.Done()=>", flag.Done()) // false

	flag.Do(func() {
		time.Sleep(time.Second)
	})

	fmt.Println("flag.Done()=>", flag.Done()) // true
}
