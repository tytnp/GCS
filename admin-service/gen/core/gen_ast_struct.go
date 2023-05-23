package core

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type TargetStruct struct {
	Name          string //结构体名称
	PkgName       string //结构体包名
	PkgImportPath string //结构体导入路径
	Methods       map[string]string
	Fields        []Field
	GenContext    *GenContext
}

func (t *TargetStruct) GenPkgImportPath() string {
	relativePath, _ := filepath.Rel(Context.DstModPath, filepath.Dir(Context.SrcFilePath))
	relativePath = filepath.Join(Context.DstModName, relativePath)
	relativePath = strings.Replace(relativePath, string(os.PathSeparator), "/", -1)
	t.PkgImportPath = relativePath
	return relativePath
}

func GetTargetStruct() []TargetStruct {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, Context.SrcFileName, nil, 0)
	if err != nil {
		panic(err)
	}
	structNames := make([]string, 0)
	targetStructs := make([]TargetStruct, 0)
	targetFunctions := []string{"TableName", "ModelName"}
	pkgName := f.Name.Name
	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.TypeSpec:
			if _, isStruct := x.Type.(*ast.StructType); isStruct {
				structNames = append(structNames, x.Name.Name)
			}
		case *ast.FuncDecl:
			if x.Recv != nil && len(x.Recv.List) > 0 {
				funcName := x.Name.Name
				recvTypeName := ""
				if r, ok := x.Recv.List[0].Type.(*ast.StarExpr); ok {
					recvTypeName = r.X.(*ast.Ident).Name
				} else if r, ok := x.Recv.List[0].Type.(*ast.Ident); ok {
					recvTypeName = r.Name
				}
				if recvTypeName != "" && contains(structNames, recvTypeName) && contains(targetFunctions, funcName) {
					returnValue := ""
					if retExpr, ok := x.Body.List[0].(*ast.ReturnStmt); ok {
						if lit, ok := retExpr.Results[0].(*ast.BasicLit); ok {

							returnValue = strings.Trim(lit.Value, "\"")
						}
					}
					addTargetMethod(recvTypeName, pkgName, funcName, returnValue, &targetStructs)
				}
			}
		}
		return true
	})

	filteredTargetStructs := make([]TargetStruct, 0)
	for _, targetStruct := range targetStructs {
		if len(targetStruct.Methods) == len(targetFunctions) {
			//fmt.Printf("%s 结构体实现了所有目标方法和它们的返回值:\n", targetStruct.Name)
			//for methodName, returnType := range targetStruct.Methods {
			//	fmt.Printf("  方法名: %s, 返回值: %s\n", methodName, returnType)
			//}
			//fmt.Println()
			targetStruct.GenPkgImportPath()
			targetStruct.GenContext = &Context
			filteredTargetStructs = append(filteredTargetStructs, targetStruct)

		}
	}
	targetStructs = filteredTargetStructs
	return targetStructs
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func addTargetMethod(name, pkgName, methodName, returnValue string, targetStructs *[]TargetStruct) {
	for i, target := range *targetStructs {
		if target.Name == name {
			if _, ok := target.Methods[methodName]; !ok {
				(*targetStructs)[i].Methods[methodName] = returnValue
			}
			return
		}
	}
	*targetStructs = append(*targetStructs, TargetStruct{Name: name, PkgName: pkgName, Methods: map[string]string{methodName: returnValue}})
}
