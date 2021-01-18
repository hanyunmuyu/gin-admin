package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/services/admins"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	messageService = admins.MessageService{}
)

type MessageController struct {
	http.BaseController
}

// @Summary 消息列表
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |title|消息标题||
// @Description |content|消息内容||
// @Tags  admin
// @version 1.0
// @Param page query int false "页码" default(1)
// @success 200 {object} utils.JSONResult{data=utils.Paginate}
// @Router /admin/message [GET]
func (mc MessageController) GetMessageList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	messageList := messageService.GetMessageList(page, 15)
	mc.Success(ctx, messageList)
}
func (mc MessageController) AddMessage() {

}

// @Summary 消息详情
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |title|消息标题||
// @Description |content|消息内容||
// @Tags  admin
// @version 1.0
// @Param id path int false "页码" default(1) minimum(1)
// @success 200 {object} utils.JSONResult{data=models.Message}
// @Router /admin/message/{id} [GET]
func (mc MessageController) GetMessageDetail(ctx *gin.Context) {
	form := struct {
		Id uint `uri:"id" binding:"required,gt=0"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := make(map[string]string)
		lang["id"] = "id"
		err := mc.Translate(err, lang)
		if err != nil {
			mc.Error(ctx, err.Error())
			return
		} else {
			mc.Error(ctx, "")
			return
		}
	}
	message := messageService.GetMessageById(form.Id)
	mc.Success(ctx, message)
}
