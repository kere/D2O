package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/libs/util"
)

var (
	// ErrDataNotFound err
	ErrDataNotFound = errors.New("没有找到相应的数据")
	// ErrLogin err
	ErrLogin = errors.New("登录验证失败")
	// ErrUserNotFound err
	ErrUserNotFound = errors.New("用户不存在")
)

const (
	// FieldIID string
	FieldIID = "iid"
	// FieldID string
	FieldID = "id"
	// FieldName string
	FieldName = "name"
	// FieldToken string
	FieldToken = "token"
	// FieldReles string
	FieldReles = "reles"
	// FieldTags string
	FieldTags = "tags"
	// FieldDateON string
	FieldDateON = "date_on"
	// FieldCreatedAt string
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt string
	FieldUpdatedAt = "updated_at"
	// FieldUserID string
	FieldUserID = "updated_at"
	// FieldOJSON string
	FieldOJSON = "o_json"
	// FieldDJSON string 统计分析字段
	FieldDJSON = "d_json"

	// LangCH  ch
	LangCH = 1
	// LangEN en
	LangEN = 2
	// LangJP  jpan
	LangJP = 3
	// LangArab  arab
	LangArab = 4
	// LangSP  西班牙 Spanish
	LangSP = 5

	// FieldNick nick
	FieldNick = "nick"
	// FieldIType itype
	FieldIType = "itype"
	// FieldStatus status
	FieldStatus = "status"
	// FieldImage image
	FieldImage = "image"
)

// IID int64
func IID(table string) int64 {
	return util.IID32(table, httpd.Site.SiteData.Secret, httpd.Site.SiteData.Nonce, fmt.Sprint(time.Now().UnixNano()))
}
