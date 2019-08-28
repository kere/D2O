package api

import (
	"encoding/json"

	"github.com/kere/gno/libs/util"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app"
	"onqee.visualstudio.com/D2O/app/model/selem"
)

// App class
type App struct {
}

// NewApp func
func NewApp() *App {
	return &App{}
}

// Auth page auth
// if require is true then do auth
func (a App) Auth(ctx *fasthttp.RequestCtx) error {
	return nil
}

// PageData func
func (a App) PageData(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	// fmt.Println(args)

	return util.MapData{"isok": true}, nil
}

// SaveSElem 保存SElem
func (a App) SaveSElem(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	itype := args.Int("itype")
	var ojson interface{}
	if itype == selem.ITypePerson {
		p := selem.Person{}
		src := util.Str2Bytes(args.String(app.FieldOJSON))
		if err := json.Unmarshal(src, &p); err != nil {
			return nil, err
		}
		ojson = p
	} else {
		p := selem.OJSON{}
		src := util.Str2Bytes(args.String(app.FieldOJSON))
		if err := json.Unmarshal(src, &p); err != nil {
			return nil, err
		}
		ojson = p
	}

	dateON := args.Time(app.FieldDateON)
	miid := args.Int64Default("m_iid", 0)
	tags := args.IntsDefault(app.FieldTags, []int{})
	reles := args.IntsDefault(app.FieldReles, []int{})

	return selem.Create(1, dateON, miid, ojson, tags, reles, itype)
}

// unsafe := blackfriday.Run(input)
// html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
