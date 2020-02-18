package data

import (
	"database/sql"
	"sync"

	"github.com/rrsoft/guestbook/utils"
)

var (
	mu  sync.Mutex // protects dbs
	dbs = make(map[string]*sql.DB)
)

func open(dsn string) (db *sql.DB, err error) {
	mu.Lock()
	defer mu.Unlock()
	var ok bool
	if db, ok = dbs[dsn]; !ok {
		settings := utils.GetSetting()
		db, err = sql.Open(settings.Driver, dsn)
		if err == nil {
			dbs[dsn] = db
		}
	}
	return
}

// Open open database connect
func Open(dsn string) (*sql.DB, error) {
	if db, ok := dbs[dsn]; ok {
		return db, nil
	}
	return open(dsn)
}

// Close release database
func Close() error {
	mu.Lock()
	defer mu.Unlock()
	var err error
	for _, db := range dbs {
		e := db.Close()
		if e != nil {
			err = e
		}
		// delete(dbs, dsn)
	}
	dbs = nil
	return err
}
