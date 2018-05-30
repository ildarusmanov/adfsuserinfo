package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// StatusController provides health-check function
type StatusController struct{}

// CreateStatusController constructor for StatusController
func CreateStatusController() *StatusController {
    return &StatusController{}
}

// Check returns ok if the service is ok
func (c *StatusController) Check(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
