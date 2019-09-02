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
	// Reles  []int  `json:"reles"`

	OJSON OJSON `json:"o_json"`
	// ReviewJSON interface{} `json:"review_json" skip:"insert"`
	// Status    int       `json:"status" skip:"insert"`
	IType     int       `json:"itype"`
	UpdatedAt time.Time `json:"updated_at" skip:"insert"`
}

// Table string
func (vo VO) Table() string {
	return Table
}

// OJSON class
type OJSON struct {
	Contents []Content `json:"contents"`
	SubForms []SubForm `json:"subforms"`
	Images   []ImageVO `json:"images"`
	Avatar   ImageVO   `json:"avatar"`
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
	Label string `json:"k"`
	Value string `json:"v"`
}

// Delete ele
func Delete(iid int64) error {
	return db.Delete(Table, "iid=?", iid)
}
