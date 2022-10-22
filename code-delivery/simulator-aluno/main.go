package main

import (
	"fmt"
	"log"

	route2 "github.com/fshreiner/Imersao-Full-Stack/tree/main/code-delivery/simulator-aluno/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
