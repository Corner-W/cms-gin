package cms

import (
	v1 "cms/api/v1"

	"github.com/gin-gonic/gin"
)

type CmsRouter struct {
}



func (s *CmsRouter) InitRouter(Router *gin.RouterGroup) {

	r := Router.Group("cms")
	var api = v1.ApiGroupApp.CmsGroup.CmsApi


	//创建一个branch
	r.POST("file", api.UploadFile)
	r.DELETE(":filename", api.DeleteFile)
	r.GET(":filename", api.PVCount)
}