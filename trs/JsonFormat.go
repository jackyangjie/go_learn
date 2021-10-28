package trs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)


func FormatJson(filepath string,writerFile string)  {
	lines := readJsonFile(filepath)
	fmt.Printf("共有%d 条数据 \n",len(lines))
	dstFile,err := os.Create(writerFile)
	defer dstFile.Close()
	if err != nil{
		fmt.Printf("创建文件失败！\n")
	}
	for _,line := range lines {

		m:= make(map[string]interface{})
		err := json.Unmarshal([]byte(line), &m)
		if err != nil {
			fmt.Printf("Unmarshal with error : %v\n",err)
		}
		recode :=Recode{}
		for key,value := range m{
             if key == "text"{
				 text := value.(string)
				 recode.Text = text
			 } else {
				//of := reflect.TypeOf(value)
				m2,ok := value.(map[string]interface{})
				if ok {
                 eventList := make([]Trigger,0)
					for _,v := range m2{
						//fmt.Printf(" %v: %v \n",k,v)
						m3,ok := v.(map[string]interface{})
						if ok {
							for k3,_ := range m3{
								trigger := Trigger{Trigger: k3,Arguments: make([]Argument,0)}
								eventList = append(eventList, trigger)

								//fmt.Printf("key3 : %v \n",k3)
							}
						}
					}
					recode.Event_list = eventList
				}

				//fmt.Printf("%v: %v ,type: %v \n",key,value,reflect.TypeOf(value))
			}

		}
		marshal, err := json.Marshal(recode)
		if err != nil{
			fmt. Println ( "error:" , err )
		}
		//fmt.Printf("%v \n",marshal)
		dstFile.WriteString(string(marshal)+ "\n" )
		//os.Stdout.Write(marshal)
	}

}



func readJsonFile(filepath string) []string  {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	lines := make([]string,0)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		lines = append(lines, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				break
			}
		}
	}
	return lines
}