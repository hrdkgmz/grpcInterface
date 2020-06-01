package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hrdkgmz/grpcInterface/def"
)

//FillStruct Define
func FillStruct(data map[string]interface{}, obj interface{}) error {
	for k, v := range data {
		err := setField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("目标结构体未定义此域: %s", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("无法为域: %s设置值", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)

	var err error

	if value != nil {

		if structFieldType != val.Type() {
			typeName := structFieldValue.Type().Name()
			var subType string
			if typeName == "" {
				typeName = structFieldValue.Type().Kind().String()
				if typeName != "array" && typeName != "chan" && typeName != "map" && typeName != "ptr" && typeName != "slice" {
					return fmt.Errorf("结构体转换异常，域：" + name + " 的kind类型异常：typeName")
				}
				subType = structFieldValue.Type().Elem().Name()
			}
			val, err = typeConversion(fmt.Sprintf("%v", value), typeName, subType)
			if err != nil {
				return err
			}
		}

		structFieldValue.Set(val)
	}
	return nil
}

func typeConversion(value string, ntype string, subType string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2019-11-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2019-11-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "ReqType" {
		i, err := strconv.Atoi(value)
		def := (def.ReqType)(i)
		return reflect.ValueOf(def), err
	} else if ntype == "slice" {
		start := strings.Index(value, "[") + 1
		end := strings.Index(value, "]")
		value = value[start:end]
		vals := strings.Split(value, " ")
		if subType == "string" {
			sli := make([]string, 0)
			for _, val := range vals {
				sli = append(sli, val)
			}
			return reflect.ValueOf(sli), nil
		}
	}

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
