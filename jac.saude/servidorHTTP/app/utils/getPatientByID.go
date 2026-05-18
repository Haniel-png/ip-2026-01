package utils

import (
	"log"
)

// GetPatientByID busca um paciente no banco de dados pelo seu ID
func GetPatientByID(id int) (*Patient, error) {
	query := `SELECT id, full_name, document, born_date, blood_type, medical_history, phone FROM patients WHERE id = $1`
	var patient Patient
	err := DB.QueryRow(query, id).Scan(&patient.ID, &patient.FullName, &patient.Document, &patient.BornDate, &patient.BloodType, &patient.MedicalHistory, &patient.Phone)
	if err != nil {
		log.Printf("Erro ao buscar paciente no banco de dados: %v", err)
		return nil, err
	}
	return &patient, nil
}
