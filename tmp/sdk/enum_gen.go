package sdk
type Color int // :8
const (
	RED Color = iota // 0
	GREEN // 1
	YELLOW = 99 // 99
)
type SEGMENTREG int // :14
const (
	SEG_DEFAULT SEGMENTREG = iota // 0
	SEG_ES // 1
	SEG_DS // 2
	SEG_FS // 3
	SEG_GS // 4
	SEG_CS // 5
	SEG_SS // 6
)
type HardwareType int // :10
const (
	HardwareAccess HardwareType = iota // 0
	HardwareWrite // 1
	HardwareExecute // 2
)
type XEDPARSE_STATUS int // :22
const (
	XEDPARSE_ERROR XEDPARSE_STATUS = iota // 0
	XEDPARSE_OK = 1 // 1
)
type TRACERECORDBYTETYPE int // :59
const (
	InstructionBody TRACERECORDBYTETYPE = iota // 0
	InstructionHeading = 1 // 1
	InstructionTailing = 2 // 2
	InstructionOverlapped = 3 // 3
	DataByte // 4
	DataWord // 5
	DataDWord // 6
	DataQWord // 7
	DataFloat // 8
	DataDouble // 9
	DataLongDouble // 10
	DataXMM // 11
	DataYMM // 12
	DataMMX // 13
	DataMixed // 14
	InstructionDataMixed // 15
)
type TRACERECORDTYPE int // :79
const (
	TraceRecordNone TRACERECORDTYPE = iota // 0
	TraceRecordBitExec // 1
	TraceRecordByteWithExecTypeAndCounter // 2
	TraceRecordWordWithExecTypeAndCounter // 3
)
type MODULESYMBOLSTATUS int // :134
const (
	MODSYMUNLOADED MODULESYMBOLSTATUS = iota // 0
	MODSYMLOADING // 1
	MODSYMLOADED // 2
)
type BP_FIELD int // :141
const (
	bpf_type BP_FIELD = iota // 0
	bpf_offset // 1
	bpf_address // 2
	bpf_enabled // 3
	bpf_singleshoot // 4
	bpf_active // 5
	bpf_silent // 6
	bpf_typeex // 7
	bpf_hwsize // 8
	bpf_hwslot // 9
	bpf_oldbytes // 10
	bpf_fastresume // 11
	bpf_hitcount // 12
	bpf_module // 13
	bpf_name // 14
	bpf_breakcondition // 15
	bpf_logtext // 16
	bpf_logcondition // 17
	bpf_commandtext // 18
	bpf_commandcondition // 19
	bpf_logfile // 20
)
type  int // :39
const (
	JSON_OBJECT  = iota // 0
	JSON_ARRAY // 1
	JSON_STRING // 2
	JSON_INTEGER // 3
	JSON_REAL // 4
	JSON_TRUE // 5
	JSON_FALSE // 6
	JSON_NULL // 7
)
type ValueType int // :239
const (
	ValueTypeNumber ValueType = iota // 0
	ValueTypeString // 1
	ValueTypeAny // 2
	ValueTypeOptionalNumber // 3
	ValueTypeOptionalString // 4
	ValueTypeOptionalAny // 5
)
type CBTYPE int // :264
const (
	CB_INITDEBUG CBTYPE = iota // 0
	CB_STOPDEBUG // 1
	CB_CREATEPROCESS // 2
	CB_EXITPROCESS // 3
	CB_CREATETHREAD // 4
	CB_EXITTHREAD // 5
	CB_SYSTEMBREAKPOINT // 6
	CB_LOADDLL // 7
	CB_UNLOADDLL // 8
	CB_OUTPUTDEBUGSTRING // 9
	CB_EXCEPTION // 10
	CB_BREAKPOINT // 11
	CB_PAUSEDEBUG // 12
	CB_RESUMEDEBUG // 13
	CB_STEPPED // 14
	CB_ATTACH // 15
	CB_DETACH // 16
	CB_DEBUGEVENT // 17
	CB_MENUENTRY // 18
	CB_WINEVENT // 19
	CB_WINEVENTGLOBAL // 20
	CB_LOADDB // 21
	CB_SAVEDB // 22
	CB_FILTERSYMBOL // 23
	CB_TRACEEXECUTE // 24
	CB_SELCHANGED // 25
	CB_ANALYZE // 26
	CB_ADDRINFO // 27
	CB_VALFROMSTRING // 28
	CB_VALTOSTRING // 29
	CB_MENUPREPARE // 30
	CB_STOPPINGDEBUG // 31
	CB_LAST // 32
)
type FORMATRESULT int // :301
const (
	FORMAT_ERROR FORMATRESULT = iota // 0
	FORMAT_SUCCESS // 1
	FORMAT_ERROR_MESSAGE // 2
	FORMAT_BUFFER_TOO_SMALL // 3
)
type FlagEnum int // :10
const (
	ZF FlagEnum = iota // 0
	OF // 1
	CF // 2
	PF // 3
	SF // 4
	TF // 5
	AF // 6
	DF // 7
	IF // 8
)
type Window int // :52
const (
	DisassemblyWindow Window = iota // 0
	DumpWindow // 1
	StackWindow // 2
	GraphWindow // 3
	MemMapWindow // 4
	SymModWindow // 5
)
type RegisterEnum int // :10
const (
	DR0 RegisterEnum = iota // 0
	DR1 // 1
	DR2 // 2
	DR3 // 3
	DR6 // 4
	DR7 // 5
	EAX // 6
	AX // 7
	AH // 8
	AL // 9
	EBX // 10
	BX // 11
	BH // 12
	BL // 13
	ECX // 14
	CX // 15
	CH // 16
	CL // 17
	EDX // 18
	DX // 19
	DH // 20
	DL // 21
	EDI // 22
	DI // 23
	ESI // 24
	SI // 25
	EBP // 26
	BP // 27
	ESP // 28
	SP // 29
	EIP // 30
	RAX // 31
	RBX // 32
	RCX // 33
	RDX // 34
	RSI // 35
	SIL // 36
	RDI // 37
	DIL // 38
	RBP // 39
	BPL // 40
	RSP // 41
	SPL // 42
	RIP // 43
	R8 // 44
	R8D // 45
	R8W // 46
	R8B // 47
	R9 // 48
	R9D // 49
	R9W // 50
	R9B // 51
	R10 // 52
	R10D // 53
	R10W // 54
	R10B // 55
	R11 // 56
	R11D // 57
	R11W // 58
	R11B // 59
	R12 // 60
	R12D // 61
	R12W // 62
	R12B // 63
	R13 // 64
	R13D // 65
	R13W // 66
	R13B // 67
	R14 // 68
	R14D // 69
	R14W // 70
	R14B // 71
	R15 // 72
	R15D // 73
	R15W // 74
	R15B // 75
	CIP // 76
	CSP // 77
	CAX // 78
	CBX // 79
	CCX // 80
	CDX // 81
	CDI // 82
	CSI // 83
	CBP // 84
	CFLAGS // 85
)
type SymbolType int // :10
const (
	Function SymbolType = iota // 0
	Import // 1
	Export // 2
)
type _IMAGEHLP_SYMBOL_TYPE_INFO int // :2679
const (
	TI_GET_SYMTAG _IMAGEHLP_SYMBOL_TYPE_INFO = iota // 0
	TI_GET_SYMNAME // 1
	TI_GET_LENGTH // 2
	TI_GET_TYPE // 3
	TI_GET_TYPEID // 4
	TI_GET_BASETYPE // 5
	TI_GET_ARRAYINDEXTYPEID // 6
	TI_FINDCHILDREN // 7
	TI_GET_DATAKIND // 8
	TI_GET_ADDRESSOFFSET // 9
	TI_GET_OFFSET // 10
	TI_GET_VALUE // 11
	TI_GET_COUNT // 12
	TI_GET_CHILDRENCOUNT // 13
	TI_GET_BITPOSITION // 14
	TI_GET_VIRTUALBASECLASS // 15
	TI_GET_VIRTUALTABLESHAPEID // 16
	TI_GET_VIRTUALBASEPOINTEROFFSET // 17
	TI_GET_CLASSPARENTID // 18
	TI_GET_NESTED // 19
	TI_GET_SYMINDEX // 20
	TI_GET_LEXICALPARENT // 21
	TI_GET_ADDRESS // 22
	TI_GET_THISADJUST // 23
	TI_GET_UDTKIND // 24
	TI_IS_EQUIV_TO // 25
	TI_GET_CALLING_CONVENTION // 26
	TI_IS_CLOSE_EQUIV_TO // 27
	TI_GTIEX_REQS_VALID // 28
	TI_GET_VIRTUALBASEOFFSET // 29
	TI_GET_VIRTUALBASEDISPINDEX // 30
	TI_GET_IS_REFERENCE // 31
	TI_GET_INDIRECTVIRTUALBASECLASS // 32
	IMAGEHLP_SYMBOL_TYPE_INFO_MAX // 33
)
type _MINIDUMP_STREAM_TYPE int // :3683
const (
	UnusedStream _MINIDUMP_STREAM_TYPE = iota // 0
	ReservedStream0 = 1 // 1
	ReservedStream1 = 2 // 2
	ThreadListStream = 3 // 3
	ModuleListStream = 4 // 4
	MemoryListStream = 5 // 5
	ExceptionStream = 6 // 6
	SystemInfoStream = 7 // 7
	ThreadExListStream = 8 // 8
	Memory64ListStream = 9 // 9
	CommentStreamA = 10 // 10
	CommentStreamW = 11 // 11
	HandleDataStream = 12 // 12
	FunctionTableStream = 13 // 13
	UnloadedModuleListStream = 14 // 14
	MiscInfoStream = 15 // 15
	MemoryInfoListStream = 16 // 16
	ThreadInfoListStream = 17 // 17
	HandleOperationListStream = 18 // 18
	TokenStream = 19 // 19
	ceStreamNull = 32768 // 32768
	ceStreamSystemInfo = 32769 // 32769
	ceStreamException = 32770 // 32770
	ceStreamModuleList = 32771 // 32771
	ceStreamProcessList = 32772 // 32772
	ceStreamThreadList = 32773 // 32773
	ceStreamThreadContextList = 32774 // 32774
	ceStreamThreadCallStackList = 32775 // 32775
	ceStreamMemoryVirtualList = 32776 // 32776
	ceStreamMemoryPhysicalList = 32777 // 32777
	ceStreamBucketParameters = 32778 // 32778
	ceStreamProcessModuleMap = 32779 // 32779
	ceStreamDiagnosisList = 32780 // 32780
	LastReservedStream = 65535 // 65535
)
type _MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE int // :4019
const (
	MiniHandleObjectInformationNone _MINIDUMP_HANDLE_OBJECT_INFORMATION_TYPE = iota // 0
	MiniThreadInformation1 // 1
	MiniMutantInformation1 // 2
	MiniMutantInformation2 // 3
	MiniProcessInformation1 // 4
	MiniProcessInformation2 // 5
	MiniHandleObjectInformationTypeMax // 6
)
type _MINIDUMP_CALLBACK_TYPE int // :4319
const (
	ModuleCallback _MINIDUMP_CALLBACK_TYPE = iota // 0
	ThreadCallback // 1
	ThreadExCallback // 2
	IncludeThreadCallback // 3
	IncludeModuleCallback // 4
	MemoryCallback // 5
	CancelCallback // 6
	WriteKernelMinidumpCallback // 7
	KernelMinidumpStatusCallback // 8
	RemoveMemoryCallback // 9
	IncludeVmRegionCallback // 10
	IoStartCallback // 11
	IoWriteAllCallback // 12
	IoFinishCallback // 13
	ReadMemoryFailureCallback // 14
	SecondaryFlagsCallback // 15
)
type _THREAD_WRITE_FLAGS int // :4370
const (
	ThreadWriteThread _THREAD_WRITE_FLAGS = iota // 1
	ThreadWriteStack = 2 // 2
	ThreadWriteContext = 4 // 4
	ThreadWriteBackingStore = 8 // 8
	ThreadWriteInstructionWindow = 16 // 16
	ThreadWriteThreadData = 32 // 32
	ThreadWriteThreadInfo = 64 // 64
)
type _MODULE_WRITE_FLAGS int // :4402
const (
	ModuleWriteModule _MODULE_WRITE_FLAGS = iota // 1
	ModuleWriteDataSeg = 2 // 2
	ModuleWriteMiscRecord = 4 // 4
	ModuleWriteCvRecord = 8 // 8
	ModuleReferencedByMemory = 16 // 16
	ModuleWriteTlsData = 32 // 32
	ModuleWriteCodeSegs = 64 // 64
)
type _MINIDUMP_TYPE int // :4554
const (
	MiniDumpNormal _MINIDUMP_TYPE = iota // 0
	MiniDumpWithDataSegs = 1 // 1
	MiniDumpWithFullMemory = 2 // 2
	MiniDumpWithHandleData = 4 // 4
	MiniDumpFilterMemory = 8 // 8
	MiniDumpScanMemory = 16 // 16
	MiniDumpWithUnloadedModules = 32 // 32
	MiniDumpWithIndirectlyReferencedMemory = 64 // 64
	MiniDumpFilterModulePaths = 128 // 128
	MiniDumpWithProcessThreadData = 256 // 256
	MiniDumpWithPrivateReadWriteMemory = 512 // 512
	MiniDumpWithoutOptionalData = 1024 // 1024
	MiniDumpWithFullMemoryInfo = 2048 // 2048
	MiniDumpWithThreadInfo = 4096 // 4096
	MiniDumpWithCodeSegs = 8192 // 8192
	MiniDumpWithoutAuxiliaryState = 16384 // 16384
	MiniDumpWithFullAuxiliaryState = 32768 // 32768
	MiniDumpWithPrivateWriteCopyMemory = 65536 // 65536
	MiniDumpIgnoreInaccessibleMemory = 131072 // 131072
	MiniDumpWithTokenInformation = 262144 // 262144
	MiniDumpValidTypeFlags = 524287 // 524287
)
type _MINIDUMP_SECONDARY_FLAGS int // :4590
const (
	MiniSecondaryWithoutPowerInfo _MINIDUMP_SECONDARY_FLAGS = iota // 1
	MiniSecondaryValidFlags = 1 // 1
)
type _LZ4_STATUS int // D:\workspace\workspace\mcp\pluginsdk\lz4\lz4file.h:4
const (
	LZ4_SUCCESS _LZ4_STATUS = iota // 0
	LZ4_FAILED_OPEN_INPUT // 1
	LZ4_FAILED_OPEN_OUTPUT // 2
	LZ4_NOT_ENOUGH_MEMORY // 3
	LZ4_INVALID_ARCHIVE // 4
	LZ4_CORRUPTED_ARCHIVE // 5
)
