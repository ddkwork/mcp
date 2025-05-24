package sdk
// ListInfo (D:\workspace\workspace\mcp\pluginsdk\bridgelist.h:4)
type ListInfo struct {
	Count int // C type: int
	Size uint // C type: size_t
	Data void * // C type: void *
}
// Car (:26)
type Car struct {
}
// _CR3_TYPE (:39)
type _CR3_TYPE struct {
}
//  (:137)
type  struct {
	Line int // C type: int
	Column int // C type: int
	Position int // C type: int
	Source [80]int8 // C type: char[80]
	Text [160]int8 // C type: char[160]
}
// PE32Struct (:338)
type PE32Struct struct {
	PE32Offset int // C type: int
	ImageBase int // C type: int
	OriginalEntryPoint int // C type: int
	BaseOfCode int // C type: int
	BaseOfData int // C type: int
	NtSizeOfImage int // C type: int
	NtSizeOfHeaders int // C type: int
	SizeOfOptionalHeaders int // C type: int
	FileAlignment int // C type: int
	SectionAligment int // C type: int
	ImportTableAddress int // C type: int
	ImportTableSize int // C type: int
	ResourceTableAddress int // C type: int
	ResourceTableSize int // C type: int
	ExportTableAddress int // C type: int
	ExportTableSize int // C type: int
	TLSTableAddress int // C type: int
	TLSTableSize int // C type: int
	RelocationTableAddress int // C type: int
	RelocationTableSize int // C type: int
	TimeDateStamp int // C type: int
	SectionNumber int // C type: int
	CheckSum int // C type: int
	SubSystem int // C type: int
	Characteristics int // C type: int
	NumberOfRvaAndSizes int // C type: int
}
// PE64Struct (:368)
type PE64Struct struct {
	PE64Offset int // C type: int
	ImageBase int // C type: int
	OriginalEntryPoint int // C type: int
	BaseOfCode int // C type: int
	BaseOfData int // C type: int
	NtSizeOfImage int // C type: int
	NtSizeOfHeaders int // C type: int
	SizeOfOptionalHeaders int // C type: int
	FileAlignment int // C type: int
	SectionAligment int // C type: int
	ImportTableAddress int // C type: int
	ImportTableSize int // C type: int
	ResourceTableAddress int // C type: int
	ResourceTableSize int // C type: int
	ExportTableAddress int // C type: int
	ExportTableSize int // C type: int
	TLSTableAddress int // C type: int
	TLSTableSize int // C type: int
	RelocationTableAddress int // C type: int
	RelocationTableSize int // C type: int
	TimeDateStamp int // C type: int
	SectionNumber int // C type: int
	CheckSum int // C type: int
	SubSystem int // C type: int
	Characteristics int // C type: int
	NumberOfRvaAndSizes int // C type: int
}
// ImportEnumData (:404)
type ImportEnumData struct {
	NewDll bool // C type: bool
	NumberOfImports int // C type: int
	ImageBase int // C type: int
	BaseImportThunk int // C type: int
	ImportThunk int // C type: int
	APIName int8 * // C type: char *
	DLLName int8 * // C type: char *
}
// THREAD_ITEM_DATA (:415)
type THREAD_ITEM_DATA struct {
	HThread int // C type: int
	DwThreadId int // C type: int
	ThreadStartAddress void * // C type: void *
	ThreadLocalBase void * // C type: void *
	TebAddress void * // C type: void *
	WaitTime int // C type: int
	Priority int // C type: int
	BasePriority int // C type: int
	ContextSwitches int // C type: int
	ThreadState int // C type: int
	WaitReason int // C type: int
}
// LIBRARY_ITEM_DATA (:430)
type LIBRARY_ITEM_DATA struct {
	HFile int // C type: int
	BaseOfDll void * // C type: void *
	HFileMapping int // C type: int
	HFileMappingView void * // C type: void *
	SzLibraryPath int8 // C type: char
	SzLibraryName int8 // C type: char
}
// LIBRARY_ITEM_DATAW (:440)
type LIBRARY_ITEM_DATAW struct {
	HFile int // C type: int
	BaseOfDll void * // C type: void *
	HFileMapping int // C type: int
	HFileMappingView void * // C type: void *
	SzLibraryPath wint8_t // C type: wchar_t
	SzLibraryName wint8_t // C type: wchar_t
}
// PROCESS_ITEM_DATA (:450)
type PROCESS_ITEM_DATA struct {
	HProcess int // C type: int
	DwProcessId int // C type: int
	HThread int // C type: int
	DwThreadId int // C type: int
	HFile int // C type: int
	BaseOfImage void * // C type: void *
	ThreadStartAddress void * // C type: void *
	ThreadLocalBase void * // C type: void *
}
// HandlerArray (:462)
type HandlerArray struct {
	ProcessId int // C type: int
	HHandle int // C type: int
}
// PluginInformation (:468)
type PluginInformation struct {
	PluginName [64]int8 // C type: char[64]
	PluginMajorVersion int // C type: int
	PluginMinorVersion int // C type: int
	PluginBaseAddress int // C type: int
	TitanDebuggingCallBack void * // C type: void *
	TitanRegisterPlugin void * // C type: void *
	TitanReleasePlugin void * // C type: void *
	TitanResetPlugin void * // C type: void *
	PluginDisabled bool // C type: bool
}
// HOOK_ENTRY (:489)
type HOOK_ENTRY struct {
}
// FILE_STATUS_INFO (:533)
type FILE_STATUS_INFO struct {
	OveralEvaluation int // C type: int
	EvaluationTerminatedByException bool // C type: bool
	FileIs64Bit bool // C type: bool
	FileIsDLL bool // C type: bool
	FileIsConsole bool // C type: bool
	MissingDependencies bool // C type: bool
	MissingDeclaredAPIs bool // C type: bool
	SignatureMZ int // C type: int
	SignaturePE int // C type: int
	EntryPoint int // C type: int
	ImageBase int // C type: int
	SizeOfImage int // C type: int
	FileAlignment int // C type: int
	SectionAlignment int // C type: int
	ExportTable int // C type: int
	RelocationTable int // C type: int
	ImportTable int // C type: int
	ImportTableSection int // C type: int
	ImportTableData int // C type: int
	IATTable int // C type: int
	TLSTable int // C type: int
	LoadConfigTable int // C type: int
	BoundImportTable int // C type: int
	COMHeaderTable int // C type: int
	ResourceTable int // C type: int
	ResourceData int // C type: int
	SectionTable int // C type: int
}
// FILE_FIX_INFO (:564)
type FILE_FIX_INFO struct {
	OveralEvaluation int // C type: int
	FixingTerminatedByException bool // C type: bool
	FileFixPerformed bool // C type: bool
	StrippedRelocation bool // C type: bool
	DontFixRelocations bool // C type: bool
	OriginalRelocationTableAddress int // C type: int
	OriginalRelocationTableSize int // C type: int
	StrippedExports bool // C type: bool
	DontFixExports bool // C type: bool
	OriginalExportTableAddress int // C type: int
	OriginalExportTableSize int // C type: int
	StrippedResources bool // C type: bool
	DontFixResources bool // C type: bool
	OriginalResourceTableAddress int // C type: int
	OriginalResourceTableSize int // C type: int
	StrippedTLS bool // C type: bool
	DontFixTLS bool // C type: bool
	OriginalTLSTableAddress int // C type: int
	OriginalTLSTableSize int // C type: int
	StrippedLoadConfig bool // C type: bool
	DontFixLoadConfig bool // C type: bool
	OriginalLoadConfigTableAddress int // C type: int
	OriginalLoadConfigTableSize int // C type: int
	StrippedBoundImports bool // C type: bool
	DontFixBoundImports bool // C type: bool
	OriginalBoundImportTableAddress int // C type: int
	OriginalBoundImportTableSize int // C type: int
	StrippedIAT bool // C type: bool
	DontFixIAT bool // C type: bool
	OriginalImportAddressTableAddress int // C type: int
	OriginalImportAddressTableSize int // C type: int
	StrippedCOM bool // C type: bool
	DontFixCOM bool // C type: bool
	OriginalCOMTableAddress int // C type: int
	OriginalCOMTableSize int // C type: int
}
// DECLSPEC_ALIGN (:782)
type DECLSPEC_ALIGN struct {
}
// YmmRegister_t (:609)
type YmmRegister_t struct {
	Low int // C type: int
	High int // C type: int
}
// x87FPURegister_t (:615)
type x87FPURegister_t struct {
	Data [10]int // C type: int[10]
	St_value int // C type: int
	Tag int // C type: int
}
// x87FPU_t (:622)
type x87FPU_t struct {
	ControlWord int // C type: int
	StatusWord int // C type: int
	TagWord int // C type: int
	ErrorOffset int // C type: int
	ErrorSelector int // C type: int
	DataOffset int // C type: int
	DataSelector int // C type: int
	Cr0NpxState int // C type: int
}
// TITAN_ENGINE_CONTEXT_t (:634)
type TITAN_ENGINE_CONTEXT_t struct {
	Cax int // C type: int
	Ccx int // C type: int
	Cdx int // C type: int
	Cbx int // C type: int
	Csp int // C type: int
	Cbp int // C type: int
	Csi int // C type: int
	Cdi int // C type: int
	R8 int // C type: int
	R9 int // C type: int
	R10 int // C type: int
	R11 int // C type: int
	R12 int // C type: int
	R13 int // C type: int
	R14 int // C type: int
	R15 int // C type: int
	Cip int // C type: int
	Eflags int // C type: int
	Gs uint16 // C type: unsigned short
	Fs uint16 // C type: unsigned short
	Es uint16 // C type: unsigned short
	Ds uint16 // C type: unsigned short
	Cs uint16 // C type: unsigned short
	Ss uint16 // C type: unsigned short
	Dr0 int // C type: int
	Dr1 int // C type: int
	Dr2 int // C type: int
	Dr3 int // C type: int
	Dr6 int // C type: int
	Dr7 int // C type: int
	RegisterArea [80]int // C type: int[80]
	X87fpu x87FPU_t // C type: x87FPU_t
	MxCsr int // C type: int
	XmmRegisters [16]int // C type: int[16]
	YmmRegisters [16]YmmRegister_t // C type: YmmRegister_t[16]
}
// XEDPARSE (:30)
type XEDPARSE struct {
}
// DBGPATCHINFO (D:\workspace\workspace\mcp\pluginsdk\_dbgfunctions.h:10)
type DBGPATCHINFO struct {
	Mod [256]int8 // C type: char[256]
	Addr duint // C type: duint
	Oldbyte byte // C type: unsigned char
	Newbyte byte // C type: unsigned char
}
// DBGCALLSTACKENTRY (:18)
type DBGCALLSTACKENTRY struct {
	Addr duint // C type: duint
	From duint // C type: duint
	To duint // C type: duint
	Comment [512]int8 // C type: char[512]
}
// DBGCALLSTACK (:26)
type DBGCALLSTACK struct {
	Total int // C type: int
	Entries DBGCALLSTACKENTRY * // C type: DBGCALLSTACKENTRY *
}
// DBGSEHRECORD (:32)
type DBGSEHRECORD struct {
	Addr duint // C type: duint
	Handler duint // C type: duint
}
// DBGSEHCHAIN (:38)
type DBGSEHCHAIN struct {
	Total duint // C type: duint
	Records DBGSEHRECORD * // C type: DBGSEHRECORD *
}
// DBGPROCESSINFO (:44)
type DBGPROCESSINFO struct {
	DwProcessId int // C type: int
	SzExeFile int8 // C type: char
	SzExeMainWindowTitle int8 // C type: char
	SzExeArgs [256]int8 // C type: char[256]
}
// DBGRELOCATIONINFO (:52)
type DBGRELOCATIONINFO struct {
	Rva int // C type: int
	Type int // C type: int
	Size int // C type: int
}
// HANDLEINFO (:87)
type HANDLEINFO struct {
	Handle duint // C type: duint
	TypeNumber byte // C type: unsigned char
	GrantedAccess uint // C type: unsigned int
}
// TCPCONNECTIONINFO (:97)
type TCPCONNECTIONINFO struct {
	RemoteAddress [50]int8 // C type: char[50]
	RemotePort uint16 // C type: unsigned short
	LocalAddress [50]int8 // C type: char[50]
	LocalPort uint16 // C type: unsigned short
	StateText [50]int8 // C type: char[50]
	State uint // C type: unsigned int
}
// WINDOW_INFO (:107)
type WINDOW_INFO struct {
	Handle duint // C type: duint
	Parent duint // C type: duint
	ThreadId int // C type: int
	Style int // C type: int
	StyleEx int // C type: int
	WndProc duint // C type: duint
	Enabled bool // C type: bool
	Position int // C type: int
	WindowTitle [512]int8 // C type: char[512]
	WindowClass [512]int8 // C type: char[512]
}
// HEAPINFO (:121)
type HEAPINFO struct {
	Addr duint // C type: duint
	Size duint // C type: duint
	Flags duint // C type: duint
}
// CONSTANTINFO (:128)
type CONSTANTINFO struct {
	Name string // C type: const char *
	Value duint // C type: duint
}
// BP_REF (:169)
type BP_REF struct {
}
// DBGFUNCTIONS_ (:564)
type DBGFUNCTIONS_ struct {
}
// BridgeCFInstruction (D:\workspace\workspace\mcp\pluginsdk\bridgegraph.h:4)
type BridgeCFInstruction struct {
	Addr duint // C type: duint
	Data [15]byte // C type: unsigned char[15]
}
// BridgeCFNodeList (:10)
type BridgeCFNodeList struct {
	ParentGraph duint // C type: duint
	Start duint // C type: duint
	End duint // C type: duint
	Brtrue duint // C type: duint
	Brfalse duint // C type: duint
	Icount duint // C type: duint
	Terminal bool // C type: bool
	Split bool // C type: bool
	Indirectcall bool // C type: bool
	Userdata void * // C type: void *
	Exits ListInfo // C type: ListInfo
	Instrs ListInfo // C type: ListInfo
}
// BridgeCFGraphList (:26)
type BridgeCFGraphList struct {
	EntryPoint duint // C type: duint
	Userdata void * // C type: void *
	Nodes ListInfo // C type: ListInfo
}
// BridgeCFNode (:41)
type BridgeCFNode struct {
}
// BridgeCFGraph (:112)
type BridgeCFGraph struct {
}
// SYMBOLPTR_ (:1004)
type SYMBOLPTR_ struct {
}
// SYMBOLINFO_ (:656)
type SYMBOLINFO_ struct {
}
// SYMBOLINFOCPP (:676)
type SYMBOLINFOCPP struct {
}
// _TYPEDESCRIPTOR (:1339)
type _TYPEDESCRIPTOR struct {
}
// json_t (:51)
type json_t struct {
}
// PLUG_INITSTRUCT (D:\workspace\workspace\mcp\pluginsdk\_plugins.h:32)
type PLUG_INITSTRUCT struct {
	PluginHandle int // C type: int
	SdkVersion int // C type: int
	PluginVersion int // C type: int
	PluginName [256]int8 // C type: char[256]
}
// PLUG_SETUPSTRUCT (:42)
type PLUG_SETUPSTRUCT struct {
	HwndDlg int // C type: int
	HMenu int // C type: int
	HMenuDisasm int // C type: int
	HMenuDump int // C type: int
	HMenuStack int // C type: int
	HMenuGraph int // C type: int
	HMenuMemmap int // C type: int
	HMenuSymmod int // C type: int
}
// PLUG_SCRIPTSTRUCT (:55)
type PLUG_SCRIPTSTRUCT struct {
	Data void * // C type: void *
}
// PLUG_CB_INITDEBUG (:61)
type PLUG_CB_INITDEBUG struct {
	SzFileName string // C type: const char *
}
// PLUG_CB_STOPDEBUG (:66)
type PLUG_CB_STOPDEBUG struct {
	Reserved void * // C type: void *
}
// PLUG_CB_STOPPINGDEBUG (:71)
type PLUG_CB_STOPPINGDEBUG struct {
	Reserved void * // C type: void *
}
// PLUG_CB_CREATEPROCESS (:76)
type PLUG_CB_CREATEPROCESS struct {
	CreateProcessInfo int * // C type: int *
	ModInfo int * // C type: int *
	DebugFileName string // C type: const char *
	FdProcessInfo int * // C type: int *
}
// PLUG_CB_EXITPROCESS (:84)
type PLUG_CB_EXITPROCESS struct {
	ExitProcess int * // C type: int *
}
// PLUG_CB_CREATETHREAD (:89)
type PLUG_CB_CREATETHREAD struct {
	CreateThread int * // C type: int *
	DwThreadId int // C type: int
}
// PLUG_CB_EXITTHREAD (:95)
type PLUG_CB_EXITTHREAD struct {
	ExitThread int * // C type: int *
	DwThreadId int // C type: int
}
// PLUG_CB_SYSTEMBREAKPOINT (:101)
type PLUG_CB_SYSTEMBREAKPOINT struct {
	Reserved void * // C type: void *
}
// PLUG_CB_LOADDLL (:106)
type PLUG_CB_LOADDLL struct {
	LoadDll int * // C type: int *
	ModInfo int * // C type: int *
	Modname string // C type: const char *
}
// PLUG_CB_UNLOADDLL (:113)
type PLUG_CB_UNLOADDLL struct {
	UnloadDll int * // C type: int *
}
// PLUG_CB_OUTPUTDEBUGSTRING (:118)
type PLUG_CB_OUTPUTDEBUGSTRING struct {
	DebugString int * // C type: int *
}
// PLUG_CB_EXCEPTION (:123)
type PLUG_CB_EXCEPTION struct {
	Exception int * // C type: int *
}
// PLUG_CB_BREAKPOINT (:128)
type PLUG_CB_BREAKPOINT struct {
	Breakpoint BRIDGEBP * // C type: BRIDGEBP *
}
// PLUG_CB_PAUSEDEBUG (:133)
type PLUG_CB_PAUSEDEBUG struct {
	Reserved void * // C type: void *
}
// PLUG_CB_RESUMEDEBUG (:138)
type PLUG_CB_RESUMEDEBUG struct {
	Reserved void * // C type: void *
}
// PLUG_CB_STEPPED (:143)
type PLUG_CB_STEPPED struct {
	Reserved void * // C type: void *
}
// PLUG_CB_ATTACH (:148)
type PLUG_CB_ATTACH struct {
	DwProcessId int // C type: int
}
// PLUG_CB_DETACH (:153)
type PLUG_CB_DETACH struct {
	FdProcessInfo int * // C type: int *
}
// PLUG_CB_DEBUGEVENT (:158)
type PLUG_CB_DEBUGEVENT struct {
	DebugEvent int * // C type: int *
}
// PLUG_CB_MENUENTRY (:163)
type PLUG_CB_MENUENTRY struct {
	HEntry int // C type: int
}
// PLUG_CB_WINEVENT (:168)
type PLUG_CB_WINEVENT struct {
	Message int * // C type: int *
	Result int32 * // C type: long *
	Retval bool // C type: bool
}
// PLUG_CB_WINEVENTGLOBAL (:175)
type PLUG_CB_WINEVENTGLOBAL struct {
	Message int * // C type: int *
	Retval bool // C type: bool
}
// PLUG_CB_LOADSAVEDB (:181)
type PLUG_CB_LOADSAVEDB struct {
	Root json_t * // C type: json_t *
	LoadSaveType int // C type: int
}
// PLUG_CB_FILTERSYMBOL (:187)
type PLUG_CB_FILTERSYMBOL struct {
	Symbol string // C type: const char *
	Retval bool // C type: bool
}
// PLUG_CB_TRACEEXECUTE (:193)
type PLUG_CB_TRACEEXECUTE struct {
	Cip duint // C type: duint
	Stop bool // C type: bool
}
// PLUG_CB_SELCHANGED (:199)
type PLUG_CB_SELCHANGED struct {
	HWindow int // C type: int
	VA duint // C type: duint
}
// PLUG_CB_ANALYZE (:205)
type PLUG_CB_ANALYZE struct {
	Graph BridgeCFGraphList // C type: BridgeCFGraphList
}
// PLUG_CB_ADDRINFO (:210)
type PLUG_CB_ADDRINFO struct {
	Addr duint // C type: duint
	Addrinfo BRIDGE_ADDRINFO * // C type: BRIDGE_ADDRINFO *
	Retval bool // C type: bool
}
// PLUG_CB_VALFROMSTRING (:217)
type PLUG_CB_VALFROMSTRING struct {
	String string // C type: const char *
	Value duint // C type: duint
	Value_size int * // C type: int *
	Isvar bool * // C type: bool *
	Hexonly bool * // C type: bool *
	Retval bool // C type: bool
}
// PLUG_CB_VALTOSTRING (:227)
type PLUG_CB_VALTOSTRING struct {
	String string // C type: const char *
	Value duint // C type: duint
	Retval bool // C type: bool
}
// PLUG_CB_MENUPREPARE (:234)
type PLUG_CB_MENUPREPARE struct {
	HMenu GUIMENUTYPE // C type: GUIMENUTYPE
}
// StringValue (:250)
type StringValue struct {
	Ptr string // C type: const char *
	IsOwner bool // C type: bool
}
// ExpressionValue (:256)
type ExpressionValue struct {
	Type ValueType // C type: ValueType
	Number duint // C type: duint
	String StringValue // C type: StringValue
}
// ArgumentInfo (:10)
type ArgumentInfo struct {
}
// BookmarkInfo (:10)
type BookmarkInfo struct {
}
// CommentInfo (:10)
type CommentInfo struct {
}
// FunctionInfo (:10)
type FunctionInfo struct {
}
// LabelInfo (:10)
type LabelInfo struct {
}
// ModuleInfo (:10)
type ModuleInfo struct {
}
// ModuleSectionInfo (:20)
type ModuleSectionInfo struct {
}
// ModuleExport (:27)
type ModuleExport struct {
}
// ModuleImport (:38)
type ModuleImport struct {
}
// SymbolInfo (:17)
type SymbolInfo struct {
}
// _LOADED_IMAGE (:110)
type _LOADED_IMAGE struct {
}
// _MODLOAD_DATA (:527)
type _MODLOAD_DATA struct {
}
// _MODLOAD_CVMISC (:536)
type _MODLOAD_CVMISC struct {
}
// _MODLOAD_PDBGUID_PDBAGE (:546)
type _MODLOAD_PDBGUID_PDBAGE struct {
}
// _tagADDRESS64 (:564)
type _tagADDRESS64 struct {
}
// _KDHELP64 (:618)
type _KDHELP64 struct {
}
// _tagSTACKFRAME64 (:765)
type _tagSTACKFRAME64 struct {
}
// API_VERSION (:909)
type API_VERSION struct {
}
// _IMAGEHLP_SYMBOL64 (:1159)
type _IMAGEHLP_SYMBOL64 struct {
}
// _IMAGEHLP_SYMBOL64_PACKAGE (:1169)
type _IMAGEHLP_SYMBOL64_PACKAGE struct {
}
// _IMAGEHLP_SYMBOLW64 (:1175)
type _IMAGEHLP_SYMBOLW64 struct {
}
// _IMAGEHLP_SYMBOLW64_PACKAGE (:1185)
type _IMAGEHLP_SYMBOLW64_PACKAGE struct {
}
// _IMAGEHLP_MODULE64 (:1242)
type _IMAGEHLP_MODULE64 struct {
}
// _IMAGEHLP_MODULEW64 (:1271)
type _IMAGEHLP_MODULEW64 struct {
}
// _IMAGEHLP_LINE64 (:1339)
type _IMAGEHLP_LINE64 struct {
}
// _IMAGEHLP_LINEW64 (:1348)
type _IMAGEHLP_LINEW64 struct {
}
// _SOURCEFILE (:1384)
type _SOURCEFILE struct {
}
// _SOURCEFILEW (:1390)
type _SOURCEFILEW struct {
}
// _IMAGEHLP_CBA_READ_MEMORY (:1414)
type _IMAGEHLP_CBA_READ_MEMORY struct {
}
// _IMAGEHLP_CBA_EVENT (:1435)
type _IMAGEHLP_CBA_EVENT struct {
}
// _IMAGEHLP_CBA_EVENTW (:1443)
type _IMAGEHLP_CBA_EVENTW struct {
}
// _IMAGEHLP_DEFERRED_SYMBOL_LOAD64 (:1451)
type _IMAGEHLP_DEFERRED_SYMBOL_LOAD64 struct {
}
// _IMAGEHLP_DEFERRED_SYMBOL_LOADW64 (:1463)
type _IMAGEHLP_DEFERRED_SYMBOL_LOADW64 struct {
}
// _IMAGEHLP_DUPLICATE_SYMBOL64 (:1494)
type _IMAGEHLP_DUPLICATE_SYMBOL64 struct {
}
// _OMAP (:1561)
type _OMAP struct {
}
// _SRCCODEINFO (:1842)
type _SRCCODEINFO struct {
}
// _SRCCODEINFOW (:1853)
type _SRCCODEINFOW struct {
}
// _IMAGEHLP_SYMBOL_SRC (:2355)
type _IMAGEHLP_SYMBOL_SRC struct {
}
// _MODULE_TYPE_INFO (:2362)
type _MODULE_TYPE_INFO struct {
}
// _SYMBOL_INFO (:2369)
type _SYMBOL_INFO struct {
}
// _SYMBOL_INFO_PACKAGE (:2388)
type _SYMBOL_INFO_PACKAGE struct {
}
// _SYMBOL_INFOW (:2394)
type _SYMBOL_INFOW struct {
}
// _SYMBOL_INFO_PACKAGEW (:2413)
type _SYMBOL_INFO_PACKAGEW struct {
}
// _IMAGEHLP_STACK_FRAME (:2419)
type _IMAGEHLP_STACK_FRAME struct {
}
// _TI_FINDCHILDREN_PARAMS (:2717)
type _TI_FINDCHILDREN_PARAMS struct {
}
// _IMAGEHLP_GET_TYPE_INFO_PARAMS (:2737)
type _IMAGEHLP_GET_TYPE_INFO_PARAMS struct {
}
// _MINIDUMP_LOCATION_DESCRIPTOR (:3609)
type _MINIDUMP_LOCATION_DESCRIPTOR struct {
}
// _MINIDUMP_LOCATION_DESCRIPTOR64 (:3615)
type _MINIDUMP_LOCATION_DESCRIPTOR64 struct {
}
// _MINIDUMP_MEMORY_DESCRIPTOR (:3622)
type _MINIDUMP_MEMORY_DESCRIPTOR struct {
}
// _MINIDUMP_MEMORY_DESCRIPTOR64 (:3633)
type _MINIDUMP_MEMORY_DESCRIPTOR64 struct {
}
// _MINIDUMP_HEADER (:3640)
type _MINIDUMP_HEADER struct {
}
// _MINIDUMP_DIRECTORY (:3660)
type _MINIDUMP_DIRECTORY struct {
}
// _MINIDUMP_STRING (:3667)
type _MINIDUMP_STRING struct {
}
// _MINIDUMP_SYSTEM_INFO (:3794)
type _MINIDUMP_SYSTEM_INFO struct {
}
// _MINIDUMP_THREAD (:3862)
type _MINIDUMP_THREAD struct {
}
// _MINIDUMP_THREAD_LIST (:3877)
type _MINIDUMP_THREAD_LIST struct {
}
// _MINIDUMP_THREAD_EX (:3884)
type _MINIDUMP_THREAD_EX struct {
}
// _MINIDUMP_THREAD_EX_LIST (:3900)
type _MINIDUMP_THREAD_EX_LIST struct {
}
// _MINIDUMP_EXCEPTION (:3911)
type _MINIDUMP_EXCEPTION struct {
}
// MINIDUMP_EXCEPTION_STREAM (:3930)
type MINIDUMP_EXCEPTION_STREAM struct {
}
// _MINIDUMP_MODULE (:3946)
type _MINIDUMP_MODULE struct {
}
// _MINIDUMP_MODULE_LIST (:3965)
type _MINIDUMP_MODULE_LIST struct {
}
// _MINIDUMP_MEMORY_LIST (:3976)
type _MINIDUMP_MEMORY_LIST struct {
}
// _MINIDUMP_MEMORY64_LIST (:3982)
type _MINIDUMP_MEMORY64_LIST struct {
}
// _MINIDUMP_EXCEPTION_INFORMATION (:3994)
type _MINIDUMP_EXCEPTION_INFORMATION struct {
}
// _MINIDUMP_EXCEPTION_INFORMATION64 (:4001)
type _MINIDUMP_EXCEPTION_INFORMATION64 struct {
}
// _MINIDUMP_HANDLE_OBJECT_INFORMATION (:4030)
type _MINIDUMP_HANDLE_OBJECT_INFORMATION struct {
}
// _MINIDUMP_HANDLE_DESCRIPTOR (:4038)
type _MINIDUMP_HANDLE_DESCRIPTOR struct {
}
// _MINIDUMP_HANDLE_DESCRIPTOR_2 (:4049)
type _MINIDUMP_HANDLE_DESCRIPTOR_2 struct {
}
// _MINIDUMP_HANDLE_DATA_STREAM (:4066)
type _MINIDUMP_HANDLE_DATA_STREAM struct {
}
// _MINIDUMP_HANDLE_OPERATION_LIST (:4080)
type _MINIDUMP_HANDLE_OPERATION_LIST struct {
}
// _MINIDUMP_FUNCTION_TABLE_DESCRIPTOR (:4093)
type _MINIDUMP_FUNCTION_TABLE_DESCRIPTOR struct {
}
// _MINIDUMP_FUNCTION_TABLE_STREAM (:4102)
type _MINIDUMP_FUNCTION_TABLE_STREAM struct {
}
// _MINIDUMP_UNLOADED_MODULE (:4120)
type _MINIDUMP_UNLOADED_MODULE struct {
}
// _MINIDUMP_UNLOADED_MODULE_LIST (:4134)
type _MINIDUMP_UNLOADED_MODULE_LIST struct {
}
// _MINIDUMP_MISC_INFO (:4157)
type _MINIDUMP_MISC_INFO struct {
}
// _MINIDUMP_MISC_INFO_2 (:4167)
type _MINIDUMP_MISC_INFO_2 struct {
}
// _MINIDUMP_MISC_INFO_3 (:4182)
type _MINIDUMP_MISC_INFO_3 struct {
}
// _MINIDUMP_MEMORY_INFO (:4214)
type _MINIDUMP_MEMORY_INFO struct {
}
// _MINIDUMP_MEMORY_INFO_LIST (:4227)
type _MINIDUMP_MEMORY_INFO_LIST struct {
}
// _MINIDUMP_THREAD_INFO (:4250)
type _MINIDUMP_THREAD_INFO struct {
}
// _MINIDUMP_THREAD_INFO_LIST (:4264)
type _MINIDUMP_THREAD_INFO_LIST struct {
}
// _MINIDUMP_TOKEN_INFO_HEADER (:4274)
type _MINIDUMP_TOKEN_INFO_HEADER struct {
}
// _MINIDUMP_TOKEN_INFO_LIST (:4281)
type _MINIDUMP_TOKEN_INFO_LIST struct {
}
// _MINIDUMP_USER_RECORD (:4293)
type _MINIDUMP_USER_RECORD struct {
}
// _MINIDUMP_USER_STREAM (:4300)
type _MINIDUMP_USER_STREAM struct {
}
// _MINIDUMP_USER_STREAM_INFORMATION (:4309)
type _MINIDUMP_USER_STREAM_INFORMATION struct {
}
// _MINIDUMP_THREAD_CALLBACK (:4340)
type _MINIDUMP_THREAD_CALLBACK struct {
}
// _MINIDUMP_THREAD_EX_CALLBACK (:4351)
type _MINIDUMP_THREAD_EX_CALLBACK struct {
}
// _MINIDUMP_INCLUDE_THREAD_CALLBACK (:4364)
type _MINIDUMP_INCLUDE_THREAD_CALLBACK struct {
}
// _MINIDUMP_MODULE_CALLBACK (:4381)
type _MINIDUMP_MODULE_CALLBACK struct {
}
// _MINIDUMP_INCLUDE_MODULE_CALLBACK (:4396)
type _MINIDUMP_INCLUDE_MODULE_CALLBACK struct {
}
// _MINIDUMP_IO_CALLBACK (:4414)
type _MINIDUMP_IO_CALLBACK struct {
}
// _MINIDUMP_READ_MEMORY_FAILURE_CALLBACK (:4423)
type _MINIDUMP_READ_MEMORY_FAILURE_CALLBACK struct {
}
// _MINIDUMP_CALLBACK_INPUT (:4432)
type _MINIDUMP_CALLBACK_INPUT struct {
}
// _MINIDUMP_CALLBACK_OUTPUT (:4451)
type _MINIDUMP_CALLBACK_OUTPUT struct {
}
// _MINIDUMP_CALLBACK_INFORMATION (:4612)
type _MINIDUMP_CALLBACK_INFORMATION struct {
}
