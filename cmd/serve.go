/*
Copyright © 2023 Nopalita <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/naufalkhz/zakat/src"
	"github.com/naufalkhz/zakat/src/controllers"
	"github.com/naufalkhz/zakat/src/gateway"
	"github.com/naufalkhz/zakat/src/repositories"
	"github.com/naufalkhz/zakat/src/services"
	"github.com/naufalkhz/zakat/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "ini untuk start api server",
	Long:  `bismillah`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := zap.NewExample()
		zap.ReplaceGlobals(logger)

		engine := gin.Default()
		engine.Use(utils.CORS())
		engine.SetTrustedProxies(nil)

		// Load env
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// Get Connection
		connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
		db, err := utils.GetConnDB(connString)
		if err != nil {
			log.Fatal("Error get connection to DB")
		}
		db.Debug()

		// Repository
		repositoryEmas := repositories.NewEmasRepository(db)
		repositoryUser := repositories.NewUserRepository(db)
		repositoryAuth := repositories.NewAuthRepository(db)

		// Gateway
		gatewayEmas := gateway.NewEmasGateway()

		// Service
		serviceEmas := services.NewEmasService(gatewayEmas, repositoryEmas)
		serviceUser := services.NewUserService(repositoryUser)
		serviceAuth := services.NewAuthService(repositoryAuth)

		// Controller
		ctrlEmas := controllers.NewEmasInterface(serviceEmas)
		ctrlUser := controllers.NewUserInterface(serviceUser)
		ctrlAuth := controllers.NewAuthInterface(serviceAuth)

		go src.CronTask(gatewayEmas, serviceEmas)

		router := src.NewRouter(ctrlEmas, ctrlUser, ctrlAuth)
		router.Init(engine).Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}