package handlers

import "net/http"

//根目录
func RootHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Asset not found\n"))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("It's Works!"))
}
