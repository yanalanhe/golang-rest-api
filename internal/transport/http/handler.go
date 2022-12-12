package transportHttp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type CommentSevice interface {

}

type Handler struct {
	Router 	*mux.Router
	Service	CommentSevice
	Server	*http.Server
}

func NewHandler(service CommentSevice) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: h.Router,
	}
	
	return h
}

func(h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}

func (h *Handler) Serve() error {
	// starts the server in non-blocking fasion
	go func() {		
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer the cancel after cancelling the above context becomes timeout 
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down gracefully")

	return nil
}