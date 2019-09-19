package upload

import (
	"fmt"
	"mime/multipart"

	"github.com/kere/gno/db"
	"github.com/kere/gno/httpd"
	"github.com/kere/gno/libs/util"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/model"
)

var (
	// StoreDir 储存目录
	StoreDir = "webroot/upload"
)

const (
	// Table string
	Table = "images"
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

var filepool bytebufferpool.Pool

// Do upload
func (m *Image) Do(ctx *fasthttp.RequestCtx, fileHeader *multipart.FileHeader) error {
	name := ctx.FormValue(model.FieldName)

	fileName, err := httpd.DoUpload(util.Bytes2Str(name), StoreDir, fileHeader, "md5")
	if err != nil {
		return err
	}

	folderB := util.Str2Bytes(StoreDir)
	folderB = folderB[7:]
	iid := util.IID32(fileName)
	row := db.MapRow{
		"iid":  iid,
		"name": fileName,
		"dir":  util.Bytes2Str(folderB),
	}
	_, err = db.CreateIfNotFound(Table, row, "iid=?", iid)
	if err != nil {
		return err
	}

	ctx.WriteString(fmt.Sprint(iid))
	ctx.WriteString(",")
	ctx.Write(folderB)
	ctx.WriteString("/")
	ctx.WriteString(fileName)

	return nil
}
