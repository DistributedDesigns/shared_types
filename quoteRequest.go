package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// QuoteRequest : Request for a quote value
type QuoteRequest struct {
	Stock      string
	UserID     string
	AllowCache bool
	ID         uint64
}

// ToCSV : Converts the QuoteRequest to a csv
func (qr *QuoteRequest) ToCSV() string {
	return fmt.Sprintf("%s,%s,%t,%d", qr.Stock, qr.UserID, qr.AllowCache, qr.ID)
}

// ParseQuoteRequest : Parses CSV as QuoteRequest
func ParseQuoteRequest(csv string) (QuoteRequest, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 4 {
		return QuoteRequest{}, errors.New("Expected 3 values in QuoteRequest csv")
	}

	allowCache, err := strconv.ParseBool(parts[2])
	if err != nil {
		return QuoteRequest{}, err
	}

	id, err := strconv.ParseUint(parts[3], 10, 64)
	if err != nil {
		return QuoteRequest{}, err
	}

	return QuoteRequest{
		Stock:      parts[0],
		UserID:     parts[1],
		AllowCache: allowCache,
		ID:         id,
	}, nil
}
