package DataStore

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)


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