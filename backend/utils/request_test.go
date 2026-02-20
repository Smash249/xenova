package utils

import (
	"testing"
)

func TestParseStringToUint(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint
		wantErr bool
	}{
		{
			name:    "Valid positive number",
			input:   "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "Valid zero",
			input:   "0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "Large number",
			input:   "4294967295",
			want:    4294967295,
			wantErr: false,
		},
		{
			name:    "Invalid - negative number",
			input:   "-123",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Invalid - non-numeric",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Invalid - empty string",
			input:   "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Invalid - float",
			input:   "123.45",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStringToUint(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStringToUint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseStringToUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
