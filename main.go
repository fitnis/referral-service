package main

import (
	"github.com/fitnis/referral-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	r := router.Group("/referrals")
	{
		handlers.RegisterReferralRoutes(r)
	}

	router.Run(":8087")
}
