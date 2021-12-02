package routers

import (
	"github.com/gin-gonic/gin"

	aphandv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1"
)

// InitRouters set endpoints and yours handlers
func InitRouters(eng *gin.Engine, h *aphandv1.Handlers) {
	apiv1 := eng.Group("/api/v1")
	{
		apiv1.POST("/books", h.Books.PostCreateBook)
		apiv1.POST("/accounts", h.Accounts.PostCreateAccount)
		apiv1.POST("/login", h.Accounts.PostLogin)
	}

}
