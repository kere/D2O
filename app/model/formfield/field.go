package formfield

import "github.com/kere/gno/db"

const (
	// Table string
	Table = "formfields"
)

// All 返回全部字段，user id 可以=0；0代表是默认字段
func All(userID int) db.DataSet {
	d, _ := db.NewQuery(Table).Where("user_id=?", userID).Query()
	return d
}
