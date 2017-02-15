package types

import (
	"reflect"
	"testing"
)

func TestQuoteRequest_ToCSV(t *testing.T) {
	type fields struct {
		Stock      string
		UserID     string
		AllowCache bool
		ID         uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{"AAPL", "jappleseed", true, 1},
			want:   "AAPL,jappleseed,true,1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qr := &QuoteRequest{
				Stock:      tt.fields.Stock,
				UserID:     tt.fields.UserID,
				AllowCache: tt.fields.AllowCache,
				ID:         tt.fields.ID,
			}
			if got := qr.ToCSV(); got != tt.want {
				t.Errorf("QuoteRequest.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseQuoteRequest(t *testing.T) {
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
			args:    args{"AAPL,jappleseed,false,1"},
			want:    QuoteRequest{"AAPL", "jappleseed", false, 1},
			wantErr: false,
		},
		{
			name:    "Empty string",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "Too many args",
			args:    args{"AAPL,jsmith,false,1,hello!"},
			wantErr: true,
		},
		{
			name:    "Too few args",
			args:    args{"AAPL"},
			wantErr: true,
		},
		{
			name:    "Accepts int as bool",
			args:    args{"AAPL,jappleseed,0,1"},
			want:    QuoteRequest{"AAPL", "jappleseed", false, 1},
			wantErr: false,
		},
		{
			name:    "Does not accept ints beside 0/1",
			args:    args{"AAPL,jappleseed,2,1"},
			wantErr: true,
		},
		{
			name:    "ID can't be negative",
			args:    args{"AAPL,jappleseed,false,-1"},
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
