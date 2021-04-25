package auth_server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"kurojs.github.com/liquid/auth_server/auth_service"
)

func Start(port int) error {
	router := mux.NewRouter()
	router.Methods("POST").Path("/login").HandlerFunc(auth_service.GetLoginHandler().Handle)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Grateful shutdown when terminate server
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-termChan
		ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer ctxCancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Failed to shutdown server: %v", err)
		}
	}()

	fmt.Printf("Starting server on %s ...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
