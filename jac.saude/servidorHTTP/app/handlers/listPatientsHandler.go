package handlers

import (
	"encoding/json"
	"net/http"
	"servidorHTTP/app/utils"
)

// ListPatientsHandler retorna a lista de todos os pacientes
func ListPatientsHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	patients, err := utils.ListPatients()
	if err != nil {
		http.Error(response, "Erro ao listar pacientes do banco de dados", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(patients)
}
