package def

type ResType int

const (
	Fail ResType = -1 + iota
	VariousResults
	Success
	SuccessWLoad
)

func (r ResType) ToString() string {
	switch r {
	case Fail:
		return "FAIL"
	case VariousResults:
		return "VARIOUSE RESULTS"
	case Success:
		return "SUCCESS"
	case SuccessWLoad:
		return "SUCCESS WITH PAYLOAD"
	default:
		return "UNKOWN"

	}
}
