package build

var mainGoTemplate = `package main

import (
	"github.com/mholt/caddy/caddy/caddymain"

	{{- range $plugin := .Plugins }}
	_ "{{ $plugin }}"
	{{- end }}
)

func main() {
	{{- if eq .Telemetry false }}
	caddymain.EnableTelemetry = false
	{{- end }}
	caddymain.Run()
}
`
