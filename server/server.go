package server

import (
	"github.com/vivcis/library-app/db"
	"github.com/vivcis/library-app/handlers"
	"github.com/vivcis/library-app/router"
	"log"
)

type Server struct {
	DB     db.PostgresDb
	Router *router.Router
}

func Start() error {
	values := db.InitDBParams()

	//Setting up the Postgres Database
	var PDB = new(db.PostgresDb)
	h := &handlers.Handler{DB: PDB}
	err := PDB.Init(values.Host, values.User, values.Password, values.DbName, values.Port)
	if err != nil {
		log.Println("error trying to initialize", err)
		return err
	}
	routes, port := router.SetupRouter(h)
	err = routes.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
