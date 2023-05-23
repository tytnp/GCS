package resolver

import (
	_ "embed"
	"fmt"
	"gcs-gen/core"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
)

//go:embed template/page_template/edit.vue.tpl
var EditTplContent string

func init() {
	core.Context.TargetResolver = append(core.Context.TargetResolver, &EditResolver{})
}

type EditResolver struct{}

func (r EditResolver) Parse(target *core.TargetStruct) {
	if tmpl, err := template.New(reflect.TypeOf(r).Name()).Parse(EditTplContent); err == nil {
		targetFile := filepath.Join(filepath.Dir(core.Context.DstModPath), "admin-web", "src", "views", target.PkgName, target.Methods["ModelName"], "edit.vue")
		if fileExists(targetFile) {
			return
		}
		targetDir := filepath.Dir(targetFile)
		err := os.MkdirAll(targetDir, 0755)
		if err != nil {
			fmt.Printf("Error creating subdirectory: %v\n", err)
			os.Exit(1)
		}
		outputFile, err := os.Create(targetFile)
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()
		// 写文件
		editTplData := getEditTplData(target)
		tmpl.Execute(outputFile, editTplData)
	}
}

type EditTplData struct {
	Target  core.TargetStruct
	Columns []core.Field
}

func getEditTplData(target *core.TargetStruct) EditTplData {
	editTplData := EditTplData{}
	editTplData.Target = *target
	columns := make([]core.Field, 0)
	for _, field := range target.Fields {
		if field.AliasName != "" {
			tag := reflect.StructTag(field.Tag)
			field.JsonName = tag.Get("json")
			if field.Name != "CreatedAt" && field.Name != "UpdatedAt" {
				columns = append(columns, field)
			}
		}
	}
	editTplData.Columns = columns
	return editTplData
}
