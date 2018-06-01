package main

import (
	"sync"
	"fmt"
	"time"
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
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		defer wg.Done()
		d.test("read")
	}()
	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()
}
