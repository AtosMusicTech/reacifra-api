package main

import (
	"cifra-api/configs"
	"cifra-api/pkg/app"
	"cifra-api/pkg/app/api/middlewares"
	"cifra-api/pkg/app/api/protocol"
	"cifra-api/pkg/app/api/requests"
	"cifra-api/pkg/app/database"
	"cifra-api/pkg/app/responses"
	"cifra-api/pkg/domain/_errors"
	"cifra-api/pkg/domain/actions"
	"cifra-api/pkg/domain/entity"
	"errors"
	"flag"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	portPtr := flag.String("p", "8082", "Port number")
	flag.Parse()

	if portPtr == nil {
		_port := viper.GetString("app.port")
		portPtr = &_port
	}

	configs.LoadConfigFile()
	db := configs.LoadDatabases()

	cifraDatabase := database.NewCifraDataBase(db)
	queueDatabase := database.NewQueueDataBase(db)
	transporteDatabase := database.NewTransporteDataBaseMemory()

	newCifraAction := actions.NewAddCifraAction(cifraDatabase)
	getCifrasAction := actions.NewGetCifrasAction(cifraDatabase)
	getCifraAction := actions.NewGetCifraAction(cifraDatabase)
	setTransporteAction := actions.NewSetTransporteAction(transporteDatabase)
	getTransporteAction := actions.NewGetTransporteAction(transporteDatabase)
	updateCifraAction := actions.NewUpdateCifraAction(cifraDatabase)
	getQueueAction := actions.NewGetQueueAction(queueDatabase)
	addQueueAction := actions.NewAddQueueAction(queueDatabase)
	removeQueueAction := actions.NewRemoveQueueAction(queueDatabase)

	engine := gin.Default()
	engine.Use(middlewares.CorsMiddleware())

	engine.GET("/cifras/import", func(ctx *gin.Context) {
		pageweb := app.NewPageWeb()
		pageweb.Load(ctx.Query("url"))
		html := pageweb.ExtractHtmlByTagName("pre")

		ctx.JSON(http.StatusOK, protocol.StatusOk(gin.H{
			"texto": html,
		}))
	})

	engine.POST("/cifras", func(c *gin.Context) {
		cifraRequest := &requests.CifraRequest{}
		if err := c.ShouldBindJSON(cifraRequest); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		cifra, err := newCifraAction.Execute(cifraRequest.ToEntity())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToCifraResponse(cifra)))
	})

	engine.GET("/cifras", func(c *gin.Context) {
		cifras, err := getCifrasAction.Execute()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToCifrasResponse(cifras)))
	})

	engine.GET("/cifras/:id", func(c *gin.Context) {
		idStr := c.Params.ByName("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		cifra, err := getCifraAction.Execute(uint(id))
		if err != nil {
			if errors.Is(err, _errors.NewCifraNaoEncontradaError()) {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToCifraResponse(cifra)))
	})

	engine.PUT("/cifras", func(c *gin.Context) {
		cifra := &requests.CifraRequest{}
		if err := c.BindJSON(cifra); err != nil {
			return
		}

		err := updateCifraAction.Execute(cifra.ToEntity())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	})

	engine.POST("/transportes", func(c *gin.Context) {
		transporte := &requests.TransporteRequest{}
		if err := c.BindJSON(transporte); err != nil {
			return
		}

		setTransporteAction.Execute(transporte.ToEntity())
	})

	engine.GET("/transportes", func(c *gin.Context) {
		transporte, _ := getTransporteAction.Execute()
		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToTransporteResponse(transporte)))
	})

	engine.GET("/anime", func(c *gin.Context) {
		transporte := entity.NewTransporte(1)
		transporte.Cifra = entity.NewCifra(100)

		setTransporteAction.Execute(transporte)
	})

	engine.GET("/queue", func(c *gin.Context) {
		queue, err := getQueueAction.Execute()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToQueueResponse(queue)))
	})

	engine.POST("/queue", func(c *gin.Context) {
		cifra := &requests.CifraRequest{}
		if err := c.BindJSON(cifra); err != nil {
			return
		}

		err := addQueueAction.Execute(cifra.ToEntity())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	})

	engine.DELETE("/queue", func(c *gin.Context) {
		cifra := &requests.CifraRequest{}
		if err := c.BindJSON(cifra); err != nil {
			return
		}

		err := removeQueueAction.Execute(cifra.ToEntity())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	})

	engine.Run(":" + *portPtr)
}
