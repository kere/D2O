package selem

import (
	"time"

	"github.com/kere/gno/db"
)

const (
	// Table main
	Table = "s_elems" // 资料从表
	// ITypePerson person
	ITypePerson = 3
)

// VO 从表
type VO struct {
	IID    int64  `json:"iid" skip:"update"`
	MIID   int64  `json:"m_iid"`   // 同步关联的 master IID
	DateON string `json:"date_on"` // 事件发生时间
	UserID int    `json:"user_id" skip:"update"`
	Tags   []int  `json:"tags"`
	Area   []int  `json:"area"`
	Reles  []int  `json:"reles"`

	OJSON      OJSON       `json:"o_json"`
	ReviewJSON interface{} `json:"review_json" skip:"insert"`

	Status    int       `json:"status" skip:"insert"`
	IType     int       `json:"itype" skip:"update"`
	UpdatedAt time.Time `json:"updated_at" skip:"insert"`
}

// Table string
func (vo VO) Table() string {
	return Table
}

// OJSON class
type OJSON struct {
	Title    string     `json:"title"`
	Text     string     `json:"text"`
	SubForms []SubForm  `json:"subforms"`
	Images   [][]string `json:"images"`
	Avatar   []string   `json:"avatar"`
}

// SubForm form
type SubForm struct {
	Title  string     `json:"title"`
	DateON string     `json:"date_on"`
	Items  []FormItem `json:"items"`
}

// FormItem form
type FormItem struct {
	Label string `json:"k"`
	Value string `json:"v"`
}

// // Create 创建从表信息
// func Create(usrid int, dateON time.Time, miid int64, ojson interface{}, tags, reles []int, itype int) (int64, error) {
// 	iid := app.IID(Table)
// 	vo := VO{
// 		IID:    iid,
// 		UserID: usrid,
// 		MIID:   miid,
// 		DateON: dateON,
// 		OJSON:  ojson,
// 		Tags:   tags,
// 		Reles:  reles,
// 		IType:  itype,
// 	}
// 	return iid, db.VOCreate(vo)
// }

// // Update data
// func Update(iid int64, dateON time.Time, miid int64, ojson interface{}, tags, reles []int) error {
// 	vo := VO{
// 		MIID:   miid,
// 		DateON: dateON,
// 		OJSON:  ojson,
// 		Tags:   tags,
// 		Reles:  reles,
// 	}
// 	return db.VOUpdate(vo, "iid=?", iid)
// }

// Delete ele
func Delete(iid int64) error {
	return db.Delete(Table, "iid=?", iid)
}
