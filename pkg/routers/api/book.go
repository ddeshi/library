package api

import (
	"github.com/ddeshi/library/pkg/Service/Book"
	"github.com/gin-gonic/gin"
)

type BookAPIController struct{}

func BookRegisterRouters(r *gin.Engine) {
	book := r.Group("/book")
	{
		book.POST("/add", Book.AddBook)
		book.POST("/delete", Book.DeleteBook)
		book.POST("/update", Book.UpdateBook)
		book.POST("/search", Book.SearchBook)
	}
}
