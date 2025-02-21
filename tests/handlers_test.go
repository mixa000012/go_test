package tests

import (
	"awesomeProject/handlers"
	"awesomeProject/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHMACSHA512(t *testing.T) {
	reqBody := service.HMACRequest{
		Text: "string",
		Key:  "string",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/sign/hmacsha512", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.HandleHMACSHA512(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ожидался код 200, получен %d", resp.StatusCode)
	}

	var responseBody service.HMACResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Errorf("Ошибка парсинга ответа: %v", err)
	}

	if responseBody.Signature == "" {
		t.Error("Пустая подпись")
	}
}

func TestHandleCreateUser(t *testing.T) {
	reqBody := service.UserRequest{
		Name: "Alice",
		Age:  30,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/postgres/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.HandleCreateUser(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ожидался код 200, получен %d", resp.StatusCode)
	}

	var responseBody service.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Errorf("Ошибка парсинга ответа: %v", err)
	}

	if responseBody.ID == 0 {
		t.Error("ID пользователя не может быть 0")
	}
}

func TestHandleIncrement(t *testing.T) {
	reqBody := service.RequestPayload{
		Key:   "counter",
		Value: 1,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/redis/incr", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.HandleIncrement(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("ожидался код 200, получен %d", resp.StatusCode)
	}

	var responseBody service.ResponsePayload
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Errorf("Ошибка парсинга ответа: %v", err)
	}

	if responseBody.Value == 0 {
		t.Error("Значение в Redis не изменилось")
	}
}
