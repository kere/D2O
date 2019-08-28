package selem

import (
	"time"

	"github.com/kere/gno/db"
	"onqee.visualstudio.com/D2O/app"
)

const (
	// Table main
	Table = "s_elems" // 资料从表
	// ITypePerson person
	ITypePerson = 3
)

// VO 从表
type VO struct {
	IID    int64     `json:"iid" skip:"update"`
	MIID   int64     `json:"m_iid"`   // 同步关联的 master IID
	DateON time.Time `json:"date_on"` // 事件发生时间
	UserID int       `json:"user_id" skip:"update"`
	Tags   []int     `json:"tags"`
	Reles  []int     `json:"reles"`

	OJSON      interface{} `json:"o_json"`
	ReviewJSON interface{} `json:"review_json" skip:"insert"`

	Status    int       `json:"status" skip:"insert"`
	IType     int       `json:"itype" skip:"update"`
	UpdatedAt time.Time `json:"updated_at" skip:"insert"`
}

// Table string
func (vo VO) Table() string {
	return Table
}

// Dat 信息
type Dat struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Lang  int    `json:"lang"`
}

// OJSON class
type OJSON struct {
	Items []Dat  `json:"items"`
	Image string `json:"image"`
}

// Person class
type Person struct {
	Image    string    `json:"image"`
	NameCH   string    `json:"name_ch"`
	NameEN   string    `json:"name_en"`
	Nick     string    `json:"nick"` // 外号
	Birthday time.Time `json:"birthday"`
}

// Create 创建从表信息
func Create(usrid int, dateON time.Time, miid int64, ojson interface{}, tags, reles []int, itype int) (int64, error) {
	iid := app.IID(Table)
	vo := VO{
		IID:    iid,
		UserID: usrid,
		MIID:   miid,
		DateON: dateON,
		OJSON:  ojson,
		Tags:   tags,
		Reles:  reles,
		IType:  itype,
	}
	return iid, db.VOCreate(vo)
}

// Update data
func Update(iid int64, dateON time.Time, miid int64, ojson interface{}, tags, reles []int) error {
	vo := VO{
		MIID:   miid,
		DateON: dateON,
		OJSON:  ojson,
		Tags:   tags,
		Reles:  reles,
	}
	return db.VOUpdate(vo, "iid=?", iid)
}

// Delete ele
func Delete(iid int64) error {
	return db.Delete(Table, "iid=?", iid)
}
