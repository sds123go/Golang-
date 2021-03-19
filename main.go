package main

import "fmt"
import "mytools/mytools"
type test struct {
	Name string
	Age  int64
}

func main() {
	testdata := test{
		Name: "sds",
		Age:  25,
	}
	result := mytools.Struct2Map(testdata)
	fmt.Println(result)
	fmt.Println("***********************")
	timestamp:=mytools.CurrentTimeStamp()
	fmt.Println(timestamp)
	fmt.Println("***********************")
	datatime:=mytools.CurrentDataTime()
	fmt.Println(datatime)
}
