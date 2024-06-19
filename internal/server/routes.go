package server

import (
	"net/http"

	"github.com/EspenS93/filterview/cmd/web/assets"

	"github.com/gin-gonic/gin"
)

func (s *GinServer) RegisterRoutes() {
	s.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messages": "pong",
		})
	})

	s.StaticFS("/assets", http.FS(assets.Files))

	RegisterPeopleRoutes(s)
}
