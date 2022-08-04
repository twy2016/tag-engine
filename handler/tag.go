package handler

import (
	"github.com/gin-gonic/gin"
	"tag-engine/model"
	response "tag-engine/pkg/resp"
	"tag-engine/service"
)

func ListTagRequest(c *gin.Context) {
	res, err := service.ListTag()
	if err != nil {
		response.DataBaseFail(c, err)
		return
	}
	response.SendResponse(c, res, "查询列表成功")
}

func GetTagRequest(c *gin.Context) {
	tag := model.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.SendResponseInvalidParameter(c, err)
		return
	}
	res, err := service.GetTag(&tag)
	if err != nil {
		response.DataBaseFail(c, err)
		return
	}
	response.SendResponse(c, res, "获取详情成功")
}

func AddTagRequest(c *gin.Context) {
	tag := model.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.SendResponseInvalidParameter(c, err)
		return
	}
	res, err := service.AddTag(&tag)
	if err != nil {
		response.DataBaseFail(c, err)
		return
	}
	response.SendResponse(c, res, "新增成功")
}

func UpdateTagRequest(c *gin.Context) {
	tag := model.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.SendResponseInvalidParameter(c, err)
		return
	}
	res, err := service.UpdateTag(&tag)
	if err != nil {
		response.DataBaseFail(c, err)
		return
	}
	response.SendResponse(c, res, "更新成功")
}

func DeleteTagRequest(c *gin.Context) {
	tag := model.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.SendResponseInvalidParameter(c, err)
		return
	}
	res, err := service.DeleteTag(tag.Id)
	if err != nil {
		response.DataBaseFail(c, err)
		return
	}
	response.SendResponse(c, res, "删除成功")
}
