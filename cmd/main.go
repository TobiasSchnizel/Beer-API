package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/TobiasSchnizel/Beer-API/internal/database"
	"github.com/TobiasSchnizel/Beer-API/internal/logs"
	"github.com/go-chi/chi"
)

const (
	migrationsRootFolder = "file/migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.initLogger()
	client := database.NewSqlClient("root:root@tcp(localhost:3308)/beers_review")
	doMigrate(client, beers_review)
	handler := web.NewCreateBeerHandler(client)
	mux := Routes(handler)
	server := NewServer(mux)
	server.Run()

}

func doMigrate(client *database.MysqlClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config())
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)
	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ :=m.Version()
	logs.Log().Infof("current migration version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}