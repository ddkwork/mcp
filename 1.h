 enum Color {
     RED,
     GREEN,
     YELLOW =99
 };

 struct Car {
     char make[50];
     int year;
 };

int add (int a, int b){
    return a+b;
};

//clang -Xclang -ast-dump=json -fsyntax-only ./1.h
/*
{
  "id": "0x275917900a0",
  "kind": "TranslationUnitDecl",
  "loc": {},
  "range": {
    "begin": {},
    "end": {}
  },
  "inner": [
    {
      "id": "0x275917908c0",
      "kind": "RecordDecl",
      "loc": {},
      "range": {
        "begin": {},
        "end": {}
      },
      "isImplicit": true,
      "name": "_GUID",
      "tagUsed": "struct",
      "inner": [
        {
          "id": "0x27591790960",
          "kind": "TypeVisibilityAttr",
          "range": {
            "begin": {},
            "end": {}
          },
          "implicit": true
        }
      ]
    },
    {
      "id": "0x275917909d8",
      "kind": "TypedefDecl",
      "loc": {},
      "range": {
        "begin": {},
        "end": {}
      },
      "isImplicit": true,
      "name": "__int128_t",
      "type": {
        "qualType": "__int128"
      },
      "inner": [
        {
          "id": "0x27591790670",
          "kind": "BuiltinType",
          "type": {
            "qualType": "__int128"
          }
        }
      ]
    },
    {
      "id": "0x27591790a48",
      "kind": "TypedefDecl",
      "loc": {},
      "range": {
        "begin": {},
        "end": {}
      },
      "isImplicit": true,
      "name": "__uint128_t",
      "type": {
        "qualType": "unsigned __int128"
      },
      "inner": [
        {
          "id": "0x27591790690",
          "kind": "BuiltinType",
          "type": {
            "qualType": "unsigned __int128"
          }
        }
      ]
    },
    {
      "id": "0x27591790d70",
      "kind": "TypedefDecl",
      "loc": {},
      "range": {
        "begin": {},
        "end": {}
      },
      "isImplicit": true,
      "name": "__NSConstantString",
      "type": {
        "qualType": "struct __NSConstantString_tag"
      },
      "inner": [
        {
          "id": "0x27591790b20",
          "kind": "RecordType",
          "type": {
            "qualType": "struct __NSConstantString_tag"
          },
          "decl": {
            "id": "0x27591790aa0",
            "kind": "RecordDecl",
            "name": "__NSConstantString_tag"
          }
        }
      ]
    },
    {
      "id": "0x27591790de0",
      "kind": "TypedefDecl",
      "loc": {},
      "range": {
        "begin": {},
        "end": {}
      },
      "isImplicit": true,
      "name": "size_t",
      "type": {
        "qualType": "unsigned long long"
      },
      "inner": [
        {
          "id": "0x27591790290",
          "kind": "BuiltinType",
          "type": {
            "qualType": "unsigned long long"
          }
        }
      ]
    },
    {
      "id": "0x27591790e88",
      "kind": "TypedefDecl",
      "loc": {},
      "range": {
        "begin": {},
        "end": {}
      },
      "isImplicit": true,
      "name": "__builtin_ms_va_list",
      "type": {
        "qualType": "char *"
      },
      "inner": [
        {
          "id": "0x27591790e40",
          "kind": "PointerType",
          "type": {
            "qualType": "char *"
          },
          "inner": [
            {
              "id": "0x27591790150",
              "kind": "BuiltinType",
              "type": {
                "qualType": "char"
              }
            }
          ]
        }
      ]
    },
    {
      "id": "0x27591790ef8",
      "kind": "TypedefDecl",
      "loc": {},
      "range": {
        "begin": {},
        "end": {}
      },
      "isImplicit": true,
      "name": "__builtin_va_list",
      "type": {
        "qualType": "char *"
      },
      "inner": [
        {
          "id": "0x27591790e40",
          "kind": "PointerType",
          "type": {
            "qualType": "char *"
          },
          "inner": [
            {
              "id": "0x27591790150",
              "kind": "BuiltinType",
              "type": {
                "qualType": "char"
              }
            }
          ]
        }
      ]
    },
    {
      "id": "0x27591790f50",
      "kind": "EnumDecl",
      "loc": {
        "offset": 6,
        "file": "./1.h",
        "line": 1,
        "col": 7,
        "tokLen": 5
      },
      "range": {
        "begin": {
          "offset": 1,
          "col": 2,
          "tokLen": 4
        },
        "end": {
          "offset": 57,
          "line": 5,
          "col": 2,
          "tokLen": 1
        }
      },
      "name": "Color",
      "inner": [
        {
          "id": "0x27591791010",
          "kind": "EnumConstantDecl",
          "loc": {
            "offset": 20,
            "line": 2,
            "col": 6,
            "tokLen": 3
          },
          "range": {
            "begin": {
              "offset": 20,
              "col": 6,
              "tokLen": 3
            },
            "end": {
              "offset": 20,
              "col": 6,
              "tokLen": 3
            }
          },
          "name": "RED",
          "type": {
            "qualType": "int"
          }
        },
        {
          "id": "0x275917da148",
          "kind": "EnumConstantDecl",
          "loc": {
            "offset": 31,
            "line": 3,
            "col": 6,
            "tokLen": 5
          },
          "range": {
            "begin": {
              "offset": 31,
              "col": 6,
              "tokLen": 5
            },
            "end": {
              "offset": 31,
              "col": 6,
              "tokLen": 5
            }
          },
          "name": "GREEN",
          "type": {
            "qualType": "int"
          }
        },
        {
          "id": "0x275917da1e8",
          "kind": "EnumConstantDecl",
          "loc": {
            "offset": 44,
            "line": 4,
            "col": 6,
            "tokLen": 6
          },
          "range": {
            "begin": {
              "offset": 44,
              "col": 6,
              "tokLen": 6
            },
            "end": {
              "offset": 52,
              "col": 14,
              "tokLen": 2
            }
          },
          "name": "YELLOW",
          "type": {
            "qualType": "int"
          },
          "inner": [
            {
              "id": "0x275917da1c8",
              "kind": "ConstantExpr",
              "range": {
                "begin": {
                  "offset": 52,
                  "col": 14,
                  "tokLen": 2
                },
                "end": {
                  "offset": 52,
                  "col": 14,
                  "tokLen": 2
                }
              },
              "type": {
                "qualType": "int"
              },
              "valueCategory": "prvalue",
              "value": "99",
              "inner": [
                {
                  "id": "0x275917da1a0",
                  "kind": "IntegerLiteral",
                  "range": {
                    "begin": {
                      "offset": 52,
                      "col": 14,
                      "tokLen": 2
                    },
                    "end": {
                      "offset": 52,
                      "col": 14,
                      "tokLen": 2
                    }
                  },
                  "type": {
                    "qualType": "int"
                  },
                  "valueCategory": "prvalue",
                  "value": "99"
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "id": "0x275917da240",
      "kind": "RecordDecl",
      "loc": {
        "offset": 71,
        "line": 7,
        "col": 9,
        "tokLen": 3
      },
      "range": {
        "begin": {
          "offset": 64,
          "col": 2,
          "tokLen": 6
        },
        "end": {
          "offset": 116,
          "line": 10,
          "col": 2,
          "tokLen": 1
        }
      },
      "name": "Car",
      "tagUsed": "struct",
      "completeDefinition": true,
      "inner": [
        {
          "id": "0x275917da390",
          "kind": "FieldDecl",
          "loc": {
            "offset": 88,
            "line": 8,
            "col": 11,
            "tokLen": 4
          },
          "range": {
            "begin": {
              "offset": 83,
              "col": 6,
              "tokLen": 4
            },
            "end": {
              "offset": 95,
              "col": 18,
              "tokLen": 1
            }
          },
          "name": "make",
          "type": {
            "qualType": "char[50]"
          }
        },
        {
          "id": "0x275917da400",
          "kind": "FieldDecl",
          "loc": {
            "offset": 108,
            "line": 9,
            "col": 10,
            "tokLen": 4
          },
          "range": {
            "begin": {
              "offset": 104,
              "col": 6,
              "tokLen": 3
            },
            "end": {
              "offset": 108,
              "col": 10,
              "tokLen": 4
            }
          },
          "name": "year",
          "type": {
            "qualType": "int"
          }
        }
      ]
    },
    {
      "id": "0x275917da5e8",
      "kind": "FunctionDecl",
      "loc": {
        "offset": 126,
        "line": 12,
        "col": 5,
        "tokLen": 3
      },
      "range": {
        "begin": {
          "offset": 122,
          "col": 1,
          "tokLen": 3
        },
        "end": {
          "offset": 164,
          "line": 14,
          "col": 1,
          "tokLen": 1
        }
      },
      "name": "add",
      "mangledName": "add",
      "type": {
        "qualType": "int (int, int)"
      },
      "inner": [
        {
          "id": "0x275917da470",
          "kind": "ParmVarDecl",
          "loc": {
            "offset": 135,
            "line": 12,
            "col": 14,
            "tokLen": 1
          },
          "range": {
            "begin": {
              "offset": 131,
              "col": 10,
              "tokLen": 3
            },
            "end": {
              "offset": 135,
              "col": 14,
              "tokLen": 1
            }
          },
          "isUsed": true,
          "name": "a",
          "type": {
            "qualType": "int"
          }
        },
        {
          "id": "0x275917da4f8",
          "kind": "ParmVarDecl",
          "loc": {
            "offset": 142,
            "col": 21,
            "tokLen": 1
          },
          "range": {
            "begin": {
              "offset": 138,
              "col": 17,
              "tokLen": 3
            },
            "end": {
              "offset": 142,
              "col": 21,
              "tokLen": 1
            }
          },
          "isUsed": true,
          "name": "b",
          "type": {
            "qualType": "int"
          }
        },
        {
          "id": "0x275917da788",
          "kind": "CompoundStmt",
          "range": {
            "begin": {
              "offset": 144,
              "col": 23,
              "tokLen": 1
            },
            "end": {
              "offset": 164,
              "line": 14,
              "col": 1,
              "tokLen": 1
            }
          },
          "inner": [
            {
              "id": "0x275917da778",
              "kind": "ReturnStmt",
              "range": {
                "begin": {
                  "offset": 151,
                  "line": 13,
                  "col": 5,
                  "tokLen": 6
                },
                "end": {
                  "offset": 160,
                  "col": 14,
                  "tokLen": 1
                }
              },
              "inner": [
                {
                  "id": "0x275917da758",
                  "kind": "BinaryOperator",
                  "range": {
                    "begin": {
                      "offset": 158,
                      "col": 12,
                      "tokLen": 1
                    },
                    "end": {
                      "offset": 160,
                      "col": 14,
                      "tokLen": 1
                    }
                  },
                  "type": {
                    "qualType": "int"
                  },
                  "valueCategory": "prvalue",
                  "opcode": "+",
                  "inner": [
                    {
                      "id": "0x275917da728",
                      "kind": "ImplicitCastExpr",
                      "range": {
                        "begin": {
                          "offset": 158,
                          "col": 12,
                          "tokLen": 1
                        },
                        "end": {
                          "offset": 158,
                          "col": 12,
                          "tokLen": 1
                        }
                      },
                      "type": {
                        "qualType": "int"
                      },
                      "valueCategory": "prvalue",
                      "castKind": "LValueToRValue",
                      "inner": [
                        {
                          "id": "0x275917da6e8",
                          "kind": "DeclRefExpr",
                          "range": {
                            "begin": {
                              "offset": 158,
                              "col": 12,
                              "tokLen": 1
                            },
                            "end": {
                              "offset": 158,
                              "col": 12,
                              "tokLen": 1
                            }
                          },
                          "type": {
                            "qualType": "int"
                          },
                          "valueCategory": "lvalue",
                          "referencedDecl": {
                            "id": "0x275917da470",
                            "kind": "ParmVarDecl",
                            "name": "a",
                            "type": {
                              "qualType": "int"
                            }
                          }
                        }
                      ]
                    },
                    {
                      "id": "0x275917da740",
                      "kind": "ImplicitCastExpr",
                      "range": {
                        "begin": {
                          "offset": 160,
                          "col": 14,
                          "tokLen": 1
                        },
                        "end": {
                          "offset": 160,
                          "col": 14,
                          "tokLen": 1
                        }
                      },
                      "type": {
                        "qualType": "int"
                      },
                      "valueCategory": "prvalue",
                      "castKind": "LValueToRValue",
                      "inner": [
                        {
                          "id": "0x275917da708",
                          "kind": "DeclRefExpr",
                          "range": {
                            "begin": {
                              "offset": 160,
                              "col": 14,
                              "tokLen": 1
                            },
                            "end": {
                              "offset": 160,
                              "col": 14,
                              "tokLen": 1
                            }
                          },
                          "type": {
                            "qualType": "int"
                          },
                          "valueCategory": "lvalue",
                          "referencedDecl": {
                            "id": "0x275917da4f8",
                            "kind": "ParmVarDecl",
                            "name": "b",
                            "type": {
                              "qualType": "int"
                            }
                          }
                        }
                      ]
                    }
                  ]
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "id": "0x275917da7a0",
      "kind": "EmptyDecl",
      "loc": {
        "offset": 165,
        "line": 14,
        "col": 2,
        "tokLen": 1
      },
      "range": {
        "begin": {
          "offset": 165,
          "col": 2,
          "tokLen": 1
        },
        "end": {
          "offset": 165,
          "col": 2,
          "tokLen": 1
        }
      }
    }
  ]
}
*/