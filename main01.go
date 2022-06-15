package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := "abcdefghijklmnopqrstuvwxyz123456789"

	//设置随机数种子
	rand.Seed(time.Now().UnixNano())

	//所有的字符串为codes，[]string类型，共10个
	codes := make([]string, 0)
	for i := 0; i < 10; i++ {
		//其中一个字符串为code，[]byte类型
		code := make([]byte, 0)
		for i := 0; i < 6; i++ {
			//设置随机数(随机索引)
			index := rand.Intn(len(s))
			code = append(code, s[index])
		}
		fmt.Printf("%s\n", code)
		codes = append(codes, string(code))
	}
	fmt.Println(codes)

	//duplicate去重
	m := make(map[string]int)
	for k, v := range codes {
		fmt.Println(k, v)
		m[v] = k
	}
	fmt.Printf("m = %+v\n", m)
	//打印新map的key值，即打印去重后的codes
	for k, _ := range m {
		fmt.Println(k)
	}
}
