package sdk
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
// ListInfo (D:\workspace\workspace\mcp\pluginsdk\bridgelist.h:4)
type ListInfo struct {
	Count int // C type: int
	Size uint // C type: size_t
	Data void * // C type: void *
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
//  (:1331)
type  struct {
	TitleHwnd void * // C type: void *
	ClassHwnd void * // C type: void *
	Title [512]int8 // C type: char[512]
	ClassName [512]int8 // C type: char[512]
}
// SYMBOLINFO_ (:656)
type SYMBOLINFO_ struct {
}
// SYMBOLINFOCPP (:676)
type SYMBOLINFOCPP struct {
}
// DECLSPEC_ALIGN (:782)
type DECLSPEC_ALIGN struct {
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
