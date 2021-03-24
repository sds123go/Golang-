package Mytools

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

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

const (
	zero  = byte('0')
	one   = byte('1')
	lsb   = byte('[') // left square brackets
	rsb   = byte(']') // right square brackets
	space = byte(' ')
)

// BytesToBinaryString get the string in binary format of a []byte or []int8.
func BytesToBinaryString(bs []byte) string {
	l := len(bs)
	if l == 0 {
		return ""
	}
	bl := l*8 + l + 1
	buf := make([]byte, 0, bl)
	buf = append(buf, lsb)
	for _, b := range bs {
		buf = appendBinaryString(buf, b)
		buf = append(buf, space)
	}
	buf[bl-1] = rsb
	return string(buf)
}

// append bytes of string in binary format.
func appendBinaryString(bs []byte, b byte) []byte {
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			bs = append(bs, zero)
		default:
			bs = append(bs, one)
		}
		b <<= 1
	}
	return bs
}
