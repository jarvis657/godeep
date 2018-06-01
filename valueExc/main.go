package main

import (
	"sync"
	"time"
	"fmt"
)

//var local sync.Mutex
type data struct {
	sync.Mutex
}

func (d data) test(s string) {
	d.Lock()
	defer func() {
		d.Unlock()
		println("success")
	}()
	for i := 0; i < 5; i++ {
		fmt.Printf("%s, %d, object addr: %p \n", s, i, &d)
		time.Sleep(time.Second)
	}
}

/*
值传递每次调用的地址都会不一样
 */
func main() {
	//var wg sync.WaitGroup
	//wg.Add(2)
	var d data
	//go func() {
	//	defer wg.Done()
	d.test("read")
	//}()
	//go func() {
	//	defer wg.Done()
	d.test("write")
	//}()
	//wg.Wait()
}
