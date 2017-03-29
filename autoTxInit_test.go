package types

import (
	"reflect"
	"testing"

	currency "github.com/distributeddesigns/currency"
)

const (
	stock    = "AAPL"
	userID   = "jappleseed"
	workerID = int(1)
)

func TestAutoTxInit_ToCSV(t *testing.T) {
	tenD, _ := currency.NewFromFloat(10.0)
	fiveD, _ := currency.NewFromFloat(5.0)
	action := "Buy"

	type fields struct {
		Amount   currency.Currency
		Trigger  currency.Currency
		Action   string
		Stock    string
		UserID   string
		WorkerID int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{tenD, fiveD, action, stock, userID, workerID},
			want:   "AAPL,jappleseed,Buy,10.00,5.00,1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aTx := &AutoTxInit{
				AutoTxKey: AutoTxKey{
					Stock:  tt.fields.Stock,
					UserID: tt.fields.UserID,
					Action: tt.fields.Action,
				},
				Amount:   tt.fields.Amount,
				Trigger:  tt.fields.Trigger,
				WorkerID: tt.fields.WorkerID,
			}
			if got := aTx.ToCSV(); got != tt.want {
				t.Errorf("AutoTxInit.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAutoTxInit(t *testing.T) {
	tenD, _ := currency.NewFromFloat(10.0)
	fiveD, _ := currency.NewFromFloat(5.0)
	action := "Buy"

	type args struct {
		csv string
	}
	tests := []struct {
		name    string
		args    args
		want    AutoTxInit
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{"AAPL,jappleseed,Buy,10.00,5.00,1"},
			want: AutoTxInit{AutoTxKey{stock, userID, action}, tenD, fiveD, workerID},
		},
		{
			name:    "Too many args",
			args:    args{"AAPL,jappleseed,Buy,10.00,5.00,1,hello!"},
			wantErr: true,
		},
		{
			name:    "Too few args",
			args:    args{"AAPL,jappleseed,Buy,10.00,5.00"},
			wantErr: true,
		},
		{
			name:    "Dollar signs in amount",
			args:    args{"AAPL,jappleseed,Buy,$10.00,5.00,1"},
			wantErr: true,
		},
		{
			name:    "Dollar signs in trigger",
			args:    args{"AAPL,jappleseed,Buy,10.00,$5.00,1"},
			wantErr: true,
		},
		{
			name:    "Worker id as key",
			args:    args{"AAPL,jappleseed,Buy,10.00,5.00,worker:1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAutoTxInit(tt.args.csv)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAutoTxInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAutoTxInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
