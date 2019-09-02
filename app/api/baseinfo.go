package api

import (
	"github.com/kere/gno/libs/util"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/model/baseinfo"
	"onqee.visualstudio.com/D2O/app/model/formfield"
	"onqee.visualstudio.com/D2O/app/model/tag"
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

// Base func
func (a *BaseInfo) Base(ctx *fasthttp.RequestCtx, args util.MapData) (interface{}, error) {
	tags := tag.All(0)
	fields := formfield.All(0)

	return util.MapData{"tags": tags, "fields": fields, "areas": baseinfo.Areas, "_data_version": baseinfo.Version}, nil
}
