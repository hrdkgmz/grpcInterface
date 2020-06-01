package request

import (
	"time"

	"github.com/hrdkgmz/grpcInterface/def"
	"github.com/hrdkgmz/grpcInterface/util"
	"fmt"
)

//Response define
type Response struct {
	TaskID   int64         `json:"task_id"`
	Status   def.ResType   `json:"status"`
	Type     def.ReqType   `json:"type"`
	DataTime time.Time     `json:"datetime"`
	Payload  []interface{} `json:"payload"`
	ErrMsg   string        `json:"err_msg"`
}

func NewResponse(status def.ResType, rtype def.ReqType, payload []interface{}, errMsg string) *Response {
	res := new(Response)
	res.Status = status
	res.Type = rtype
	res.Payload = payload
	res.ErrMsg = errMsg

	return res
}

func NewBlankResponse() *Response {
	res := &Response{}
	res.Payload = make([]interface{}, 0)
	return res
}

func (r *Response) GetStatus() def.ResType {
	return r.Status
}

func (r *Response) GetType() def.ReqType {
	return r.Type
}

func (r *Response) GetPayload() []interface{} {
	return r.Payload
}

func (r *Response) GetErrMsg() string {
	return r.ErrMsg
}

func (r *Response) SetStatus(args def.ResType) *Response {
	r.Status = args
	return r
}

func (r *Response) SetType(args def.ReqType) *Response {
	r.Type = args
	return r
}
func (r *Response) SetPayload(args []interface{}) *Response {
	r.Payload = args
	return r
}

func (r *Response) SetErrMsg(args string) *Response {
	r.ErrMsg = args
	return r
}
func (r *Response) AppendToPayload(val interface{}) *Response {
	if r.Payload == nil {
		r.Payload = make([]interface{}, 0)
	}
	r.Payload = append(r.Payload, val)
	return r
}

func (r *Response) ToJSONString() string {
	str, err := util.EncodeJSON(r)
	if err != nil {
		fmt.Println("Response JSON序列化失败", err)
		return ""
	}
	return str
}
