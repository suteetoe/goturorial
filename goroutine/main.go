package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// start := time.Now()
	// go doSomething()
	// go doSomething()
	// go doSomething()
	// time.Sleep(120 * time.Millisecond)
	// fmt.Println(time.Since(start))

	// example()

	// // การใช้ waitgroup
	// wg := &sync.WaitGroup{}

	// wg.Add(3)

	// start := time.Now()
	// go doSomethingWaitDone(wg)
	// go doSomethingWaitDone(wg)
	// go doSomethingWaitDone(wg)

	// wg.Wait()
	// fmt.Println(time.Since(start))

	testRaceCondition()
}

func doSomething() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
}

func doSomethingWaitDone(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
	wg.Done()
}

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Print("+")
		}
	}()
	time.Sleep(time.Second)

}

// race conditin

var i int
var mux sync.Mutex // มาเพื่อแก้ไข race condition

// // โดยปรกติ จะเอา i กับ mutex ไปไว้ด้วยกัน
// type thing struct {
// 	i   int
// 	mux sync.Mutex
// }

func testRaceCondition() {
	go func() {
		for {
			fmt.Println(read())
		}
	}()

	for i := 0; i < 10; i++ {
		write(i)
	}
}

func write(n int) {
	mux.Lock() // ถ้าไม่ lock อาจจะมี race condition เกิดขึ้น
	i = n
	mux.Unlock() // set ค่าเสร็จแล้ว unlock
}

func read() int {
	mux.Lock()
	defer mux.Unlock()
	return i
}
