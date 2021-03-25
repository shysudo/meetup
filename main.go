package main

import (
	"context"
	"github.com/go-chi/chi"
	mc "github.com/shysudo/meetup/common"
	"github.com/shysudo/meetup/handlers/add"
	"github.com/shysudo/meetup/handlers/get"
	"github.com/shysudo/meetup/handlers/update"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	mc.InitDb()
}

func router() http.Handler {

	r := chi.NewMux()
	r.Route("/meetup", func(r chi.Router) {
		r.Post("/participants", add.RegisterParticipantHandler)
		r.Put("/participants", update.UpdateParticipantHandler)
		r.Get("/participants", get.GetParticipantListHandler)
	})
	return r
}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	server := http.Server{
		Addr:    ":9001",
		Handler: router(),
	}

	go func() {
		if err := http.ListenAndServe(server.Addr, server.Handler); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	return
}
