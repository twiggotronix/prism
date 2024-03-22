package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	config "prism/proxy/config"
	models "prism/proxy/models"
)

type ConfigController interface {
	Get(context *gin.Context)
	Set(context *gin.Context)
}

type configController struct {
	appConfig config.AppConfig
}

func NewConfigController(appConfig config.AppConfig) ConfigController {
	return &configController{
		appConfig: appConfig,
	}
}

func (cc *configController) Get(context *gin.Context) {
	config := cc.appConfig.Get()
	context.JSON(http.StatusOK, config)
}

func (cc *configController) Set(context *gin.Context) {
	var config models.Config
	if err := context.ShouldBindJSON(&config); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cc.appConfig.Set(config)
	context.JSON(http.StatusCreated, config)
}
