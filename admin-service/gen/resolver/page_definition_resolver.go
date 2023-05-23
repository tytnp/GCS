package resolver

import (
	_ "embed"
	"fmt"
	"gcs-gen/core"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

//go:embed template/page_template/definition.js.tpl
var DefJsTplContent string

func init() {
	core.Context.TargetResolver = append(core.Context.TargetResolver, &DefJsResolver{})
}

type DefJsResolver struct{}

func (r DefJsResolver) Parse(target *core.TargetStruct) {
	if tmpl, err := template.New(reflect.TypeOf(r).Name()).Parse(DefJsTplContent); err == nil {
		targetFile := filepath.Join(filepath.Dir(core.Context.DstModPath), "admin-web", "src", "views", target.PkgName, target.Methods["ModelName"], "definition.js")
		if fileExists(targetFile) {
			return
		}
		targetDir := filepath.Dir(targetFile)
		if e := os.MkdirAll(targetDir, 0755); e != nil {
			fmt.Printf("Error creating subdirectory: %v\n", e)
			os.Exit(1)
		}
		outputFile, err := os.Create(targetFile)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
		// 写文件
		defJsTplData := getTmplData(target)
		tmpl.Execute(outputFile, defJsTplData)
	}
}

type DefJsTplData struct {
	ColumnsMetaData []core.Field
	Columns         []core.Field
}

func getTmplData(target *core.TargetStruct) DefJsTplData {
	columnsMetaData := make([]core.Field, 0)
	columns := make([]core.Field, 0)
	for _, field := range target.Fields {
		if field.AliasName != "" {
			tag := reflect.StructTag(field.Tag)
			field.JsonName = parseColumnFromGormTag(tag.Get("gorm"))
			columnsMetaData = append(columnsMetaData, field)
		}
	}
	for _, field := range target.Fields {
		if field.AliasName != "" {
			tag := reflect.StructTag(field.Tag)
			field.JsonName = tag.Get("json")
			columns = append(columns, field)
		}
	}
	defJsTplData := DefJsTplData{}
	defJsTplData.ColumnsMetaData = columnsMetaData
	defJsTplData.Columns = columns
	return defJsTplData
}

func parseColumnFromGormTag(gormTag string) string {
	parts := strings.Split(gormTag, ";")
	for _, part := range parts {
		kv := strings.Split(part, ":")
		if len(kv) == 2 && kv[0] == "column" {
			return kv[1]
		}
	}
	return ""
}
