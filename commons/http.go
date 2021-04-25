package commons

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResp(resp http.ResponseWriter, code int, message string) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(code)
	_ = json.NewEncoder(resp).Encode(map[string]string{
		"data": message,
	})
}
