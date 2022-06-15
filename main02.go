package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {

	//生成10个短code
	s := "abcdefghijklmnopqrstuvwxyz123456789"
	rand.Seed(time.Now().UnixNano())
	codes := make([]string, 0)
	for i := 0; i < 10; i++ {
		code := make([]byte, 0)
		for i := 0; i < 6; i++ {
			index := rand.Intn(len(s))
			code = append(code, s[index])
		}
		codes = append(codes, string(code))
	}
	//duplicate去重
	m := make(map[string]int)
	for k, v := range codes {
		fmt.Println(k, v)
		m[v] = k
	}
	for k, v := range m {
		fmt.Println(v, k)
	}

	i := 0
	a := make([]string, 10)
	for k, _ := range m {

		a[i] = k
		i++
	}
	fmt.Println(a)

	//创建excel文件
	xlsx := excelize.NewFile()

	//创建新表单
	index := xlsx.NewSheet("code")

	for k, v := range a {
		col := "A" + strconv.FormatInt(int64(k+1), 10)
		err := xlsx.SetCellValue("code", col, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	err := xlsx.SaveAs("code.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
