package smarthome

import (
	"log"

	"github.com/hashicorp/go-memdb"
)

type Database struct{}

type DBDevice struct {
	Key   string
	Name  string
	Ip    string
	Alias string
	State string
}

var db *memdb.MemDB

func (d *Database) Init() error {
	log.Print("Initializing Database")
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"device": &memdb.TableSchema{
				Name: "device",
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

	tmpdb, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}
	db = tmpdb

	return nil
}

func (d *Database) Insert(key, name, ip, alias, state string) error {
	log.Printf("Inserting into database: %s - %s - %s - %s - %s", key, name, ip, alias, state)
	data := &DBDevice{key, name, ip, alias, state}
	txn := db.Txn(true)
	if err := txn.Insert("device", data); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (d *Database) HasIp(ip string) (bool, error) {
	txn := db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("device", "id", ip)
	if err != nil {
		return false, err
	}
	log.Print("Database HasIP Returned: ", raw)
	if raw == nil {
		return false, nil
	}

	return true, nil
}

func (d *Database) ReadAll() (memdb.ResultIterator, error) {
	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("device", "id")
	if err != nil {
		return nil, err
	}
	return it, nil
}
