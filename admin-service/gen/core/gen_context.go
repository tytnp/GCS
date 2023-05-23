package core

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GenContext struct {
	DstModName     string
	DstModPath     string
	SrcFileName    string
	SrcFilePath    string
	TargetStruct   []TargetStruct
	TargetResolver []TargetResolver
}

var Context GenContext

func InitContext() {
	// 获取当前文件路径
	currentFile := os.Getenv("GOFILE")
	//currentFile = "D:\\codebase\\ideaProjects\\GCS\\admin-service\\model\\gpt\\gpt_user.go"
	// 获取包目录
	pkgDir := filepath.Dir(currentFile)
	// 创建 go list 命令以获取模块信息
	goListCmd := exec.Command("go", "list", "-m", "-f", "{{.Path}} {{.Dir}}")
	// 设置 go list 命令的工作目录
	goListCmd.Dir = pkgDir
	// 运行 go list 命令并捕获输出
	modListOutput, err := goListCmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	// 通过 go list 命令输出获取模块名和模块路径等模块信息
	outputFields := strings.Fields(string(modListOutput))
	Context.DstModName, Context.DstModPath = outputFields[0], outputFields[1]
	Context.SrcFileName = currentFile
	Context.SrcFilePath, _ = filepath.Abs(currentFile)

	// 解析 目标结构体
	Context.TargetStruct = GetTargetStruct()
	for i, ts := range Context.TargetStruct {
		fields := GetTargetStructField(ts.Name)
		Context.TargetStruct[i].Fields = fields
	}
}
