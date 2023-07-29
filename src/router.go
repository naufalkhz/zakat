package src

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalkhz/zakat/src/controllers"
	"github.com/naufalkhz/zakat/utils"
)

type Router interface {
	Init(e *gin.Engine) *gin.Engine
}

type router struct {
	ctrlEmas controllers.EmasInterface
	ctrlUser controllers.UserInterface
	ctrlAuth controllers.AuthInterface
}

func NewRouter(
	ctrlEmas controllers.EmasInterface,
	ctrlUser controllers.UserInterface,
	ctrlAuth controllers.AuthInterface,
) Router {
	// if ctrlEmas == nil || ctrlLogin == nil {
	if ctrlEmas == nil {
		panic("ctrlEmas or ctrlLogin is nil")
	}

	return &router{
		ctrlEmas: ctrlEmas,
		ctrlUser: ctrlUser,
		ctrlAuth: ctrlAuth,
	}
}

func (r *router) Init(e *gin.Engine) *gin.Engine {
	v1 := e.Group("/api")

	safe := v1.Group("/safe").Use(utils.Auth())
	{
		safe.GET("/emas", r.ctrlEmas.Get)
		safe.PUT("/user", r.ctrlUser.Edit)
		safe.GET("/user", r.ctrlUser.GetUserSessionRest)
	}

	v1.POST("/auth/login", r.ctrlAuth.SignIn)

	v1.POST("/user", r.ctrlUser.Create)
	v1.GET("/user", r.ctrlUser.Get)

	return e
}

// func test() {
// 	svcEmas := services.NewEmasService()
// 	ctrlEmas := controllers.NewEmasInterface(svcEmas)
// 	router := NewRouter(ctrlEmas, controllers.NewLoginService())
// 	router.Init(gin.Default())
// }
