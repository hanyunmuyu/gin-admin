package admins

import (
	"fmt"
	"gin-admin/app/http"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

type UploadController struct {
	http.BaseController
}

// @Summary 单文件上传
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Accept mpfd
// @Tags  admin
// @version 1.0
// @Param file formData file true "文件"
// @success 200 {object} utils.JSONResult{data=map[string]string} "成功"
// @Router /admin/v1/upload [POST]
func (upload *UploadController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		upload.Error(ctx, err.Error())
		return
	}
	ext := path.Ext(file.Filename)
	if !upload.checkExt(ext) {
		upload.Error(ctx, "文件类型不符合要求")
		return
	}
	filePath := fmt.Sprintf("./upload/%v", file.Filename)
	err = ctx.SaveUploadedFile(file, filePath)
	filePath = strings.TrimLeft(filePath, ".")
	if err != nil {
		upload.Error(ctx, err.Error())
		return
	}
	upload.Success(ctx, gin.H{"path": filePath})

}

// @Summary 多文件上传【暂时不支持多文件上传文档，因为插件无法生成文档，file[]作为key上传多文件】
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @accept x-www-form-urlencoded
// @Tags  admin
// @version 1.0
// @success 200 object utils.JSONResult{data=[]string} "成功"
// @Router /admin/v1/upload/multi [POST]
func (upload *UploadController) UploadMulti(ctx *gin.Context) {

	form, err := ctx.MultipartForm()
	if err != nil {
		upload.Error(ctx, err.Error())
		return
	}
	files := form.File["file[]"]
	filePathList := make([]string, len(files)-1)
	for _, file := range files {
		ext := path.Ext(file.Filename)
		if !upload.checkExt(ext) {
			upload.Error(ctx, "文件类型不符合要求")
			return
		}
		filePath := fmt.Sprintf("./upload/%v", file.Filename)
		err = ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			upload.Error(ctx, err.Error())
		}
		filePath = strings.TrimLeft(filePath, ".")

		filePathList = append(filePathList, filePath)
	}
	upload.Success(ctx, filePathList)
}
func (upload *UploadController) checkExt(ext string) bool {
	fileExtList := [...]string{".png", ".gif", ".jpeg", ".jpg"}

	for _, v := range fileExtList {
		if v == ext {
			return true
		}
	}
	return false

}
