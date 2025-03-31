package models

type Referral struct {
	PatientID  string `json:"patientId"`
	Department string `json:"department"`
	Reason     string `json:"reason"`
}

type GenericResponse struct {
	Message string `json:"message"`
}
