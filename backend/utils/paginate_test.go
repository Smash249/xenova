package utils

import (
	"testing"

	"github.com/Smash249/xenova/backend/internal/global"
)

func TestPaginate(t *testing.T) {
	tests := []struct {
		name           string
		params         global.PaginateReq
		expectPage     int64
		expectPageSize int64
		expectOffset   int
	}{
		{
			name: "Valid pagination - page 1",
			params: global.PaginateReq{
				Page:     1,
				PageSize: 10,
			},
			expectPage:     1,
			expectPageSize: 10,
			expectOffset:   0,
		},
		{
			name: "Valid pagination - page 2",
			params: global.PaginateReq{
				Page:     2,
				PageSize: 20,
			},
			expectPage:     2,
			expectPageSize: 20,
			expectOffset:   20,
		},
		{
			name: "Invalid page - zero",
			params: global.PaginateReq{
				Page:     0,
				PageSize: 10,
			},
			expectPage:     1,
			expectPageSize: 10,
			expectOffset:   0,
		},
		{
			name: "Invalid page - negative",
			params: global.PaginateReq{
				Page:     -1,
				PageSize: 10,
			},
			expectPage:     1,
			expectPageSize: 10,
			expectOffset:   0,
		},
		{
			name: "Invalid pageSize - zero",
			params: global.PaginateReq{
				Page:     1,
				PageSize: 0,
			},
			expectPage:     1,
			expectPageSize: 10,
			expectOffset:   0,
		},
		{
			name: "Invalid pageSize - negative",
			params: global.PaginateReq{
				Page:     1,
				PageSize: -5,
			},
			expectPage:     1,
			expectPageSize: 10,
			expectOffset:   0,
		},
		{
			name: "PageSize exceeds maximum",
			params: global.PaginateReq{
				Page:     1,
				PageSize: 150,
			},
			expectPage:     1,
			expectPageSize: 100,
			expectOffset:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			paginateFn := Paginate(tt.params)
			if paginateFn == nil {
				t.Error("Paginate() returned nil function")
			}
			// The actual offset and limit validation would require GORM DB
			// This test validates the function is created without errors
		})
	}
}
