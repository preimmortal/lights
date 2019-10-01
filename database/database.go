package database

import (
	"github.com/hashicorp/go-memdb"
)

type Scan struct {
	Name string
	Ip   string
	Port string
}

var database *memdb.MemDB

func Init() error {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"scan": &memdb.TableSchema{
				Name: "scan",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Ip"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}

	database = db

	return nil
}

func Insert(name, ip, port string) error {
	data := &Scan{name, ip, port}
	txn := database.Txn(true)
	if err := txn.Insert("scan", data); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func ReadAll() (memdb.ResultIterator, error) {
	txn := database.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("scan", "id")
	if err != nil {
		return nil, err
	}
	return it, nil
}
