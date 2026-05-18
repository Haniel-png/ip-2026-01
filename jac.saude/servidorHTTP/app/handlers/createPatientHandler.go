package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
)

// CreatePatientHandler é responsável por processar os dados enviados para criar um novo paciente
func CreatePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	fullName := request.FormValue("fullName")
	document := request.FormValue("cpf")
	bornDate := request.FormValue("birthDate")
	bloodType := request.FormValue("bloodType")
	medicalHistory := request.FormValue("medicalHistory")
	phone := request.FormValue("phone")

	err := utils.InsertPatient(fullName, document, bornDate, bloodType, medicalHistory, phone)
	if err != nil {
		http.Error(response, "Erro ao salvar os dados no banco de dados", http.StatusInternalServerError)
		return
	}

	// Redireciona após o sucesso
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
