package main

import (
	transaltor "github.com/chipx/go-translator"
	"github.com/chipx/go-translator/api"
	"github.com/chipx/go-translator/datasource"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"time"
)
import _ "github.com/jackc/pgx/v4/stdlib"

func main() {

	db, err := sqlx.Open("pgx", "postgres://root:123@127.0.0.1/postgres")
	if err != nil {
		panic(err)
	}

	db.Ping()
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	ds := datasource.NewSqlDataSource(db, "translations", time.Minute)

	connString := "localhost:8080"
	lis, err := net.Listen("tcp", connString)
	if err != nil {
		log.WithError(err).Fatalf("failed to listen: %s", connString)
	}

	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_logrus.UnaryServerInterceptor(log.NewEntry(log.StandardLogger())),
	),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(log.NewEntry(log.StandardLogger())),
		))
	api.RegisterTranslatorServer(grpcServer, transaltor.NewApiServer(ds, false))
	grpcServer.Serve(lis)
}
