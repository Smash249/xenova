package request

import (
	"github.com/Smash249/xenova/backend/internal/global"
)

type CreateJournalismSeriesReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateJournalismSeriesReq struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteJournalismSeriesReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}

type GetJournalismReq struct {
	global.PaginateReq
	Title    string `query:"title" json:"title"`
	SeriesID uint   `query:"seriesId" json:"seriesId" validate:"required"`
}

type CreateJournalismReq struct {
	Title    string   `json:"title" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	Cover    string   `json:"cover" validate:"required"`
	Author   string   `json:"author"`
	Source   string   `json:"source"`
	Tags     []string `json:"tags"`
	SeriesID uint     `json:"seriesId" validate:"required"`
}

type UpdateJournalismReq struct {
	ID       uint     `json:"id" validate:"required"`
	Title    string   `json:"title" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	Cover    string   `json:"cover" validate:"required"`
	Author   string   `json:"author"`
	Source   string   `json:"source"`
	Tags     []string `json:"tags"`
	SeriesID uint     `json:"seriesId" validate:"required"`
}

type DeleteJournalismReq struct {
	IdList []uint `json:"id_list" validate:"required"`
}
