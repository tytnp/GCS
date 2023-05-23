package core

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"reflect"
	"strings"
)

type Field struct {
	Name      string
	AliasName string
	JsonName  string
	Type      string
	Tag       string
}

func GetTargetStructField(targetStruct string) []Field {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, Context.SrcFileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fields := make([]Field, 0)
	packagePaths := make([]string, 0)
	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.TypeSpec:
			if _, isStruct := x.Type.(*ast.StructType); isStruct && x.Name.Name == targetStruct {
				structType := x.Type.(*ast.StructType)
				for _, field := range structType.Fields.List {
					if ident, ok := field.Type.(*ast.Ident); ok {
						if len(field.Names) > 0 {
							fields = append(fields, Field{Name: field.Names[0].Name, Type: ident.Name, Tag: strings.Trim(field.Tag.Value, "`")})
						} else {
							if typeSpec, ok1 := ident.Obj.Decl.(*ast.TypeSpec); ok1 {
								if st, ok2 := typeSpec.Type.(*ast.StructType); ok2 {
									GetChildStructField(st, &fields)
								}
							}
						}
					} else if selectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {
						if !GetSelectorStructField(packagePaths, selectorExpr.Sel.Name, &fields) {
							if len(field.Names) > 0 {
								typeName := ""
								if selectorExprX, selectorExprOk := selectorExpr.X.(*ast.Ident); selectorExprOk {
									typeName = selectorExprX.Name
								}
								typeName = typeName + "." + selectorExpr.Sel.Name
								fields = append(fields, Field{Name: field.Names[0].Name, Type: typeName, Tag: strings.Trim(field.Tag.Value, "`")})
							}
						}
					} else if arrayType, ok := field.Type.(*ast.ArrayType); ok {
						typeName := "[]" + arrayType.Elt.(*ast.Ident).Name
						fields = append(fields, Field{Name: field.Names[0].Name, Type: typeName, Tag: strings.Trim(field.Tag.Value, "`")})
					}
				}
			}
		case *ast.ImportSpec:
			packagePaths = append(packagePaths, strings.Trim(x.Path.Value, "\""))
		}
		return true
	})
	for index, field := range fields {
		AliasName := getAliasName(field.Tag)
		if field.Type == "admin-service/model/base.Date" {
			fields[index].Type = "time.Time"
		}
		if AliasName != "" {
			fields[index].AliasName = AliasName
		}
		fmt.Printf("Field Name: %s | Field Type: %s | Field Tag: %s | Field AliasName: %s\n", fields[index].Name, fields[index].Type, fields[index].Tag, fields[index].AliasName)
		//fmt.Println("json:", reflect.StructTag(field.Tag).Get("json"))
	}
	return fields
}

func GetChildStructField(structType *ast.StructType, fields *[]Field) {
	for _, field := range structType.Fields.List {
		if ident, ok := field.Type.(*ast.Ident); ok {
			if len(field.Names) > 0 {
				*fields = append(*fields, Field{Name: field.Names[0].Name, Type: ident.Name, Tag: strings.Trim(field.Tag.Value, "`")})
			} else {
				if typeSpec, ok1 := ident.Obj.Decl.(*ast.TypeSpec); ok1 {
					if st, ok2 := typeSpec.Type.(*ast.StructType); ok2 {
						GetChildStructField(st, fields)
					}
				}
			}
		}
	}
}

func GetSelectorStructField(packagePaths []string, typeName string, fields *[]Field) bool {
	conf := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo}
	var typeNameObject types.Object
	for _, packagePath := range packagePaths {
		pkgs, err := packages.Load(conf, packagePath)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(pkgs) == 0 {
			fmt.Println("Not found package:", packagePath)
			continue
		}
		for _, pkg := range pkgs {
			for _, def := range pkg.TypesInfo.Defs {
				if def != nil && def.Name() == typeName {
					typeNameObject = def
					break
				}
			}
		}
	}
	if typeNameObject != nil {
		structType, ok := typeNameObject.Type().Underlying().(*types.Struct)
		if ok {
			for i := 0; i < structType.NumFields(); i++ {
				field := structType.Field(i)
				tag := structType.Tag(i)
				*fields = append(*fields, Field{Name: field.Name(), Type: field.Type().String(), Tag: tag})
			}
		} else {
			return false
		}
	}
	return true
}

func getAliasName(tagValue string) string {
	tag := reflect.StructTag(tagValue)
	aliasNameValue := tag.Get("AliasName")
	return aliasNameValue
}
