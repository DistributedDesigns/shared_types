package types

import (
	"errors"
	"fmt"
	"strings"

	"strconv"

	currency "github.com/distributeddesigns/currency"
	"github.com/petar/GoLLRB/llrb"
)

// AutoTxInit : Request for auto transaction init
type AutoTxInit struct {
	Amount   currency.Currency
	Trigger  currency.Currency
	Action   string
	Stock    string
	UserID   string
	WorkerID int
}

// ToCSV : Converts the QuoteRequest to a csv
func (aTx *AutoTxInit) ToCSV() string {
	parts := make([]string, 6)

	parts[0] = fmt.Sprintf("%0.2f", aTx.Amount.ToFloat())
	parts[1] = fmt.Sprintf("%0.2f", aTx.Trigger.ToFloat())
	parts[2] = aTx.Action
	parts[3] = aTx.Stock
	parts[4] = aTx.UserID
	parts[5] = strconv.FormatUint(aTx.WorkerID, 10)

	return strings.Join(parts, ",")
}

// Less : Interface function for llrb
func (aTx AutoTxInit) less(than AutoTxInit) bool {
	return (aTx.Trigger.ToFloat() < than.Trigger.ToFloat())
}

// Less : Interface function for llrb
func (aTx AutoTxInit) Less(than llrb.Item) bool {
	return aTx.less(than.(AutoTxInit))
}

// ParseAutoTxInit : Parses CSV as QuoteRequest
func ParseAutoTxInit(csv string) (AutoTxInit, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 6 {
		return AutoTxInit{}, errors.New("Expected number of values in AutoTxInit csv")
	}

	currAmount, err := currency.NewFromString(parts[0])
	if err != nil {
		return AutoTxInit{}, err
	}
	currTrigger, err := currency.NewFromString(parts[1])
	if err != nil {
		return AutoTxInit{}, err
	}
	workerNum, err := strconv.ParseInt(parts[5], 10, 64)
	if err != nil {
		return AutoTxInit{}, err
	}

	return AutoTxInit{
		Amount:   currAmount,
		Trigger:  currTrigger,
		Action:   parts[2],
		Stock:    parts[3],
		UserID:   parts[4],
		WorkerID: int(workerNum),
	}, nil
}
