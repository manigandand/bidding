package bidder

import (
	"bidding/pkg/errors"
	"bidding/pkg/respond"
	"bidding/schema"
	"bidding/utils"
	"fmt"
	"net/http"
)

func getBiddersHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	respond.OK(w, store.Bidder().List())
	return nil
}

func bidderRegisterHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var req schema.BidderReq

	if err := utils.Decode(r, &req); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}

	// NOTE: considering the bidder request come from localhost
	bidder := &schema.Bidder{
		ID:    fmt.Sprintf("bidder_%d", store.Bidder().Count()+1),
		Name:  req.Name,
		Host:  fmt.Sprintf("http://%s", r.Host),
		Delay: req.Delay,
	}
	store.Bidder().Add(bidder)

	respond.Created(w, bidder)
	return nil
}
