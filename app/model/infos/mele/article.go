package article

import (
	"fmt"
	"time"

	"github.com/kere/gno/db"
	"github.com/kere/gno/libs/util"
)

const (
	tableM  = "m_eles"  // 资料主表
	tableMd = "m_datas" // 资料主表，内容
	tableS  = "s_eles"  // 资料从表
	tableH  = "h_eles"  // 资料历史表
	tableHd = "h_datas" // 资料历史，内容
)

// SEleVO 从表
type SEleVO struct {
	IID       int64     `json:"iid" skip:"update"`
	MIID      int64     `json:"iid"`     // 关联master IID
	DateON    string    `json:"date_on"` // 事件发生时间
	UserID    int       `json:"user_id"`
	Status    int       `json:"status"`
	UpdatedAt time.Time `json:"updated_at" skip:"insert"`

	OJSON      interface{} `json:"o_json"`
	ReviewJSON interface{} `json:"review_json"`
}

// CreateS 创建从表信息
func CreateS(usrid int, dateON string, title, text string, forumID int) error {
	now := time.Now()
	row := db.MapRow{
		"iid":      util.IID32(fmt.Sprint(usrid, forumID, '-', now.UnixNano())),
		"user_id":  usrid,
		"title":    title,
		"text":     text,
		"forum_id": forumID,
	}
	return db.Create(table, row)
}

// Update article
func Update(iid int64, title, text string) error {
	row := db.MapRow{
		"title":      title,
		"text":       text,
		"updated_at": time.Now(),
	}
	return db.Update(table, row, "iid=?", iid)
}

// Delete article
func Delete(iid int64) error {
	return db.Delete(table, "iid=?", iid)
}
