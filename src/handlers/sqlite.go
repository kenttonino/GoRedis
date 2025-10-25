package handlers

import (
	"GoRedis/src/services"
	"GoRedis/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func SQLiteReadyHandler(w http.ResponseWriter, r *http.Request) {
	route := utils.SQLReadyRoute + " "
	err := services.SQLCreateTable()

	if err != nil {
		fmt.Println(utils.TextRed(err.Error()))
		errResponse := &utils.HTTPResponse{Message: "Something is wrong", Status: 500}
		errResponseJson, _ := json.Marshal(errResponse)
		log.Print(utils.TextRed(route + strconv.Itoa(errResponse.Status)))
		fmt.Fprint(w, string(errResponseJson))
		return
	}

	response := &utils.HTTPResponse{Message: "SQLite is ready", Status: 200}
	responseJson, _ := json.Marshal(response)
	log.Print(utils.TextGreen(route + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}
