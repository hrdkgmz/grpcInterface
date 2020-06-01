package request

import (
	"github.com/hrdkgmz/grpcInterface/def"
)

//QueryChainCodeReq define
type QueryChainCodeReq struct {
	Fn        def.ReqType `json:"fn"`
	OrgName   string      `json:"org_name"`
	ChannelID string      `json:"channel_id"`
	OrgUser   string      `json:"org_user"`
	CCID      string      `json:"cc_id"`
	Fcn       string      `json:"fcn"`
	Params    []string    `json:"params"`
	Peers     []string    `json:"peers"`
}

func NewQueryChainCodeReq(orgName string, channelID string, orgUser string, ccID string, fcn string, params []string, peers []string) *QueryChainCodeReq {
	r := new(QueryChainCodeReq)
	r.Fn = def.QueryChainCode
	r.OrgName = orgName
	r.ChannelID = channelID
	r.OrgUser = orgUser
	r.CCID = ccID
	r.Fcn = fcn
	r.Params = params
	r.Peers = peers

	return r

}

func NewBlankQueryChainCodeReq() *QueryChainCodeReq {
	r := new(QueryChainCodeReq)
	r.Fn = def.QueryChainCode

	return r
}

func (r *QueryChainCodeReq) GetFn() def.ReqType {
	return r.Fn
}

func (r *QueryChainCodeReq) GetOrgName() string {
	return r.OrgName
}

func (r *QueryChainCodeReq) GetChannelID() string {
	return r.ChannelID
}

func (r *QueryChainCodeReq) GetOrgUser() string {
	return r.OrgUser
}

func (r *QueryChainCodeReq) GetCCID() string {
	return r.CCID
}

func (r *QueryChainCodeReq) GeFcn() string {
	return r.Fcn
}

func (r *QueryChainCodeReq) GetParams() []string {
	return r.Params
}

func (r *QueryChainCodeReq) GetPeers() []string {
	return r.Peers
}

func (r *QueryChainCodeReq) SetFn(args def.ReqType) *QueryChainCodeReq {
	r.Fn = args
	return r
}

func (r *QueryChainCodeReq) SetOrgName(args string) *QueryChainCodeReq {
	r.OrgName = args
	return r
}

func (r *QueryChainCodeReq) SetChannelID(args string) *QueryChainCodeReq {
	r.ChannelID = args
	return r
}

func (r *QueryChainCodeReq) SetOrgUser(args string) *QueryChainCodeReq {
	r.OrgUser = args
	return r
}

func (r *QueryChainCodeReq) SetCCID(args string) *QueryChainCodeReq {
	r.CCID = args
	return r
}

func (r *QueryChainCodeReq) SetFcn(args string) *QueryChainCodeReq {
	r.Fcn = args
	return r
}

func (r *QueryChainCodeReq) SetParams(args []string) *QueryChainCodeReq {
	r.Params = args
	return r
}

func (r *QueryChainCodeReq) SetPeers(args []string) *QueryChainCodeReq {
	r.Peers = args
	return r
}

//GetFuncName define
func (r *QueryChainCodeReq) GetFuncName() string {
	return r.Fn.ToString()
}
