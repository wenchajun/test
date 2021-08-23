package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	flag "github.com/spf13/pflag"
)

func init() {
	rand.Seed(time.Now().Unix())
	//file := "/dist/test.log";
	file := "./test.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件

}

var N = flag.IntP("cycles", "N", 10, "Cycles per second")
var M = flag.IntP("logSize", "M", 5, "Output log size")
var C2 = flag.IntP("time", "C", 60, "time")
var serial = 0 //流水号
var lock sync.Mutex

//var quit chan int

func main() {
	flag.Parse()
	var c = make(chan int, *N)
	for o := 0; o < *C2; o++{

		//fmt.Println(*N)
		//fmt.Println(*M)

		start := time.Now().UnixNano() //.UnixNano()//获取现在的时间

		//如果每秒发送2000条内就不用开启多线程
		//超过2000条，开启n/2000个线程/协程
		if *N < 1000 {
			Log(*N, *M)
		} else {
			threadN := (*N) / 500

			for i := 0; i < *N; i++ {
				serial++
				c <- serial
			}
			for i := 0; i < threadN; i++ {
				//go concurrencylog(cap(c),c,*M)
				go concurrencylog(c, *M)
			}

		}
		end := time.Now().UnixNano()
		elapsed := end - start
		//elapsed := time.Since(t1).Nanoseconds()//计算过去了多久的时间

		//fmt.Println("-------------")
		//fmt.Println(elapsed)
		//	fmt.Println("-------------")
		//	fmt.Println(elapsed)
		//	log.Println(elapsed)
		if elapsed < 1000000000 {
			time.Sleep(time.Duration(1000000000-elapsed) * time.Nanosecond)
		}
	}

}

//日志输出函数

func Log(cycles int, logSize int) {
	for i := 0; i < cycles; i++ {
		func() {
			a := "fluent-bit-test"
			//date := now.Format("2006-01-02 15:04:05")
			//userId := rand.Intn(10000)

			//lock.Lock()
			serial = serial + 1
			//log.Println("----------"+ "--------------",serial)
			//lock.Unlock()








			result := "name:eloncheng|" + "|" + a + "|" + "The current number is" + strconv.Itoa(serial)
            _=result
			fmt.Println("Dec 14 06:41:08 Exception : wrong !")
			fmt.Println(" com.myproject.2)")
			fmt.Println(" at com.myproject)")
			log.Println("Dec 14 06:41:08 Exception in thread main java.lang.RuntimeException: Something has gone wrong, aborting!")
			log.Println(" com.myproject.module.MyProject.badMethod(MyProject.java:22)")
			log.Println(" at com.myproject.module.MyProject.oneMoreMethod(MyProject.java:18)")
			log.Println(" at com.myproject.module.MyProject.anotherMethod(MyProject.java:14)")
			log.Println(" at com.myproject.module.MyProject.someMethod(MyProject.java:10)")
			log.Println(" at com.myproject.module.MyProject.main(MyProject.java:6)")


			//for m := 0; m < logSize; m++ {
			//	log.Println(result) // log 还是可以作为输出的前缀
			//	fmt.Println(result)
			//}

			//log.Println(result) // log 还是可以作为输出的前缀
			//fmt.Println(result)


		}()
	}
}

func concurrencylog(c chan int, logSize int) {
	for {
		lock.Lock()
		No, ok := <-c
		if !ok {
			break
		}
		a := "fluent-bit-test"
		// 发消息：我执行完啦！
		result := "name:eloncheng|" + "|" + a + "|" + "The current number is" + strconv.Itoa(No)
		for m := 0; m < logSize; m++ {
			//result = result + result
			log.Println(result)
			fmt.Println(result)
		}
		//log.Println(result)
		//fmt.Println(result)
		lock.Unlock()

	}
	close(c)
}
