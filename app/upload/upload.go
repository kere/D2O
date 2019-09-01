package upload

import (
	"fmt"
	"time"

	"github.com/kere/gno/db"
	"github.com/kere/gno/libs/util"
	"github.com/valyala/fasthttp"
)

const (
	// Table string
	Table = "images"

	// StoreDir 储存目录
	StoreDir = "webroot/u/m/"
)

// Image class
type Image struct {
}

// NewImage class
func NewImage() *Image {
	return &Image{}
}

// Auth a
func (m *Image) Auth(ctx *fasthttp.RequestCtx) error {
	return nil
}

// Success f
func (m *Image) Success(ctx *fasthttp.RequestCtx, token, ext, folder string, now time.Time) error {
	name := token + ext
	iid := util.IID32(name)
	row := db.MapRow{
		"iid":     iid,
		"name":    name,
		"date_on": now.Format("200601"),
	}
	_, err := db.CreateIfNotFound(Table, row, "iid=?", iid)
	if err != nil {
		return err
	}

	ctx.WriteString(fmt.Sprint(iid))
	ctx.WriteString(",")
	folderB := util.Str2Bytes(folder)
	ctx.Write(folderB[7:])
	ctx.WriteString("/")
	ctx.WriteString(name)

	return nil
}

// StoreDir f
func (m *Image) StoreDir(now time.Time) string {
	return StoreDir + now.Format("200601")
}
