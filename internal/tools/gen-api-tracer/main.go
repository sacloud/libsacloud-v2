package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/libsacloud-v2/internal/define"
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/tools"
)

const destination = "sacloud/trace/zz_api_tracer.go"

func init() {
	log.SetFlags(0)
	log.SetPrefix("gen-api-tracer: ")
}

func main() {
	schema.IsOutOfSacloudPackage = true

	tools.WriteFileWithTemplate(&tools.TemplateConfig{
		OutputPath: filepath.Join(tools.ProjectRootPath(), destination),
		Template:   tmpl,
		Parameter:  define.Resources,
	})
	log.Printf("generated: %s\n", filepath.Join(destination))
}

const tmpl = `// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-tracer'; DO NOT EDIT

package sacloud

import (
{{- range .ImportStatements "context" "log" }}
	{{ . }}
{{- end }}
)

{{ range . }} {{ $typeName := .TypeName }}

/************************************************* 
* {{ $typeName }}Tracer
*************************************************/

// {{ $typeName }}Tracer is for trace {{ $typeName }}Op operations
type {{ $typeName }}Tracer struct {
	Internal sacloud.{{$typeName}}API
}

// New{{ $typeName}}Tracer creates new {{ $typeName}}Tracer instance
func New{{ $typeName}}Tracer(in sacloud.{{$typeName}}API) sacloud.{{$typeName}}API {
	return &{{ $typeName}}Tracer {
		Internal: in,
	}
}

{{ range .Operations }}{{$returnErrStatement := .ReturnErrorStatement}}{{ $operationName := .MethodName }}
// {{ .MethodName }} is API call with trace log
func (t *{{ $typeName }}Tracer) {{ .MethodName }}(ctx context.Context{{ range .AllArguments }}, {{ .ArgName }} {{ .TypeName }}{{ end }}) {{.ResultsStatement}} {
	log.Println("[TRACE] {{ $typeName }}Tracer.{{ .MethodName }} start:	args => ["{{ range .AllArguments }},"{{.ArgName}}=", {{ .ArgName }}{{ end }} ,"]")
	defer func() {
		log.Println("[TRACE] {{ $typeName }}Tracer.{{ .MethodName }}: end")
	}()

	return t.Internal.{{ .MethodName }}(ctx{{ range .AllArguments }}, {{ .ArgName }}{{ end }})
}
{{- end -}}

{{ end }}
`
