package respond

import (
	"bidding/pkg/errors"
	"bidding/pkg/trace"

	"compress/gzip"
	"encoding/json"
	"net/http"
)

// Response holds the handlerfunc response
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta Meta        `json:"meta"`
}

// PageResponse holds the paginated handlerfunc response
type PageResponse struct {
	Data interface{} `json:"data"`
	Meta MetaPage    `json:"meta"`
}

// Meta holds the status of the request informations
type Meta struct {
	Status  int    `json:"status_code"`
	Message string `json:"error_message,omitempty"`
}

// MetaPage holds the paginated data inforamtions
type MetaPage struct {
	Meta
	Total    int    `json:"total,omitempty"`
	Count    int    `json:"count,omitempty"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
}

// Page holds the paginate informations
type Page struct {
	Offset int `schema:"offset" url:"offset"`
	Limit  int `schema:"limit" url:"limit"`
}

// Format customize the http response
func Format(w http.ResponseWriter, status int, data interface{}) {
	var res Response
	res.Data = data
	res.Meta = Meta{Status: status}
	With(w, status, res)
}

// OK send the 200 http response
func OK(w http.ResponseWriter, data interface{}) {
	var res Response
	res.Data = data
	res.Meta = Meta{Status: http.StatusOK}
	With(w, http.StatusOK, res)
}

// Created send the 200 http response
func Created(w http.ResponseWriter, data interface{}) {
	var res Response
	res.Data = data
	res.Meta = Meta{Status: http.StatusCreated}
	With(w, http.StatusCreated, res)
}

// Fail write the error response
func Fail(w http.ResponseWriter, e *errors.AppError) {
	var res Response
	res.Meta = Meta{Status: e.Status, Message: e.Message}
	With(w, e.Status, res)
}

// With sets the response headers, and write response to client
func With(w http.ResponseWriter, status int, data interface{}) {
	gz := gzip.NewWriter(w)
	defer gz.Close()
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(status)
	if status != http.StatusNoContent {
		if _, err := gz.Write(buf); err != nil {
			trace.Log.Error("respond.With.error: ", err)
		}
	}
}
