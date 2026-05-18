package handlers

import (
	"encoding/json"
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

// GetPatientHandler busca e retorna um paciente específico pelo ID
func GetPatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	idStr := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	patient, err := utils.GetPatientByID(id)
	if err != nil {
		http.Error(response, "Paciente não encontrado ou erro no banco", http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(patient)
}
