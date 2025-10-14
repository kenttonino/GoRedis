package handlers

import (
	services "GoRedis/src/services/redis_service"
	"GoRedis/src/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func RedisHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	redisClient := services.NewRedisClient()
	_, err := redisClient.Ping(ctx).Result()

	if err != nil {
		log.Print(utils.TextRed("/redis " + strconv.Itoa(500)))
		errorResponse := &utils.HTTPResponse{Message: "Failed redis connection.", Status: 500}
		errorResponseJson, _ := json.Marshal(errorResponse)
		fmt.Fprint(w, string(errorResponseJson))
		return
	}

	response := &utils.HTTPResponse{Message: "Redis is ready.", Status: 200}
	responseJson, _ := json.Marshal(response)
	log.Print(utils.TextGreen("/redis " + strconv.Itoa(response.Status)))
	fmt.Fprint(w, string(responseJson))
}
