package handlers

import (
	"Go-Backend/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	response := &utils.HTTPResponse{Message: "Server is ready.", Status: 200}
	responseJson, _ := json.Marshal(response)

	log.Print(utils.TextGreen("/ready " + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}
