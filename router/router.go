package router

import (
	"cloud-md-zk/handler/check"
	"cloud-md-zk/handler/zk"
	"cloud-md-zk/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Load loads the middleware's, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middleware's
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// pprof router
	//pprof.Register(g)

	clusterGroup := g.Group("/api/v1/zk/")
	{
		clusterGroup.PUT("/cluster/:clusterName", zk.CreateZKCluster)
		clusterGroup.DELETE("/cluster/:clusterName", zk.DeleteZKCluster)
		clusterGroup.GET("/cluster/:clusterName", zk.GetZKCluster)
		clusterGroup.POST("/cluster/:clusterName", zk.UpdateZKCluster)
	}

	systemCheck := g.Group("/check")
	{
		systemCheck.GET("/health", check.HealthCheck)
		systemCheck.GET("/disk", check.DiskCheck)
		systemCheck.GET("/cpu", check.CPUCheck)
		systemCheck.GET("/ram", check.RAMCheck)
	}

	return g
}
