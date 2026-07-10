package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	"github.com/Davethompson01/School_Paddy_golang/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *sql.DB
}

func main() {
	// automatically load env
	godotenv.Load(".env")

	// gets port number from .env
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Failed to Load Port")
	}

	//get postgres url from env
	postGres_Url := os.Getenv("DB_URL")
	if postGres_Url == "" {
		log.Fatal("Failed to load Url")
	}

	// Instanctiate database connection
	conn, err := database.DatabaseConnection()
	if err != nil {
		log.Fatal("Failed to Load Database connection")
	}

	cfg := config.ApiConfig{
		DB: conn,
	}
	defer cfg.DB.Close()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{},
		ExposedHeaders:   []string{},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	// v1Router.POST("/")
	router.Mount("/v1", v1Router)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	server_err := server.ListenAndServe()
	if server_err != nil {
		log.Fatal(server_err)
	}

}
