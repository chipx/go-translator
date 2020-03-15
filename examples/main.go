package main

import (
	"fmt"
	"github.com/chipx/go-translator"
	"github.com/chipx/go-translator/datasource"
	"github.com/chipx/go-translator/store"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)
import _ "github.com/jackc/pgx/v4/stdlib"

func main() {
	db, err := sqlx.Open("pgx", "postgres://root:123@127.0.0.1/postgres")
	if err != nil {
		panic(err)
	}

	db.Ping()

	logrus.SetReportCaller(true)

	ds := datasource.NewSqlDataSource(db, "translations", time.Minute)
	tStore, err := store.NewMemoryStore(ds, time.Minute)
	if err != nil {
		panic(err)
	}
	transaltor.InitTranslator(tStore)
	fmt.Printf("T(movie.3)=%s\n", transaltor.Translate("ru", "movie.3"))
	fmt.Printf("T(movie.4)=%s\n", transaltor.Translatef("ru", "movie.4%d", 1))
}
