package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/lyracampos/quake-log-parser/internal/handlers"
	"github.com/lyracampos/quake-log-parser/internal/logger"
	"github.com/lyracampos/quake-log-parser/internal/usecases"

	"github.com/gorilla/mux"
)

const defatulHTMLPath = "../static/html"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var htmlFilePATH string

	flag.StringVar(&htmlFilePATH, "h", defatulHTMLPath, "path with html file.")
	flag.Parse()

	logger.InitLog()

	logger.Log.Info("run: starting application")

	router := mux.NewRouter()

	parseLogUseCase := usecases.NewParseLogUseCase()
	handler := handlers.NewHandler(parseLogUseCase)
	router.HandleFunc("/parse", handler.Parse)

	if _, err := os.Stat(htmlFilePATH); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("THE FILE DOES NOT EXIST")
		}
	}

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(htmlFilePATH))))

	address := ":8080"
	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	// run API
	go func() {
		logger.Log.Infof("running API HTTP server at: %s", address)

		if err := server.ListenAndServe(); err != nil {
			logger.Log.Errorf("listen and server - failed on running server: %v", err)
		}
	}()

	// Graceful shutting down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	sig := <-c
	logger.Log.Infof("Got signal:", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error("Error shutting down server: %w", err)
	}

	logger.Log.Info("shutting down")
	return nil
}
