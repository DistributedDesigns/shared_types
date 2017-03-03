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
	unixTime := time.Unix(123456, 789*1e6)
	cryptokey := "abc123="
	var id uint64 = 1

	type fields struct {
		Price     currency.Currency
		Stock     string
		UserID    string
		Timestamp time.Time
		Cryptokey string
		ID        uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{tenD, stock, userID, unixTime, cryptokey, id},
			want:   "10.00,AAPL,jappleseed,123456789,abc123=,1",
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
				ID:        tt.fields.ID,
			}
			if got := q.ToCSV(); got != tt.want {
				t.Errorf("Quote.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseQuote(t *testing.T) {
	tenD, _ := currency.NewFromFloat(10.0)
	stock := "AAPL"
	userID := "jappleseed"
	unixTime := time.Unix(123456, 789*1e6)
	cryptokey := "abc123="
	var id uint64 = 1

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
			args: args{"10.00,AAPL,jappleseed,123456789,abc123=,1"},
			want: Quote{tenD, stock, userID, unixTime, cryptokey, id},
		},
		{
			name: "Strips linebreaks",
			args: args{"10.00,AAPL,jappleseed,123456789,abc123=\n,1"},
			want: Quote{tenD, stock, userID, unixTime, cryptokey, id},
		},
		{
			name:    "Too few args",
			args:    args{"10.00,AAPL,jappleseed,123456789,abc123="},
			wantErr: true,
		},
		{
			name:    "Too many args",
			args:    args{"10.00,AAPL,jappleseed,123456789,abc123=,1,hello!"},
			wantErr: true,
		},
		{
			name:    "Price stored as string",
			args:    args{"$10.00,AAPL,jappleseed,123456789,abc123=,1"},
			wantErr: true,
		},
		{
			name:    "ID is negative",
			args:    args{"10.00,AAPL,jappleseed,123456789,abc123=,-1"},
			wantErr: true,
		},
		{
			name:    "Time stored as formatted date",
			args:    args{"10.00,AAPL,jappleseed,1970-01-02 10:17:36.789 +0000 UTC,abc123=,1"},
			wantErr: true,
		},
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
