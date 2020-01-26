package schema

// BidResponse holds basic details about the bid
type BidResponse struct {
	BidderID string  `json:"bidder_id"`
	Amount   float64 `json:"amount"`
}
