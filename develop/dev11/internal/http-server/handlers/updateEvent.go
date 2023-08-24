package handlers

import (
	"dev11/internal/storage"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateEvent(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			SendResponse(w, http.StatusMethodNotAllowed, Response{Error: "method error"})
			return
		}

		// Получаем параметры из тела запроса
		err := r.ParseForm()
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "parse error"})
			return
		}

		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "convert error"})
			return
		}

		eventID, err := strconv.Atoi(r.FormValue("event_id"))
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "convert error"})
			return
		}

		date := r.FormValue("date")

		// Валидируем параметры
		if userID <= 0 || eventID <= 0 || date == "" {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Обновляем новые данные в хранилище
		if err = store.Update(userID, eventID, date); err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "event update error"})
			return
		}

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: "event update successful!"})

		fmt.Println("UPDATE:", store)

	}
}
