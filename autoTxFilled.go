package types

import (
	"errors"
	"fmt"
	"strings"

	"strconv"

	currency "github.com/distributeddesigns/currency"
)

// AutoTxFilled : Request for auto transaction init
type AutoTxFilled struct {
	AutoTxKey AutoTxKey
	AddFunds  currency.Currency
	AddStocks uint
}

// ToCSV : Converts the QuoteRequest to a csv
func (aTx *AutoTxFilled) ToCSV() string {
	parts := make([]string, 2)

	parts[0] = fmt.Sprintf("%0.2f", aTx.AddFunds.ToFloat())
	parts[1] = fmt.Sprintf("%d", aTx.AddStocks)

	return fmt.Sprintf("%s,%s", aTx.AutoTxKey.ToCSV(), strings.Join(parts, ","))
}

// ParseAutoTxFilled : Parses CSV as QuoteRequest
func ParseAutoTxFilled(csv string) (AutoTxFilled, error) {
	parts := strings.Split(csv, ",")

	if len(parts) != 5 {
		return AutoTxFilled{}, errors.New("Expected number of values in AutoTxFilled csv")
	}

	currAmount, err := currency.NewFromString(parts[3])
	if err != nil {
		return AutoTxFilled{}, err
	}
	numStocks, err := strconv.ParseUint(parts[4], 10, 64)
	if err != nil {
		return AutoTxFilled{}, err
	}

	return AutoTxFilled{
		AutoTxKey: AutoTxKey{
			Stock:  parts[0],
			UserID: parts[1],
			Action: parts[2],
		},
		AddFunds:  currAmount,
		AddStocks: uint(numStocks),
	}, nil
}
