package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/article_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/global"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 文章api
type ArticleController struct {
}

// @Title 文章
// @Description 文章
// @Success 200 {object} app.Response
// @router /weixin/article/info/:id [get]
// @Tags Admin
func (e *ArticleController) Get(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	global.LOG.Info(id)
	articleService := article_service.Article{
		Id: id,
	}
	vo := articleService.Get()
	response.OkWithData(vo, c)
}

// @Title 文章列表
// @Description 文章列表
// @Success 200 {object} app.Response
// @router /weixin/article [get]
// @Tags Admin
func (e *ArticleController) GetAll(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	articleService := article_service.Article{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := articleService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 文章添加
// @Description 文章添加
// @Success 200 {object} app.Response
// @router /weixin/article [post]
// @Tags Admin
func (e *ArticleController) Post(c *gin.Context) {
	var (
		model models.WechatArticle
	)

	paramErr := app.BindAndValidate(c, &model)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	articleService := article_service.Article{
		M: &model,
	}

	if err := articleService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 文章修改
// @Description 文章修改
// @Success 200 {object} app.Response
// @router /weixin/article [put]
// @Tags Admin
func (e *ArticleController) Put(c *gin.Context) {
	var (
		model models.WechatArticle
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	articleService := article_service.Article{
		M: &model,
	}

	if err := articleService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 文章删除
// @Description 文章删除
// @Success 200 {object} app.Response
// @router /weixin/article/:id [delete]
// @Tags Admin
func (e *ArticleController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	articleService := article_service.Article{Ids: ids}

	if err := articleService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 发布文章
// @Description 发布文章
// @Success 200 {object} app.Response
// @router /weixin/article/publish/:id [get]
// @Tags Admin
func (e *ArticleController) Pub(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	global.LOG.Info(id)
	articleService := article_service.Article{
		Id: id,
	}
	if err := articleService.Pub(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData(nil, c)
}
