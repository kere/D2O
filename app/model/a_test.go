package model

import (
	"fmt"
	"testing"

	"github.com/kere/gno"
	"github.com/kere/gno/libs/util"
)

func init() {
	gno.Init("../app.conf")
}

func TestCrypto(t *testing.T) {
	str := "hello world!!"
	ints := Dasit(str)

	b64 := UnDasit(ints)
	if string(b64) != str {
		fmt.Println(b64)
		fmt.Println(util.Str2Bytes(str))
		t.Fatal(string(b64), len(b64), len(str))
	}
}
