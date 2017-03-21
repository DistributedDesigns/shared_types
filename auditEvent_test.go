package types

import (
	"reflect"
	"testing"
)

func TestAuditEvent_ToCSV(t *testing.T) {
	type fields struct {
		UserID    string
		ID        uint64
		EventType string
		Content   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{"jappleseed", uint64(1), "command", "DO THIS NOW"},
			want:   "jappleseed,1,command,DO THIS NOW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ae := &AuditEvent{
				UserID:    tt.fields.UserID,
				ID:        tt.fields.ID,
				EventType: tt.fields.EventType,
				Content:   tt.fields.Content,
			}
			if got := ae.ToCSV(); got != tt.want {
				t.Errorf("AuditEvent.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAuditEvent(t *testing.T) {
	type args struct {
		csv string
	}
	tests := []struct {
		name    string
		args    args
		want    AuditEvent
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{"jappleseed,1,command,DO THIS NOW"},
			want: AuditEvent{"jappleseed", uint64(1), "command", "DO THIS NOW"},
		},
		{
			name:    "Too many args",
			args:    args{"jappleseed,1,command,DO THIS NOW,what?"},
			wantErr: true,
		},
		{
			name:    "Too few args",
			args:    args{"jappleseed,1,command"},
			wantErr: true,
		},
		{
			name:    "ID is string",
			args:    args{"jappleseed,DEAD-BEEF,command,DO THIS NOW"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAuditEvent(tt.args.csv)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAuditEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAuditEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
