package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"websocket/storage"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"shortUrl"`
}

type handler struct {
	schema  string
	host    string
	storage storage.Service
}

func New(schema string, host string, storage storage.Service) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//h := handler{schema, host, storage}
	//r.POST("/encode", responseHandler(h.encode))
	//r.GET("/:shortLink", h.redirect)
	//r.GET("/info/:shortLink", responseHandler(h.decode))

	return r
}

func responseHandler(h func(ctx *gin.Context) (interface{}, int, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, status, err := h(ctx)
		if err != nil {
			data = err.Error()
		}
		ctx.JSON(status, response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}
