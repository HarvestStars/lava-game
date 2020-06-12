package protocol

// SlotInfo 存在redis中的全量信息
type SlotInfo struct {
	Total     float64 `json:"total"`
	LongInfo  account `json:"long"`
	ShortInfo account `json:"short"`
}

type account struct {
	Addr   string `json:"address"`
	Amount int64  `json:"amount"`
}

// ChainInfo ...
type ChainInfo struct {
	Height    int32 `json:"height"`
	SlotIndex int32 `json:"slot"`
}
