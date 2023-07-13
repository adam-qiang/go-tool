/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-07 00:43:57
 */

package gotool

import (
	"encoding/json"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"time"
	"unsafe"
)

// MapToString MAP数据转字符串
// @param mapData interface{}
// @return stringData string
func MapToString(mapData interface{}) (stringData string) {
	dataJson, err := json.Marshal(mapData)
	if err != nil {
		return
	}

	stringData = string(dataJson)

	return
}

// StringToMap MAP数据转成字符串数据后再转回MAP
// @param stringData string
// @return mapData map[string]interface{}
func StringToMap(stringData string) (mapData map[string]interface{}) {
	mapData = make(map[string]interface{})
	byteData := StringToBytes(stringData)

	err := json.Unmarshal(byteData, &mapData)
	if err != nil {
		return nil
	}
	return
}

// StringToIntMap MAP数据转成字符串数据后再转回int类型KEY的map
// @param stringData string
// @return mapData map[int]string
func StringToIntMap(stringData string) (mapData map[int]string) {
	mapData = make(map[int]string)
	byteData := StringToBytes(stringData)

	err := json.Unmarshal(byteData, &mapData)
	if err != nil {
		return nil
	}
	return
}

// ArrStringToString 数组（值为string）转字符串
// @param arrData []string
// @return stringData string
func ArrStringToString(arrData []string) (stringData string) {
	dataJson, err := json.Marshal(arrData)
	if err != nil {
		return ""
	}

	stringData = string(dataJson)

	return stringData
}

// StringToArrString 字符串转数组（值为string）
// @param stringData string
// @return mapData []string
func StringToArrString(stringData string) (mapData []string) {
	mapData = make([]string, 0)
	byteData := StringToBytes(stringData)

	err := json.Unmarshal(byteData, &mapData)
	if err != nil {
		return nil
	}
	return
}

// InterfaceMapToMap interface类型map转为[string]interface{}类型
// @param data interface{}
// @return mapData map[string]interface{}
func InterfaceMapToMap(data interface{}) (mapData map[string]interface{}) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	stringData := string(dataJson)

	mapData = make(map[string]interface{})

	err = json.Unmarshal([]byte(stringData), &mapData)
	if err != nil {
		return nil
	}
	return
}

// JsonToMap JSON（byte）转MAP
// @param stuObj struct{}
// @return map[string]interface{}
// @return error
func JsonToMap(jsonByte []byte) (map[string]interface{}, error) {
	var mRet map[string]interface{}
	err1 := json.Unmarshal(jsonByte, &mRet)
	if err1 != nil {
		return nil, err1
	}
	return mRet, nil
}

// InterfaceToStrVal interface类型值转为字符串
// @param value interface{}
// @return string
func InterfaceToStrVal(value interface{}) string {
	// interface 转 string
	var val string
	if value == nil {
		return val
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		val = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		val = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		val = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		val = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		val = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		val = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		val = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		val = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		val = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		val = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		val = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		val = strconv.FormatUint(it, 10)
	case string:
		val = value.(string)
	case []byte:
		val = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		val = string(newValue)
	}

	return val
}

// StringToBytes 字符串转byte
// @param data string
// @return []byte
func StringToBytes(data string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&data))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// DayToWeek 将时间转为星期
// @param time.Time
// @return string
func DayToWeek(time time.Time) string {
	weeks := make(map[string]string)
	weeks["Monday"] = "星期一"
	weeks["Tuesday"] = "星期二"
	weeks["Wednesday"] = "星期三"
	weeks["Thursday"] = "星期四"
	weeks["Friday"] = "星期五"
	weeks["Saturday"] = "星期六"
	weeks["Sunday"] = "星期日"

	week := time.Weekday().String()

	return weeks[week]
}

// MonthToQuarter 将月份转为季度
// @param month int
// @return string
func MonthToQuarter(month int) string {
	var quarter string
	if month >= 1 && month <= 3 {
		//1月1号
		quarter = "Q1"
	} else if month >= 4 && month <= 6 {
		quarter = "Q2"
	} else if month >= 7 && month <= 9 {
		quarter = "Q3"
	} else {
		quarter = "Q4"
	}
	return quarter
}

// TimeRemark 目标时间距离当前时间时长
// @param in string
// @return out string
func TimeRemark(in string) (out string) {
	//当前时间戳
	timeUnix := time.Now().Unix()
	inn, _ := strconv.ParseInt(in, 10, 64)
	outt := timeUnix - inn

	f := map[string]string{
		"1":        "秒",
		"60":       "分钟",
		"3600":     "小时",
		"86400":    "天",
		"604800":   "星期",
		"2592000":  "个月",
		"31536000": "年",
	}
	var keys []string
	for k := range f {
		keys = append(keys, k)
	}
	//sort.Strings(keys)
	//数字字符串 排序
	sort.Slice(keys, func(i, j int) bool {
		numA, _ := strconv.Atoi(keys[i])
		numB, _ := strconv.Atoi(keys[j])
		return numA < numB
	})
	for _, k := range keys {
		v2, _ := strconv.Atoi(k)
		cc := math.Floor(float64(int(outt) / int(v2)))
		if 0 != cc {
			out = strconv.FormatFloat(cc, 'f', -1, 64) + f[k] + "前"
		}
	}
	return
}

// RandomStr 随机生成指定长度字符串
// @param length int
// @return string
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
