package main

import (
	"runtime"
	"orderbook/internal/orderbook"
	_ "orderbook/internal/orderbook/store/localStore/OrderBooksStore"
	"github.com/gin-gonic/gin"
	"fmt"
)


func main() {
	fmt.Println("MAIN Program running....")
	runtime.GOMAXPROCS(4)

	// Config goes here
	app := gin.Default()

	//Routes
	orderbook.RouterMain(app)

	err := app.Run()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("up and running ....")
}
