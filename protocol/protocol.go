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
	Height       int `json:"height"`
	SlotIndex    int `json:"slot"`
	BlocksInSlot int `json:"blocksinslot"`
}

// LiquidInfo ...
type LiquidInfo struct {
	Long  map[string]int64 `json:"LongBenefi"`
	Short map[string]int64 `json:"ShortBenefi"`
}

// Participate ...
type Participate struct {
	PoolEntrySet []MemPoolEntryWithHash `json:"pooltxs"`
}

// MemPoolEntryWithHash ...
type MemPoolEntryWithHash struct {
	Hash    string           `json:"txid"`
	Account map[string]int64 `json:"account"`
}
