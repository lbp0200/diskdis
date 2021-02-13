package diskdis

import "github.com/dgraph-io/badger/v3"

func (db *DataBase) Get(key string) (value string, err error) {
	if db.localDb == nil || db.localDb.IsClosed() {
		return
	}

	var valCopy []byte
	err = db.localDb.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return nil
	})
	value = string(valCopy)
	return
}

func (db *DataBase) Set(key, value string) (err error) {
	if db.localDb == nil || db.localDb.IsClosed() {
		return
	}

	err = db.localDb.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), []byte(value))
		err := txn.SetEntry(e)
		return err
	})
	return
}
