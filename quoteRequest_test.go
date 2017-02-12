package types

import (
	"reflect"
	"testing"
)

func TestQuoteRequest_ToCSV(t *testing.T) {
	const (
		stock  = "AAPL"
		userID = "jappleseed"
	)

	type fields struct {
		Stock  string
		UserID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{stock, userID},
			want:   stock + "," + userID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr := &QuoteRequest{
				Stock:  tt.fields.Stock,
				UserID: tt.fields.UserID,
			}
			if got := qr.ToCSV(); got != tt.want {
				t.Errorf("QuoteRequest.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseQuoteRequest(t *testing.T) {
	const (
		stock  = "AAPL"
		userID = "jappleseed"
	)

	type args struct {
		csv string
	}
	tests := []struct {
		name    string
		args    args
		want    QuoteRequest
		wantErr bool
	}{
		{
			name:    "Happy path",
			args:    args{stock + "," + userID},
			want:    QuoteRequest{stock, userID},
			wantErr: false,
		},
		{
			name:    "Empty string",
			args:    args{""},
			want:    QuoteRequest{},
			wantErr: true,
		},
		{
			name:    "Too many args",
			args:    args{"AAPL,jsmith,12345"},
			want:    QuoteRequest{},
			wantErr: true,
		},
		{
			name:    "Too few args",
			args:    args{"AAPL"},
			want:    QuoteRequest{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseQuoteRequest(tt.args.csv)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseQuoteRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseQuoteRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
