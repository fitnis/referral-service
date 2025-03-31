package services

import (
	"github.com/fitnis/referral-service/models"
)

var referrals []models.Referral

func CreateReferral(r models.Referral) models.GenericResponse {
	referrals = append(referrals, r)
	return models.GenericResponse{Message: "Referral note created"}
}

func ReferToSpecialist(r models.Referral) models.GenericResponse {
	referrals = append(referrals, r)
	return models.GenericResponse{Message: "Referred to specialist"}
}
