package server

import (
	"net/http"

	"github.com/EspenS93/filterview/internal/provider"

	"github.com/gin-gonic/gin"
)

func RegisterPeopleRoutes(s *GinServer) {
	provider := provider.New("people")

	s.GET("/people", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", provider.GetPage(c))
	})
	s.GET("/person/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", provider.GetDetail(c))
	})
}
