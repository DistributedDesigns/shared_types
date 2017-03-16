package types

import (
	"errors"
	"fmt"
	"strings"

	"strconv"

	currency "github.com/distributeddesigns/currency"
)

// AutoTxInit : Request for auto transaction init
type AutoTxInit struct {
	Amount   currency.Currency
	Trigger  currency.Currency
	Stock    string
	UserID   string
	workerID uint64
}

// ToCSV : Converts the QuoteRequest to a csv
func (aTx *AutoTxInit) ToCSV() string {
	return fmt.Sprintf("%s,%s,%s,%s,%d", aTx.Amount.String(), aTx.Trigger.String(), aTx.Stock, aTx.UserID, aTx.workerID)
}

// ParseQuoteRequest : Parses CSV as QuoteRequest
func ParseAutoTxInit(csv string) (AutoTxInit, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 5 {
		return AutoTxInit{}, errors.New("Expected number of values in AutoTxInit csv")
	}

	// allowCache, err := strconv.ParseBool(parts[2])
	// if err != nil {
	// 	return QuoteRequest{}, err
	// }

	// id, err := strconv.ParseUint(parts[3], 10, 64)
	// if err != nil {
	// 	return QuoteRequest{}, err
	// }

	currAmount, err := currency.NewFromString(parts[0])
	if err != nil {
		return AutoTxInit{}, err
	}
	currTrigger, err := currency.NewFromString(parts[1])
	if err != nil {
		return AutoTxInit{}, err
	}
	workerNum, err := strconv.ParseUint(parts[4], 10, 64)
	if err != nil {
		return AutoTxInit{}, err
	}

	return AutoTxInit{
		Amount:   currAmount,
		Trigger:  currTrigger,
		Stock:    parts[2],
		UserID:   parts[3],
		workerID: workerNum,
	}, nil
}
