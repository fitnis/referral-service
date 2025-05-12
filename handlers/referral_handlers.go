package handlers

import (
	"net/http"

	"github.com/fitnis/referral-service/models"
	"github.com/fitnis/referral-service/services"

	"github.com/gin-gonic/gin"
)

func RegisterReferralRoutes(rg *gin.RouterGroup) {
	ref := rg.Group("/referrals")
	ref.POST("", refer)
	ref.POST("/specialist", refer)
}

func refer(c *gin.Context) {
	var req models.Referral
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.ReferToSpecialist(req))
}
