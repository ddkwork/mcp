package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/tidwall/gjson"
)

// ----------------- 完整数据结构定义 -----------------
type (
	EnumMember struct {
		Name          string
		ExplicitValue string
		ComputedValue int
	}

	EnumInfo struct {
		Name    string
		Loc     string
		Members []EnumMember
	}

	StructField struct {
		Name     string
		Type     string
		TypeDecl string
	}

	StructInfo struct {
		Name    string
		Loc     string
		IsImpl  bool
		Fields  []StructField
		Methods []FunctionInfo
	}

	FunctionParam struct {
		Name string
		Type string
	}

	FunctionInfo struct {
		Name         string
		Loc          string
		ReturnType   string
		Params       []FunctionParam
		IsMethod     bool
		ReceiverType string
	}

	Result struct {
		Enums     []EnumInfo
		Structs   []StructInfo
		Functions []FunctionInfo
		Typedefs  map[string]string
	}
)

// ----------------- 完整解析逻辑 -----------------
func TestName(t *testing.T) {
	// 完整测试数据（包含所有节点类型）
	jsonData := mylog.Check2(os.ReadFile("2.json"))
	//jsonData := mylog.Check2(os.ReadFile("D:\\workspace\\workspace\\mcp\\bridgemain.h.json"))

	root := gjson.Parse(string(jsonData))
	results := traverseNode(root)

	// 生成所有代码
	fmt.Println("// Auto-generated code")
	generateAllCode(results)
}

func traverseNode(node gjson.Result) (result Result) {
	result.Typedefs = make(map[string]string)

	var processNode func(gjson.Result)
	processNode = func(n gjson.Result) {
		switch kind := n.Get("kind").String(); kind {
		case "EnumDecl":
			result.Enums = append(result.Enums, parseEnum(n))
		case "RecordDecl":
			if n.Get("tagUsed").String() == "struct" {
				result.Structs = append(result.Structs, parseStruct(n))
			}
		case "FunctionDecl", "CXXMethodDecl":
			result.Functions = append(result.Functions, parseFunction(n))
		case "TypedefDecl":
			name := n.Get("name").String()
			if target := n.Get("type.qualType").String(); target != "" {
				result.Typedefs[name] = target
			}
		}

		n.Get("inner").ForEach(func(_, child gjson.Result) bool {
			processNode(child)
			return true
		})
	}

	processNode(node)
	return
}

func parseEnum(node gjson.Result) EnumInfo {
	info := EnumInfo{
		Name: node.Get("name").String(),
		Loc:  formatLoc(node.Get("loc")),
	}

	nextValue := 0
	node.Get("inner").ForEach(func(_, member gjson.Result) bool {
		if member.Get("kind").String() == "EnumConstantDecl" {
			explicit, computed := resolveEnumValue(member, nextValue)
			if explicit != "" {
				if val, err := parseNumber(explicit); err == nil {
					nextValue = val + 1
				}
			} else {
				nextValue++
			}

			info.Members = append(info.Members, EnumMember{
				Name:          member.Get("name").String(),
				ExplicitValue: explicit,
				ComputedValue: computed,
			})
		}
		return true
	})

	return info
}

func parseStruct(node gjson.Result) StructInfo {
	info := StructInfo{
		Name:   node.Get("name").String(),
		Loc:    formatLoc(node.Get("loc")),
		IsImpl: node.Get("isImplicit").Bool(),
	}

	node.Get("inner").ForEach(func(_, child gjson.Result) bool {
		switch kind := child.Get("kind").String(); kind {
		case "FieldDecl":
			info.Fields = append(info.Fields, StructField{
				Name:     child.Get("name").String(),
				Type:     child.Get("type.qualType").String(),
				TypeDecl: resolveType(child.Get("type")),
			})
		case "CXXMethodDecl":
			fn := parseFunction(child)
			fn.IsMethod = true
			fn.ReceiverType = info.Name
			info.Methods = append(info.Methods, fn)
		}
		return true
	})

	return info
}

func parseFunction(node gjson.Result) FunctionInfo {
	info := FunctionInfo{
		Name:       node.Get("name").String(),
		Loc:        formatLoc(node.Get("loc")),
		ReturnType: resolveType(node.Get("type.resultType")),
	}

	node.Get("inner").ForEach(func(_, param gjson.Result) bool {
		if param.Get("kind").String() == "ParmVarDecl" {
			info.Params = append(info.Params, FunctionParam{
				Name: param.Get("name").String(),
				Type: resolveType(param.Get("type")),
			})
		}
		return true
	})

	return info
}

// ----------------- 完整工具函数 -----------------
func resolveEnumValue(node gjson.Result, defaultVal int) (string, int) {
	var explicit string
	node.Get("inner").ForEach(func(_, child gjson.Result) bool {
		if child.Get("value").Exists() {
			explicit = child.Get("value").String()
			return false
		}
		return true
	})

	if explicit == "" {
		explicit = node.Get("init.value").String()
	}

	if explicit != "" {
		val, err := parseNumber(explicit)
		if err != nil {
			log.Printf("WARN: Invalid value %q, using %d", explicit, defaultVal)
			return explicit, defaultVal
		}
		return explicit, val
	}
	return "", defaultVal
}

func parseNumber(s string) (int, error) {
	base := 10
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
		base = 16
		i, err := strconv.ParseInt(s, base, 32)
		return int(i), err
	} else if strings.HasPrefix(s, "0") && len(s) > 1 {
		base = 8
		i, err := strconv.ParseInt(s, base, 32)
		return int(i), err
	}
	return strconv.Atoi(s)
}

func resolveType(typeNode gjson.Result) string {
	typeMappings := map[string]string{
		"unsigned long": "uint32",
		"char *":        "string",
		"int":           "int32",
	}

	qualType := typeNode.Get("qualType").String()
	if mapped, ok := typeMappings[qualType]; ok {
		return mapped
	}

	// 处理指针类型
	if ptr := typeNode.Get("inner.#(kind==PointerType)"); ptr.Exists() {
		return "*" + resolveType(ptr.Get("type"))
	}

	// 处理嵌套声明
	if decl := typeNode.Get("decl"); decl.Exists() {
		return decl.Get("name").String()
	}

	return qualType
}

func formatLoc(loc gjson.Result) string {
	if !loc.Exists() {
		return ""
	}
	return fmt.Sprintf("%s:%d",
		loc.Get("file").String(),
		loc.Get("line").Int(),
	)
}

// ----------------- 完整代码生成器 -----------------
func generateAllCode(results Result) {
	// 生成枚举
	for _, e := range results.Enums {
		fmt.Printf("type %s int // %s\nconst (\n", e.Name, e.Loc)
		for i, m := range e.Members {
			line := fmt.Sprintf("\t%s", m.Name)
			if i == 0 {
				line += " = iota"
			} else if m.ExplicitValue != "" {
				line += " = " + m.ExplicitValue
			}
			fmt.Printf("%s // %d\n", line, m.ComputedValue)
		}
		fmt.Println(")")
	}

	// 生成结构体
	for _, s := range results.Structs {
		fmt.Printf("// %s (%s)\n", s.Name, s.Loc)
		fmt.Printf("type %s struct {\n", s.Name)
		for _, f := range s.Fields {
			fmt.Printf("\t%s %s // C type: %s\n",
				strings.Title(f.Name),
				f.TypeDecl,
				f.Type)
		}
		fmt.Println("}")

		// 生成方法
		for _, m := range s.Methods {
			fmt.Printf("func (this *%s) %s(", s.Name, m.Name)
			params := make([]string, len(m.Params))
			for i, p := range m.Params {
				params[i] = fmt.Sprintf("%s %s", p.Name, p.Type)
			}
			fmt.Printf("%s) %s {\n\t// TODO: implement\n}\n\n",
				strings.Join(params, ", "),
				m.ReturnType)
		}
	}

	// 生成函数
	for _, f := range results.Functions {
		fmt.Printf("// %s (%s)\n", f.Name, f.Loc)
		fmt.Printf("func %s(", f.Name)
		params := make([]string, len(f.Params))
		for i, p := range f.Params {
			params[i] = fmt.Sprintf("%s %s", p.Name, p.Type)
		}
		fmt.Printf("%s) %s {\n\t// TODO: implement\n}\n\n",
			strings.Join(params, ", "),
			f.ReturnType)
	}
}
