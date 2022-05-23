package types

import (
	"reflect"
	"testing"
	"time"
)

func TestDateTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		dt      Time
		want    []byte
		wantErr bool
	}{
		{
			name:    "test1",
			dt:      Date(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)),
			want:    []byte(`"2020-01-01T00:00:00Z"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dt.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		dt      Time
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			dt:      Date(time.Date(2022, 2, 1, 10, 20, 30, 0, time.UTC)),
			args:    args{b: []byte(`"2022-02-01T10:20:30.000Z"`)},
			wantErr: false,
		},
		{
			name:    "test2",
			dt:      Date(time.Date(1977, 2, 5, 0, 0, 0, 0, time.UTC)),
			args:    args{b: []byte(`"Sat Feb 05 1977 00:00:00 GMT+0000 (Coordinated Universal Time)"`)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &Time{}
			if err := dt.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if dt.Time.Day() != tt.dt.Time.Day() ||
				dt.Time.Hour() != tt.dt.Time.Hour() ||
				dt.Time.Minute() != tt.dt.Time.Minute() ||
				dt.Time.Month() != tt.dt.Time.Month() ||
				dt.Time.Year() != tt.dt.Time.Year() {
				t.Errorf("DateTime.UnmarshalJSON() = %v, want %v", dt.Time, tt.dt.Time)
			}
		})
	}
}
