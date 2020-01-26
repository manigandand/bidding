package auctioneer

import (
	"bidding/pkg/errors"
	"bidding/pkg/respond"
	"bidding/schema"
	"bidding/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"
)

type bidds []*schema.BidResponse

func (a bidds) Len() int           { return len(a) }
func (a bidds) Less(i, j int) bool { return a[i].Amount > a[j].Amount }
func (a bidds) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// - gets all the bidders
// - start a bid section by making req to all the bidders with ad auction details
// - collects the response from the bidders, anounce the winner
func auctionHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var (
		input schema.AuctionReq
		wg    sync.WaitGroup
	)
	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}
	bidders := store.Bidder().List()
	if len(bidders) == 0 {
		return errors.NotFound("no bidders available for this auction")
	}
	wg.Add(len(bidders))

	data := make(chan *schema.BidResponse, len(bidders))
	for _, b := range bidders {
		go collectBidResponse(input.AuctionID, b.Host, &wg, data)
	}

	var bidRes bidds
	for i := 0; i < len(bidders); i++ {
		if d := <-data; d != nil {
			bidRes = append(bidRes, d)
		}
	}

	wg.Wait()
	close(data)
	sort.Sort(bidRes)
	if len(bidRes) == 0 {
		return errors.BadRequest("bidders not responding with in time")
	}

	respond.OK(w, bidRes[0])
	return nil
}

func collectBidResponse(auctionID, host string, wg *sync.WaitGroup,
	data chan *schema.BidResponse) {
	var err error
	body := bytes.NewBuffer(nil)
	json.NewEncoder(body).Encode(map[string]interface{}{
		"auction_id": auctionID,
	})

	defer func() {
		wg.Done()
		if err != nil {
			fmt.Println("null data, ", err.Error())
			data <- nil
		}
	}()

	url := host + "/v1/bid"
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 190*time.Millisecond)
	defer cancel()
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var res struct {
		Data *schema.BidResponse `json:"data"`
		Meta respond.Meta        `json:"meta"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	data <- res.Data
}
