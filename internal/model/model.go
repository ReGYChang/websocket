package model

type Message struct {
	EventType             string `json:"e"` // EventType represents the update type
	Time                  uint64 `json:"E"` // Time represents the event time
	Symbol                string `json:"s"` // Symbol represents the symbol related to the update
	TradeID               uint64 `json:"a"` // TradeID is the aggregated trade ID
	Price                 string `json:"p"` // Price is the trade price
	Quantity              string `json:"q"` // Quantity is the trade quantity
	FirstBreakDownTradeID uint64 `json:"f"` // FirstBreakDownTradeID is the first breakdown trade ID
	LastBreakDownTradeID  uint64 `json:"l"` // LastBreakDownTradeID is the last breakdown trade ID
	TradeTime             uint64 `json:"T"` // Time is the trade time
	Maker                 bool   `json:"m"` // Maker indicates whether buyer is a maker
}

type StreamMsg struct {
	Stream string  `json:"stream"`
	Data   Message `json:"data"`
}
