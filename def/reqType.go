package def

type ReqType int

const (
	CreateChannel ReqType = 1 + iota
	JoinChannel
	InstallCC
	InstantiateCC
	InvokeChainCode
	QueryChainCode
	UpgradeCC
	QueryRC
	QueryLC
	RegisterBlockEvent
	RegisterFilteredBlockEvent
	RegisterChainCodeEvent
	RegisterTxStatusEvent
	CreateIdentity
	ModifyIdentity
	RemoveIdentity
	Enroll
	Reenroll
	Register
	Revoke
	AddAffiliation
	ModifyAffiliation
	GetAffiInfo
	GetIdenInfo
	RemoveAffiliation
	CreateSdkClient
	MultipleRequests
)

func (r ReqType) ToString() string {
	switch r {
	case CreateChannel:
		return "CreateChannel"
	case JoinChannel:
		return "JoinChannel"
	case InstallCC:
		return "InstallCC"
	case InstantiateCC:
		return "InstantiateCC"
	case InvokeChainCode:
		return "InvokeChainCode"
	case QueryChainCode:
		return "QueryChainCode"
	case UpgradeCC:
		return "UpgradeCC"
	case QueryRC:
		return "QueryRC"
	case QueryLC:
		return "QueryLC"
	case RegisterBlockEvent:
		return "RegisterBlockEvent"
	case RegisterFilteredBlockEvent:
		return "RegisterFilteredBlockEvent"
	case RegisterTxStatusEvent:
		return "RegisterTxStatusEvent"
		// ===MSP===//
	case CreateIdentity:
		return "CreateIdentity"
	case ModifyIdentity:
		return "ModifyIdentity"
	case RemoveIdentity:
		return "RemoveIdentity"
	case Enroll:
		return "Enroll"
	case Reenroll:
		return "Reenroll"
	case Register:
		return "Register"
	case Revoke:
		return "Revoke"
	case MultipleRequests:
		return "MultipleRequests"
	case AddAffiliation:
		return "AddAffiliation"
	case ModifyAffiliation:
		return "ModifyAffiliation"
	case GetAffiInfo:
		return "GetAffiInfo"
	case GetIdenInfo:
		return "GetIdenInfo"
	case RemoveAffiliation:
		return "RemoveAffiliation"
	case CreateSdkClient:
		return "CreateSdkClient"
	default:
		return "UNKOWN REQUEST"

	}
}
