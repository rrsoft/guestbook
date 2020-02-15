package data

import (
	"database/sql"
	"sync"
)

var (
	dbs = make(map[string]*sql.DB)
	mu  sync.Mutex // protects dbs
)

func Open(dsn string) (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()
	db, ok := dbs[dsn]
	if !ok {
		var err error
		db, err = sql.Open(AppStting.Driver, dsn)
		if err != nil {
			return nil, err
		}
		dbs[dsn] = db
	}
	return db, nil
}

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
