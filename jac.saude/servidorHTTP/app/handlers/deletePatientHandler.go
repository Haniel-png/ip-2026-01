package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

// DeletePatientHandler processa a exclusão de um paciente pelo ID
func DeletePatientHandler(response http.ResponseWriter, request *http.Request) {
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

	err = utils.DeletePatient(id)
	if err != nil {
		http.Error(response, "Erro ao apagar o registro no banco de dados", http.StatusInternalServerError)
		return
	}

	// Redireciona após o sucesso
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
