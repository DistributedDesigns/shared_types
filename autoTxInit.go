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
	AutoTxKey AutoTxKey
	Amount    currency.Currency
	Trigger   currency.Currency
	WorkerID  int
}

// ToCSV : Converts the QuoteRequest to a csv
func (aTx *AutoTxInit) ToCSV() string {
	parts := make([]string, 3)

	parts[0] = fmt.Sprintf("%0.2f", aTx.Amount.ToFloat())
	parts[1] = fmt.Sprintf("%0.2f", aTx.Trigger.ToFloat())
	parts[2] = fmt.Sprintf("%d", aTx.WorkerID)

	return fmt.Sprintf("%s,%s", aTx.AutoTxKey.ToCSV(), strings.Join(parts, ","))
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

	currAmount, err := currency.NewFromString(parts[3])
	if err != nil {
		return AutoTxInit{}, err
	}
	currTrigger, err := currency.NewFromString(parts[4])
	if err != nil {
		return AutoTxInit{}, err
	}
	workerNum, err := strconv.ParseInt(parts[5], 10, 64)
	if err != nil {
		return AutoTxInit{}, err
	}

	return AutoTxInit{
		AutoTxKey: AutoTxKey{
			Stock:  parts[0],
			UserID: parts[1],
			Action: parts[2],
		},
		Amount:   currAmount,
		Trigger:  currTrigger,
		WorkerID: int(workerNum),
	}, nil
}
