package main

import (
	"log"
	"net/http"
	"time"

	"api-go/configs"
	"api-go/flags"
	"api-go/models"
	"api-go/controllers"
	"api-go/routes"
	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cfgPath, err := flags.ParseFlags()
	handleError(err)

	cfg, err := config.LoadConfig(cfgPath)
	handleError(err)

	conn, err := config.ParseConnectionURL(cfg)
	handleError(err)

	db, err := sql.Open(cfg.Database.Driver, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := setupApp(cfg, db)
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      app,
	}

	log.Println("Application running on " + addr)
	log.Fatal(srv.ListenAndServe())
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func setupApp(cfg *config.Config, db *sql.DB) http.Handler {
	r := models.NewSuperRepository(db)
	api := models.NewSuperHeroApi(cfg.API.Endpoint, cfg.API.Token)
	s := models.NewSuperHeroService(api, r)
	ctrl := controllers.NewSuperHeroCtrl(s)
	return routes.NewHeroRouter(ctrl)
}
