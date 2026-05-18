package utils

// Patient representa a entidade de um Paciente no sistema de saúde
type Patient struct {
	ID             int    `json:"id"`
	FullName       string `json:"fullName"`
	Document       string `json:"document"` // CPF / Documento
	BornDate       string `json:"bornDate"`
	BloodType      string `json:"bloodType"`
	MedicalHistory string `json:"medicalHistory"`
	Phone          string `json:"phone"`
}
