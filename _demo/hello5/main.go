package main

import (
	"sync"
	"fmt"
)

func main() {
	// WaitGroup 对象内部有一个计数器,最初从0开始,它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
	// Add(n) 把计数器设置为n ,Done() 每次把计数器-1 ,wait() 会阻塞代码的运行,直到计数器地值减为0
	wg := sync.WaitGroup{}
	chTom := make(chan string)
	chJerry := make(chan string)

	wg.Add(1)
	// Tom task
	go func() {
		for {
			w := <-chTom
			if w == "Hello Tom" {
				fmt.Println("Tom : Yeah, Hello Jerry.")
				chJerry <- "Hello Jerry"
			} else {
				fmt.Println("Tom : Nice Day.")
				chJerry <- "Nice Day"
				wg.Done()
				break
			}
		}
	}()

	wg.Add(1)
	// Jerry task
	go func() {
		for {
			w := <-chJerry
			if w == "Hello Jerry" {
				fmt.Println("Jerry : Fine, Tom.")
				chTom <- "Fine Tom"
			} else {
				wg.Done()
				break
			}
		}
	}()

	//fmt.Println("Jerry : Hello Tom.")
	//chTom <- "Hello Tom"
	fmt.Println("Tom : Hello Jerry.")
	chJerry <- "Hello Jerry"
	wg.Wait()
}

// func main() {
//     wg := sync.WaitGroup{}
//     wg.Add(100)
//     for i := 0; i < 100; i++ {
//         go f(i, &wg)
//     }
//     wg.Wait()
// }

// 一定要通过指针传值，不然进程会进入死锁状态
// func f(i int, wg *sync.WaitGroup) { 
//     fmt.Println(i)
//     wg.Done()
// }