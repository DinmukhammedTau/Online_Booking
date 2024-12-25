package handlers

import (
	"encoding/json"
	"net/http"
	"unicode"
)

type JSONResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	// Обработка GET-запроса
	if r.Method == http.MethodGet {
		response := JSONResponse{Status: "success", Message: "GET request received successfully"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Обработка POST-запроса
	if r.Method == http.MethodPost {
		var request map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response := JSONResponse{Status: "fail", Message: "Invalid JSON format"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		message, ok := request["message"].(string)
		if !ok {
			response := JSONResponse{Status: "fail", Message: "Invalid JSON message"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Проверка на наличие чисел в строке
		for _, char := range message {
			if unicode.IsDigit(char) {
				response := JSONResponse{Status: "fail", Message: "Message cannot contain numbers"}
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}
		}

		// Если всё в порядке
		response := JSONResponse{Status: "success", Message: "Data successfully received"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Если метод не поддерживается
	response := JSONResponse{Status: "fail", Message: "Method not allowed"}
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(response)
}
