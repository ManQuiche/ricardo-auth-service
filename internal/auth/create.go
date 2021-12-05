package auth

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CreateUser(writer http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var ur CreateUserRequest

	err := json.NewDecoder(req.Body).Decode(&ur)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(writer).Encode(&ur)
}
