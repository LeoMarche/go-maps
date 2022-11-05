package linestringops

import (
	"database/sql"
	"fmt"
	"math"

	"github.com/twpayne/go-geom/encoding/wkb"
	"github.com/yourbasic/graph"
)

type Intersection struct {
	I    [][2]int
	X, Y float64
}

type Key struct {
	X, Y float64
}

func LoadTroncons(db *sql.DB, columnName, tabName string) (*[]wkb.LineString, error) {

	query := fmt.Sprintf("SELECT %s FROM %s", columnName, tabName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lineStringTab := []wkb.LineString{}

	for rows.Next() {
		var b []byte
		var mp wkb.LineString

		if err := rows.Scan(&b); err != nil {
			return nil, err
		}

		if err := mp.Scan(b[40:]); err != nil {
			return nil, err
		}

		lineStringTab = append(lineStringTab, mp)
	}
	return &lineStringTab, nil
}

func CreateGraph(lineStringsTab *[]wkb.LineString) (*graph.Mutable, *map[Key]int, *map[int]Key, *map[Key][][2]int, error) {
	// Map that stores X, Y and related indices
	m := make(map[Key][][2]int)
	mapping := make(map[Key]int)
	reverseMapping := make(map[int]Key)

	ind := 0
	for i, ls := range *lineStringsTab {
		for i_c, c := range ls.Coords() {
			k := Key{c.X(), c.Y()}
			m[k] = append(m[k], [2]int{i, i_c})
			if _, ok := mapping[k]; !ok {
				mapping[k] = ind
				reverseMapping[ind] = k
				ind++
			}
		}
	}

	g := graph.New(len(m))

	for key, el := range m {
		for _, ls := range el {
			x := (*lineStringsTab)[ls[0]].Coords()[ls[1]].X()
			y := (*lineStringsTab)[ls[0]].Coords()[ls[1]].Y()
			if ls[1] > 0 {
				prevX := (*lineStringsTab)[ls[0]].Coords()[ls[1]-1].X()
				prevY := (*lineStringsTab)[ls[0]].Coords()[ls[1]-1].Y()
				g.AddCost(mapping[key], mapping[Key{prevX, prevY}], int64(math.Sqrt(math.Pow(x-prevX, 2)+math.Pow(y-prevY, 2))))
			}
			if ls[1] < len((*lineStringsTab)[ls[0]].Coords())-1 {
				nextX := (*lineStringsTab)[ls[0]].Coords()[ls[1]+1].X()
				nextY := (*lineStringsTab)[ls[0]].Coords()[ls[1]+1].Y()
				g.AddCost(mapping[key], mapping[Key{nextX, nextY}], int64(math.Sqrt(math.Pow(x-nextX, 2)+math.Pow(y-nextY, 2))))
			}
		}
	}
	return g, &mapping, &reverseMapping, &m, nil
}

func MatchFromToToNodes(from, to Key, mapping *map[Key]int) (int, int, error) {

	if len(*mapping) == 0 {
		return 0, 0, fmt.Errorf("the supplied mapping is empty")
	}

	// TODO: find a more efficient way to do the matching

	minDistFrom := -1.0
	NodeFrom := -1

	minDistTo := -1.0
	NodeTo := -1

	for key, elem := range *mapping {
		distFrom := math.Pow(key.X-from.X, 2) + math.Pow(key.Y-from.Y, 2)
		if minDistFrom == -1.0 || distFrom < minDistFrom {
			minDistFrom = distFrom
			NodeFrom = elem
		}
		distTo := math.Pow(key.X-to.X, 2) + math.Pow(key.Y-to.Y, 2)
		if minDistTo == -1.0 || distTo < minDistTo {
			minDistTo = distTo
			NodeTo = elem
		}
	}

	return NodeFrom, NodeTo, nil
}

func SolvePath(g *graph.Mutable, NodeFrom, NodeTo int, reverseMapping *map[int]Key, m *map[Key][][2]int) (int64, *[]int, *[]Key, *[]int) {

	sp, dst := graph.ShortestPath(g, NodeFrom, NodeTo)

	roads := []int{}
	geopoints := []Key{}
	lastAvailableRoads := []int{}

	for _, n := range sp {
		k := (*reverseMapping)[n]
		if len(lastAvailableRoads) > 0 {
			newAvailableRoads := []int{}
			for _, v := range (*m)[k] {
				newAvailableRoads = append(newAvailableRoads, v[0])
			}
			road := IntersectLists(lastAvailableRoads, newAvailableRoads)
			roads = append(roads, road[0])
			lastAvailableRoads = newAvailableRoads
		} else {
			for _, v := range (*m)[k] {
				lastAvailableRoads = append(lastAvailableRoads, v[0])
			}
		}
		geopoints = append(geopoints, k)
	}
	return dst, &roads, &geopoints, &sp
}

func IntersectLists(a []int, b []int) []int {
	set := make([]int, 0)
	hash := make(map[int]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
