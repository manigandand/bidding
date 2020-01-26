package schema

import (
	"bidding/pkg/errors"
	"strings"
	"time"
)

// Bidder holds the basic details about the bidder
type Bidder struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Host  string        `json:"host"`
	Delay time.Duration `json:"delay"`
}

// BidderReq holds the basic details about the bidder registration req
type BidderReq struct {
	Name  string        `json:"name"`
	Delay time.Duration `json:"delay"`
}

// Ok validates the request
func (b *BidderReq) Ok() error {
	switch {
	case strings.TrimSpace(b.Name) == "":
		return errors.IsRequiredErr("name")
	}

	return nil
}
