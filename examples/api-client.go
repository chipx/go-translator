package main

import (
	"context"
	"fmt"
	"github.com/chapsuk/grace"
	transaltor "github.com/chipx/go-translator"
	"github.com/chipx/go-translator/datasource"
	"github.com/chipx/go-translator/store"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	apiDataSource := datasource.NewApiDataSource("localhost:8080", time.Second*60, &[]grpc.DialOption{
		grpc.WithInsecure(),
	})

	//store := store2.NewDirectStore(apiDataSource)
	redisClient := redis.NewClient(&redis.Options{Addr: "localhost:6379", DB: 5})
	tStore, err := store.NewRedisStore(redisClient, apiDataSource, time.Second*30)
	if err != nil {
		log.Fatal(err)
	}
	transaltor.InitTranslator(tStore)
	fmt.Println(transaltor.Translate("ru", "Hello world"))
	ctx := grace.ShutdownContext(context.Background())
	<-ctx.Done()
}
