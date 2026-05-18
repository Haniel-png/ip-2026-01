package utils

import (
	"log"
)

// InsertPatient insere um novo paciente no banco de dados
func InsertPatient(fullName, document, bornDate, bloodType, medicalHistory, phone string) error {
	query := `INSERT INTO patients (full_name, document, born_date, blood_type, medical_history, phone) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := DB.Exec(query, fullName, document, bornDate, bloodType, medicalHistory, phone)
	if err != nil {
		log.Printf("Erro ao inserir paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente inserido com sucesso!")
	return nil
}
