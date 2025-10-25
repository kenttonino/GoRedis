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

func SQLReadyHandler(w http.ResponseWriter, r *http.Request) {
	route := utils.SQLReadyRoute + " "

	if r.Method == "GET" {
		db, err := services.CreateSQLInstance()

		if err != nil {
			log.Println(utils.TextRed(err.Error()))
			errResponse := &utils.HTTPResponse{Message: err.Error(), Status: http.StatusBadRequest}
			errResponseJson, _ := json.Marshal(errResponse)
			fmt.Fprint(w, string(errResponseJson))
			return
		}
		defer db.Close()

		response := &utils.HTTPResponse{Message: "SQLite is ready", Status: http.StatusOK}
		responseJson, _ := json.Marshal(response)
		log.Print(utils.TextGreen(route + strconv.Itoa(response.Status)))
		fmt.Fprint(w, string(responseJson))
		return
	}

	log.Print(utils.TextRed(route + strconv.Itoa(http.StatusMethodNotAllowed)))
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func SQLCreateUserTableHandler(w http.ResponseWriter, r *http.Request) {
	route := utils.SQLCreateUserTableRoute + " "

	if r.Method == "GET" {
		err := services.CreateSQLUserTable()

		if err != nil {
			log.Println(utils.TextRed(err.Error()))
			errResponse := &utils.HTTPResponse{Message: err.Error(), Status: http.StatusBadRequest}
			errResponseJson, _ := json.Marshal(errResponse)
			fmt.Fprint(w, string(errResponseJson))
			return
		}

		response := &utils.HTTPResponse{Message: "Successfully created the user table.", Status: http.StatusOK}
		responseJson, _ := json.Marshal(response)
		log.Print(utils.TextGreen(route + strconv.Itoa(response.Status)))
		fmt.Fprint(w, string(responseJson))
		return
	}

	log.Print(utils.TextRed(route + strconv.Itoa(http.StatusMethodNotAllowed)))
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
