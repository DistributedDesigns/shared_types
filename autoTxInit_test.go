package types

import (
	"reflect"
	"testing"

	currency "github.com/distributeddesigns/currency"
)

const (
	stock    = "AAPL"
	userID   = "jappleseed"
	workerID = uint64(1)
)

func TestAutoTxInit_ToCSV(t *testing.T) {
	tenD, _ := currency.NewFromFloat(10.0)
	fiveD, _ := currency.NewFromFloat(5.0)

	type fields struct {
		Amount   currency.Currency
		Trigger  currency.Currency
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
			fields: fields{tenD, fiveD, stock, userID, workerID},
			want:   "10.00,5.00,AAPL,jappleseed,1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aTx := &AutoTxInit{
				Amount:   tt.fields.Amount,
				Trigger:  tt.fields.Trigger,
				Stock:    tt.fields.Stock,
				UserID:   tt.fields.UserID,
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
			args: args{"10.00,5.00,AAPL,jappleseed,1"},
			want: AutoTxInit{tenD, fiveD, stock, userID, workerID},
		},
		{
			name:    "Too many args",
			args:    args{"10.00,5.00,AAPL,jappleseed,1,hello!"},
			wantErr: true,
		},
		{
			name:    "Too few args",
			args:    args{"10.00,5.00,AAPL,jappleseed"},
			wantErr: true,
		},
		{
			name:    "Dollar signs in amount",
			args:    args{"$10.00,5.00,AAPL,jappleseed,1"},
			wantErr: true,
		},
		{
			name:    "Dollar signs in trigger",
			args:    args{"10.00,$5.00,AAPL,jappleseed,1"},
			wantErr: true,
		},
		{
			name:    "Worker id as key",
			args:    args{"10.00,5.00,AAPL,jappleseed,worker:1"},
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
