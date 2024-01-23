package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting the order microservice")
	ctx := cmd.Context()

	r, closeFn := createOrderMicroservice()

	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDER_SERVICE_BIND_ADDR"), Handler: r}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}

		<-ctx.Done()
		fmt.Println("closing order microservice")

		if err := server.Close(); err != nil{
			panic(err)
		}

	}()
}

func createOrderMicroservice()(router *chi.Mux, closeFn func()){

}
