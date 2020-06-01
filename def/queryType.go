package def

type QueryType int

const (
	QueryChannel QueryType = 1 + iota
	QueryInstalledCC
	QueryInstantiatedCC
	QueryCfgFromOrderer
	QueryBlockByTxID
	QueryBlockByHash
	QueryBlockByNum
	QueryBlockTx
	QueryConfig
	QueryCfgBlock
	QueryBlockInfo
)

func (r QueryType) ToString() string {
	switch r {
	case QueryChannel:
		return "QueryChannel"
	case QueryInstalledCC:
		return "QueryInstalledCC"
	case QueryInstantiatedCC:
		return "QueryInstantiatedCC"
	case QueryCfgFromOrderer:
		return "QueryCfgFromOrderer"
	case QueryBlockByTxID:
		return "QueryBlockByTxID"
	case QueryBlockByHash:
		return "QueryBlockByHash"
	case QueryBlockByNum:
		return "QueryBlockByNum"
	case QueryBlockTx:
		return "QueryBlockTx"
	case QueryConfig:
		return "QueryConfig"
	case QueryCfgBlock:
		return "QueryCfgBlock"
	case QueryBlockInfo:
		return "QueryBlockInfo"
	default:
		return "UNKOWN QUERY REQUEST"
	}
}
