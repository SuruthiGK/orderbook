package OrderBooksStore

// WsDepthEvent define websocket depth event
type WsDepthEvent struct {
	Event         string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	UpdateID      int64  `json:"u"`
	FirstUpdateID int64  `json:"U"`
	Bids          []Bid  `json:"b"`
	Asks          []Ask  `json:"a"`
}

// Bid define bid info with price and quantity
type Bid struct {
	Price    string
	Quantity string
}

// Ask define ask info with price and quantity
type Ask struct {
	Price    string
	Quantity string
}
