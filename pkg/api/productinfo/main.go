package main

import (
	"fmt"
	v1 "github.com/hugomatus/order-system/pkg/api/productinfo/v1"
)

func main() {

	fmt.Println("Starting New Server")
	v1.NewServer()
}
