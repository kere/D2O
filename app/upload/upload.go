package upload

import "github.com/valyala/fasthttp"

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

// Do it
func (m *Image) Do(ctx *fasthttp.RequestCtx) error {
	return nil
}

// StoreDir f
func (m *Image) StoreDir(last []byte) string {
	return "webroot/upload/images"
}
