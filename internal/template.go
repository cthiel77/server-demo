package internal

import (
	"bytes"
	"html/template"
)

// RenderTemplate renders a template with injected data
func RenderTemplate(tplContent string, data any) (s string) {
	var buf bytes.Buffer
	t, _ := template.New("page_content").Parse(tplContent)
	t.Execute(&buf, data)
	s = buf.String()
	return
}
