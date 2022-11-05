package gomapsapi

import (
	"github.com/LeoMarche/go-maps/pkg/linestringops"
	"github.com/yourbasic/graph"
)

type ReturnCode struct {
	Status string `json:"Status"`
}

type ReturnPath struct {
	Status    string              `json:"Status"`
	Distance  int64               `json:"Distance"`
	Routes    []int               `json:"Routes"`
	GeoPoints []linestringops.Key `json:"GeoPoints"`
}

//WorkingSet contains variables for main to work
type WorkingSet struct {
	G              *graph.Mutable
	Mapping        *map[linestringops.Key]int
	ReverseMapping *map[int]linestringops.Key
	M              *map[linestringops.Key][][2]int
}
