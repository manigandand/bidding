package auctioneer

import (
	"bidding/api"
	"net/http"

	"github.com/go-chi/chi"
)

// Init initializes all the v1 routes
func Init(r chi.Router) {
	r.Method(http.MethodPost, "/", api.Handler(auctionHandler))
}
