package store

import "bidding/schema"

var bidders = make([]*schema.Bidder, 0)

// BidderStore implements the Store interface
type BidderStore struct {
	*Conn
}

// NewBidderStore returns new store object
func NewBidderStore(st *Conn) *BidderStore {
	return &BidderStore{st}
}

// Add register the new bidder into the list
func (b *BidderStore) Add(bidder *schema.Bidder) {
	bidders = append(bidders, bidder)
}

// List returns all the regisytered bidders
func (b *BidderStore) List() []*schema.Bidder {
	return bidders
}

// Count returns the count of regisytered bidders
func (b *BidderStore) Count() int {
	return len(bidders)
}
