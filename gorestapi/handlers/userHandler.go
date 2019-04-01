package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	userpkg "../user"
)

func bodyToUser(req *http.Request, u *userpkg.User) error {
	if req.Body == nil {
		return errors.New("request body is empty")
	}
	if u == nil {
		return errors.New("a user is empty")
	}
	bd, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bd, u)
}

func userGetAll(res http.ResponseWriter, req *http.Request) {
	users, err := userpkg.All()
	if err != nil {
		postError(res, http.StatusInternalServerError)
		return
	}
	postBodyResponse(res, http.StatusOK, jsonResponse{"Users": users})
}

func usersPostOne(res http.ResponseWriter, req *http.Request) {
	u := new(userpkg.User)
	err := bodyToUser(req, u)

	if err != nil {
		postError(res, http.StatusBadRequest)
		return
	}
	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil {

	}
}
