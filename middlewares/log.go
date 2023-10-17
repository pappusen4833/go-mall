package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/models"
	"go-mall/pkg/jwt"
	"go-mall/pkg/logging"
	"regexp"
	"strings"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.Path
		logging.Info(url)
		method := strings.ToLower(c.Request.Method)
		user, err := jwt.GetAdminUser(c)
		logging.Error(err)
		if err != nil {
			c.Next()
			return
		}

		reg := regexp.MustCompile(`[0-9]+`)
		newUrl := reg.ReplaceAllString(url, "*")
		menu := models.FindMenuByRouterAndMethod(newUrl, method)
		log := models.SysLog{
			Description: menu.Name,
			Method:      method,
			RequestIp:   c.ClientIP(),
			Username:    user.Username,
			Address:     newUrl,
			Browser:     "",
			Type:        0,
			Uid:         user.Id,
		}
		now := time.Now()
		c.Next()
		consume := time.Now().Sub(now)
		log.Time = consume.Microseconds()
		models.AddLog(&log)
	}
}
