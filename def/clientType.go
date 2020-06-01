package def

type ClientType int

const (
	ResmgmtClient ClientType = 1 + iota
	ChannelClient
	LedgerClient
	EventClient
	// MspClient 一个组织一个MSPClient
	AllClient
)

func (r ClientType) ToString() string {
	switch r {
	case ResmgmtClient:
		return "ResmgmtClient"
	case ChannelClient:
		return "ChannelClient"
	case LedgerClient:
		return "LedgerClient"
	case EventClient:
		return "EventClient"
	// case MspClient:
	// 	return "MspClient"
	case AllClient:
		return "AllClient"
	default:
		return "UNKOWN CLIENT"
	}
}
