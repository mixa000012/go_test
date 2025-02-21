package handlers

import (
	"awesomeProject/service"
	"net/http"
)

// HandleIncrement увеличивает значение в Redis
// @Summary Инкрементировать значение в Redis
// @Description Увеличивает значение по ключу на переданное значение
// @Accept  json
// @Produce  json
// @Param request body service.RequestPayload true "Данные запроса"
// @Success 200 {object} service.ResponsePayload "Ответ с новым значением"
// @Failure 400 {string} string "Некорректные данные"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /redis/incr [post]
func HandleIncrement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST разрешен", http.StatusMethodNotAllowed)
		return
	}

	service.IncrementValue(w, r)
}
