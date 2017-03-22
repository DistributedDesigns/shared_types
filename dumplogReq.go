package types

import (
	"errors"
	"strings"
)

// DumplogRequest Triggers a ledger dump for the user to the specified file.
// If UserID == "admin" then a dump is generated with all ledger entries.
type DumplogRequest struct {
	UserID   string
	Filename string
}

// ToCSV serializes the DumpLogRequest object as a CSV
func (dr *DumplogRequest) ToCSV() string {
	return dr.UserID + "," + dr.Filename
}

// ParseDumplogRequest deserializes a CSV into a DumplogRequest object
func ParseDumplogRequest(csv string) (DumplogRequest, error) {
	parts := strings.Split(csv, ",")
	if len(parts) != 2 {
		return DumplogRequest{}, errors.New("Expected 2 values in DumplogRequest csv")
	}

	return DumplogRequest{
		UserID:   parts[0],
		Filename: parts[1],
	}, nil
}
