package app

import (
	"fmt"
	"time"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/libs/util"
)

const (
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
	// FieldOJSON ojson
	FieldOJSON = "ojson"
	// FieldIType itype
	FieldIType = "itype"
	// FieldStatus status
	FieldStatus = "status"
	// FieldImage image
	FieldImage = "image"
	// FieldDateON date_on
	FieldDateON = "date_on"
	// FieldIID iid
	FieldIID = "iid"
	// FieldTags tags
	FieldTags = "tags"
	// FieldReles reles
	FieldReles = "reles"
)

// IID int64
func IID(table string) int64 {
	return util.IID32(table, httpd.Site.Secret, httpd.Site.Nonce, fmt.Sprint(time.Now().UnixNano()))
}
