package sele

import (
	"time"

	"github.com/kere/gno/db"
	"onqee.visualstudio.com/D2O/app"
)

const (
	// Table main
	Table = "s_eles" // 资料从表
)

// VO 从表
type VO struct {
	IID       int64     `json:"iid" skip:"update"`
	MIID      int64     `json:"m_iid"`   // 同步关联的 master IID
	DateON    string    `json:"date_on"` // 事件发生时间
	UserID    int       `json:"user_id"`
	Tags      []int     `json:"tags"`
	Reles     []int     `json:"reles"`
	UpdatedAt time.Time `json:"updated_at" skip:"insert"`

	OJSON      interface{} `json:"o_json"`
	ReviewJSON interface{} `json:"review_json" skip:"insert"`
}

// Table string
func (vo VO) Table() string {
	return Table
}

// DatVO 信息
type DatVO struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Lang  int    `json:"lang"`
}

// Create 创建从表信息
func Create(usrid int, dateON string, miid int64, datas []DatVO, tags, reles []int) (int64, error) {
	iid := app.IID(Table)
	vo := VO{
		IID:    iid,
		UserID: usrid,
		MIID:   miid,
		DateON: dateON,
		OJSON:  datas,
		Tags:   tags,
		Reles:  reles,
	}
	return iid, db.VOCreate(vo)
}

// Update data
func Update(iid int64, usrid int, dateON string, miid int64, datas []DatVO, tags, reles []int) error {
	vo := VO{
		MIID:   miid,
		DateON: dateON,
		OJSON:  datas,
		Tags:   tags,
		Reles:  reles,
	}
	return db.VOUpdate(vo, "iid=?", iid)
}

// Delete ele
func Delete(iid int64) error {
	return db.Delete(Table, "iid=?", iid)
}
