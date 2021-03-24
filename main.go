package main

import "fmt"
import "Mytools/Mytools"
type test struct {
	Name string
	Age  int64
}

func main() {
	// testdata := test{
	// 	Name: "sds",
	// 	Age:  25,
	// }
	// result := mytools.Struct2Map(testdata)
	// fmt.Println(result)
	// fmt.Println("***********************")
	// timestamp:=mytools.CurrentTimeStamp()
	// fmt.Println(timestamp)
	// fmt.Println("***********************")
	// datatime:=mytools.CurrentDataTime()
	// fmt.Println(datatime)
	bytedata:=[]byte{}
	//intdata:=258
	//result:=mytools.IntToBytes(intdata)
	result:=mytools.BytesToBinaryString(bytedata)
	fmt.Print(result)
}
