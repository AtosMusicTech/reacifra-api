package main

import (
	"cifra-api/configs"
	"cifra-api/pkg/app/api/middlewares"
	"cifra-api/pkg/app/api/protocol"
	"cifra-api/pkg/app/api/requests"
	"cifra-api/pkg/app/database"
	"cifra-api/pkg/app/responses"
	"cifra-api/pkg/domain/_errors"
	"cifra-api/pkg/domain/actions"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfigFile()
	db := configs.LoadDatabases()

	musicaDatabase := database.NewMusicaDataBase(db)
	transporteDatabase := database.NewTransporteDataBaseMemory()

	newMusicaAction := actions.NewNewMusicaAction(musicaDatabase)
	getMusicasAction := actions.NewGetMusicasAction(musicaDatabase)
	getMusicaAction := actions.NewGetMusicaAction(musicaDatabase)
	setTransporteAction := actions.NewSetTransporteAction(transporteDatabase)
	getTransporteAction := actions.NewGetTransporteAction(transporteDatabase)
	updateMusicaAction := actions.NewUpdateMusicaAction(musicaDatabase)

	engine := gin.Default()
	engine.Use(middlewares.CorsMiddleware())

	engine.POST("/musicas", func(c *gin.Context) {
		musicaRequest := &requests.MusicaRequest{}
		if err := c.ShouldBindJSON(musicaRequest); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		musica, err := newMusicaAction.Execute(musicaRequest.ToEntity())
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToMusicaResponse(musica)))
	})

	engine.GET("/musicas", func(c *gin.Context) {
		musicas, err := getMusicasAction.Execute()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToMusicasResponse(musicas)))
	})

	engine.GET("/musicas/:id", func(c *gin.Context) {
		idStr := c.Params.ByName("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		musica, err := getMusicaAction.Execute(uint(id))
		if err != nil {
			if errors.Is(err, _errors.NewMusicaNaoEncontradaError()) {
				c.AbortWithError(http.StatusNotFound, err)
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, protocol.StatusOk(responses.ConvertToMusicaResponse(musica)))
	})

	engine.PUT("/musicas", func(c *gin.Context) {
		musica := &requests.MusicaRequest{}
		if err := c.BindJSON(musica); err != nil {
			return
		}

		err := updateMusicaAction.Execute(musica.ToEntity())
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

	engine.Run(":8082")
}
