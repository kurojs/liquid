package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"kurojs.github.com/liquid/commons"
)

func Start(port int) error {
	router := mux.NewRouter()
	router.Methods("POST").Path("/verify").HandlerFunc(VerifyAccessToken)

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

	fmt.Printf("Starting client on %s ...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

type VerifyAccessTokenPayload struct {
	Token string
}

func VerifyAccessToken(resp http.ResponseWriter, req *http.Request) {
	payload := &VerifyAccessTokenPayload{}
	if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
		commons.WriteJSONResp(resp, http.StatusInternalServerError, "an error has been occurred")
		return
	}

	authToken, err := commons.ClaimToken(payload.Token)
	if err != nil {
		fmt.Printf("claim token failed %s\n", err)
		commons.WriteJSONResp(resp, http.StatusInternalServerError, err.Error())
		return
	}

	if err := authToken.Valid(); err != nil {
		commons.WriteJSONResp(resp, http.StatusUnauthorized, err.Error())
		return
	}

	commons.WriteJSONResp(resp, http.StatusOK, fmt.Sprintf("Welcome %s! Read for challenge?", authToken.UserName))
}
