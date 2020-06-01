package def

type TaskType int

const (
	InvokeTask ReqType = 1 + iota
	QueryTask
)