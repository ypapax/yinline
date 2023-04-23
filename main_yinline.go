package yinline

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func Inline(s string) string {
	return InlineSep(s, "     ")
}

func InlineSep(s, sep string) string {
	ll := strings.Split(s, "\n")
	return strings.Join(ll, sep)
}

func InlineErr(err error) string {
	return InlineSep(fmt.Sprintf("%+v", err), "_______")
}

func InlineStack() string {
	return InlineSep(string(debug.Stack()), " sssssss ")
}

func InlineObj(s interface{}) string {
	return InlineSep(fmt.Sprintf("%+v", s), " oooooooo ")
}
