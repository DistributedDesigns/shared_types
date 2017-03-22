package types

import (
	"reflect"
	"testing"
)

func TestDumplogRequest_ToCSV(t *testing.T) {
	type fields struct {
		UserID   string
		Filename string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Happy path",
			fields: fields{"jappleseed", "ja-dumplog"},
			want:   "jappleseed,ja-dumplog",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := &DumplogRequest{
				UserID:   tt.fields.UserID,
				Filename: tt.fields.Filename,
			}
			if got := dr.ToCSV(); got != tt.want {
				t.Errorf("DumplogRequest.ToCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDumplogRequest(t *testing.T) {
	type args struct {
		csv string
	}
	tests := []struct {
		name    string
		args    args
		want    DumplogRequest
		wantErr bool
	}{
		{
			name: "Happy path",
			args: args{"jappleseed,ja-dumplog"},
			want: DumplogRequest{"jappleseed", "ja-dumplog"},
		},
		{
			name:    "Too many args",
			args:    args{"jappleseed,ja-dumplog,what"},
			wantErr: true,
		},
		{
			name:    "Too few args",
			args:    args{"jappleseed"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDumplogRequest(tt.args.csv)
			if (err != nil) != tt.wantErr {
				t.Errorf("DumplogRequest.ParseDumplogRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DumplogRequest.ParseDumplogRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
