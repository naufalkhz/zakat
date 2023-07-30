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
	ctrlInfaq controllers.InfaqInterface
}

func NewRouter(
	ctrlEmas controllers.EmasInterface,
	ctrlUser controllers.UserInterface,
	ctrlAuth controllers.AuthInterface,
	ctrlBank controllers.BankInterface,
	ctrlZakat controllers.ZakatInterface,
	ctrlInfaq controllers.InfaqInterface,
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
		ctrlInfaq: ctrlInfaq,
	}
}

func (r *router) Init(e *gin.Engine) *gin.Engine {
	v1 := e.Group("/v1")

	v1.POST("/auth/login", r.ctrlAuth.SignIn)
	auth := v1.Group("/auth").Use(utils.Auth())
	{
		auth.GET("/me", r.ctrlAuth.GetUserSessionRest)
	}

	v1.GET("/emas", r.ctrlEmas.Get)

	v1.POST("/user", r.ctrlUser.Create)
	user := v1.Group("/user").Use(utils.Auth())
	{
		user.PUT("", r.ctrlUser.Edit)
		user.GET("/riwayat", r.ctrlUser.GetRiwayatUser)
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
		zakat.POST("/tabungan", r.ctrlZakat.CreateZakatTabungan)
		zakat.POST("/perdagangan", r.ctrlZakat.CreateZakatPerdagangan)
		zakat.POST("/emas", r.ctrlZakat.CreateZakatEmas)
	}

	infaq := v1.Group("/infaq").Use(utils.Auth())
	{
		infaq.POST("", r.ctrlInfaq.CreateInfaq)
		infaq.GET("/list", r.ctrlInfaq.GetListInfaq)
		infaq.POST("/payment", r.ctrlInfaq.CreateInfaqRiwayat)
		infaq.GET("/riwayat/:limit", r.ctrlInfaq.GetListInfaqRiwayatLastLimit)
	}

	return e
}
