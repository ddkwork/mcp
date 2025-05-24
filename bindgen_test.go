package main

import (
	_ "embed"
	"github.com/ddkwork/encoding/jsontree"
	"github.com/ddkwork/ux"
	"testing"
)

func TestName(t *testing.T) {
	//fakeError.Walk(".")
	//jsonData := mylog.Check2(os.ReadFile("2.h.json"))
	//bind("2.h", jsonData)

	bind("2.h", runClangASTDump("2.h"))

}

func TestWalk(t *testing.T) {
	//t.Skip("耗时，缓存ast，不建议使用")
	Walk()
}

//go:embed pluginsdk/_dbgfunctions.h.json
var productInfo []byte

func TestJsonTree(t *testing.T) {
	ux.Run("jsonTree", jsontree.Layout(productInfo))
}
