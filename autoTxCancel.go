package types

import (
	"errors"
	"fmt"
	"strings"

	"strconv"
)

// AutoTxCancel : Request for auto transaction init
type AutoTxCancel struct {
	Stock    string
	UserID   string
	WorkerID uint64
}

// ToCSV : Converts the QuoteRequest to a csv
func (aTx *AutoTxCancel) ToCSV() string {
	return fmt.Sprintf("%s,%s,%d", aTx.Stock, aTx.UserID, aTx.WorkerID)
}

// ParseAutoTxCancel : Parses CSV as QuoteRequest
func ParseAutoTxCancel(csv string) (AutoTxCancel, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 3 {
		return AutoTxCancel{}, errors.New("Expected number of values in AutoTxCancel csv")
	}

	workerNum, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return AutoTxCancel{}, err
	}

	return AutoTxCancel{
		Stock:    parts[0],
		UserID:   parts[1],
		WorkerID: workerNum,
	}, nil
}
