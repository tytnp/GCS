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

//go:embed template/go_template/api.go.tpl
var apiTplContent string

func init() {
	core.Context.TargetResolver = append(core.Context.TargetResolver, &ApiResolver{})
}

type ApiResolver struct{}

func (r ApiResolver) Parse(target *core.TargetStruct) {
	if tmpl, err := template.New(reflect.TypeOf(r).Name()).Parse(apiTplContent); err == nil {
		targetFile := filepath.Join(core.Context.DstModPath, "api", target.PkgName, target.Methods["TableName"]+"_api.go")
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
