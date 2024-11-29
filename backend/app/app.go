package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"sirichai-bank/apis"
	"sirichai-bank/config"
	"sirichai-bank/drivers"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type App interface {
	Start(context.Context) error
}

type app struct {
	firebase *drivers.FirebaseClient
	gin      *gin.Engine
}

func NewApp(conf *config.AppConfig, ctx context.Context) (App, error) {
	// mongoDB, err := drivers.MongoDBConn(conf.DBConfig)
	// if err != nil {
	// 	return nil, err
	// }
	//Initialize Firebase
	firebase, err := drivers.FirebaseConn(conf.FirebaseCred)
	if err != nil {
		return nil, err
	}

	
	



	// srvAuth := services.NewAuthService(mongoDB.Collection(dtos.DBCollectionUsers))
	// srvList := services.NewListService(mongoDB.Collection(dtos.DBCollectionListings))
	// srvMsg := services.NewMessageService(mongoDB.Collection(dtos.DBCollectionMessages))

	// f := apis.NewFiberAPI(conf.SecretKey, mongoDB, srvAuth, srvList, srvMsg)
	gin := apis.NewGinAPI(conf.SecretKey, firebase)

	return &app{
		firebase: firebase,
		gin:      gin,
	}, nil
}

func (a *app) Start(ctx context.Context) error {
	g, c := errgroup.WithContext(ctx)

	server := &http.Server{
		Addr:    ":8080",
		Handler: a.gin,
	}

	g.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		slog.Info(fmt.Sprintf("Server started on %s", server.Addr))
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		<-c.Done()

		log.Println("Shutting down the server...")
		return server.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return g.Wait()
}
