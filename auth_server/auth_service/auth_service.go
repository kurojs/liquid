package auth_service

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"kurojs.github.com/liquid/auth_server/store"
	"kurojs.github.com/liquid/commons"
)

var (
	once         sync.Once
	loginHandler *LoginHandler
)

type LoginHandler struct {
	Store store.Store
}

func GetLoginHandler() *LoginHandler {
	once.Do(func() {
		loginHandler = &LoginHandler{
			// TODO: use config instead of hardcode db env
			Store: store.GetMySQLStore("root", "h@rd2h@ck", "liquid", "localhost", 3307),
		}
	})

	return loginHandler
}

func (h *LoginHandler) Handle(resp http.ResponseWriter, req *http.Request) {
	user := &store.User{}
	if err := json.NewDecoder(req.Body).Decode(user); err != nil {
		commons.WriteJSONResp(resp, http.StatusInternalServerError, "an internal server error occurred")
		return
	}

	isValid, err := h.Store.ValidateUser(req.Context(), user)
	if err != nil {
		log.Printf("validate user err: %s\n", err)

		commons.WriteJSONResp(resp, http.StatusInternalServerError, "an internal server error occurred")
		return
	}

	if !isValid {
		commons.WriteJSONResp(resp, http.StatusUnauthorized, "wrong user name or password")
		return
	}

	token, err := commons.CreateToken(user.Username)
	if err != nil {
		commons.WriteJSONResp(resp, http.StatusInternalServerError, "an internal server error occurred")
		return
	}

	commons.WriteJSONResp(resp, http.StatusOK, token)
}
