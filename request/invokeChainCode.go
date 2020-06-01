package request

import (
	"github.com/hrdkgmz/grpcInterface/def"
)

//InvokeCCReq define
type InvokeChainCodeReq struct {
	Fn        def.ReqType `json:"fn"`
	OrgName   string      `json:"org_name"`
	ChannelID string      `json:"channel_id"`
	OrgUser   string      `json:"org_user"`
	CCID      string      `json:"cc_id"`
	Fcn       string      `json:"fcn"`
	Params    []string    `json:"params"`
	Peers     []string    `json:"peers"`
	SysUser   string      `json:"sys_user"`
}

func NewInvokeChainCodeReq(orgName string, channelID string, orgUser string, ccID string, fcn string, params []string, peers []string) *InvokeChainCodeReq {
	r := new(InvokeChainCodeReq)
	r.Fn = def.InvokeChainCode
	r.OrgName = orgName
	r.ChannelID = channelID
	r.OrgUser = orgUser
	r.CCID = ccID
	r.Fcn = fcn
	r.Params = params
	r.Peers = peers

	return r

}

func NewBlankInvokeChainCodeReq() *InvokeChainCodeReq {
	r := new(InvokeChainCodeReq)
	r.Fn = def.InvokeChainCode

	return r
}

func (r *InvokeChainCodeReq) GetFn() def.ReqType {
	return r.Fn
}

func (r *InvokeChainCodeReq) GetOrgName() string {
	return r.OrgName
}

func (r *InvokeChainCodeReq) GetChannelID() string {
	return r.ChannelID
}

func (r *InvokeChainCodeReq) GetOrgUser() string {
	return r.OrgUser
}

func (r *InvokeChainCodeReq) GetCCID() string {
	return r.CCID
}

func (r *InvokeChainCodeReq) GeFcn() string {
	return r.Fcn
}

func (r *InvokeChainCodeReq) GetParams() []string {
	return r.Params
}

func (r *InvokeChainCodeReq) GetPeers() []string {
	return r.Peers
}

func (r *InvokeChainCodeReq) SetFn(args def.ReqType) *InvokeChainCodeReq {
	r.Fn = args
	return r
}

func (r *InvokeChainCodeReq) SetOrgName(args string) *InvokeChainCodeReq {
	r.OrgName = args
	return r
}

func (r *InvokeChainCodeReq) SetChannelID(args string) *InvokeChainCodeReq {
	r.ChannelID = args
	return r
}

func (r *InvokeChainCodeReq) SetOrgUser(args string) *InvokeChainCodeReq {
	r.OrgUser = args
	return r
}

func (r *InvokeChainCodeReq) SetCCID(args string) *InvokeChainCodeReq {
	r.CCID = args
	return r
}

func (r *InvokeChainCodeReq) SetFcn(args string) *InvokeChainCodeReq {
	r.Fcn = args
	return r
}

func (r *InvokeChainCodeReq) SetParams(args []string) *InvokeChainCodeReq {
	r.Params = args
	return r
}

func (r *InvokeChainCodeReq) SetPeers(args []string) *InvokeChainCodeReq {
	r.Peers = args
	return r
}

//GetFuncName define
func (r *InvokeChainCodeReq) GetFuncName() string {
	return r.Fn.ToString()
}
