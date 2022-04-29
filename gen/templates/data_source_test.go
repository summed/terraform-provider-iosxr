//go:build ignore
{{if .ExcludeTest}}//go:build testAll{{end}}
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxr{{camelCase .Name}}(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if .TestPrerequisites}}testAccDataSourceIosxr{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccDataSourceIosxr{{camelCase .Name}}Config,
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (ne .Id true) (ne .Reference true) (ne .WriteOnly true) (ne .ExcludeTest true)}}
					resource.TestCheckResourceAttr("data.iosxr_{{snakeCase $name}}.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

{{- if .TestPrerequisites}}
const testAccDataSourceIosxr{{camelCase .Name}}PrerequisitesConfig = `
{{- range $index, $item := .TestPrerequisites}}
resource "iosxr_gnmi" "PreReq{{$index}}" {
  path = "{{.Path}}"
  {{- if .NoDelete}}
  delete = false
  {{- end}}
  attributes = {
    {{- range  .Attributes}}
      {{.Name}} = {{if .Reference}}{{.Reference}}{{else}}"{{.Value}}"{{end}}
    {{- end}}
  }
  {{- if .Dependencies}}
  depends_on = [{{range .Dependencies}}iosxr_gnmi.PreReq{{.}}, {{end}}]
  {{- end}}
}
{{ end}}
`
{{- end}}

const testAccDataSourceIosxr{{camelCase .Name}}Config = `
resource "iosxr_{{snakeCase $name}}" "test" {
{{- range  .Attributes}}
{{- if ne .ExcludeTest true}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
{{- if .TestPrerequisites}}
  depends_on = [{{range $index, $item := .TestPrerequisites}}iosxr_gnmi.PreReq{{$index}}, {{end}}]
{{- end}}
}

data "iosxr_{{snakeCase .Name}}" "test" {
{{- range  .Attributes}}
{{- if or (eq .Id true) (eq .Reference true)}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
  depends_on = [iosxr_{{snakeCase $name}}.test]
}
`
