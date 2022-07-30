package amountdecimal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GetType(variable interface{}) string {
	return reflect.TypeOf(variable).Kind().String()
}

func StrToUpper(str string) string {
	return strings.ToUpper(str)
}

func Float64Decimal(f float64) int {
	numstr := fmt.Sprint(f)
	tmp := strings.Split(numstr, ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}

func InterfaceToStr(val interface{}) string {
	var str string

	if val == nil {
		return str
	}

	switch val.(type) {
	case float64:
		ft := val.(float64)
		str = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := val.(float32)
		str = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := val.(int)
		str = strconv.Itoa(it)
	case uint:
		it := val.(uint)
		str = strconv.Itoa(int(it))
	case int8:
		it := val.(int8)
		str = strconv.Itoa(int(it))
	case uint8:
		it := val.(uint8)
		str = strconv.Itoa(int(it))
	case int16:
		it := val.(int16)
		str = strconv.Itoa(int(it))
	case uint16:
		it := val.(uint16)
		str = strconv.Itoa(int(it))
	case int32:
		it := val.(int32)
		str = strconv.Itoa(int(it))
	case uint32:
		it := val.(uint32)
		str = strconv.Itoa(int(it))
	case int64:
		it := val.(int64)
		str = strconv.FormatInt(it, 10)
	case uint64:
		it := val.(uint64)
		str = strconv.FormatUint(it, 10)
	case string:
		str = val.(string)
	case []byte:
		str = string(val.([]byte))
	default:
		newValue, _ := json.Marshal(val)
		str = string(newValue)
	}

	return str
}
