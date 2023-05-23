package main

import (
	"flag"
	"fmt"
	"gcs-gen/core"
	_ "gcs-gen/resolver"
	"reflect"
)

var hasPage bool

//go:generate go install
func main() {
	flag.BoolVar(&hasPage, "hasPage", true, "has page file")
	flag.Parse()
	fmt.Printf("has page file: %v\n", hasPage)
	core.InitContext()
	for i, _ := range core.Context.TargetStruct {
		for _, r := range core.Context.TargetResolver {
			//r.Parse(&core.Context.TargetStruct[i])
			value := reflect.TypeOf(r).String()
			switch value {
			case "*resolver.DefJsResolver":
				if hasPage {
					r.Parse(&core.Context.TargetStruct[i])
				}
			case "*resolver.IndexResolver":
				if hasPage {
					r.Parse(&core.Context.TargetStruct[i])
				}
			case "*resolver.ListResolver":
				if hasPage {
					r.Parse(&core.Context.TargetStruct[i])
				}
			case "*resolver.EditResolver":
				if hasPage {
					r.Parse(&core.Context.TargetStruct[i])
				}
			default:
				r.Parse(&core.Context.TargetStruct[i])
			}
		}
	}
	//pretty.Println(core.Context.TargetStruct)
}
