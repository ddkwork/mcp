package sdk
type TSInputEncoding int // :60
const (
	TSInputEncodingUTF8 TSInputEncoding = iota // 0
	TSInputEncodingUTF16LE // 1
	TSInputEncodingUTF16BE // 2
	TSInputEncodingCustom // 3
)
type TSSymbolType int // :67
const (
	TSSymbolTypeRegular TSSymbolType = iota // 0
	TSSymbolTypeAnonymous // 1
	TSSymbolTypeSupertype // 2
	TSSymbolTypeAuxiliary // 3
)
type TSLogType int // :104
const (
	TSLogTypeParse TSLogType = iota // 0
	TSLogTypeLex // 1
)
type TSQuantifier int // :140
const (
	TSQuantifierZero TSQuantifier = iota // 0
	TSQuantifierZeroOrOne // 1
	TSQuantifierZeroOrMore // 2
	TSQuantifierOne // 3
	TSQuantifierOneOrMore // 4
)
type TSQueryPredicateStepType int // :155
const (
	TSQueryPredicateStepTypeDone TSQueryPredicateStepType = iota // 0
	TSQueryPredicateStepTypeCapture // 1
	TSQueryPredicateStepTypeString // 2
)
type TSQueryError int // :166
const (
	TSQueryErrorNone TSQueryError = iota // 0
	TSQueryErrorSyntax // 1
	TSQueryErrorNodeType // 2
	TSQueryErrorField // 3
	TSQueryErrorCapture // 4
	TSQueryErrorStructure // 5
	TSQueryErrorLanguage // 6
)
type  int // :1377
const (
	TSWasmErrorKindNone  = iota // 0
	TSWasmErrorKindParse // 1
	TSWasmErrorKindCompile // 2
	TSWasmErrorKindInstantiate // 3
	TSWasmErrorKindAllocate // 4
)
