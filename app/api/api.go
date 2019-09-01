package api

import (
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
