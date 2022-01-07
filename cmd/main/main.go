package main

import (
	"github.com/TobiasSchnizel/Beer-API/beers/web"
	"github.com/TobiasSchnizel/Beer-API/internal/database"
	"github.com/TobiasSchnizel/Beer-API/internal/logs"
	reviews "github.com/TobiasSchnizel/Beer-API/reviews/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationsRootFolder     = "file/migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()

	client := database.NewSqlClient("root:root@tcp(localhost:3306)/beer")
	doMigrate(client, "beer")

	mongoClient := database.NewMongoClient("localhost")

	reviewHandler := reviews.NewReviewHandler(mongoClient)

	handler := web.NewCreateBeerHandler(client)
	mux := Routes(handler, reviewHandler)
	server := NewServer(mux)
	server.Run()
}

func doMigrate(client *database.MySqlClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)
	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migration version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}
