package service

import (
	"awesomeProject/config"
	"awesomeProject/db"
	"context"
	"encoding/json"
	"net/http"
)

type RequestPayload struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type ResponsePayload struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func IncrementValue(w http.ResponseWriter, r *http.Request) {
	var req RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	redisClient := db.InitRedis(config.Cfg.RedisAddr)
	defer redisClient.Close()

	ctx := context.Background()
	newValue, err := redisClient.IncrBy(ctx, req.Key, int64(req.Value)).Result()
	if err != nil {
		http.Error(w, "Ошибка работы с Redis", http.StatusInternalServerError)
		return
	}

	resp := ResponsePayload{Key: req.Key, Value: int(newValue)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
