package mytools

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"time"
)

//Struct2Map 结构体转字典
func Struct2Map(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		result[objT.Field(i).Name] = objV.Field(i).Interface()
	}
	return result
}

//CurrentTimeStamp 当前时间戳
func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

//CurrentDataTime 当前日期和时间
func CurrentDataTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//BytesToInt 字节转64位整型
func BytesToInt64(bys []byte) int64 {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	err := binary.Read(bytebuff, binary.BigEndian, &data)
	if err != nil {
		panic(err)
	}
	return data
}
