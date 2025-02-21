package handlers

import (
	"awesomeProject/service"
	"net/http"
)

// HandleCreateUser создает пользователя в PostgreSQL
// @Summary Создать пользователя в БД
// @Description Добавляет нового пользователя в PostgreSQL
// @Accept  json
// @Produce  json
// @Param request body service.UserRequest true "Данные пользователя"
// @Success 200 {object} service.UserResponse "ID нового пользователя"
// @Failure 400 {string} string "Некорректные данные"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /postgres/users [post]
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST разрешен", http.StatusMethodNotAllowed)
		return
	}

	service.CreateUser(w, r)
}
