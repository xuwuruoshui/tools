package main

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"os"
	"strings"
	"text/template"
	"unicode"
)

const (
	BizModelInputPath = "./template/end/model.txt"
	BizModelOutPath   = "../end/model/"

	BizRouterInputPath = "./template/end/router.txt"
	BizRouterOutPath   = "../end/api/router/"

	BizHandlerInputPath = "./template/end/handler.txt"
	BizHandlerOutPath   = "../end/api/handler/"

	BizServiceInputPath  = "./template/end/service.txt"
	BizServiceOutputPath = "../end/service/"

	BizRepositoryInputPath  = "./template/end/repository.txt"
	BizRepositoryOutputPath = "../end/repository/"

	Model      = ".go"
	Router     = "_router.go"
	Handler    = "_handler.go"
	Service    = "_service.go"
	Repository = "_repository.go"
)

type Biz struct {
	MethodName  string       // 方法名
	ValName     string       // 变量名
	FileName    string       // 文件名
	TableName    string       // 表名
	Description string       // 变量名
	Conditions  []*Condition // 条件查询参数
}

type Condition struct {
	Name string // 名字
	LowerName string // 小写
	Type string // 类型
	Tags string //标签
}

func NewBiz(fileInputPath, fileSuffix string) *Biz {
	file, _ := os.ReadFile(fileInputPath)
	f, err := decorator.Parse(string(file))
	if err != nil {
		panic(err)
	}

	var biz Biz

	dst.Inspect(f, func(node dst.Node) bool {

		switch tmp := node.(type) {
		case *dst.TypeSpec:
			tmpName := tmp.Name.Name
			biz.MethodName = tmpName
			biz.ValName = strings.ToLower(tmpName[:1]) + tmpName[1:]
			for index, r := range tmpName {
				if unicode.IsUpper(r) && index != 0 {
					biz.TableName += "_" + strings.ToLower(string(r))
				} else {
					biz.TableName += strings.ToLower(string(r))
				}
			}

			biz.FileName = biz.TableName+fileSuffix
		case *dst.StructType:
			for _, field := range tmp.Fields.List {
				field1 := field
				condtion := &Condition{
					Name: field1.Names[0].Name,
					Tags: field1.Tag.Value,
				}
				for _, r := range field1.Names[0].Name {
					if unicode.IsUpper(r){
						condtion.LowerName += "_" + strings.ToLower(string(r))
					} else {
						condtion.LowerName += strings.ToLower(string(r))
					}
				}
				switch tmp2 := field1.Type.(type) {
				case *dst.StarExpr:
					condtion.Type = "*" + tmp2.X.(*dst.Ident).Name
				case *dst.Ident:
					condtion.Type = tmp2.Name
				}

				biz.Conditions = append(biz.Conditions, condtion)
			}
		}
		return true
	})

	return &biz
}

type BizFile struct {
	SuffixName string
	InputPath  string
	OutPath    string
}

// 业务生成
func BizGenerate(modelPath string) {

	// 生成业务
	inputPath := []BizFile{
		{SuffixName: Model, InputPath: BizModelInputPath, OutPath: BizModelOutPath},
		{SuffixName: Router, InputPath: BizRouterInputPath, OutPath: BizRouterOutPath},
		{SuffixName: Handler, InputPath: BizHandlerInputPath, OutPath: BizHandlerOutPath},
		{SuffixName: Service, InputPath: BizServiceInputPath, OutPath: BizServiceOutputPath},
		{SuffixName: Repository, InputPath: BizRepositoryInputPath, OutPath: BizRepositoryOutputPath},
	}

	for _, v := range inputPath {
		biz := NewBiz(modelPath, v.SuffixName)
		file, err := os.OpenFile(v.OutPath+biz.FileName, os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		tmpl, err := template.ParseFiles(v.InputPath)
		if err != nil {
			fmt.Println("create template failed, err:", err)
			return
		}
		// 通过ast获取结构体内的文件类型
		tmpl.Execute(file, &biz)
	}
}
