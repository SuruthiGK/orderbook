package OrderBookSnapStore

// DepthService show depth info - Request params - URL encoded
type DepthService struct {
	symbol string
	limit  *int
}


// DepthResponse define depth info with bids and asks
type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
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