package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/pkg/response"
	"github.com/Smash249/xenova/backend/utils"
	"gorm.io/gorm"
)

type JournalismService struct{}

// ==================== JournalismSeries 新闻系列 CRUD ====================

// GetJournalismSeriesList 获取新闻系列列表
func (j *JournalismService) GetJournalismSeriesList(params request.GetJournalismSeriesReq) (*global.PaginatorResp[models.JournalismSeries], error) {
	query := global.DB.Model(&models.JournalismSeries{})
	result, err := utils.Paginator[models.JournalismSeries](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateJournalismSeries 创建新闻系列
func (j *JournalismService) CreateJournalismSeries(params request.CreateJournalismSeriesReq) error {
	return global.DB.Create(&models.JournalismSeries{
		Name:        params.Name,
		Description: params.Description,
	}).Error
}

// UpdateJournalismSeries 更新新闻系列
func (j *JournalismService) UpdateJournalismSeries(params request.UpdateJournalismSeriesReq) error {
	return global.DB.Model(&models.JournalismSeries{}).Where("id = ?", params.ID).Updates(models.JournalismSeries{
		Name:        params.Name,
		Description: params.Description,
	}).Error
}

// DeleteJournalismSeries 删除新闻系列
func (j *JournalismService) DeleteJournalismSeries(params request.DeleteJournalismSeriesReq) error {
	if len(params.IdList) == 0 {
		return nil
	}
	return global.DB.Where("id in ?", params.IdList).Delete(&models.JournalismSeries{}).Error
}

// ==================== Journalism 新闻 CRUD ====================

// GetJournalismList 获取新闻列表
func (j *JournalismService) GetJournalismList(params request.GetJournalismReq) (*global.PaginatorResp[models.Journalism], error) {
	query := global.DB.Model(&models.Journalism{})
	if params.SeriesID != 0 {
		query = query.Where("series_id = ?", params.SeriesID)
	}
	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}
	result, err := utils.Paginator[models.Journalism](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetJournalismDetail 根据ID获取新闻
func (j *JournalismService) GetJournalismDetail(id uint) (response.JournalismDetailResp, error) {
	var result response.JournalismDetailResp

	tx := global.DB.Begin()
	if tx.Error != nil {
		return result, tx.Error
	}

	if err := tx.Model(&models.Journalism{}).
		Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error; err != nil {
		tx.Rollback()
		return result, err
	}

	err := tx.Table("journalism AS j").
		Select(`
            j.id,
            j.series_id,
            js.name AS series_name,
            j.title,
            j.content,
            j.view_count,
            j.created_at,
            j.updated_at
        `).
		Joins("LEFT JOIN journalism_series AS js ON j.series_id = js.id").
		Where("j.id = ?", id).
		Scan(&result).Error

	if err != nil {
		tx.Rollback()
		return result, err
	}

	if err := tx.Commit().Error; err != nil {
		return result, err
	}

	return result, nil
}

// CreateJournalism 创建新闻
func (j *JournalismService) CreateJournalism(params request.CreateJournalismReq) error {
	return global.DB.Create(&models.Journalism{
		Title:    params.Title,
		Content:  params.Content,
		SeriesID: params.SeriesID,
	}).Error
}

// UpdateJournalism 更新新闻
func (j *JournalismService) UpdateJournalism(params request.UpdateJournalismReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(models.Journalism{
		Title:    params.Title,
		Content:  params.Content,
		SeriesID: params.SeriesID,
	}).Error
}

// DeleteJournalism 删除新闻
func (j *JournalismService) DeleteJournalism(params request.DeleteJournalismReq) error {
	return global.DB.Where("id in ?", params.IdList).Delete(&models.Journalism{}).Error
}
