package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CountMiddleValue(matrix [][]int) float64 {
	var answer = 0.0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			num := float64(matrix[i][j])
			answer += num
		}
	}
	return answer / float64(len(matrix)*len(matrix[0]))
}

func handleMatrix(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Читаем JSON тело запроса
	var matrix [][]int
	err := json.NewDecoder(r.Body).Decode(&matrix)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Вычисляем среднее арифметическое
	result := CountMiddleValue(matrix)

	// Отправляем результат клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"average": result})
}

func main() {
	http.HandleFunc("/calculate", handleMatrix) // Маршрут для обработки запросов
	fmt.Println("Server is running on http://localhost:8087")
	http.ListenAndServe(":8087", nil) // Запускаем сервер
}
