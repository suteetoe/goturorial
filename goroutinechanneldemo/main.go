package main

import "fmt"

func main() {
	// ch := make(chan int)

	// go fibonacci(ch)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }

	// เอา fibonacci มาปรับทำ selected statement

	ch := make(chan int)
	qCh := make(chan struct{}) // ประกาศ channel โดยไม่ให้กิน พื้นที่ memory คือให้เป็น nil pointer แหละ หรือใช้ [0]func ก็ได้

	// [0]func

	go fibonacciQuiteChannel(ch, qCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	qCh <- struct{}{}
	fmt.Println("end")
}

func fibonacci(ch chan int) {
	a, b := 0, 1

	for {
		ch <- a
		a, b = b, a+b
	}
}

func fibonacciQuiteChannel(ch chan int, qCh chan struct{}) {
	a, b := 0, 1

	for {
		select {
		case <-qCh:
			qCh <- struct{}{}
			return
		case ch <- a:
			a, b = b, a+b

		}
	}
}
