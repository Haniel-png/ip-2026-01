package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

// UpdatePatientHandler processa a atualização dos dados de um paciente
func UpdatePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	idStr := request.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	fullName := request.FormValue("fullName")
	document := request.FormValue("cpf")
	bornDate := request.FormValue("birthDate")
	bloodType := request.FormValue("bloodType")
	medicalHistory := request.FormValue("medicalHistory")
	phone := request.FormValue("phone")

	err = utils.UpdatePatient(id, fullName, document, bornDate, bloodType, medicalHistory, phone)
	if err != nil {
		http.Error(response, "Erro ao atualizar os dados no banco de dados", http.StatusInternalServerError)
		return
	}

	// Redireciona após o sucesso
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
