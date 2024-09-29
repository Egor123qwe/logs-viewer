package auth

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	Auth(c *gin.Context)
	Login(c *gin.Context)
}

type service struct {
}

func New(router *gin.RouterGroup) Service {
	srv := &service{}

	router.POST("/login", srv.Login)

	return srv
}

func (s service) Auth(c *gin.Context) {

}

func (s service) Login(c *gin.Context) {

}
