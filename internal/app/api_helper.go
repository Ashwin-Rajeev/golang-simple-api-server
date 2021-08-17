package app

import (
	"encoding/json"
	"net/http"
)

// Response is a JSON response.
type Response struct {
	Status string          `json:"status"`
	Error  string          `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

// send sends a JSON response.
func send(w http.ResponseWriter, status int, result interface{}, err error) {
	var j []byte
	if err == nil {

		res, err := json.Marshal(result)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		r := &Response{
			Status: http.StatusText(status),
			Result: res,
		}
		j, err = json.Marshal(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		r := &Response{
			Status: http.StatusText(status),
			Error:  err.Error(),
		}
		j, err = json.Marshal(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
