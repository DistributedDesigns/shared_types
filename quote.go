package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/distributeddesigns/currency"
)

// Quote : Response from the quoteserver
type Quote struct {
	Price     currency.Currency
	Stock     string
	UserID    string
	Timestamp time.Time
	Cryptokey string
}

// ToCSV : Serialize the Quote as a CSV
func (q *Quote) ToCSV() string {
	parts := make([]string, 5)

	parts[0] = fmt.Sprintf("%.02f", q.Price.ToFloat())
	parts[1] = q.Stock
	parts[2] = q.UserID
	parts[3] = fmt.Sprintf("%d", q.Timestamp.UnixNano()/1000)
	parts[4] = q.Cryptokey

	return strings.Join(parts, ",")
}

// ParseQuote : Attempt to parse the CSV as a Quote
func ParseQuote(csv string) (Quote, error) {
	// remove messiness
	csv = strings.TrimSpace(csv)

	parts := strings.Split(csv, ",")

	if len(parts) != 5 {
		return Quote{}, errors.New("Expected 5 values in Quote csv")
	}

	price, err := currency.NewFromString(parts[0])
	if err != nil {
		return Quote{}, err
	}

	// Unix time has to be converted string -> int -> Time
	milliSec, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return Quote{}, err
	}

	sec := milliSec / 1e6
	nano := milliSec % 1e6 * 1000

	quote := Quote{
		Price:     price,
		Stock:     parts[1],
		UserID:    parts[2],
		Timestamp: time.Unix(sec, nano),
		Cryptokey: parts[4],
	}

	return quote, nil
}
