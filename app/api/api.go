package api

import (
	"errors"

	"github.com/kere/gno/db"
	"github.com/kere/gno/libs/util"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/model"
	"onqee.visualstudio.com/D2O/app/model/selem"
)

// App class
type App struct{}

// NewApp func
func NewApp() *App {
	return &App{}
}

// Auth page auth
// if require is true then do auth
func (a *App) Auth(ctx *fasthttp.RequestCtx) error {
	return nil
}

// PageData func
func (a *App) PageData(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	return util.MapData{"isok": true}, nil
}

// SElemByIID get SElem
func (a *App) SElemByIID(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	iid := args.Int64(model.FieldIID)
	return selem.PageData(iid)
}

// SElems get SElem list
func (a *App) SElems(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	q := db.NewQuery(selem.Table)
	dat, err := q.Query()
	db.DataSetStrf(dat)
	return dat, err
}

// LoadSElem get SElem
func (a *App) LoadSElem(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	iid := args.Int64(model.FieldIID)
	row, err := db.NewQuery(selem.Table).Where("iid=?", iid).QueryOne()
	if err != nil {
		return row, err
	}

	if row.IsEmpty() {
		return row, errors.New("没有找到相应的数据")
	}

	db.StrictDBMapRow(row, selem.VO{})
	return row, nil
}

// SaveSElem 保存SElem
func (a *App) SaveSElem(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	vo := selem.VO{}
	row := (db.MapRow)(args)
	db.Row2VO(row, &vo)

	iid := args.Int64("iid")
	var err error
	if iid == 0 {
		vo.IID = model.IID(selem.Table)
		err = db.VOCreate(vo)
	} else {
		err = db.VOUpdate(vo, "iid=?", vo.IID)
	}

	return 1, err
}
