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

//go:embed template/go_template/service.go.tpl
var serviceTplContent string

func init() {
	core.Context.TargetResolver = append(core.Context.TargetResolver, &ServiceResolver{})
}

type ServiceResolver struct{}

func (r ServiceResolver) Parse(target *core.TargetStruct) {
	if tmpl, err := template.New(reflect.TypeOf(r).Name()).Parse(serviceTplContent); err == nil {
		targetFile := filepath.Join(core.Context.DstModPath, "service", target.PkgName, target.Methods["TableName"]+"_service.go")
		targetDir := filepath.Dir(targetFile)
		if fileExists(targetFile) {
			return
		}
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
