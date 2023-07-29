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
	ctrlEmas  controllers.EmasInterface
	ctrlUser  controllers.UserInterface
	ctrlAuth  controllers.AuthInterface
	ctrlBank  controllers.BankInterface
	ctrlZakat controllers.ZakatInterface
}

func NewRouter(
	ctrlEmas controllers.EmasInterface,
	ctrlUser controllers.UserInterface,
	ctrlAuth controllers.AuthInterface,
	ctrlBank controllers.BankInterface,
	ctrlZakat controllers.ZakatInterface,
) Router {
	// if ctrlEmas == nil || ctrlLogin == nil {
	if ctrlEmas == nil {
		panic("ctrlEmas or ctrlLogin is nil")
	}

	return &router{
		ctrlEmas:  ctrlEmas,
		ctrlUser:  ctrlUser,
		ctrlAuth:  ctrlAuth,
		ctrlBank:  ctrlBank,
		ctrlZakat: ctrlZakat,
	}
}

func (r *router) Init(e *gin.Engine) *gin.Engine {
	v1 := e.Group("/v1")

	v1.POST("/auth/login", r.ctrlAuth.SignIn)

	v1.GET("/emas", r.ctrlEmas.Get)

	v1.POST("/user", r.ctrlUser.Create)
	user := v1.Group("/user").Use(utils.Auth())
	{
		user.PUT("", r.ctrlUser.Edit)
		user.GET("", r.ctrlUser.GetUserSessionRest)
	}

	bank := v1.Group("/bank").Use(utils.Auth())
	{
		bank.POST("", r.ctrlBank.Create)
		bank.GET("/list", r.ctrlBank.GetListBank)
		bank.GET("/:id", r.ctrlBank.GetBankById)
	}

	zakat := v1.Group("/zakat").Use(utils.Auth())
	{
		zakat.POST("/penghasilan", r.ctrlZakat.CreateZakatPenghasilan)
	}

	return e
}

// func test() {
// 	svcEmas := services.NewEmasService()
// 	ctrlEmas := controllers.NewEmasInterface(svcEmas)
// 	router := NewRouter(ctrlEmas, controllers.NewLoginService())
// 	router.Init(gin.Default())
// }
