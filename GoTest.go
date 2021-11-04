package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"trs.com/trs/trs"
)

func main() {

	get, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}

	//err = xml.NewDecoder(get.Body).Decode(&doc)
	buf := new(strings.Builder)
	io.Copy(buf, get.Body)
	body := buf.String()
	//fmt.Printf(body)
	log.Printf(body)
	fmt.Println("Hello word!")
	trs.Test()
	fmt.Println(comp(10, 43))
	trs.ForTest()

	trs.ForSimple()

	trs.MapTest()

	trs.ForMap()

	trs.StutcTest()
	var writeFile = "C:\\Users\\71908\\Downloads\\train.txt"
	var readFile = "C:\\Users\\71908\\Downloads\\train.json"
	trs.FormatJson(readFile, writeFile) // train 数据

	var writeDevFile = "C:\\Users\\71908\\Downloads\\dev.txt"
	var readDevFile = "C:\\Users\\71908\\Downloads\\dev.json"
	trs.FormatJson(readDevFile, writeDevFile)

	var writeTestFile = "C:\\Users\\71908\\Downloads\\test.txt"
	var readTestFile = "C:\\Users\\71908\\Downloads\\test.json"
	trs.FormatJson(readTestFile, writeTestFile)
}

func comp(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
