package handlers

import (
	"net/http"

	"github.com/fitnis/referral-service/models"
	"github.com/fitnis/referral-service/services"

	"github.com/gin-gonic/gin"
)

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","department":"Neurology","reason":"Headache"}' http://localhost:8087/referrals
func CreateReferral(c *gin.Context) {
	var req models.Referral
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.CreateReferral(req))
}

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","department":"Neurology","reason":"Headache"}' http://localhost:8087/referrals/specialist
func ReferToSpecialist(c *gin.Context) {
	var req models.Referral
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.ReferToSpecialist(req))
}
