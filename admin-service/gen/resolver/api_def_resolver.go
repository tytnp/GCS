package resolver

import (
	"fmt"
	"gcs-gen/core"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func init() {
	core.Context.TargetResolver = append(core.Context.TargetResolver, &DefinitionResolver{})
}

type DefinitionResolver struct {
	file        *ast.File
	ImportSpecs map[string]string
	target      *core.TargetStruct
}

func (r *DefinitionResolver) Parse(target *core.TargetStruct) {
	r.target = target
	r.ImportSpecs = map[string]string{
		fmt.Sprintf("\"%s/api/base\"", core.Context.DstModName):                   "api",
		fmt.Sprintf("\"%s/model/%s\"", core.Context.DstModName, target.PkgName):   "model",
		fmt.Sprintf("\"%s/service/%s\"", core.Context.DstModName, target.PkgName): "service",
	}
	fset := token.NewFileSet()
	filename := filepath.Join(core.Context.DstModPath, "api", target.PkgName, "definition.go")
	exists := fileExists(filename)
	var err error
	if exists {
		r.file, err = parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}
	} else {
		fileDir := filepath.Dir(filename)
		err := os.MkdirAll(fileDir, 0755)
		if err != nil {
			fmt.Printf("Error creating subdirectory: %v\n", err)
			os.Exit(1)
		}
		r.file = &ast.File{
			Name: ast.NewIdent(target.PkgName),
		}
	}
	//写import
	for k, v := range r.ImportSpecs {
		if r.SearchImportSpec(k) == false {
			r.AddImportSpec(k, v)
		}
	}
	//写struct
	if r.SearchStructType() == false {
		r.AddStructType()
	}
	//写func
	if isFuncExist, isParamExist := r.SearchInitFuncByParam(target.Name + "Api"); isFuncExist == false || isParamExist == false {
		r.AddInitFuncByParam(isFuncExist, isParamExist)
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err1 := format.Node(f, fset, r.file); err1 != nil {
		panic(err1)
	}
}

func (r *DefinitionResolver) SearchInitFuncByParam(funcParam string) (bool, bool) {
	isFuncExist := false
	isParamExist := false
	ast.Inspect(r.file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == "init" {
				isFuncExist = true
				ast.Inspect(x.Body, func(n ast.Node) bool {
					assignStmt, ok := n.(*ast.AssignStmt)
					if !ok {
						return true
					}
					callExpr, ok := assignStmt.Rhs[0].(*ast.CallExpr)
					if !ok {
						return true
					}
					if fun, ok := callExpr.Fun.(*ast.Ident); ok && fun.Name == "append" {
						compositeLit := callExpr.Args[1].(*ast.UnaryExpr).X.(*ast.CompositeLit)
						ident, ok := compositeLit.Type.(*ast.Ident)
						if ok && ident.Name == funcParam {
							isParamExist = true
							//fmt.Println("找到了 base.ApiContext = append(base.ApiContext, &SysApiApi{})")
						}
					}
					return true
				})
				return false
			}
		}
		return true
	})
	return isFuncExist, isParamExist
}

func (r *DefinitionResolver) AddInitFuncByParam(isFuncExist bool, isParamExist bool) {
	if isFuncExist == false {
		r.file.Decls = append(r.file.Decls, &ast.FuncDecl{
			Name: ast.NewIdent("init"),
			Type: &ast.FuncType{},
			Body: &ast.BlockStmt{},
		})
	}
	for _, decl := range r.file.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			if decl.Name.Name == "init" {
				// 第四步：添加到 init 函数中
				decl.Body.List = append(decl.Body.List, &ast.AssignStmt{
					Lhs: []ast.Expr{ast.NewIdent("api.ApiContext")},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: ast.NewIdent("append"),
							Args: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent("api"),
									Sel: ast.NewIdent("ApiContext"),
								},
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: ast.NewIdent(r.target.Name + "Api"),
									},
								},
							},
						},
					},
				})
			}
		}
	}
}

func (r *DefinitionResolver) SearchStructType() bool {
	isExist := false
	ast.Inspect(r.file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if _, ok := x.Type.(*ast.StructType); ok {
				if x.Name.Name == (r.target.Name + "Api") {
					isExist = true
					return false
				}
			}
		}
		return true
	})
	return isExist
}

func (r *DefinitionResolver) AddStructType() {
	r.file.Decls = append(r.file.Decls, &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(r.target.Name + "Api"),
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.IndexExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("api"),
										Sel: ast.NewIdent("Api"),
									},
									Index: &ast.SelectorExpr{
										X:   ast.NewIdent("model"),
										Sel: ast.NewIdent(r.target.Name),
									},
								},
							},
							{
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("service"),
									Sel: ast.NewIdent(r.target.Name + "Service"),
								},
							},
						},
					},
				},
			},
		},
	})
}

func (r *DefinitionResolver) SearchImportSpec(importName string) bool {
	isExist := false
	ast.Inspect(r.file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			if x.Path.Value == importName {
				isExist = true
				return false
			}
		}
		return true
	})
	return isExist
}

func (r *DefinitionResolver) AddImportSpec(importPath string, alias string) {
	var name *ast.Ident
	if alias != "" {
		name = ast.NewIdent(alias)
	}
	r.file.Decls = append(r.file.Decls, &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Name: name,
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: importPath,
				},
			},
		},
	})
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
