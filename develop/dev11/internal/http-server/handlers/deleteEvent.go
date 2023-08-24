package handlers

import (
	"dev11/internal/storage"
	"net/http"
	"strconv"
)

func DeleteEvent(store *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			SendResponse(w, http.StatusMethodNotAllowed, Response{Error: "method error"})
			return
		}

		// Получаем параметр из тела запроса
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

		// Валидируем параметр
		if userID <= 0 {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Удаляем в хранилище
		if err = store.Delete(userID); err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "delete in storage error"})
			return
		}

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: "event deleted successful!"})
	}
}
