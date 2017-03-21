package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// AuditEvent is a serialization object that can be easily stored
// in a database or ledger.
type AuditEvent struct {
	UserID    string
	ID        uint64
	EventType string
	Content   string
}

const auditEventLen = 4

// ToCSV serializes the AuditEvent as a CSV
func (ae *AuditEvent) ToCSV() string {
	parts := make([]string, auditEventLen)

	parts[0] = ae.UserID
	parts[1] = strconv.FormatUint(ae.ID, 10)
	parts[2] = ae.EventType
	parts[3] = ae.Content // TODO: Escape commas somehow?

	return strings.Join(parts, ",")
}

// ParseAuditEvent attempts to parse a csv as an AuditEvent
func ParseAuditEvent(csv string) (AuditEvent, error) {
	parts := strings.Split(csv, ",")

	if len(parts) != auditEventLen {
		msg := fmt.Sprintf("Expected %d values in AuditEvent csv", auditEventLen)
		return AuditEvent{}, errors.New(msg)
	}

	id, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return AuditEvent{}, err
	}

	return AuditEvent{
		UserID:    parts[0],
		ID:        id,
		EventType: parts[2],
		Content:   parts[3],
	}, nil
}
