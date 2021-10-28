package trs

import (
	"fmt"
);

//since2 := []string{"a","c","ab","dd"}
var since = []string{"a","c","ab","dd"}

func  ForTest(){
	//since := []string{"a","c","ab","dd"}
	 for i,value := range since{
		 fmt.Printf("index:%d,value:%s\n",i,value)
	 }

 }


 func  ForSimple(){
	 for i:=0;i<len(since);i++ {
		 fmt.Printf("simple-> index:%d,value:%s\n",i,since[i])
	 }
 	
 }

var colors =	map[string]string{
"AliceBlue": "#f0f8ff",
"Coral": "#ff7F50",
"DarkGray": "#a9a9a9",
"ForestGreen": "#228b22",
}

func MapTest(){
 	 dict := make(map[string]string)
     dict["abc"]="sfasf";
     value,_ := dict["abc"]
     fmt.Printf("map value:%s \n",value)
 }

 func ForMap(){
 	for key,value := range colors{
 		fmt.Printf("map key:%s,value:%s \n",key,value)

	}

 }

func StutcTest()  {
	var user User
	fmt.Printf(user.Id,user.Name,user.Sex,user.Address,user.Age);
}