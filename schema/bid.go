package schema

// Bid holds basic details about the bid
type Bid struct {
	AuctionID string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}
