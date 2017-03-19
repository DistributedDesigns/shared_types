package types

import (
	"reflect"
	"testing"
)

func TestAutoTxCancel_ToCSV(t *testing.T) {
	type fields struct {
		Stock    string
		UserID   string
		WorkerID uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{"AAPL", "jappleseed", 1},
			want:   "AAPL,jappleseed,1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aTx := &AutoTxCancel{
				Stock:    tt.fields.Stock,
				UserID:   tt.fields.UserID,
				WorkerID: tt.fields.WorkerID,
			}
			if got := aTx.ToCSV(); got != tt.want {
				t.Errorf("AutoTxCancel.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAutoTxCancel(t *testing.T) {
	type args struct {
		csv string
	}
	tests := []struct {
		name    string
		args    args
		want    AutoTxCancel
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{"AAPL,jappleseed,1"},
			want: AutoTxCancel{"AAPL", "jappleseed", 1},
		},
		{
			name:    "Empty string",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "Too many args",
			args:    args{"AAPL,jappleseed,1,hello!"},
			wantErr: true,
		},
		{
			name:    "Too few args",
			args:    args{"AAPL,jappleseed"},
			wantErr: true,
		},
		{
			name:    "Doesn't accept string workerIDs",
			args:    args{"AAPL,jappleseed,worker:1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAutoTxCancel(tt.args.csv)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAutoTxCancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAutoTxCancel() = %v, want %v", got, tt.want)
			}
		})
	}
}
