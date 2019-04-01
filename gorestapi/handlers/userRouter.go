package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouter(res http.ResponseWriter, req *http.Request) {
	path := strings.TrimSuffix(req.URL.Path, "/")

	if path == "/users" {
		switch req.Method {
		case http.MethodGet:
			userGetAll(res, req)
			return
		case http.MethodPost:
			return
		default:
			postError(res, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/users")
	if !bson.IsObjectIdHex(path) {
		postError(res, http.StatusMethodNotAllowed)
		return
	}

	// id := bson.ObjectIdHex(path)
	switch req.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(res, http.StatusMethodNotAllowed)
		return
	}
}
