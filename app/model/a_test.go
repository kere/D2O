package model

import (
	"testing"

	"github.com/kere/gno"
	"github.com/kere/gno/db"
	"onqee.visualstudio.com/D2O/app"
	"onqee.visualstudio.com/D2O/app/model/infos/sele"
)

func init() {
	gno.Init("../app.conf")
}

func TestSEle(t *testing.T) {
	db.Current().Exec("truncate table " + sele.Table)

	datas := make([]sele.DatVO, 2)
	datas[0] = sele.DatVO{Title: "title1", Text: "text1", Lang: app.LangEN}
	datas[1] = sele.DatVO{Title: "title2", Text: "text2", Lang: app.LangCH}

	tags := []int{1, 2, 3}
	reles := []int{10, 20, 30}

	iid, err := sele.Create(1, "2019-01-02", 12345, datas, tags, reles)
	if err != nil {
		t.Fatal(err)
	}

	row, _ := db.NewQuery(sele.Table).Where("iid=?", iid).QueryOne()
	rtags := row.Ints("tags")
	if len(rtags) != 3 {
		t.Fatal(rtags)
	}
	rreles := row.Ints("reles")
	if len(rreles) != 3 {
		t.Fatal(rreles)
	}
}
