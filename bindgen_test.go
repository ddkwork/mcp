package main

import (
	_ "embed"
	"github.com/ddkwork/encoding/jsontree"
	"github.com/ddkwork/ux"
	"github.com/tidwall/gjson"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	//fakeError.Walk(".")
	//bind("2.h")
	//bind("pluginsdk/bridgemain.h")
}

func TestWalk(t *testing.T) {
	Walk()
}

//go:embed cache/_dbgfunctions.h.json
var productInfo []byte

func TestJsonTree(t *testing.T) {
	ux.Run("jsonTree", jsontree.Layout(productInfo))
}

func TestSkipWindowsSdkDir(t *testing.T) {
	t.Skip("ok")
	raw := gjson.Parse(skipSdkJson).Raw
	if strings.Contains(raw, "Program Files") {
		//if strings.Contains(raw, "C:\\Program Files") {
		panic("skip windows sdk dir")
	}
}

var skipSdkJson = `
{
  "id": "0x17647c05198",
  "kind": "FunctionDecl",
  "loc": {
    "spellingLoc": {
      "offset": 112988,
      "line": 3015,
      "col": 32,
      "tokLen": 23,
      "includedFrom": {
        "file": "C:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared\\minwindef.h"
      }
    },
    "expansionLoc": {
      "offset": 120386,
      "line": 3148,
      "col": 1,
      "tokLen": 22,
      "includedFrom": {
        "file": "C:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared\\minwindef.h"
      }
    }
  },
  "range": {
    "begin": {
      "offset": 120379,
      "line": 3147,
      "col": 1,
      "tokLen": 5,
      "includedFrom": {
        "file": "C:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared\\minwindef.h"
      }
    },
    "end": {
      "offset": 120474,
      "line": 3150,
      "col": 5,
      "tokLen": 1,
      "includedFrom": {
        "file": "C:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared\\minwindef.h"
      }
    }
  },
  "previousDecl": "0x17647c04ff0",
  "name": "_InterlockedIncrement16",
  "mangledName": "_InterlockedIncrement16",
  "type": {
    "qualType": "short (volatile short *)"
  },
  "inner": [
    {
      "id": "0x17647c04ed8",
      "kind": "ParmVarDecl",
      "loc": {
        "offset": 120462,
        "line": 3149,
        "col": 51,
        "tokLen": 6,
        "includedFrom": {
          "file": "C:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared\\minwindef.h"
        }
      },
      "range": {
        "begin": {
          "offset": 120446,
          "col": 35,
          "tokLen": 5,
          "includedFrom": {
            "file": "C:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared\\minwindef.h"
          }
        },
        "end": {
          "offset": 120462,
          "col": 51,
          "tokLen": 6,
          "includedFrom": {
            "file": "C:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared\\minwindef.h"
          }
        }
      },
      "name": "Addend",
      "type": {
        "qualType": "volatile SHORT *"
      }
    }
  ]
}
`
