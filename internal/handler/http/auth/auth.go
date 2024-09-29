package auth

import (
	"net/http"

	"github.com/Egor123qwe/logs-viewer/internal/handler/model"
	"github.com/Egor123qwe/logs-viewer/internal/handler/model/auth"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Auth(c *gin.Context)
	Login(c *gin.Context)
}

type service struct{}

func New(router *gin.RouterGroup) Service {
	srv := &service{}

	router.POST("/login", srv.Login)

	return srv
}

func (s service) Auth(c *gin.Context) {
	c.Next()
}

func (s service) Login(c *gin.Context) {
	var credentials auth.Login

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResp{Error: err.Error()})
		return
	}

}
