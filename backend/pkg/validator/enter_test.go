package validator

import (
	"testing"
)

func TestInitCustomValidator(t *testing.T) {
	cv := InitCustomValidator()
	
	if cv == nil {
		t.Error("InitCustomValidator should return a non-nil validator")
	}
	
	if cv.validator == nil {
		t.Error("CustomValidator should have a non-nil validator field")
	}
}

func TestCustomValidatorValidate(t *testing.T) {
	cv := InitCustomValidator()
	
	// Test with a valid struct
	type TestStruct struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}
	
	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name: "Valid struct",
			input: TestStruct{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			wantErr: false,
		},
		{
			name: "Invalid - missing required field",
			input: TestStruct{
				Name:  "",
				Email: "john@example.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid - bad email format",
			input: TestStruct{
				Name:  "John Doe",
				Email: "invalid-email",
			},
			wantErr: true,
		},
		{
			name: "Invalid - all fields empty",
			input: TestStruct{
				Name:  "",
				Email: "",
			},
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cv.Validate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
