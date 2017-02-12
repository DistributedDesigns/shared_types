package types

import (
	"errors"
	"strings"
)

// QuoteRequest : Request for a quote value
type QuoteRequest struct {
	Stock  string
	UserID string
}

// ToCSV : Converts the QuoteRequest to a csv
func (qr *QuoteRequest) ToCSV() string {
	return qr.Stock + "," + qr.UserID
}

// ParseQuoteRequest : Parses CSV as QuoteRequest
func ParseQuoteRequest(csv string) (QuoteRequest, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 2 {
		return QuoteRequest{}, errors.New("Expected 2 values in QuoteRequest csv")
	}

	return QuoteRequest{Stock: parts[0], UserID: parts[1]}, nil
}
