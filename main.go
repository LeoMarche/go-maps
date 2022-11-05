package main

import (
	"log"

	"github.com/go-spatial/geom/encoding/gpkg"
)

func main() {
	h, err := gpkg.Open("cities.gpkg")
	if err != nil {
		log.Println("err:", err)
		return
	}
	defer h.Close()
}
