package main

import (
	"github.com/ddkwork/golibrary/fakeError"
	"github.com/ddkwork/golibrary/mylog"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	fakeError.Walk(".")
	jsonData := mylog.Check2(os.ReadFile("2.json"))
	bind("2.h", jsonData)
}
