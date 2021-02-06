package mytools

import "reflect"

//Struct2Map 结构体转字典
func Struct2Map(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV:=reflect.ValueOf(obj)
	for i:=0;i<objT.NumField();i++{
		result[objT.Field(i).Name]=objV.Field(i).Interface()
	}
	return result
}
