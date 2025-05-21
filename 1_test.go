package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"iter"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestName(t *testing.T) {
	//f := mylog.Check2(os.Open("bridgemain.h.json"))
	f := mylog.Check2(os.Open("2.json"))
	decoder := json.NewDecoder(f)
	i := 0
	for decoder.More() {
		var node map[string]any
		mylog.Check(decoder.Decode(&node))
		mylog.Call(func() {
			bind(node, "bridgemain.h"+strconv.Itoa(i))
		})
		i++ //2 time, why?
		//mylog.Struct(node)
	}
	//var root map[string]any
	//mylog.Check(json.Unmarshal(mylog.Check2(os.ReadFile("bridgemain.h.json")), &root))
}

func TestClangExecution(t *testing.T) {
	//fakeError.Walk(".")
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".h" {
			if filepath.Base(path) != "bridgemain.h" {
				return nil
			}
			output := runClangASTDump(path)
			rootNode := parseAST(output)
			bind(rootNode, path)

		}
		return err
	})
}

func bind(node map[string]any, path string) {
	info := collectCodeInfo(node)
	//mylog.Struct(info)

	// todo 不合并头文件的话，应该在下面的操作之前map去重一波，最终的输出在第一次发现的文件中?
	g := stream.NewGeneratedFile()
	g.P("package main")
	for _, enum := range info.Enums {
		for j, v := range enum.Values {
			switch j {
			case 0:
				g.P("const (")
				g.P(v.Name, " ", enum.Type, " = iota +"+strconv.Itoa(v.Value))
			default:
				g.P(v.Name, " ", "=", v.Value)
			}
			if j == len(enum.Values)-1 {
				g.P(")")
			}
		}
	}

	for _, object := range info.Structs {
		for j, field := range object.Fields {
			switch j {
			case 0:
				g.P("type ", object.Name, " struct {")
				g.P(field.Name, " ", field.Type)
			default:
				g.P(field.Name, " ", field.Type)
			}
			if j == len(object.Fields)-1 {
				g.P("}")
			}
		}
	}

	for i, function := range info.Functions {
		switch i {
		case 0:
			g.P("func "+
				stream.ToCamelUpper(function.Name),
				function.Params,
				function.ReturnType,
				" {")
		default:
			g.P(function.Name, function.ReturnType, "(", function.Params, ")")
		}
		if i == len(info.Functions)-1 {
			g.P("panic(", strconv.Quote("implement me"), ")")
			g.P("}")
		}
	}
	stream.WriteGoFile(filepath.Join("tmp", filepath.Base(path)+"_gen.go"), g.String())
}

// 数据结构定义
type CodeInfo struct {
	Enums     []Enum
	Structs   []Struct
	Functions []Function
}

type EnumValue struct {
	Name  string
	Value int
}

type Enum struct {
	Name   string
	Type   string
	Values []EnumValue
}

type Struct struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
}

type Function struct {
	Name       string
	ReturnType string
	Params     []Param
}

type Param struct {
	Name string
	Type string
}

var Flags = `
#include <intrin.h>
		"-fno-builtin"
		"-nostdinc"
		//"-DVOID=void"
#define _WIN32_WINNT 0x0601
`

func extractFlags() []string {
	f := make([]string, 0)
	for s := range strings.Lines(Flags) {
		switch {
		case !strings.HasPrefix(s, "//"):
			s = strings.TrimSpace(s)
		case strings.HasPrefix(s, "#define"):
			s = strings.TrimSpace(strings.TrimPrefix(s, "#define"))
		case strings.HasPrefix(s, "#include"): //todo

		default:
			s = strings.TrimSpace(s)
		}
	}
	return f
}

func runClangASTDump(file string) []byte {
	arg := []string{
		"clang-cl",
		"-Xclang",
		"-ast-dump=json",
		"-fsyntax-only",
	}
	// includes := vswhere.New().VisualStudio().Includes
	// for _, include := range includes {
	arg = append(arg, "-I"+"D:\\fork\\fakeWindows\\MiniSDK\\inc\\sdk")
	arg = append(arg, "-I"+"D:\\fork\\fakeWindows\\MiniSDK\\inc\\sdk\\crt")
	// }

	//arg = append(arg, extractFlags()...)//todo test flags
	arg = append(arg, file)
	out := stream.RunCommandArgs(arg...)
	stream.WriteTruncate(file+".json", out.Output)
	return out.Output.Bytes()
}

func parseAST(data []byte) map[string]any {
	var root map[string]any
	mylog.Check(json.Unmarshal(data, &root))
	return root
}

func collectCodeInfo(root map[string]any) CodeInfo {
	var info CodeInfo
	var walk func(node map[string]any)
	walk = func(node map[string]any) {
		for k, v := range node {
			switch k {
			case "kind":
				for s := range asString(v) {
					switch s {
					case "EnumDecl":
						mylog.Call(func() {
							for enum := range parseEnum(node) {
								info.Enums = append(info.Enums, enum)
							}
						})
					case "CXXRecordDecl", "RecordDecl":
						for object := range parseStruct(node) {
							if object.Name != "_GUID" {
								info.Structs = append(info.Structs, object)
							}
						}
					case "FunctionDecl":
						for function := range parseFunction(node) {
							info.Functions = append(info.Functions, function)
						}
					}
				}
			case "inner":
				inner, ok := v.([]any) //asInner(k, v)
				if !ok {
					return
				}
				for _, value := range inner {
					m, ok := value.(map[string]any)
					if !ok {
						return
					}
					walk(m) //asNode(v)
				}
			}
		}
	}
	walk(root)
	return info
}

func hasInner(elem any) ([]any, bool) {
	node, ok := elem.(map[string]any)
	if !ok {
		return nil, false
	}
	v, b := node["inner"]
	return v.([]any), b
}

// Traverse the object recursively and print every path / value pair.
func traverse(path string, obj interface{}) {
	switch obj.(type) {
	case []interface{}:
		for i, subnode := range obj.([]interface{}) {
			traverse(fmt.Sprintf("%s[%d]", path, i), subnode)
		}
	case map[string]interface{}:
		for k, v := range obj.(map[string]interface{}) {
			traverse(fmt.Sprintf("%s[%q]", path, k), v)
		}
	case string:
		fmt.Printf("%s => %q\n", path, obj)
	case float64:
		f := obj.(float64)
		num := int(f)
		if f == float64(num) {
			// the number is an integer
			fmt.Printf("%s => %d\n", path, num)
		} else {
			fmt.Printf("%s => %f\n", path, f)
		}
	default:
		fmt.Printf("%s => %v\n", path, obj)
	}
}

func parseEnum(node map[string]any) iter.Seq[Enum] {
	return func(yield func(Enum) bool) {

		//todo json path, node keys
		//"name": "Color" -->
		//"name": "Color" --> inner --> "name": "RED",  "type": { "qualType": "int"
		//"name": "Color" --> inner --> "name": "YELLOW",  "type": { "qualType": "int"  "inner": [  "value": "99",

		currentValue := 0
		//kind:=node["kind"].(string)

		//

		e := Enum{Name: node["name"].(string), Values: []EnumValue{}}

		for k, v := range asNode(node["inner"]) {
			Type := bindType(node["type"])
			value := EnumValue{
				Name:  node["name"].(string),
				Value: 0,
			}
			valueStr := ""
			inner, b := hasInner(v)
			if b {
				for i, a := range inner {
					v, ok := node["value"]
					if ok {
						value = v.(string)
					}
				}

			}
			e.Values = append(e.Values, value)
		}

		switch kind {
		case "EnumConstantDecl":
			intValue := mylog.Check2(strconv.Atoi(value))
			e.Values = append(e.Values, EnumValue{
				Name:  name,
				Value: intValue,
			})
			currentValue = intValue + 1
		case "ConstantExpr":
			e.Values = append(e.Values, EnumValue{Name: name, Value: currentValue})
			currentValue++
		default:
			panic("not found kind:" + kind)
		}

		for k, v := range node {
			for k, v := range asInner(k, v) {

				name := ""
				switch k {
				case "kind":
					for s := range asString(v) {

					}
				case "name":
					for s := range asString(v) {
						name = s
					}
				case "type":
					e.Type = bindType(v)
				}
				if !yield(e) {
					return
				}
			}
		}
	}
}

func parseStruct(node map[string]any) iter.Seq[Struct] {
	return func(yield func(Struct) bool) {
		name := ""
		for s := range asString(node["name"]) {
			name = s
		}
		for k, v := range node {
			obj := Struct{Name: name}
			switch k {
			//case "name":
			//	for s := range asString(v) {
			//		obj.Name = s
			//	}
			case "inner":
				for k, v := range asInner(k, v) {
					switch k {
					case "kind":
						for s := range asString(v) {
							if s != "FieldDecl" {
								continue //TypeVisibilityAttr
								panic("not found kind:" + s)
							}
						}
					case "name":
						for s := range asString(v) {
							name = s
						}
					case "type":
						obj.Fields = append(obj.Fields, Field{
							Name: name,
							Type: bindType(v),
						})
					}
				}
			}
			if !yield(obj) {
				return
			}
		}
	}
}

func parseFunction(node map[string]any) iter.Seq[Function] {
	return func(yield func(Function) bool) {
		for k, v := range node {
			f := Function{
				Name:       node["name"].(string),
				ReturnType: bindType(node["type"]),
			}
			for k, v := range asInner(k, v) {
				name := ""
				switch k {
				case "kind":
					for s := range asString(v) {
						if s != "ParmVarDecl" {
							continue
							panic("not found kind:" + s) //CompoundStmt todo make unit test
						}
					}
				case "name":
					for s := range asString(v) {
						name = s
					}
				case "type":
					f.Params = append(f.Params, Param{
						Name: name,
						Type: bindType(v),
					})
				}
				if !yield(f) {
					return
				}
			}
		}
	}
}

func asInner(k string, v any) iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		if k == "inner" {
			inner, ok := v.([]any)
			if !ok {
				return
			}
			for _, v := range inner {
				for k, v := range asNode(v) {
					if !yield(k, v) {
						return
					}
				}
			}
		}
	}
}

func asNode(value any) iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		m, ok := value.(map[string]any)
		if !ok {
			return
		}
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func asString(value any) iter.Seq[string] {
	return func(yield func(string) bool) {
		s, ok := value.(string)
		if !ok {
			return
		}
		if !yield(s) {
			return
		}
	}
}

func bindType(value any) string {
	for k, v := range asNode(value) {
		if k == "qualType" {
			for s := range asString(v) {
				return strings.NewReplacer(
					"double", "float64",
					"long double", "float128",
					"int", "int",
					"DWORD", "uint32",
					"DWORD32", "uint32",
					"DWORD64", "uint64",
					"Long", "int32",
					"ULong", "uint32",
					"ULONGLONG", "uint64",
					"ULong32", "uint32",
					"ULong64", "uint64",
					"UShort32", "uint16",
					"ULongLong", "uint64",
					"UShort", "uint16",
					"UChar", "byte",
					"UByte", "byte",
					"unsigned int", "uint",
					"long long", "int64",
					"unsigned long long", "uint64",
					"long", "int32",
					"unsigned long", "uint32",
					"short", "int16",
					"unsigned short", "uint16",
					"char", "int8",
					"unsigned char", "byte",
					"float", "float32",
					"double", "float64",
					"bool", "bool",
					"void", "void",
					"const char *", "string", //todo byte* ?
				).Replace(s) //todo bind mcp cpp json type check and convert code gen
			}
		}
	}
	panic("not found qualType")
}
