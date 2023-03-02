// Package stdfunc contains language-agnostic functions for rtcg templates.
package gencommon

import (
	"strings"
	"text/template"

	"github.com/UoY-RoboStar/rtcg/internal/strmanip"
)

// Funcs adds the standard function map to base.
func Funcs(base *template.Template) *template.Template {
	return base.Funcs(template.FuncMap{
		"toUpper":            strings.ToUpper,
		"toLowerUnderscored": strmanip.ToLowerUnderscored,
	})
}