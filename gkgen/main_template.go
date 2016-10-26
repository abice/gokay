package gkgen

const (
	masterTemplate = `{{- define "header"}}
	package {{.package}}

	import (
	"github.com/zencoder/gokay"
	)
	{{end -}}

	{{- define "struct"}}
	{{$ruleMap := .rules -}}
	func (s *{{.name}}) Validate () (error) {
	  {{/* Only make the errors map if we're going to use it. */}}
	  {{ if gt (len .rules) 0 }}
	  errors := make(gokay.ErrorMap)
	  {{ end }}
	  {{/* Iterate through the fields of the struct */}}
	  {{range $index, $field := .st.Fields.List -}}
	    {{ $name := (index $field.Names 0).Name -}}
	    {{ $rules := index $ruleMap $name -}}
	    {{ if gt (len $rules) 0 -}}
	    // Begin {{ $name }} Validations
	    {{- range $rIndex, $rule := $rules }}
	      // {{ $rule -}}
	      {{ CallTemplate $rule $field }}
	    {{- end -}}
	    // End {{ $name }} Validations

	    {{ end }}
	  {{- end}}
	  {{/* Only check the errors map if we have rules. */}}
	  {{ if gt (len .rules) 0 }}
	  if len(errors >0) {
	    return errors
	  }
	  {{- end}}
	  return nil
	}
	{{end}}
`
)
