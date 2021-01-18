package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type MessageService struct {
}

func (ms MessageService) GetMessageList(page, limit int) *utils.Paginate {
	var messageList []models.Message
	var total int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&messageList).Offset(-1).Count(&total)
	return utils.NewPaginate(total, page, limit, messageList)
}
func (ms MessageService) DeleteMessage(messageId uint) {

}
func (ms MessageService) GetMessageById(messageId uint) (msg models.Message) {
	db.DB.First(&msg, messageId)
	return
}
