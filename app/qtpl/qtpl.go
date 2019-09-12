package qtpl

import (
	"regexp"

	"github.com/microcosm-cc/bluemonday"
)

var (
	policy  = bluemonday.UGCPolicy()
	linkReg = regexp.MustCompile(`^(?:(\S+)\s*\|\s*)?(http[s]?:\/\/\S+)`)
)

func init() {
	policy.AllowAttrs("data-echo").OnElements("img")
}
