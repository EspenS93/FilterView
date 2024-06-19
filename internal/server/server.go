package server

import (
	"github.com/EspenS93/filterview/internal/gintemplrenderer"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	*gin.Engine
}

func New() *GinServer {
	server := &GinServer{
		Engine: gin.Default(),
	}
	ginHtmlRenderer := server.HTMLRender
	server.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	server.Engine.SetTrustedProxies(nil)
	return server
}
