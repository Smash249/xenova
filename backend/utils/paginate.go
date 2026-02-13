package utils

import (
	"math"

	"github.com/Smash249/xenova/backend/internal/global"
	"gorm.io/gorm"
)

func Paginate(params global.PaginateReq) func(db *gorm.DB) *gorm.DB {
	page := params.Page
	pageSize := params.PageSize
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

func Paginator[T any](query *gorm.DB, params global.PaginateReq, orderParams ...string) (*global.PaginatorResp[T], error) {
	var (
		count        int64
		data         []T
		defaultOrder = "created_at DESC"
	)

	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}

	if len(orderParams) > 0 && orderParams[0] != "" {
		defaultOrder = orderParams[0]
	}
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if err := query.Scopes(Paginate(params)).Order(defaultOrder).Find(&data).Error; err != nil {
		return nil, err
	}
	return &global.PaginatorResp[T]{
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalCount: count,
		TotalPage:  int64(math.Ceil(float64(count) / float64(params.PageSize))),
		Data:       data,
	}, nil
}
