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
	ID        uint64
}

// ToCSV : Serialize the Quote as a CSV
func (q *Quote) ToCSV() string {
	parts := make([]string, 6)

	parts[0] = fmt.Sprintf("%.02f", q.Price.ToFloat())
	parts[1] = q.Stock
	parts[2] = q.UserID
	parts[3] = fmt.Sprintf("%d", q.Timestamp.UnixNano()/1e6)
	parts[4] = q.Cryptokey
	parts[5] = strconv.FormatUint(q.ID, 10)

	return strings.Join(parts, ",")
}

// ParseQuote : Attempt to parse the CSV as a Quote
func ParseQuote(csv string) (Quote, error) {
	// remove messiness
	csv = strings.TrimSpace(csv)
	csv = strings.Replace(csv, "\n", "", -1)

	parts := strings.Split(csv, ",")

	if len(parts) != 6 {
		return Quote{}, errors.New("Expected 6 values in Quote csv")
	}

	price, err := currency.NewFromString(parts[0])
	if err != nil {
		return Quote{}, err
	}

	// Unix time has to be converted string -> int -> Time
	unixTimeMillisec, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return Quote{}, err
	}
	seconds := unixTimeMillisec / 1e3
	nano := (unixTimeMillisec % 1e3) * 1e6

	id, err := strconv.ParseUint(parts[5], 10, 64)
	if err != nil {
		return Quote{}, err
	}

	quote := Quote{
		Price:     price,
		Stock:     parts[1],
		UserID:    parts[2],
		Timestamp: time.Unix(seconds, nano),
		Cryptokey: parts[4],
		ID:        id,
	}

	return quote, nil
}
