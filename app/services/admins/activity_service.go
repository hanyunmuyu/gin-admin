package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type ActivityService struct{}

func NewActivityService() *ActivityService {
	return &ActivityService{}
}

func (a ActivityService) GetActivityList(page, limit int) *utils.Paginate {
	var activityList []models.Activity
	var count int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&activityList).Offset(-1).Count(&count)
	return utils.NewPaginate(count, page, limit, activityList)
}
