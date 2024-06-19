package providerInterface

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

type Provider interface {
	GetPage(c *gin.Context) templ.Component
	GetDetail(c *gin.Context) templ.Component
}
