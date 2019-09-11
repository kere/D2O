package page

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/kere/gno/httpd"
	"github.com/kere/gno/libs/util"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"onqee.visualstudio.com/D2O/app/model"
	"onqee.visualstudio.com/D2O/app/model/user"
)

const (
	cookieMaxAge = 3600 * 24 * 30
	// CookieUIID uiid
	CookieUIID = "_uiid"
	// CookieAccess _token
	CookieAccess = "_token"
	// CookieNick _nick
	CookieNick = "_nick"
)

// Auth page
func Auth(ctx *fasthttp.RequestCtx) error {
	nick := ctx.Request.Header.Cookie(CookieNick)
	uiid := ctx.Request.Header.Cookie(CookieUIID)
	val := ctx.Request.Header.Cookie(CookieAccess)
	if len(nick) == 0 {
		return model.ErrLogin
	}

	row := user.LoginUser(util.Bytes2Str(nick))
	if row.IsEmpty() {
		return model.ErrLogin
	}

	token := model.UnDasit(util.Bytes2Str(val))
	arr := bytes.Split(token, httpd.BDote)
	if len(arr) != 2 {
		return model.ErrLogin
	}

	dbToken := row.Bytes(model.FieldToken)

	if util.Bytes2Str(arr[1]) != accessToken(arr[0], uiid, dbToken) {
		return model.ErrLogin
	}

	ctx.SetUserValue(model.FieldUserID, row.Int(model.FieldID))
	ctx.SetUserValue(model.FieldNick, row.String(model.FieldNick))
	return nil
}

// DoLogin user
func DoLogin(ctx *fasthttp.RequestCtx, nick string, srcb, signb []byte) (string, error) {
	ts := ctx.Request.Header.Peek(httpd.APIFieldTS)
	pageToken := ctx.Request.Header.Peek(httpd.APIFieldPageToken)
	uiid := ctx.Request.Header.Cookie(CookieUIID)
	if len(uiid) == 0 {
		return "", model.ErrLogin
	}

	src := model.UnDasit(util.Bytes2Str(srcb))
	sign := model.UnDasit(util.Bytes2Str(signb))

	if len(src) == 0 {
		return "", model.ErrLogin
	}

	// 判断签名是否正确
	signNew := md5.Sum(append(src, ts...))
	if fmt.Sprintf("%x", signNew) != util.Bytes2Str(sign) {
		return "", model.ErrLogin
	}

	row := user.LoginUser(nick)
	if row.IsEmpty() {
		return "", model.ErrUserNotFound
	}

	dbToken := row.Bytes(model.FieldToken)
	tokenNew := authToken(ts, pageToken, uiid, dbToken)
	if tokenNew != util.Bytes2Str(src) {
		return "", model.ErrLogin
	}

	expire := time.Now().AddDate(0, 1, 0)
	// set cookie access token
	accToken := accessToken(ts, uiid, dbToken)
	// println("addcookie:", string(ts), string(uiid), string(dbToken))

	acc := model.Dasit(util.Bytes2Str(ts) + "." + accToken)
	cook := fasthttp.Cookie{}
	cook.SetKey(CookieAccess)
	cook.SetValue(acc)
	// cook.SetMaxAge(cookieMaxAge)
	cook.SetExpire(expire)
	// cook.SetHTTPOnly(true)
	cook.SetPath("/")
	cook.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	ctx.Response.Header.SetCookie(&cook)

	cook2 := fasthttp.Cookie{}
	cook2.SetKey(CookieNick)
	cook2.SetValue(nick)
	// cook2.SetMaxAge(cookieMaxAge)
	cook.SetExpire(expire)
	// cook2.SetHTTPOnly(true)
	cook2.SetPath("/")
	cook2.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	ctx.Response.Header.SetCookie(&cook2)

	return accToken, nil
}

func accessToken(ts, uiid, dbToken []byte) string {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	buf.Write(ts)
	buf.Write(util.Str2Bytes(httpd.Site.SiteData.Secret))
	buf.Write(dbToken)
	buf.Write(ts)
	buf.Write(uiid)

	b := md5.Sum(buf.Bytes())
	return fmt.Sprintf("%x", b)
}

func authToken(ts, pageToken, uiid, md5pwd []byte) string {
	// ts + md5(pwd) + ts + pageToken + uiid + ts
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	buf.Write(ts)
	buf.Write(md5pwd)
	buf.Write(ts)
	buf.Write(pageToken)
	buf.Write(uiid)
	buf.Write(ts)

	b := md5.Sum(buf.Bytes())
	return fmt.Sprintf("%x", b)
}
