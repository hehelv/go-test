package main

import (
	"fmt"
)

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// 每个协程的prime是不变的, 都是上一个协程输出的第一个数, 如 第一个协程的prime是2 ， 第二个协程是 3，
func filter(in chan int, out chan int, prime int) {
	for {
		num := <-in
		if num%prime != 0 {
			out <- num
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	// 没有wait， 循环完了就会退出
	for i := 0; i < 6; i++ {
		prime := <-ch
		fmt.Printf("prime:%d\n", prime)
		// 每次循环都会创建一个新的out, 这里一共会开6个协程, 第一个协程的out 会成为第二个协程的in 以此类推,形成一个类似这样的结构
		/*
			    				2			   3			5				7			  11
		for循环一直推数-->in-->filter-->out-->filter-->out-->filter-->out-->filter-->out-->filter

		这样每个filter都是素数，并且都是连续的
		 */
		out := make(chan int)
		go filter(ch, out, prime)
		ch = out
	}
}