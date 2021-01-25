// @Title  对一些常见数据类型的处理函数
// @Description  包括类型转换，判断数组中是否存在某个值
package common

import (
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
	"time"
)

//接口转string
func Interface2String(data interface{}) string {
	if data == nil {
		return ""
	}
	res, ok := data.(string)
	if !ok {
		return ""
	} else {
		return res
	}
}

//接口转int
func Interface2Int(data interface{}) int32 {
	if data == nil {
		return 0
	}
	res, ok := data.(int32)
	if ok {
		return res
	} else {
		return 0
	}
}

//时间转string
func Time2String(times time.Time, showHour bool) string {
	var cstZone = time.FixedZone("CST", 8*3600)
	if showHour {
		return times.In(cstZone).Format("2006-01-02 15:04:05")
	}
	return times.In(cstZone).Format("2006/01/02")
}

//字符串转时间
func Str2time(t string) time.Time {
	loc, _ := time.LoadLocation("Local")
	ti, err := time.ParseInLocation("2006-01-02 15:04:05", t, loc)
	if err == nil {
		return ti
	}
	return time.Now()
}

//primitive.D数据转结构体
func Primitive2Struct(data interface{}, result interface{}) error {
	bsonBytes, err := bson.Marshal(data)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bsonBytes, result)
	if err != nil {
		return err
	}
	return nil
}

//结构体数据转bson.m
func Struct2Bson(data interface{}, result interface{}) error {
	return Primitive2Struct(data, result)
}

//字符串转int
func Str2Int(dataS string) int {
	data, err := strconv.Atoi(dataS)
	if err != nil {
		return 0
	}
	return data
}

//int转字符串
func Int2Str(data int) string {
	return strconv.Itoa(data)
}

//判断int 数组中是否存在某个值
func IsInIntArray(data []int, key int) bool {
	for _, v := range data {
		if key == v {
			return true
		}
	}
	return false
}

//判断string 数组中是否存在某个值
func IsInStringArray(data []string, key string) bool {
	for _, v := range data {
		if key == v {
			return true
		}
	}
	return false
}

// 接口转float64
func Interface2Float(data interface{}) float64 {
	if data == nil {
		return 0
	}
	res, ok := data.(float64)
	if !ok {
		return 0
	} else {
		return res
	}
}

//float64接口转int
func InterfaceFloat2Int(data interface{}) int {
	if data == nil {
		return 0
	}
	res, ok := data.(float64)
	if !ok {
		return 0
	} else {
		return int(res)
	}
}

//float64接口转字符串
func InterfaceFloat2String(data interface{}) string {
	if data == nil {
		return ""
	}
	res, ok := data.(float64)
	if !ok {
		return ""
	} else {
		return strconv.Itoa(int(res))
	}
}

// 把字符串批量转换为int数组
func String2IntArray(data string, sep string) []int {
	strArray := strings.Split(data, sep)
	var returnData []int
	for _, v := range strArray {
		returnData = append(returnData, Str2Int(v))
	}
	return returnData
}

// 字符串转bool值
func Str2Bool(data string) bool {
	if data == "true" {
		return true
	}
	return false
}
