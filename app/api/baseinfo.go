package api

import (
	"github.com/kere/gno/libs/util"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/model"
	"onqee.visualstudio.com/D2O/app/model/baseinfo"
	"onqee.visualstudio.com/D2O/app/page"
)

// BaseInfo class
type BaseInfo struct {
}

// NewBaseInfo func
func NewBaseInfo() *BaseInfo {
	return &BaseInfo{}
}

// Auth page auth
// if require is true then do auth
func (a *BaseInfo) Auth(ctx *fasthttp.RequestCtx) error {
	return nil
}

// IsLogin user
func (a *BaseInfo) IsLogin(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	err := page.Auth(ctx)
	if err != nil {
		return 0, nil
	}
	return 1, nil
}

// DoUserLogin user
func (a *BaseInfo) DoUserLogin(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	nick := args.String(model.FieldNick)
	err := page.DoLogin(ctx, nick, args.Bytes("src"), args.Bytes("sign"))
	if err != nil {
		return util.MapData{"value": 0, "message": err.Error()}, nil
	}

	return util.MapData{"value": 1, "message": "success"}, nil
}

// Base func
func (a *BaseInfo) Base(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	// tags := tag.All(0)
	// fields := formfield.All(0)

	return util.MapData{"formfields": baseinfo.FormFields, "areas": baseinfo.Areas, "_data_version": baseinfo.Version}, nil
}
