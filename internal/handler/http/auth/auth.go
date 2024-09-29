package auth

import (
	"net/http"

	"github.com/Egor123qwe/logs-viewer/internal/handler/model"
	"github.com/Egor123qwe/logs-viewer/internal/handler/model/auth"
	srvmodel "github.com/Egor123qwe/logs-viewer/internal/model/auth"
	authsrv "github.com/Egor123qwe/logs-viewer/internal/service/auth"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const (
	sessionKey = "logs_viewer_session"
)

type Handler interface {
	Auth(c *gin.Context)

	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type handler struct {
	srv   authsrv.Service
	store *sessions.CookieStore
}

func New(router *gin.RouterGroup, srv authsrv.Service) Handler {
	h := &handler{
		srv:   srv,
		store: sessions.NewCookieStore([]byte("secret")),
	}

	router.POST("/login", h.Login)
	router.DELETE("/logout", h.Logout)

	return h
}

func (h handler) Login(c *gin.Context) {
	var credentials auth.Login

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResp{Error: err.Error()})
		return
	}

	session, err := h.store.Get(c.Request, sessionKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResp{Error: err.Error()})
		return
	}

	srvReq := srvmodel.Credentials{
		Username: credentials.Username,
		Password: credentials.Password,
	}

	if err := h.srv.Auth(c, srvReq); err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResp{Error: err.Error()})
		return
	}

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResp{Error: err.Error()})
		return
	}
}

func (h handler) Logout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:   sessionKey,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(c.Writer, cookie)
}

func (h handler) Auth(c *gin.Context) {
	session, err := h.store.Get(c.Request, sessionKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResp{Error: err.Error()})
		c.Abort()
		return
	}

	if session.IsNew {
		c.JSON(http.StatusUnauthorized, model.ErrorResp{Error: "unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
