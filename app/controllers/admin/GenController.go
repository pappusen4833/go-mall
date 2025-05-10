package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/params/admin"
	"go-mall/app/services/gen_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 代码生成器api
type GenController struct {
}

// @Title 获取所有表
// @Description 获取所有表
// @Success 200 {object} app.Response
// @router /tools/gen/tables [get]
// @Tags Admin
func (e *GenController) GetAllDBTables(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	genService := gen_service.Gen{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := genService.GetDBTablesAll()
	response.OkWithData(vo, c)
}

// @Title 导入数据库表
// @Description 导入数据库表
// @Success 200 {object} app.Response
// @router /tools/gen/import [post]
// @Tags Admin
func (e *GenController) ImportTable(c *gin.Context) {
	var (
		param admin.GenTableParan
	)
	httpCode, errCode := app.BindAndValid(c, &param)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	genService := gen_service.Gen{
		GenTableParan: &param,
	}

	if err := genService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 获取已经导入的表
// @Description 获取已经导入的表
// @Success 200 {object} app.Response
// @router /tools/gen/systables [get]
// @Tags Admin
func (e *GenController) GetAllTables(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	genService := gen_service.Gen{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := genService.GetTablesAll()
	response.OkWithData(vo, c)
}

// @Title 获取表的信息
// @Description 获取表的信息
// @Success 200 {object} app.Response
// @router /tools/gen/config/:name [get]
// @Tags Admin
func (e *GenController) GetTableInfo(c *gin.Context) {
	name := c.Param("name")
	genService := gen_service.Gen{
		Name: name,
	}
	vo := genService.GetTableInfo()
	response.OkWithData(vo, c)
}

// @Title 获取表的列信息
// @Description 获取表的列信息
// @Success 200 {object} app.Response
// @router /tools/gen/columns [get]
// @Tags Admin
func (e *GenController) GetTableColumns(c *gin.Context) {
	name := c.DefaultQuery("tableName", "")
	genService := gen_service.Gen{
		Name: name,
	}
	vo := genService.GetTableColumns()
	response.OkWithData(vo, c)
}

// @Title 保存配置
// @Description 保存配置
// @Success 200 {object} app.Response
// @router /gen/config [put]
// @Tags Admin
func (e *GenController) ConfigPut(c *gin.Context) {
	var (
		model models.SysTables
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	genService := gen_service.Gen{
		Table: &model,
	}

	if err := genService.TableSave(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 保存列配置
// @Description 保存列配置
// @Success 200 {object} app.Response
// @router /gen/columns [put]
// @Tags Admin
func (e *GenController) ColumnsPut(c *gin.Context) {
	var (
		model []models.SysColumns
	)
	if err := c.ShouldBindJSON(&model); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	genService := gen_service.Gen{
		Columns: model,
	}

	if err := genService.ColumnSave(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 代码预览
// @Description 代码预览
// @Success 200 {object} app.Response
// @router /tools/gen/preview [get]
// @Tags Admin
func (e *GenController) Preview(c *gin.Context) {
	name := c.Param("name")
	genService := gen_service.Gen{
		Name: name,
	}
	vo := genService.Preview()
	response.OkWithData(vo, c)
}

// @Title 代码生产
// @Description 代码生产
// @Success 200 {object} app.Response
// @router /tools/gen/code [get]
// @Tags Admin
func (e *GenController) GenCode(c *gin.Context) {
	name := c.Param("name")
	genService := gen_service.Gen{
		Name: name,
	}
	genService.GenCode()
	//appG.Response(http.StatusOK, "代码已经成功生成在根目录template下", nil)
	response.OkWithData("代码已经成功生成在根目录template下", c)
}
