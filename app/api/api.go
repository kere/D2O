package api

import (
	"errors"

	"github.com/kere/gno/db"
	"github.com/kere/gno/libs/util"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app"
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
	// fmt.Println(args)

	return util.MapData{"isok": true}, nil
}

// LoadSElem get SElem
func (a *App) LoadSElem(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	iid := args.Int64(app.FieldIID)
	row, err := db.NewQuery(selem.Table).Where("iid=?", iid).QueryOne()
	if err != nil {
		return row, err
	}
	if row.IsEmpty() {
		return row, errors.New("没有找到相应的数据")
	}

	vo := selem.VO{}
	db.Row2VO(row, &vo)
	return vo, nil
}

// SaveSElem 保存SElem
func (a *App) SaveSElem(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	vo := selem.VO{}
	err := util.ConvertA2B(args, &vo)
	if err != nil {
		return 0, err
	}

	iid := args.Int64("iid")
	if iid == 0 {
		vo.IID = app.IID(selem.Table)
		err = db.VOCreate(vo)
	} else {
		err = db.VOUpdate(vo, "iid=?", vo.IID)
	}

	return 1, err
}

// unsafe := blackfriday.Run(input)
// html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
