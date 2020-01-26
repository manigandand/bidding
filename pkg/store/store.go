package store

import (
	"bidding/schema"
)

// Store global store interface
type Store interface {
	Bidder() Bidders
}

// Bidders store interface
type Bidders interface {
	Add(bidder *schema.Bidder)
	List() []*schema.Bidder
	Count() int
}

// Conn struct holds the store connection
type Conn struct {
	BidderConn Bidders
}

// NewStore inits new store connection
func NewStore() *Conn {
	conn := new(Conn)
	conn.BidderConn = NewBidderStore(conn)

	return conn
}

// Bidder implements the store interface
func (s *Conn) Bidder() Bidders {
	return s.BidderConn
}
