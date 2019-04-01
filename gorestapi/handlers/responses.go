package handlers

import (
	"encoding/json"
	"net/http"
)

type jsonResponse map[string]interface{}

func postError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func postBodyResponse(res http.ResponseWriter, code int, content jsonResponse) {
	if content != nil {
		js, err := json.Marshal(content)
		if err != nil {
			postError(res, http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-type", "application/json")
		res.WriteHeader(code)
		res.Write(js)
		return
	}
	res.WriteHeader(code)
	res.Write([]byte(http.StatusText(code)))
}
