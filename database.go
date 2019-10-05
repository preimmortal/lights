package smarthome

import (
	"github.com/golang/glog"
	"github.com/hashicorp/go-memdb"
)

type Database struct{}

type DBScan struct {
	Name string
	Ip   string
	Port string
}

var db *memdb.MemDB

func (d *Database) Init() error {
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

	tmpdb, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}
	db = tmpdb

	return nil
}

func (d *Database) Insert(name, ip, port string) error {
	data := &DBScan{name, ip, port}
	txn := db.Txn(true)
	if err := txn.Insert("scan", data); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (d *Database) HasIp(ip string) (bool, error) {
	txn := db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("scan", "id", ip)
	if err != nil {
		return false, err
	}
	glog.Info("Database HasIP Returned: ", raw)
	if raw == nil {
		return false, nil
	}

	return true, nil
}

func (d *Database) ReadAll() (memdb.ResultIterator, error) {
	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("scan", "id")
	if err != nil {
		return nil, err
	}
	return it, nil
}
