package main

import (
	"fmt"
	"trs.com/trs/trs"
)



func main() {
	fmt.Println("Hello word!")
	trs.Test();
	fmt.Println(comp(10,43))
	trs.ForTest();

	trs.ForSimple()

	trs.MapTest();

	trs.ForMap();

	trs.StutcTest();
	var writeFile = "C:\\Users\\71908\\Downloads\\train.txt"
	var readFile ="C:\\Users\\71908\\Downloads\\train.json"
    trs.FormatJson(readFile,writeFile)  // train 数据

	var writeDevFile = "C:\\Users\\71908\\Downloads\\dev.txt"
	var readDevFile ="C:\\Users\\71908\\Downloads\\dev.json"
    trs.FormatJson(readDevFile,writeDevFile)

	var writeTestFile = "C:\\Users\\71908\\Downloads\\test.txt"
	var readTestFile ="C:\\Users\\71908\\Downloads\\test.json"
	trs.FormatJson(readTestFile,writeTestFile)
}

func comp(x int ,y int) (int)  {
	if x> y {
		return x;
	}else{
		return y;
	}
}