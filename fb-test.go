package main

import (
	"fmt"
	//"errors"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
	file := "/dist/test.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[elk-test]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func main() {

	var b = 0

	for {
		go func() {
			strs := []string{"fb-test1", "fb-test2", "fb-test3", "fb-test4", "fb-test5", "fb-test6", "fb-test7", "fb-test8", "fb-test9"}

			//a, _ := Random(strs, 1)
			a := strs[1]

			//now := time.Now()
			//date := now.Format("2006-01-02 15:04:05")

			userId := rand.Intn(10000)
			b = b + 1
			result := "elon|" + strconv.Itoa(userId+1) + "|" + a + "|" + strconv.Itoa(b)
			fmt.Println(result)
			//log.Println(result ) // log 还是可以作为输出的前缀

		}()
		time.Sleep(1 * time.Second / 100)
		//time.Sleep(1*time.Second)
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
