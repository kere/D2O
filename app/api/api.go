package api

import (
	"github.com/kere/gno/db"
	"github.com/kere/gno/libs/util"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/model"
	"onqee.visualstudio.com/D2O/app/model/selem"
	"onqee.visualstudio.com/D2O/app/page"
	"onqee.visualstudio.com/D2O/app/page/home"
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
	return page.Auth(ctx)
}

// LoadSElem get SElem
func (a *App) LoadSElem(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	iid := args.Int64(model.FieldIID)
	v, err := selem.PageData(iid)
	if !v.IsEmpty() {
		v[model.FieldUserID] = 0
	}
	return v, err
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
		userID := ctx.UserValue(model.FieldUserID)
		if userID == nil {
			return nil, model.ErrUserNotFound
		}
		vo.UserID = userID.(int)
		err = db.VOCreate(vo)
	} else {
		err = db.VOUpdate(vo, "iid=?", vo.IID)

		// clear page cache
		home.GetCellView().ClearCache(vo.IID)
	}

	return vo.IID, err
}
