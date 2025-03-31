package main

import (
	"github.com/fitnis/referral-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	r := router.Group("/referrals")
	{
		r.POST("", handlers.CreateReferral)
		r.POST("/specialist", handlers.ReferToSpecialist)
	}

	router.Run(":8087")
}
