package auctioneer

import (
	"bidding/pkg/errors"
	"net/http"
)

// - gets all the bidders
// - start a bid section by making req to all the bidders with ad auction details
// - collects the response from the bidders, anounce the winner
func auctionHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	return nil
}
