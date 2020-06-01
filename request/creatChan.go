package request

//CreateChanReq define
type CreateChanReq struct {
	TaskID    int64  `json:"task_id"`
	SysUser   string `json:"sys_user"`
	ChannelID string `json:"channel_id"`
}

//JoinChanReq define
type JoinChanReq struct {
	TaskID    int64    `json:"task_id"`
	SysUser   string   `json:"sys_user"`
	ChannelID string   `json:"channel_id"`
	WithPeer  []string `json:"with_peer"`
}

//InsCCReq define
type InsCCReq struct {
	TaskID   int64    `json:"task_id"`
	SysUser  string   `json:"sys_user"`
	CCID     string   `json:"cc_id"`
	WithPeer []string `json:"with_peer"`
}

//InstantCCReq define
type InstantCCReq struct {
	TaskID    int64    `json:"task_id"`
	SysUser   string   `json:"sys_user"`
	CCID      string   `json:"cc_id"`
	WithPeer  []string `json:"with_peer"`
	Policy    string   `json:"policy"`
	InitArgs  []string `json:"init_args"`
	ChannelID string   `json:"channel_id"`
}

//UpgrdCCReq define
type UpgrdCCReq struct {
	TaskID    int64    `json:"task_id"`
	SysUser   string   `json:"sys_user"`
	ChannelID string   `json:"channel_id"`
	CCID      string   `json:"cc_id"`
	InitArgs  []string `json:"init_args"`
	WithPeer  []string `json:"with_peer"`
	Policy    string   `json:"policy"`
}

//EnrollAdminReq define
type EnrollAdminReq struct {
	TaskID  int64  `json:"task_id"`
	SysUser string `json:"sys_user"`
	Secret  string  `json:"secret"`
}

//QueryCCReq define
type QueryCCReq struct {
	TaskID    int64    `json:"task_id"`
	SysUser   string   `json:"sys_user"`
	ChannelID string   `json:"channel_id"`
	CCID      string   `json:"cc_id"`
	Fcn       string   `json:"f_cn"`
	Param     []string `json:"param"`
	WithPeer  []string `json:"with_peer"`
}

//InvokeCCReq define
type InvokeCCReq struct {
	TaskID    int64    `json:"task_id"`
	SysUser   string   `json:"sys_user"`
	ChannelID string   `json:"channel_id"`
	CCID      string   `json:"cc_id"`
	Fcn       string   `json:"f_cn"`
	Param     []string `json:"param"`
	WithPeer  []string `json:"with_peer"`
}

//RegBlkEventReq define
type RegBlkEventReq struct {
	TaskID    int64  `json:"task_id"`
	SysUser   string `json:"sys_user"`
	ChannelID string `json:"channel_id"`
}

//RegFltBlkEventReq define
type RegFltBlkEventReq struct {
	TaskID    int64  `json:"task_id"`
	SysUser   string `json:"sys_user"`
	ChannelID string `json:"channel_id"`
}

//RegCCEventReq define
type RegCCEventReq struct {
	TaskID      int64  `json:"task_id"`
	SysUser     string `json:"sys_user"`
	ChannelID   string `json:"channel_id"`
	CCID        string `json:"cc_id"`
	EventFilter string `json:"event_filter"`
}

//RegTxEventReq define
type RegTxEventReq struct {
	TaskID    int64  `json:"task_id"`
	SysUser   string `json:"sys_user"`
	ChannelID string `json:"channel_id"`
	TxID      string `json:"tx_id"`
}
