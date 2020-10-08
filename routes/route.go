package routes

import (
	"github.com/divisi-developer-poros/poros-web-backend/controllers/tag"
	"github.com/divisi-developer-poros/poros-web-backend/middleware"
	userController "github.com/divisi-developer-poros/poros-web-backend/controllers/user"
	"net/http"
	"github.com/gin-gonic/gin"
	test "github.com/divisi-developer-poros/poros-web-backend/controllers/testing"


)

var (
	TestingHandlers test.Cobs
	TagHandlers     tag.HandlerTag
	TokenMiddleware middleware.TokenMiddleware
	UserHandlers	userController.UserHandler
)

type Test struct {
	message	string
	status int
}

// Start inisialisasi route yang digunakan
func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// Example Of JWT Middleware
	r.GET("/guest", TestingHandlers.Guest)
	r.POST("/login", TestingHandlers.Login)
	r.GET("/home", TokenMiddleware.AuthorizeToken, TestingHandlers.Home)

	r.GET("/users", UserHandlers.GetAll)
	r.GET("/users/:id", UserHandlers.Get)
	r.POST("/users", UserHandlers.Create)
	r.PUT("/users/:id", UserHandlers.Update)
	r.DELETE("/users/:id", UserHandlers.Delete)

	// tag routes
	r.GET("/tags", TokenMiddleware.AuthorizeToken, TagHandlers.GetTags)
	r.GET("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.GetTagByID)
	r.POST("/tags", TokenMiddleware.AuthorizeToken, TagHandlers.CreateTag)
	r.PUT("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.UpdateTagByID)
	r.DELETE("/tags/:id", TokenMiddleware.AuthorizeToken, TagHandlers.DeleteTag)

	r.Run()
}
