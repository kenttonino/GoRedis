package handlers

import (
	"GoRedis/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func SQLiteReadyHandler(w http.ResponseWriter, r *http.Request) {
	response := &utils.HTTPResponse{Message: "SQLite is ready", Status: 200}
	responseJson, _ := json.Marshal(response)
	log.Print(utils.TextGreen("/sqlite/ready " + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}
