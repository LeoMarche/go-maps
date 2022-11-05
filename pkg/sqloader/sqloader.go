package sqloader

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

const (
	SPATIALITE = "spatialite"
)

func drivername() string {
	type entrypoint struct {
		lib  string
		proc string
	}

	var libs = []entrypoint{
		{"mod_spatialite", "sqlite3_modspatialite_init"},
		{"mod_spatialite.dylib", "sqlite3_modspatialite_init"},
		{"libspatialite.so", "sqlite3_modspatialite_init"},
		{"libspatialite.so.5", "spatialite_init_ex"},
		{"libspatialite.so", "spatialite_init_ex"},
	}

	for _, s := range sql.Drivers() {
		if s == SPATIALITE {
			return SPATIALITE
		}
	}

	sql.Register(SPATIALITE, &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			for _, v := range libs {
				if err := conn.LoadExtension(v.lib, v.proc); err == nil {
					return nil
				}
			}
			return errors.New("spatialite extension not found")
		},
	})

	return SPATIALITE
}

func OpenGPKG(filename string) (*sql.DB, error) {
	return sql.Open(drivername(), "BDC_5-0_GPKG_LAMB93_D038-ED2022-03-15.gpkg")
}
