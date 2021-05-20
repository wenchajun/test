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
	log.SetPrefix("[elk-test]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

var N = flag.IntP("cycles", "N", 10, "Cycles per second")
var M = flag.IntP("logSize", "M", 2, "Output log size")
var serial = 0 //流水号
var lock sync.Mutex

func main() {
	flag.Parse()

	for {
		fmt.Println(*N)
		fmt.Println(*M)

		start := time.Now().UnixNano() //.UnixNano()//获取现在的时间

		//如果每秒发送2000条内就不用开启多线程
		//超过2000条，开启n/2000个线程/协程
		if *N < 2000 {
			Log(*N, *M)
		} else {
			threadN := (*N) / 20000
			for i := 0; i < threadN; i++ {
				go Log(*N, *M)
			}
		}
		end := time.Now().UnixNano()
		elapsed := end - start
		//elapsed := time.Since(t1).Nanoseconds()//计算过去了多久的时间

		fmt.Println("-------------")
		fmt.Println(elapsed)
		fmt.Println("-------------")

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

			lock.Lock()
			serial = serial + 1
			//log.Println("----------"+ "--------------",serial)
			//lock.Unlock()

			result := "elon|"  + "|" + a + "|" +"本次流水号是"+ strconv.Itoa(serial)

			for m := 0; m < logSize; m++ {
				result = result + result
			}
			//fmt.Println(result)
            //fmt.Println("----------"+ "--------------",serial)
			log.Println(result) // log 还是可以作为输出的前缀
			lock.Unlock()
		}()
	}
}
