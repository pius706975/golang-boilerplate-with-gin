package serve

import (
	"go-gin/api/routes"
	envConfig "go-gin/config"
	"go-gin/package/database"
	"go-gin/package/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCMD = &cobra.Command{
	Use:   "serve",
	Short: "For Running api server",
	RunE:  serve,
}

func corsHandler() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	return c
}

func serve(cmd *cobra.Command, args []string) error {
	envCfg := envConfig.LoadConfig()

	errorLogger, debugLogger := utils.InitLogger()
	debugLogger.Println("Starting server...")
	
	db, err := database.NewDB()
	if err != nil {
		errorLogger.Println("DB connection failed: ", err)
		return err
	}

	mainRoute := gin.Default()
	if err := routes.RouteApp(mainRoute, db); err != nil {
		errorLogger.Println("Failed to initialize route: ", err)
		return err
	}

	c := corsHandler()
	handler := c.Handler(mainRoute)

	debugLogger.Printf("Server running on port %s", envCfg.Port)
	if err := http.ListenAndServe(":"+envCfg.Port, handler); err != nil {
		errorLogger.Println("Failed to start server: ", err)
		return err
	}

	return nil
}
