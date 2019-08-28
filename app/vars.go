package app

import (
	"fmt"
	"time"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/libs/util"
)

const (
	// LangEN en
	LangEN = 1
	// LangCH  ch
	LangCH = 2
	// LangSP  西班牙 Spanish
	LangSP = 3
)

// IID int64
func IID(table string) int64 {
	return util.IID32(table, httpd.Site.Secret, httpd.Site.Nonce, fmt.Sprint(time.Now().UnixNano()))
}
