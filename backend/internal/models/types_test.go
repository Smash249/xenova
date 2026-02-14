package models

import (
	"testing"
)

func TestStringListValue(t *testing.T) {
	tests := []struct {
		name    string
		list    StringList
		wantErr bool
	}{
		{
			name:    "Empty list",
			list:    StringList{},
			wantErr: false,
		},
		{
			name:    "Single item",
			list:    StringList{"item1"},
			wantErr: false,
		},
		{
			name:    "Multiple items",
			list:    StringList{"item1", "item2", "item3"},
			wantErr: false,
		},
		{
			name:    "Nil list",
			list:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := tt.list.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("StringList.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && value == nil {
				t.Error("StringList.Value() should not return nil value")
			}
		})
	}
}

func TestStringListScan(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name:    "Valid JSON string",
			input:   `["item1","item2"]`,
			wantErr: false,
		},
		{
			name:    "Valid JSON bytes",
			input:   []byte(`["item1","item2"]`),
			wantErr: false,
		},
		{
			name:    "Empty JSON array",
			input:   `[]`,
			wantErr: false,
		},
		{
			name:    "Nil value",
			input:   nil,
			wantErr: false,
		},
		{
			name:    "Invalid type",
			input:   123,
			wantErr: true,
		},
		{
			name:    "Invalid JSON",
			input:   `invalid json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list StringList
			err := list.Scan(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringList.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
