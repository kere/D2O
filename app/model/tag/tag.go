package tag

import "github.com/kere/gno/db"

const (
	// Table string
	Table = "tags"
)

// VO data
type VO struct {
	ID     int    `json:"id" skip:"all"`
	Name   string `json:"name"`
	CataID int    `json:"cata_id"`
}

// TxCreate tags
func TxCreate(tx *db.Tx, name string, cataID int) (bool, error) {
	isE, err := tx.Exists(Table, "name=?", name)
	if tx.DoError(err) {
		return false, err
	}
	if isE {
		return false, nil
	}

	ins := db.NewInsert(Table)
	_, err = ins.TxInsert(tx, db.MapRow{"name": name, "cata_id": cataID})
	if tx.DoError(err) {
		return false, err
	}

	return true, nil
}

// All 返回全部
func All() db.DataSet {
	dat, _ := db.NewQuery(Table).Select("id,name").Where("cata_id=?", 1).Query()
	return dat
}
