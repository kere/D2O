package selem

import (
	"time"

	"github.com/kere/gno/db"
	"onqee.visualstudio.com/D2O/app/model"
	"onqee.visualstudio.com/D2O/app/model/baseinfo"
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
	MIID   int64     `json:"m_iid"`                                       // 同步关联的 master IID
	DateON time.Time `json:"date_on" skipempty:"all" format:"2006-01-02"` // 事件发生时间
	UserID int       `json:"user_id" skip:"update"`
	Nick   string    `json:"nick" skip:"all"`
	Tags   []string  `json:"tags" skipempty:"all"`
	Area   []int     `json:"area" skipempty:"all"`
	// Reles  []int  `json:"reles"`

	OJSON OJSON `json:"o_json"`
	// ReviewJSON interface{} `json:"review_json" skip:"insert"`
	// Status    int       `json:"status" skip:"insert"`
	IType     int       `json:"itype"`
	UpdatedAt time.Time `json:"updated_at" type:"autotime"`
}

// Table string
func (vo VO) Table() string {
	return Table
}

// OJSON class
type OJSON struct {
	Contents []Content       `json:"contents"`
	SubForms []SubForm       `json:"subforms"`
	Images   []ImageVO       `json:"images"`
	Avatar   ImageVO         `json:"avatar"`
	Area     []baseinfo.Area `json:"area"`
}

// Content class
type Content struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Lang  string `json:"lang"`
}

// ImageVO class
type ImageVO struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// SubForm form
type SubForm struct {
	Title  string     `json:"title"`
	DateON string     `json:"date_on"`
	Items  []FormItem `json:"items"`
}

// FormItem form
type FormItem struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

// Delete ele
func Delete(iid int64) error {
	return db.Delete(Table, "iid=?", iid)
}

// PageData by iid
func PageData(iid int64) (db.MapRow, error) {
	row, err := db.NewQuery(Table).Where("iid=?", iid).QueryOne()
	if err != nil {
		return nil, err
	}

	if row.IsEmpty() {
		return row, model.ErrDataNotFound
	}

	vo := VO{}
	db.StrictDBMapRow(row, vo)
	return row, nil
}

// PageViewData by iid
func PageViewData(iid int64) (db.MapRow, error) {
	q := db.NewQuery("s_elems a left join users b on (a.user_id=b.id)").Select("a.iid,a.m_iid,a.date_on,b.nick, a.o_json,a.area,a.tags,a.reles,a.status,a.itype,a.updated_at")
	row, err := q.Where("a.iid=?", iid).QueryOne()
	if err != nil {
		return nil, err
	}
	if row.IsEmpty() {
		return row, model.ErrDataNotFound
	}

	vo := VO{}
	db.StrictDBMapRow(row, vo)

	return row, nil
}
