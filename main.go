package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/LeoMarche/go-maps/pkg/linestringops"
	"github.com/LeoMarche/go-maps/pkg/sqloader"
)

func main() {

	db, err := sqloader.OpenGPKG("BDC_5-0_GPKG_LAMB93_D038-ED2022-03-15.gpkg")
	if err != nil {
		log.Fatal(err)
	}

	lineStringTab, err := linestringops.LoadTroncons(db, "geometrie", "troncon_de_route")

	if err != nil {
		log.Fatal(err)
	}

	ops, err := linestringops.CalculateCommonPoints(lineStringTab)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reflect.TypeOf((*lineStringTab)[0].Coords()))

	fmt.Println((*ops)[1])
}
