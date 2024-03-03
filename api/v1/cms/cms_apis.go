package cms

import (

	"cms/global"
	"cms/common/consts"
	"cms/model/common/response"
	"cms/model/cms/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CmsApi struct {
}



// @Summary 创建文章
// @x-creator	"wangdanfeng"
// @Produce json
// @Param tag_id body string true "标签ID"
// @Param title body string true "文章标题"
// @Param desc body string false "文章简述"
// @Param cover_image_url body string true "封面图片地址"
// @Param content body string true "文章内容"
// @Param created_by body int true "创建者"
// @Param state body int false "状态"
// @Success 200 {object} request.Article "成功"
// @Failure 100 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/v1/cms/article [post]
func (a CmsApi) CreateArticle(c *gin.Context) {

	var r request.Article
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error(consts.ParseParamsErrMsg, zap.Any("err", err.Error()))
		response.FailWithMessage(consts.ParseParamsErrMsg, c)
		return
	}
	global.GVA_LOG.Debug("Enter CreateArticle", zap.Any("r", r))


	// 进行一些列内部业务逻辑处理


	response.OkWithData(consts.OperateSucMsg, c)

}


// UploadFile 上传文件
//	@Tags		codeHosts 代码源
//	@x-creator	"wangdanfeng"
//	@Summary	上传文件
//  @Param file formData file true "文件上传"
//	@Accept		application/json
//	@Produce	application/json
//	@Success	200	{object}	string	"操作成功"
//	@Router		/api/v1/cms/file [post]
func (api *CmsApi) UploadFile(c *gin.Context) {

	// 单文件
	file, _ := c.FormFile("file")
	// log.Println(file.Filename)
	global.GVA_LOG.Debug("UploadFile", zap.Any("filename", file))
	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	response.OkWithMessage(consts.OperateSucMsg, c)
}


// UploadFile 删除文件
//	@Tags		codeHosts 代码源
//	@x-creator	"wangdanfeng"
//	@Summary	删除文件
//	@Param		filename path string true "file name"
//	@Accept		application/json
//	@Produce	application/json
//	@Success	200	
//	@Router		/api/v1/cms/:filename [delete]
func (api *CmsApi) DeleteFile(c *gin.Context) {


	

}



//	@x-creator	"wangdanfeng"
//	@Summary	文件或者网页的访问统计
//	@Param		visitor body string true "访客标识 ip或者用户名"
//	@Param		action body  string true "访客行为，点赞，点踩等等"
//	@Param		visite_date body  string true "访问时间"
//	@Accept		application/json
//	@Produce	application/json
//	@Success	200	
//	@Router		/api/v1/cms/pageview [post]
func (api *CmsApi) PVCount(c *gin.Context) {


	var r request.PVCount
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error(consts.ParseParamsErrMsg, zap.Any("err", err.Error()))
		response.FailWithMessage(consts.ParseParamsErrMsg, c)
		return
	}
	global.GVA_LOG.Debug("Enter PVCount", zap.Any("r", r))


	// 进行一些列内部业务逻辑处理


	response.OkWithData(consts.OperateSucMsg, c)

}

