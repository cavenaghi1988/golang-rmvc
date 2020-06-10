package controller

import (
	"encoding/json"
	"mvc/services"
	"mvc/util"
	"net/http"
	"strconv"
)

//GetUser get User ID
func GetUser(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseFloat(req.URL.Query().Get("user_id"), 64)
	if err != nil {
		apiErr := &util.ApplicationError{
			Meassage: "User id need be a number",
			Status:   http.StatusBadRequest,
			Code:     "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write(jsonValue)
		// Handel the error
		return
	}

	// Return user to clien

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
