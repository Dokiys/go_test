package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

// 需要先 go install . 安装到本地
func main() {
	v := "1"
	options := protogen.Options{
		ParamFunc: func(name, value string) error {
			if name == "version" {
				v = value
			}
			return nil
		},
	}

	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			g := gen.NewGeneratedFile("aaa.go", f.GoImportPath)
			g.P("// Code generated by protoc-gen-my. DO NOT EDIT.")
			g.P("// Version: ", v)
			g.P()
			g.P("package ", "gen")
			g.P()
			g.P("// ", protogen.GoImportPath("context").Ident(""))
		}
		return nil
	})
}
