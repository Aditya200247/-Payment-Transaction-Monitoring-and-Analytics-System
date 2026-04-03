package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"payment_monitor/handlers"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().UTC().Format(time.RFC3339),
		})
	})

	// ── API v1 group ─────────────────────────────────────────────────────────
	v1 := router.Group("/api/v1")

	v1.Use(requestLogger())
	{
		transactions := v1.Group("/transactions")
		{
			transactions.POST("/", handlers.CreateTransaction)
		}
	}
}

func requestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next() 
		duration := time.Since(start)
		status := ctx.Writer.Status()

		level := "INFO"
		if status >= 500 {
			level = "ERROR"
		} else if status >= 400 {
			level = "WARN"
		}

	
		println("["+level+"]",
			ctx.Request.Method,
			ctx.Request.URL.Path,
			status,
			duration.String(),
		)
	}
}