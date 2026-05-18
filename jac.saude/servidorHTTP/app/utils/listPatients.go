package utils

import (
	"log"
)

// ListPatients lista todos os pacientes cadastrados no banco de dados
func ListPatients() ([]Patient, error) {
	query := `SELECT id, full_name, document, born_date, blood_type, medical_history, phone FROM patients`
	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("Erro ao listar pacientes no banco de dados: %v", err)
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var patient Patient
		err := rows.Scan(&patient.ID, &patient.FullName, &patient.Document, &patient.BornDate, &patient.BloodType, &patient.MedicalHistory, &patient.Phone)
		if err != nil {
			log.Printf("Erro ao ler dados do paciente: %v", err)
			continue
		}
		patients = append(patients, patient)
	}
	return patients, nil
}
