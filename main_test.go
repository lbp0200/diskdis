package diskdis

import (
	"testing"
)

func TestStart(t *testing.T) {
	db, err := NewDataBase("/tmp/badger")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer func() {
		err = db.Close()
		if err != nil {
			t.Fatal(err.Error())
		}
	}()
	err = db.Set("a", "1")
	if err != nil {
		t.Fatal(err.Error())
	}
	var v string
	v, err = db.Get("a")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(v)
	if v != "1" {
		t.Fatal()
	}
}
