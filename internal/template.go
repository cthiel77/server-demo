package internal

import (
	"bytes"
	"html/template"

	defaultError "github.com/cthiel77/server-demo/internal/error"
)

// RenderTemplate renders a template with injected data
func RenderTemplate(tplContent string, data any) (s string, err *defaultError.Error) {
	var buf bytes.Buffer
	t, e := template.New("page_content").Parse(tplContent)
	if e != nil {
		err = defaultError.FromGoError("1700227900", defaultError.StatusNotAcceptableData, e)
		return
	}
	t.Execute(&buf, data)
	s = buf.String()
	return
}
