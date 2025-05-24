package main

import (
	"bytes"
	"fmt"
	"github.com/ddkwork/ddk/vswhere"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/safemap"
	"github.com/ddkwork/golibrary/stream"
	"github.com/tidwall/gjson"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Walk(root, tagetDir, pkgName string, yield func(string) bool) {
	var paths []string
	mylog.Check(filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if !yield(path) {
			return err
		}
		switch filepath.Ext(path) {
		case ".cpp", ".c", ".h":
			paths = append(paths, path)
		}
		return err
	}))
	bind(tagetDir, pkgName, paths...)
}

func bind(tagetDir, pkgName string, paths ...string) {
	result := Result{
		Enums:     new(safemap.M[string, EnumInfo]),
		Structs:   new(safemap.M[string, StructInfo]),
		Functions: new(safemap.M[string, FunctionInfo]),
		Typedefs:  new(safemap.M[string, string]),
	}
	for _, path := range paths {
		root := gjson.Parse(string(runClangASTDump(path)))
		traverseNode(root, &result)
	}
	os.RemoveAll("tmp")
	os.RemoveAll(tagetDir)
	generateAllCode(result, tagetDir, pkgName)
}

var Flags = `
#define _WIN32_WINNT 0x0601
`

func extractFlags() []string {
	f := make([]string, 0)
	for s := range strings.Lines(Flags) {
		switch {
		case !strings.HasPrefix(s, "//"):
			s = strings.TrimSpace(s)
		case strings.HasPrefix(s, "#define"):
			s = strings.TrimSpace(strings.TrimPrefix(s, "#define")) //todo -D
		case strings.HasPrefix(s, "#include"): //todo

		default:
			s = strings.TrimSpace(s)
		}
	}
	return f
}

func runClangASTDump(path string) []byte {
	path = mylog.Check2(filepath.Abs(path))
	jsonPath := filepath.Join("cache", filepath.Base(path)+".json")
	if stream.FileExists(jsonPath) {
		return mylog.Check2(os.ReadFile(jsonPath))
	}

	arg := []string{
		"clang",
		`-x`, `c++`,
		"-Xclang", "-ast-dump=json", "-fsyntax-only", "-fno-builtin", "-nostdinc",
	}
	includes := vswhere.New().VisualStudio().Includes
	for _, include := range includes {
		arg = append(arg, "-I", include)
	}
	arg = append(arg, "-I", filepath.Dir(path))

	//arg = append(arg, extractFlags()...)//todo test flags
	arg = append(arg, path)
	out := stream.RunCommandArgs(arg...)
	mylog.Success(path)
	stream.WriteTruncate(jsonPath, out.Output)
	return out.Output.Bytes()
}

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
		Enums     *safemap.M[string, EnumInfo]
		Structs   *safemap.M[string, StructInfo]
		Functions *safemap.M[string, FunctionInfo]
		Typedefs  *safemap.M[string, string]
	}
)

// todo msvc sal 注解，add callback
func traverseNode(node gjson.Result, result *Result) {
	var processNode func(gjson.Result)
	processNode = func(n gjson.Result) {
		//mylog.Warning(n.Get("inner.0.kind").String())
		//skips := []string{
		//	n.Get("loc.file").String(),
		//	n.Get("loc.includedFrom.file").String(),
		//	n.Get("inner.0.loc.expansionLoc.file").String(),
		//	n.Get("inner.0.loc.expansionLoc.includedFrom.file").String(),
		//	n.Get("inner.0.loc.includedFrom.file").String(),
		//}
		kind := n.Get("kind").String()
		if kind == "BuiltinType" {
			return
		}
		if !strings.Contains(n.Raw, "TranslationUnitDecl") && strings.Contains(n.Raw, "Program Files") {
			return
		}

		s := n.Get("loc.includedFrom.file").String()
		if s != "" {
			//mylog.Warning(filepath.Base(path), filepath.Base(s))
			//if filepath.Base(path) != filepath.Base(s) {
			//	return
			//}
		}

		//if slices.ContainsFunc(skips, func(s string) bool {
		//	return strings.HasPrefix(s, "C:\\Program Files")
		//	if s != "" {
		//		return filepath.Base(path) != filepath.Base(s)
		//	}
		//	return strings.HasPrefix(s, "C:\\Program Files")
		//}) {
		//	return
		//}
		switch kind {
		case "EnumDecl":
			info := parseEnum(n)
			if info.Name == "" {
				info.Name = FindAnonymousName(node, n)
			}
			result.Enums.Update(info.Name, info)
		case "RecordDecl", "CXXRecordDecl":
			if n.Get("name").String() != "_GUID" {
				if n.Get("tagUsed").String() == "struct" {
					object := parseStruct(n)
					if object.Name == "" {
						object.Name = FindAnonymousName(node, n)
					}
					result.Structs.Update(object.Name, object)
				}
			}
		case "FunctionDecl", "CXXMethodDecl":
			function := parseFunction(n)
			if function.Name == "" {
				function.Name = FindAnonymousName(node, n)
			}
			result.Functions.Update(function.Name, function)
		case "TypedefDecl":
			//qualType := n.Get("type.qualType").String() //todo use id get name
			//mylog.Success(qualType)
			//case strings.HasPrefix(qualType, "enum "):
			//	info.Name = strings.TrimPrefix(qualType, "enum ")
			//	result.Enums = append(result.Enums, info)
			//name := n.Get("name").String()
			//if target := n.Get("type.qualType").String(); target != "" {
			//	result.Typedefs[name] = target
			//}
		}

		n.Get("inner").ForEach(func(_, child gjson.Result) bool {
			processNode(child)
			return true
		})
	}

	processNode(node)
}

func FindAnonymousName(root, n gjson.Result) (name string) {
	id := n.Get("id").String()
	root.Get("inner").ForEach(func(_, child gjson.Result) bool { //todo 性能测试，但是这样似乎是准确的,对于处理typedef的匿名写法，通过id找到对应的名称
		//mylog.Json("row", child.Raw)//nice way for debug
		if id == child.Get("inner.0.ownedTagDecl.id").String() {
			name = child.Get("name").String()
			return false //break we can use iter.seq2
		}
		return true
	})
	if name == "" {
		//mylog.Json("row", n.Raw) //nice way for debug
		mylog.Todo("not found Anonymous Name")
		//panic("not found name")
	}
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
				if val, e := parseNumber(explicit); e == nil {
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
		val, e := parseNumber(explicit)
		if e != nil {
			mylog.CheckIgnore(e)
			return "", defaultVal
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
		i, e := strconv.ParseInt(s, base, 32)
		return int(i), e
	} else if strings.HasPrefix(s, "0") && len(s) > 1 {
		base = 8
		i, e := strconv.ParseInt(s, base, 32)
		return int(i), e
	}
	return strconv.Atoi(s)
}

func resolveType(typeNode gjson.Result) string {
	s := strings.NewReplacer(
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
		"UINT64", "uint64",
		"UShort32", "uint16",
		"ULongLong", "uint64",
		"UShort", "uint16",
		"UChar", "byte",
		"UByte", "byte",
		"CHAR", "int8",
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
		"BOOL", "bool",
		"void", "void",
		"PVOID", "uintptr",
		"PCHAR", "int8*",
		"size_t", "uint", //todo test size_t
		//"void *", "uintptr", //todo test uintptr
		//"void", "uintptr", //todo test uintptr
		"const char *", "string",
	).Replace(typeNode.Get("qualType").String()) //todo bind mcp cpp json type check and convert code gen

	if strings.Contains(s, "[") {
		index := strings.Index(s, "[")
		arrayType := s[:index]
		arrayLength := s[index:]
		return arrayLength + arrayType
	}
	return s

	//////////////////////////////////////
	typeMappings := map[string]string{
		"unsigned long": "uint32",
		"char *":        "string",
		"int":           "int32",
		"char[50]":      "[50]int8",
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
func generateAllCode(results Result, tagetDir, pkgName string) {
	buffer := bytes.NewBufferString("")
	buffer.WriteString("package " + pkgName + "\n")

	for _, e := range results.Enums.Range() {
		buffer.WriteString(fmt.Sprintf("type %s int // %s\nconst (\n", e.Name, e.Loc))
		for i, m := range e.Members {
			line := fmt.Sprintf("\t%s", m.Name)
			if i == 0 {
				line += " " + e.Name + " = iota"
			} else if m.ExplicitValue != "" {
				line += " = " + m.ExplicitValue
			}
			buffer.WriteString(fmt.Sprintf("%s // %d\n", line, m.ComputedValue))
		}
		buffer.WriteString(")\n")
	}
	stream.WriteGoFile(filepath.Join(tagetDir, "enum_gen.go"), buffer.String())

	// 生成结构体
	buffer.Reset()
	buffer.WriteString("package " + pkgName + "\n")

	for _, s := range results.Structs.Range() {
		buffer.WriteString(fmt.Sprintf("// %s (%s)\n", s.Name, s.Loc))
		buffer.WriteString(fmt.Sprintf("type %s struct {\n", s.Name))
		for _, f := range s.Fields {
			buffer.WriteString(fmt.Sprintf("\t%s %s // C type: %s\n",
				strings.Title(f.Name),
				f.TypeDecl,
				f.Type))
		}
		buffer.WriteString("}\n")

		// 生成方法
		for _, m := range s.Methods {
			buffer.WriteString(fmt.Sprintf("func (this *%s) %s(", s.Name, m.Name))
			params := make([]string, len(m.Params))
			for i, p := range m.Params {
				params[i] = fmt.Sprintf("%s %s", p.Name, p.Type)
			}
			buffer.WriteString(fmt.Sprintf("%s) %s {\n\t// TODO: implement make json packet\n}\n\n",
				strings.Join(params, ", "),
				m.ReturnType))
		}
	}
	stream.WriteGoFile(filepath.Join(tagetDir, "struct_gen.go"), buffer.String())

	// 生成函数
	buffer.Reset()
	buffer.WriteString("package " + pkgName + "\n")

	for _, f := range results.Functions.Range() {
		buffer.WriteString(fmt.Sprintf("// %s (%s)\n", f.Name, f.Loc))
		buffer.WriteString(fmt.Sprintf("func %s(", f.Name))
		params := make([]string, len(f.Params))
		for i, p := range f.Params {
			params[i] = fmt.Sprintf("%s %s", p.Name, p.Type)
		}
		buffer.WriteString(fmt.Sprintf("%s) %s {\n\t// TODO: implement make json packet\n}\n\n",
			strings.Join(params, ", "),
			f.ReturnType))
	}
	stream.WriteGoFile(filepath.Join(tagetDir, "func_gen.go"), buffer.String())
}
