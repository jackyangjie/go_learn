package wordcount

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func wordCount(text string) (count int) {
	//count = len(strings.Fields(text))
	count = strings.Count(text, "")
	return count
}

func Count() (c int) {
	var fileName = "./word.txt"
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(file)
	c = wordCount(text)
	fmt.Println("文件", fileName, "共有", c, "个字符")
	return
}
