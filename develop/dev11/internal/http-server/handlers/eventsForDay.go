package handlers

import (
	"dev11/internal/storage"
	"log"
	"net/http"
	"strconv"
)

// EventsForDay ...
func EventsForDay(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
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

		eventDay, err := strconv.Atoi(r.FormValue("event_day"))
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "convert error"})
			return
		}

		// Валидируем параметры
		if userID <= 0 || eventDay <= 0 {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Ищем данные в хранилище
		dates, err := store.ForDay(userID, eventDay)
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "event for day error"})
			return
		}

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: dates})

		log.Println("event for day successful!")
	}
}
