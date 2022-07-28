package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting, i)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 睡眠片刻（随机0到2.5秒）
	}
	wg.Done() // 通知当前任务已完成
}

// 延迟调用函数可以修改包含此延迟调用函数的最内层函数的返回值
func delayCall(n int) (r int) {
	defer func() {
		r += n
	}()
	return n + n
}

// 恐慌（panic）和恢复（recover）
func test() {
	defer func() {
		fmt.Println("正常退出了")
	}()
	fmt.Println("hello")
	defer func() {
		//v := recover()
		fmt.Println("Panic被恢复了")
	}()
	panic("我要制造恐慌")
	fmt.Println("我这里执行不到")
}

func main() {
	test()
	//var result = delayCall(5)
	//fmt.Println(result)

	//延迟调用函数
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)

	//fmt.Println(runtime.GOMAXPROCS(10))

	//fmt.Println(time.Second)
	//rand.Seed(time.Now().UnixNano())
	//log.SetFlags(0)
	//// 注册两个新任务
	//wg.Add(2)
	//go SayGreetings("hi!", 10)
	//go SayGreetings("hello!", 10)
	//wg.Wait() // 阻塞，知道当前任务完成
	//fmt.Println("2121")
}
