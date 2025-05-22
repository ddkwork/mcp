package main

import (
	"github.com/ddkwork/golibrary/mylog"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	//fakeError.Walk(".")
	jsonData := mylog.Check2(os.ReadFile("2.h.json"))
	bind("2.h", jsonData)
}

func TestWalk(t *testing.T) {
	Walk()
}
