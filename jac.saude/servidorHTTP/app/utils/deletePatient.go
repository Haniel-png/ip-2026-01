package utils

import (
	"log"
)

// DeletePatient remove um paciente do banco de dados pelo seu ID
func DeletePatient(id int) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao apagar paciente do banco de dados: %v", err)
		return err
	}
	log.Println("Paciente apagado com sucesso!")
	return nil
}
