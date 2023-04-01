package route

import (
	"belajar-SQL/handler"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	api := r.Group("/books") // prefix
	{
		api.POST("", server.CreateBook)       // /Books
		api.GET("/:id", server.GetBookByID)   // /Book/:id
		api.PUT("/:id", server.UpdateBook)    // /Book/:id
		api.DELETE("/:id", server.DeleteBook) // /Book/:id
	}
}
