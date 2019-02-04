package main

import (
	"data"
	"fmt"
	"log"
	"net/http"
	"router"
)

func main() {
	db := data.InitDb()
	defer db.Close()
	r := router.CreateRouterHandler()
	http.Handle("/", r)
	fmt.Println("Now serving on localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
