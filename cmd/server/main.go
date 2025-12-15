package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pubudu2003060/proxy_project/internal/config"
	"github.com/pubudu2003060/proxy_project/internal/db"
	"github.com/pubudu2003060/proxy_project/internal/server"

	_ "github.com/lib/pq"
)

func main() {
	envConfig := config.NewEnvConfig()
	log.Println("server run in port:", envConfig.Port)

	db, err := db.Connect(envConfig.DBURL)
	if err != nil {
		log.Fatal("database connection error:", err)
	}
	defer db.Close()

	router := server.NewRouter(db)

	srv := &http.Server{
		Addr:         ":" + envConfig.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("server running in port:", envConfig.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server failed")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("shutdown initiated")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
	log.Println("server stopped")

}
