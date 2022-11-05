package main

import (
	"log"
	"net/http"

	"github.com/LeoMarche/go-maps/pkg/gomapsapi"
	"github.com/LeoMarche/go-maps/pkg/linestringops"
	"github.com/LeoMarche/go-maps/pkg/sqloader"
	"github.com/gorilla/mux"
)

func handleRequests(ws *gomapsapi.WorkingSet) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/ping", ws.Ping)
	myRouter.HandleFunc("/solvePath", ws.SolvePath)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	// TODO: find a way to precompute and store the graph, instead of recomputing it at start
	// TODO: retrieve route names and return it instead of route numbers

	// Load the GPKG file
	db, err := sqloader.OpenGPKG("BDC_5-0_GPKG_LAMB93_D038-ED2022-03-15.gpkg")
	if err != nil {
		log.Fatal(err)
	}

	// Load the roads
	lineStringTab, err := linestringops.LoadTroncons(db, "geometrie", "troncon_de_route")
	if err != nil {
		log.Fatal(err)
	}

	// Create graph from roads
	g, mapping, reverseMapping, m, err := linestringops.CreateGraph(lineStringTab)

	if err != nil {
		log.Fatal(err)
	}

	ws := gomapsapi.WorkingSet{
		G:              g,
		Mapping:        mapping,
		ReverseMapping: reverseMapping,
		M:              m,
	}

	handleRequests(&ws)

}
