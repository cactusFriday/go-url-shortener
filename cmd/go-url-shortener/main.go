package main

import (
	"flag"
	"net/http"

	"github.com/cactusfriday/go-url-shortener/internal/api"
	"github.com/cactusfriday/go-url-shortener/internal/common"
)

func main() {
	var storageType string
	flag.StringVar(&storageType, "s", "ram", "Storage type to use <sql || ram>")
	flag.Parse()

	common.InitConfig(storageType)

	server := api.NewServer()

	http.ListenAndServe(":8080", server)
}
