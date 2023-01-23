package main

import (
	"fmt"
	"generator/utils"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/token"
	"os"
	"strings"
)

type EnumGenerator struct {
	Name        string
	Description string
	Data        []*EnumItem
}

type EnumItem struct {
	Key   string
	Value string
}

// const生成map
func ConstGenerateMap(inputPath,outputPath string) {

	file, _ := os.ReadFile(inputPath)
	f, err := decorator.Parse(string(file))
	if err != nil {
		panic(err)
	}

	mg := &EnumGenerator{}
	dst.Inspect(f, func(node dst.Node) bool {
		switch tmp := node.(type) {
		case *dst.GenDecl:
			// 获取颜色参数
			for _, data := range tmp.Decs.NodeDecs.Start {
				str := strings.TrimSpace(data)
				annotations := strings.Fields(str)
				if len(annotations) == 3 {
					switch annotations[1] {
					case "@Name":
						if len(annotations) != 3 {
							panic("参数必须为3个")
						}
						mg.Name = annotations[2]
					case "@Description":
						if len(annotations) != 3 {
							panic("参数必须为3个")
						}
						mg.Description = annotations[2]
					}
				}
			}
			tmp.Decs.NodeDecs.After = dst.EmptyLine
		case *dst.ValueSpec:
			fmt.Println()
			key := tmp.Names[0].Name
			value := ""
			values := strings.Fields(tmp.Decs.NodeDecs.End[0])
			if len(values) == 2 {
				value = values[1]
			}
			mg.Data = append(mg.Data, &EnumItem{Key: key, Value: value})
		}
		return true
	})

	f.Decls = append(f.Decls, mg.generate())
	utils.OutFile(f, inputPath, outputPath)

}

func (m *EnumGenerator) generate() *dst.GenDecl {
	var genDecl = &dst.GenDecl{
		Specs: []dst.Spec{
			&dst.ValueSpec{
				Names:  []*dst.Ident{&dst.Ident{}},
				Values: []dst.Expr{&dst.CompositeLit{}},
			},
		},
	}
	genDecl.Tok = token.VAR
	spec := genDecl.Specs[0].(*dst.ValueSpec)
	spec.Names[0].Name = m.Name
	lit := spec.Values[0].(*dst.CompositeLit)
	lit.Type = &dst.MapType{Key: &dst.Ident{Name: "int"}, Value: &dst.Ident{Name: "string"}}

	for _, data := range m.Data {
		tmp := data
		// 赋值
		lit.Elts = append(lit.Elts, &dst.KeyValueExpr{
			Key:   &dst.Ident{Name: tmp.Key},
			Value: &dst.BasicLit{Kind: token.STRING, Value: "\"" + tmp.Value + "\""},
			Decs: dst.KeyValueExprDecorations{
				NodeDecs: dst.NodeDecs{Before: dst.NewLine, After: dst.NewLine},
			},
		})
	}

	return genDecl
}