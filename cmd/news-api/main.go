package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/IvanKyrylov/news-api/internal/config"
	"github.com/IvanKyrylov/news-api/internal/news"
	"github.com/IvanKyrylov/news-api/internal/news/db"
	"github.com/IvanKyrylov/news-api/pkg/logging"
	"github.com/IvanKyrylov/news-api/pkg/postgres"
	"github.com/IvanKyrylov/news-api/pkg/shutdown"
)

func main() {
	logger := logging.Init()
	logger.Println("logger init")
	cfg := config.GetConfig(logger)
	logger.Println("config init")

	logger.Println("router init")
	router := http.NewServeMux()

	client, err := postgres.NewClient(context.Background(), cfg.Postgres.Host, cfg.Postgres.Port,
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SslMode)
	if err != nil {
		logger.Fatal(err)
	}

	newsRepository := db.NewRepository(client, logger)

	newService := news.NewService(newsRepository, logger)

	userHandler := news.Handler{
		Logger:      logger,
		NewsService: newService,
	}

	userHandler.Register(router)

	logger.Println("Start application")
	start(router, logger, cfg, client)
}

func start(router http.Handler, logger *log.Logger, cfg *config.Config, client postgres.Client) {
	var server *http.Server
	var listener net.Listener

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		socketPath := path.Join(appDir, "app.sock")
		logger.Printf("socket path: %s", socketPath)

		logger.Println("create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		logger.Printf("bind application to host: %s and port: %s", cfg.Listen.BindIP, cfg.Listen.Port)
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		if err != nil {
			logger.Fatal(err)
		}
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, logger, server)

	logger.Println("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			client.Close()
			logger.Println("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}
