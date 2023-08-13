package app

import (
	"errors"
	"fmt"
	"github.com/avtara/boilerplate-go/utils"
	"github.com/gearintellix/u2"
	"github.com/jmoiron/sqlx"
	"time"
)

func (cfg *App) InitPostgres() (err error) {
	sqlConn := utils.GetConfig("database.connection_string", `
		host=__host__
		port=__port__
		user=__user__
		password=__password__
		dbname=__name__
		sslmode=__sslMode__
		application_name=__appKey__
	`)
	sqlConn = u2.Binding(sqlConn, map[string]string{
		"host":     utils.GetConfig("database.host", "127.0.0.1"),
		"port":     utils.GetConfig("database.port", "5432"),
		"user":     utils.GetConfig("database.user", "boilerplate-go"),
		"password": utils.GetConfig("database.password", "p4swOrd"),
		"name":     utils.GetConfig("database.name", "boilerplate-go"),
		"sslMode":  utils.GetConfig("database.sslMode", "disable"),
		"appKey":   utils.GetConfig("database.appKey", "boilerplate-go"),
	})

	db, err := sqlx.Connect(utils.GetConfig("database.engine", "postgres"), sqlConn)
	if err != nil {
		return errors.Join(err, errors.New(fmt.Sprintf("Failed connect to database %s",
			utils.GetConfig("database.name", "boilerplate-go"))))
	}

	db.SetConnMaxLifetime(time.Minute *
		time.Duration(utils.StringToInt(utils.GetConfig("database.connection_lifetime", "15"), 15)))
	db.SetMaxIdleConns(int(utils.StringToInt(utils.GetConfig("database.maximum_idle", "5"), 5)))
	db.SetMaxOpenConns(int(utils.StringToInt(utils.GetConfig("database.connection_max_open", "0"), 0)))

	cfg.DB = db

	return nil
}
