package v1

import (
	"bidding/api/v1/auctioneer"
	"bidding/api/v1/bidder"

	"github.com/go-chi/chi"
)

// Init initializes all the v1 routes
func Init(r chi.Router) {
	r.Route("/auction", auctioneer.Init)
	r.Route("/bidder", bidder.Init)
}
