package main

import (
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
