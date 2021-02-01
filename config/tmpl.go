package config

import (
	"html/template"
)

// TPL stores parsed templates
var TPL *template.Template

var funcMap = template.FuncMap{
	"increment": increment,
}

func increment(index int) int {
	return index + 1
}

func init() {
	TPL = template.Must(template.New("").Funcs(funcMap).ParseGlob("web/template/*.gohtml"))
}
