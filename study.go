package main

import (
	"fmt"
	"time"

	"math/rand"
	"strconv"
)

//func init() {
//	rand.Seed(time.Now().Unix())
//	file := "/dist/test.log";
//	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
//	if err != nil {
//		panic(err)
//	}
//	log.SetOutput(logFile) // 将文件设置为log输出的文件
//	log.SetPrefix("[elk-test]")
//	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
//}
//控制日志输出的大小和速率，速度控制为每秒发送N条，发送后sleep一秒钟。忽略了执行时间，控制大小则是for循坏添加字段，字母b是流水号，记载打印了多少

func main() {

	var b = 0
	var N int
	var M int
	fmt.Scanln(&N)
	fmt.Scanln(&M)
	for {
		fmt.Println(N)
		for i := 0; i < N; i++ {
			func() {
				strs := []string{"fb-test1", "fb-test2", "fb-test3", "fb-test4", "fb-test5", "fb-test6", "fb-test7", "fb-test8", "fb-test9"}

				//a, _ := Random(strs, 1)
				a := strs[1]

				//now := time.Now()
				//date := now.Format("2006-01-02 15:04:05")

				userId := rand.Intn(10000)
				b = b + 1
				result := "elon|" + strconv.Itoa(userId+1) + "|" + a + "|" + strconv.Itoa(b)
				for m := 0; m < M; m++ {
					result = result + result
				}
				fmt.Println(result)
				//log.Println(result ) // log 还是可以作为输出的前缀

			}()
		}
		//time.Sleep(1*time.Second/100)
		time.Sleep(2 * time.Second)
	}

}

//func Random(strings []string, length int) (string, error) {
//	if len(strings) <= 0 {
//		return "", errors.New("the length of the parameter strings should not be less than 0")
//	}
//
//	if length <= 0 || len(strings) <= length {
//		return "", errors.New("the size of the parameter length illegal")
//	}
//
//	for i := len(strings) - 1; i > 0; i-- {
//		num := rand.Intn(i + 1)
//		strings[i], strings[num] = strings[num], strings[i]
//	}
//
//	str := ""
//	for i := 0; i < length; i++ {
//		str += strings[i]
//	}
//	return str, nil
//}
