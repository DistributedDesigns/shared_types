package types

import (
	"errors"
	"fmt"
	"strings"
)

// AutoTxKey : Key for accessing aTx in the aTx store.
type AutoTxKey struct {
	Stock  string
	UserID string
	Action string
}

// ToCSV : Converts the QuoteRequest to a csv
func (aTxKey *AutoTxKey) ToCSV() string {
	return fmt.Sprintf("%s,%s,%s", aTxKey.Stock, aTxKey.UserID, aTxKey.Action)
}

// ParseAutoTxKey : Parses CSV as AutoTxKey
func ParseAutoTxKey(csv string) (AutoTxKey, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 3 {
		return AutoTxKey{}, errors.New("Expected number of values in AutoTxCancel csv")
	}

	return AutoTxKey{
		Stock:  parts[0],
		UserID: parts[1],
		Action: parts[2],
	}, nil
}
