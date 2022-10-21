package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wb_l0/publisher"
	"wb_l0/service"
	"wb_l0/subscriber"

	"github.com/gin-gonic/gin"
)

var o_service *service.OrderService

func getByUID(ctx *gin.Context) {
	uid := ctx.Query("order_uid")
	order, err := o_service.GetOrderByUID(&uid)
	if err != nil {
		_, err = ctx.Writer.Write([]byte(fmt.Sprintf("Can't find order with this UID: %v", err)))
		if err != nil {
			fmt.Printf("ERROR: can't find order with this UID: %v\n", err)
		}
		return
	}
	bytes, err := json.Marshal(order)
	if err != nil {
		fmt.Printf("ERROR: can't marshal order: %v\n", err)
		return
	}
	_, err = ctx.Writer.Write(bytes)
	if err != nil {
		fmt.Printf("ERROR: can't write bytes of json: %v\n", err)
		return
	}
}

func main() {
	order_service := service.NewOrderService()

	o_service = order_service

	router := gin.Default()
	// load html and static files..
	router.LoadHTMLFiles("web/html/index.html")
	router.Static("/css", "web/css")
	router.GET("/wb/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "main page",
		})
	})
	router.GET("/wb/get/", getByUID)

	go subscriber.Subscribe(order_service)
	go publisher.Publish()

	router.Run(":8000")
}
