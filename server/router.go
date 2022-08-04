package server

import (
	"github.com/gin-gonic/gin"
	"tag-engine/handler"
)

func initRouter(engine *gin.Engine) {
	router := engine.Group("/")
	router.POST("/listTag", handler.ListTagRequest)
	router.POST("/getTag", handler.GetTagRequest)
	router.POST("/addTag", handler.AddTagRequest)
	router.POST("/updateTag", handler.UpdateTagRequest)
	router.POST("/deleteTag", handler.DeleteTagRequest)
}
