package linestringops

import (
	"database/sql"
	"fmt"

	"github.com/twpayne/go-geom/encoding/wkb"
)

type Intersection struct {
	I    []int
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

func CalculateCommonPoints(lineStringsTab *[]wkb.LineString) (*[]Intersection, error) {
	// Map that stores X, Y and related indices
	m := make(map[Key][]int)

	for i, ls := range *lineStringsTab {
		for _, c := range ls.Coords() {
			m[Key{c.X(), c.Y()}] = append(m[Key{c.X(), c.Y()}], i)
		}
	}
	intersectionTab := []Intersection{}
	for key, el := range m {
		if len(el) > 1 {
			intersectionTab = append(intersectionTab, Intersection{el, key.X, key.Y})
		}
	}
	return &intersectionTab, nil
}
