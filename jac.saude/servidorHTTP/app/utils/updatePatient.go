package utils

import (
	"log"
)

// UpdatePatient atualiza os dados de um paciente existente no banco de dados
func UpdatePatient(id int, fullName, document, bornDate, bloodType, medicalHistory, phone string) error {
	query := `UPDATE patients SET full_name = $1, document = $2, born_date = $3, blood_type = $4, medical_history = $5, phone = $6 WHERE id = $7`
	_, err := DB.Exec(query, fullName, document, bornDate, bloodType, medicalHistory, phone, id)
	if err != nil {
		log.Printf("Erro ao atualizar paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente atualizado com sucesso!")
	return nil
}
