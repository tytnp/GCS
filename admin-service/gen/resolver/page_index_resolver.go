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

//go:embed template/page_template/index.vue.tpl
var IndexTplContent string

func init() {
	core.Context.TargetResolver = append(core.Context.TargetResolver, &IndexResolver{})
}

type IndexResolver struct{}

func (r IndexResolver) Parse(target *core.TargetStruct) {
	if tmpl, err := template.New(reflect.TypeOf(r).Name()).Parse(IndexTplContent); err == nil {
		targetFile := filepath.Join(filepath.Dir(core.Context.DstModPath), "admin-web", "src", "views", target.PkgName, target.Methods["ModelName"], "index.vue")
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
		tmpl.Execute(outputFile, target)
	}
}
