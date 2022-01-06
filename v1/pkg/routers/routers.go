package routers

import (
	"github.com/gin-gonic/gin"

	apauthtokenabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth/access-token/abstract"
	aphandv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1"
	apmiddlewares "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/middlewares"
)

// Router ...
type Router struct {
	Token apauthtokenabstract.Token
	Eng   *gin.Engine
	H     *aphandv1.Handlers
}

// InitRouters set endpoints with yours handlers
func (r *Router) InitRouters() {
	// init v1 routers
	r.initV1Routers()
}

// initV1Routers set endpoints with yours handlers
func (r *Router) initV1Routers() {
	r.Eng.Static("/upload/media", "/media/vallin/data/Projects/Go/web/demo-gin-rest-api-library/v1/gitignore/store/")
	// middleware for cors
	r.Eng.Use(apmiddlewares.Cors())

	apiv1 := r.Eng.Group("/api/v1")
	// middleware for authorization
	apiv1.Use(apmiddlewares.AuthJwt(r.Token))

	{
		apiv1.POST("/accounts", r.H.Accounts.PostCreateAccount)
		apiv1.POST("/login", r.H.Accounts.PostLogin)
		apiv1.POST("/generate-access-token", r.H.Accounts.PostGenerateAccessToken)

		apiv1.POST("/books", r.H.Books.PostAddBook)
		apiv1.GET("/books", r.H.Books.GetListBooks)
		apiv1.GET("/books/:id", r.H.Books.GetRetrieveBook)
		apiv1.PUT("/books/:mid", r.H.Books.PutEditBook)
		apiv1.DELETE("/books/:id", r.H.Books.DeleteBook)
	}
}
