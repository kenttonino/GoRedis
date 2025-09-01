package routes

import (
	"Go-Backend/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	response := &utils.HTTPResponse{Message: "Server is ready.", Status: 200}
	responseJson, _ := json.Marshal(response)
	fmt.Fprint(w, string(responseJson))
}
