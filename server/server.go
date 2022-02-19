package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerType interface {
	ShutDownServer()
	Serve()
}

type server struct {
	srv *http.Server
}

func NewServer(port string, s *gin.Engine) ServerType {
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      s,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return &server{srv}
}

func (s *server) Serve() {
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", s.srv.Addr)
	log.Fatal(s.srv.ListenAndServe())
}

func (s *server) ShutDownServer() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	sig := <-ch
	log.Println("Shutting down server in 30s ", sig)

	ctx, _ := context.WithTimeout(context.TODO(), 30*time.Second)
	s.srv.Shutdown(ctx)
}

func (*server) LimitRate(c chan time.Time) {
	// 	for {
	// 		select {
	// 		case <-time.NewTicker(time.Second):
	// 			if len(c) == 0 {
	// 				log.Println("New Requests")
	// 				for i := 0; i < 10; i++ {
	// 					c <- time.Now()
	// 				}
	// 			}
	// 		}
	// 	}
}
