package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/damasukma/sandbox_auth/infrastructure/persistence"
	"github.com/damasukma/sandbox_auth/interfaces"
	"github.com/damasukma/sandbox_auth/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load(".env")

	m := flag.Bool("migrate", false, "Migrate Database")
	flag.Parse()

	connection := &utils.Connection{
		User:     os.Getenv("DBUSER"),
		Host:     os.Getenv("DBHOST"),
		Password: os.Getenv("DBPASSWORD"),
		Port:     os.Getenv("DBPORT"),
		DBName:   os.Getenv("DBNAME"),
	}
	config, err := utils.NewConnectionDB(connection)

	if err != nil {
		panic(err.Error())
	}

	services := persistence.NewRepositories(config)

	if *m {
		services.Migrate()
		return
	}

	userHandler := interfaces.NewUsers(services.User)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/auth", userHandler.Login).Methods(http.MethodPost)
	r.HandleFunc("/auth/register", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", userHandler.Find).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", userHandler.FindID).Methods(http.MethodGet)
	r.HandleFunc("/users", userHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", userHandler.Delete).Methods(http.MethodDelete)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Http server error - %v \n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error %v", err.Error())
	}
}
