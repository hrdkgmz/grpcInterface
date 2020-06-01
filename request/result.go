package request

import (
	"fmt"
	"github.com/hrdkgmz/grpcInterface/util"
)

type Result struct {
	Code  string      `json:"code"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func NewReslut(code string, name string, value interface{}) *Result {
	res := &Result{code, name, value}
	return res
}

func NewBlankReslut() *Result {
	res := &Result{}
	return res
}

func (r *Result) GetCode() string {
	return r.Code
}

func (r *Result) GetName() string {
	return r.Name
}

func (r *Result) GetValue() interface{} {
	return r.Value
}

func (r *Result) SetCode(args string) *Result {
	r.Code = args
	return r
}

func (r *Result) SetName(args string) *Result {
	r.Name = args
	return r
}
func (r *Result) SetValue(args interface{}) *Result {
	r.Value = args
	return r
}

func (r *Result) ToJSONString() string {
	str, err := util.EncodeJSON(r)
	if err != nil {
		fmt.Println("Result JSON序列化失败", err)
		return ""
	}
	return str
}
