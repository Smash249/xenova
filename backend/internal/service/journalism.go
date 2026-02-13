package service

import (
	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/Smash249/xenova/backend/pkg/request"
	"github.com/Smash249/xenova/backend/utils"
)

var JournalismServiceApp = new(JournalismService)

type JournalismService struct{}

// ==================== JournalismSeries 新闻系列 CRUD ====================

// GetJournalismSeriesList 获取新闻系列列表
func (j *JournalismService) GetJournalismSeriesList() ([]models.JournalismSeries, error) {
	var seriesList []models.JournalismSeries
	err := global.DB.Find(&seriesList).Error
	return seriesList, err
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
	query := global.DB.Model(&models.Journalism{}).Where("series_id = ?", params.SeriesID)
	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}
	result, err := utils.Paginator[models.Journalism](query, params.PaginateReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetJournalismById 根据ID获取新闻
func (j *JournalismService) GetJournalismById(id uint) (models.Journalism, error) {
	var journalism models.Journalism
	err := global.DB.First(&journalism, id).Error
	return journalism, err
}

// CreateJournalism 创建新闻
func (j *JournalismService) CreateJournalism(params request.CreateJournalismReq) error {
	return global.DB.Create(&models.Journalism{
		Title:    params.Title,
		Content:  params.Content,
		Cover:    params.Cover,
		Author:   params.Author,
		Source:   params.Source,
		Tags:     params.Tags,
		SeriesID: params.SeriesID,
	}).Error
}

// UpdateJournalism 更新新闻
func (j *JournalismService) UpdateJournalism(params request.UpdateJournalismReq) error {
	return global.DB.Where("id = ?", params.ID).Updates(models.Journalism{
		Title:    params.Title,
		Content:  params.Content,
		Cover:    params.Cover,
		Author:   params.Author,
		Source:   params.Source,
		Tags:     params.Tags,
		SeriesID: params.SeriesID,
	}).Error
}

// DeleteJournalism 删除新闻
func (j *JournalismService) DeleteJournalism(params request.DeleteJournalismReq) error {
	return global.DB.Where("id in ?", params.IdList).Delete(&models.Journalism{}).Error
}
