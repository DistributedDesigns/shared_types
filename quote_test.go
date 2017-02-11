package types

import (
	"reflect"
	"testing"
	"time"

	"github.com/distributeddesigns/currency"
)

func TestQuote_ToCSV(t *testing.T) {
	userID := "jappleseed"
	stock := "AAPL"
	tenD, _ := currency.NewFromFloat(10.0)
	unixTime := time.Unix(123, 123456789)
	cryptokey := "abc123="

	type fields struct {
		UserID    string
		Stock     string
		Price     currency.Currency
		Timestamp time.Time
		Cryptokey string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{userID, stock, tenD, unixTime, cryptokey},
			want:   "jappleseed,AAPL,10.00,123123456,abc123=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quote{
				UserID:    tt.fields.UserID,
				Stock:     tt.fields.Stock,
				Price:     tt.fields.Price,
				Timestamp: tt.fields.Timestamp,
				Cryptokey: tt.fields.Cryptokey,
			}
			if got := q.ToCSV(); got != tt.want {
				t.Errorf("Quote.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseQuote(t *testing.T) {
	userID := "jappleseed"
	stock := "AAPL"
	tenD, _ := currency.NewFromFloat(10.0)
	unixTime := time.Unix(123, 123456000)
	cryptokey := "abc123="

	type args struct {
		csv string
	}
	tests := []struct {
		name    string
		args    args
		want    Quote
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{"jappleseed,AAPL,10.00,123123456,abc123="},
			want: Quote{userID, stock, tenD, unixTime, cryptokey},
		},
		{
			name:    "Too few args",
			args:    args{"jappleseed,AAPL,10.00,123123456"},
			wantErr: true,
		},
		{
			name:    "Too many args",
			args:    args{"jappleseed,AAPL,10.00,123123456,abc123=,hello!"},
			wantErr: true,
		},
		{
			name:    "Price stored as string",
			args:    args{"jappleseed,AAPL,$10.00,123123456,abc123="},
			wantErr: true,
		},
		{
			name:    "Price stored as string",
			args:    args{"jappleseed,AAPL,$10.00,123123456,abc123="},
			wantErr: true,
		},
		{
			name:    "Time stored as string",
			args:    args{"jappleseed,AAPL,10.00,1970-01-01 00:02:03.123456789 +0000 UTC,abc123="},
			wantErr: true,
		},
		// TODO: Pass this test!
		// {
		// 	name: "Time stored with extra precision",
		// 	args: args{"jappleseed,AAPL,10.00,123123456789,abc123="},
		// 	want: Quote{userID, stock, tenD, unixTime, cryptokey},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseQuote(tt.args.csv)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseQuote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
