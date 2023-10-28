package respond

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// JSON maps go type to a JSON struct
type JSON map[string]interface{}

// Response struct contains all the fields needed to respond
// to a particular request
type Response struct {
	StatusCode int
	Data       interface{}
	Headers    map[string]string
}

func OK(w http.ResponseWriter, data, meta interface{}) error {
	return SendResponse(w, http.StatusOK, WrapPayload(data, meta), nil)
}

// WrapPayload is used to create a generic payload for the data
// and the metadata passed
func WrapPayload(data, meta interface{}) JSON {
	x := make(JSON)
	if data != nil {
		x["data"] = data
	}

	if meta != nil {
		x["meta"] = meta
	}

	return x
}

// Fail write the error response
// Common func to send all the error response
func Fail(w http.ResponseWriter, e error) {
	log.Error().Err(e).Msgf("Failed to process request: %v", e)
	SendResponse(w, http.StatusBadRequest, WrapPayload(nil, e.Error()), nil)
}

// SendResponse is a helper function which sends a response with the passed data
func SendResponse(w http.ResponseWriter, statusCode int, data interface{}, headers map[string]string) error {
	return NewResponse(statusCode, data, headers).Send(w)
}

// NewResponse returns a new response object.
func NewResponse(statusCode int, data interface{}, headers map[string]string) *Response {
	return &Response{
		StatusCode: statusCode,
		Data:       data,
		Headers:    headers,
	}
}

// Send sends data encoded to JSON
func (res *Response) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	if res.Headers != nil {
		for key, value := range res.Headers {
			w.Header().Set(key, value)
		}
	}
	w.WriteHeader(res.StatusCode)

	if res.StatusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(res.Data); err != nil {
			log.Warn().Err(err).Msg("respond.send.error: ")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}
