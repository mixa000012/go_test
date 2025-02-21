package handlers

import (
	"awesomeProject/service"
	"net/http"
)

// HandleHMACSHA512 генерирует HMAC-SHA512 подпись
// @Summary Генерация HMAC-SHA512 подписи
// @Description Возвращает HMAC-SHA512 подпись переданного текста с указанным ключом
// @Accept  json
// @Produce  json
// @Param request body service.HMACRequest true "Данные запроса"
// @Success 200 {object} service.HMACResponse "Ответ с подписью"
// @Failure 400 {string} string "Некорректные данные"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /sign/hmacsha512 [post]
func HandleHMACSHA512(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST разрешен", http.StatusMethodNotAllowed)
		return
	}

	service.GenerateHMAC(w, r)
}
