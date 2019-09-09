package user

import (
	"github.com/kere/gno/db"
)

const (
	//Table users
	Table = "users"
)

// LoginUser 当前登录用户
func LoginUser(nick string) db.MapRow {
	row, _ := db.NewQuery(Table).Select("iid,nick,token,status").Where("nick=?", nick).QueryOne()
	return row
}
