package app

import (
	"github.com/gin-gonic/gin"
	"go-mall/pkg/constant"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	ErrCode int         `json:"code"`
	Code    int         `json:"status"`
	Msg     string      `json:"msg"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePage struct {
	Code      int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, errCode interface{}, data interface{}) {
	switch errCode.(type) {
	case int:
		intCode := errCode.(int)
		theErrCode := intCode
		if intCode == 200 {
			theErrCode = 0
		}
		g.C.JSON(httpCode, Response{
			Code:    intCode,
			Msg:     constant.GetMsg(intCode),
			Data:    data,
			ErrCode: theErrCode,
		})
	case string:
		strCode := errCode.(string)
		g.C.JSON(httpCode, Response{
			Code: 9999,
			Msg:  strCode,
			Data: data,
		})
	}

	return
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "success",
		"data":    data,
	})
	return
}

func (g *Gin) Error(errCode int, message string) {
	g.C.JSON(http.StatusOK, map[string]interface{}{
		"code":    errCode,
		"message": message,
		"data":    nil,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) ResponsePage(httpCode int, errCode interface{}, data interface{}, total, totalPage int) {
	intCode := errCode.(int)
	g.C.JSON(httpCode, ResponsePage{
		Code:      intCode,
		Msg:       constant.GetMsg(intCode),
		Data:      data,
		Total:     total,
		TotalPage: totalPage,
	})
	return
}
