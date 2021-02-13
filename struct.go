package diskdis

import (
	"github.com/dgraph-io/badger/v3"
)

type DataBase struct {
	localDb *badger.DB
}

func NewDataBase(path string) (db *DataBase, err error) {
	db = &DataBase{}
	db.localDb, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return
	}
	return
}

func (db *DataBase) Close() (err error) {
	if db.localDb == nil || db.localDb.IsClosed() {
		return
	}
	return db.localDb.Close()
}
