package provider

import (
	"database/sql"
	"fmt"
	"github.com/shellrean-playground/hotel-be-common-util/config"
)

func GetDatabaseConnection(conf *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		conf.Database.Host,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.DBName,
		conf.Database.Port,
		conf.Database.Timezone,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
