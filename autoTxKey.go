package types

import "fmt"

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
