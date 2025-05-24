package sdk
// GetField (:308)
func GetField(field BP_FIELD, value int &)  {
	// TODO: implement make json packet
}

// SetField (:316)
func SetField(field BP_FIELD, value const int &)  {
	// TODO: implement make json packet
}

// operator() (:0)
func operator()(str string, userdata void *)  {
	// TODO: implement make json packet
}

// __invoke (:310)
func __invoke(str string, userdata void *)  {
	// TODO: implement make json packet
}

// BridgeInit (:52)
func BridgeInit()  {
	// TODO: implement make json packet
}

// BridgeStart (:62)
func BridgeStart()  {
	// TODO: implement make json packet
}

// BridgeAlloc (:69)
func BridgeAlloc(size uint)  {
	// TODO: implement make json packet
}

// BridgeFree (:75)
func BridgeFree(ptr void *)  {
	// TODO: implement make json packet
}

// BridgeSettingGet (:84)
func BridgeSettingGet(section string, key string, value int8 *)  {
	// TODO: implement make json packet
}

// BridgeSettingGetUint (:93)
func BridgeSettingGetUint(section string, key string, value duint *)  {
	// TODO: implement make json packet
}

// BridgeSettingSet (:102)
func BridgeSettingSet(section string, key string, value string)  {
	// TODO: implement make json packet
}

// BridgeSettingSetUint (:111)
func BridgeSettingSetUint(section string, key string, value duint)  {
	// TODO: implement make json packet
}

// BridgeSettingFlush (:117)
func BridgeSettingFlush()  {
	// TODO: implement make json packet
}

// BridgeSettingRead (:124)
func BridgeSettingRead(errorLine int *)  {
	// TODO: implement make json packet
}

// BridgeGetDbgVersion (:130)
func BridgeGetDbgVersion()  {
	// TODO: implement make json packet
}

// BridgeIsProcessElevated (:136)
func BridgeIsProcessElevated()  {
	// TODO: implement make json packet
}

// BridgeGetNtBuildNumber (:142)
func BridgeGetNtBuildNumber()  {
	// TODO: implement make json packet
}

// BridgeUserDirectory (:147)
func BridgeUserDirectory()  {
	// TODO: implement make json packet
}

// BridgeIsARM64Emulated (:152)
func BridgeIsARM64Emulated()  {
	// TODO: implement make json packet
}

// Data (:46)
func Data()  {
	// TODO: implement make json packet
}

// Count (:55)
func Count()  {
	// TODO: implement make json packet
}

// Cleanup (:65)
func Cleanup()  {
	// TODO: implement make json packet
}

// operator& (:78)
func operator&()  {
	// TODO: implement make json packet
}

// operator[] (:89)
func operator[](index uint)  {
	// TODO: implement make json packet
}

// CopyData (:102)
func CopyData(listInfo ListInfo *, listData const int &)  {
	// TODO: implement make json packet
}

// Free (:119)
func Free(graphList const BridgeCFGraphList *)  {
	// TODO: implement make json packet
}

// ToVector (:131)
func ToVector(listInfo const ListInfo *, listData int &, freedata bool)  {
	// TODO: implement make json packet
}

// ToNodeList (:93)
func ToNodeList()  {
	// TODO: implement make json packet
}

// __debugbreak (:0)
func __debugbreak()  {
	// TODO: implement make json packet
}

// AddNode (:148)
func AddNode(node const BridgeCFNode &)  {
	// TODO: implement make json packet
}

// AddParent (:155)
func AddParent(child duint, parent duint)  {
	// TODO: implement make json packet
}

// ToGraphList (:169)
func ToGraphList()  {
	// TODO: implement make json packet
}

// operator= (:0)
func operator=( GuiDisableScope &&)  {
	// TODO: implement make json packet
}

// DbgInit (:1011)
func DbgInit()  {
	// TODO: implement make json packet
}

// DbgExit (:1012)
func DbgExit()  {
	// TODO: implement make json packet
}

// DbgMemRead (:1013)
func DbgMemRead(va duint, dest void *, size duint)  {
	// TODO: implement make json packet
}

// DbgMemWrite (:1014)
func DbgMemWrite(va duint, src const void *, size duint)  {
	// TODO: implement make json packet
}

// DbgMemGetPageSize (:1015)
func DbgMemGetPageSize(base duint)  {
	// TODO: implement make json packet
}

// DbgMemFindBaseAddr (:1016)
func DbgMemFindBaseAddr(addr duint, size duint *)  {
	// TODO: implement make json packet
}

// DbgCmdExec (:1027)
func DbgCmdExec(cmd string)  {
	// TODO: implement make json packet
}

// DbgCmdExecDirect (:1037)
func DbgCmdExecDirect(cmd string)  {
	// TODO: implement make json packet
}

// DbgMemMap (:1038)
func DbgMemMap(memmap MEMMAP *)  {
	// TODO: implement make json packet
}

// DbgIsValidExpression (:1039)
func DbgIsValidExpression(expression string)  {
	// TODO: implement make json packet
}

// DbgIsDebugging (:1040)
func DbgIsDebugging()  {
	// TODO: implement make json packet
}

// DbgIsJumpGoingToExecute (:1041)
func DbgIsJumpGoingToExecute(addr duint)  {
	// TODO: implement make json packet
}

// DbgGetLabelAt (:1042)
func DbgGetLabelAt(addr duint, segment SEGMENTREG, text int8 *)  {
	// TODO: implement make json packet
}

// DbgSetLabelAt (:1043)
func DbgSetLabelAt(addr duint, text string)  {
	// TODO: implement make json packet
}

// DbgClearLabelRange (:1044)
func DbgClearLabelRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgGetCommentAt (:1045)
func DbgGetCommentAt(addr duint, text int8 *)  {
	// TODO: implement make json packet
}

// DbgSetCommentAt (:1046)
func DbgSetCommentAt(addr duint, text string)  {
	// TODO: implement make json packet
}

// DbgClearCommentRange (:1047)
func DbgClearCommentRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgGetBookmarkAt (:1048)
func DbgGetBookmarkAt(addr duint)  {
	// TODO: implement make json packet
}

// DbgSetBookmarkAt (:1049)
func DbgSetBookmarkAt(addr duint, isbookmark bool)  {
	// TODO: implement make json packet
}

// DbgClearBookmarkRange (:1050)
func DbgClearBookmarkRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgGetModuleAt (:1051)
func DbgGetModuleAt(addr duint, text int8 *)  {
	// TODO: implement make json packet
}

// DbgGetBpxTypeAt (:1052)
func DbgGetBpxTypeAt(addr duint)  {
	// TODO: implement make json packet
}

// DbgValFromString (:1053)
func DbgValFromString(string string)  {
	// TODO: implement make json packet
}

// DbgGetRegDumpEx (:1054)
func DbgGetRegDumpEx(regdump REGDUMP *, size uint)  {
	// TODO: implement make json packet
}

// DbgValToString (:1055)
func DbgValToString(string string, value duint)  {
	// TODO: implement make json packet
}

// DbgMemIsValidReadPtr (:1056)
func DbgMemIsValidReadPtr(addr duint)  {
	// TODO: implement make json packet
}

// DbgGetBpList (:1057)
func DbgGetBpList(type BPXTYPE, list BPMAP *)  {
	// TODO: implement make json packet
}

// DbgGetFunctionTypeAt (:1058)
func DbgGetFunctionTypeAt(addr duint)  {
	// TODO: implement make json packet
}

// DbgGetLoopTypeAt (:1059)
func DbgGetLoopTypeAt(addr duint, depth int)  {
	// TODO: implement make json packet
}

// DbgGetBranchDestination (:1060)
func DbgGetBranchDestination(addr duint)  {
	// TODO: implement make json packet
}

// DbgScriptLoad (:1061)
func DbgScriptLoad(filename string)  {
	// TODO: implement make json packet
}

// DbgScriptUnload (:1062)
func DbgScriptUnload()  {
	// TODO: implement make json packet
}

// DbgScriptRun (:1063)
func DbgScriptRun(destline int)  {
	// TODO: implement make json packet
}

// DbgScriptStep (:1064)
func DbgScriptStep()  {
	// TODO: implement make json packet
}

// DbgScriptBpToggle (:1065)
func DbgScriptBpToggle(line int)  {
	// TODO: implement make json packet
}

// DbgScriptBpGet (:1066)
func DbgScriptBpGet(line int)  {
	// TODO: implement make json packet
}

// DbgScriptCmdExec (:1067)
func DbgScriptCmdExec(command string)  {
	// TODO: implement make json packet
}

// DbgScriptAbort (:1068)
func DbgScriptAbort()  {
	// TODO: implement make json packet
}

// DbgScriptGetLineType (:1069)
func DbgScriptGetLineType(line int)  {
	// TODO: implement make json packet
}

// DbgScriptSetIp (:1070)
func DbgScriptSetIp(line int)  {
	// TODO: implement make json packet
}

// DbgScriptGetBranchInfo (:1071)
func DbgScriptGetBranchInfo(line int, info SCRIPTBRANCH *)  {
	// TODO: implement make json packet
}

// DbgSymbolEnum (:1072)
func DbgSymbolEnum(base duint, cbSymbolEnum CBSYMBOLENUM, user void *)  {
	// TODO: implement make json packet
}

// DbgSymbolEnumFromCache (:1073)
func DbgSymbolEnumFromCache(base duint, cbSymbolEnum CBSYMBOLENUM, user void *)  {
	// TODO: implement make json packet
}

// DbgSymbolEnumRange (:1074)
func DbgSymbolEnumRange(start duint, end duint, symbolMask uint, cbSymbolEnum CBSYMBOLENUM, user void *)  {
	// TODO: implement make json packet
}

// DbgAssembleAt (:1075)
func DbgAssembleAt(addr duint, instruction string)  {
	// TODO: implement make json packet
}

// DbgModBaseFromName (:1076)
func DbgModBaseFromName(name string)  {
	// TODO: implement make json packet
}

// DbgDisasmAt (:1077)
func DbgDisasmAt(addr duint, instr DISASM_INSTR *)  {
	// TODO: implement make json packet
}

// DbgStackCommentGet (:1078)
func DbgStackCommentGet(addr duint, comment STACK_COMMENT *)  {
	// TODO: implement make json packet
}

// DbgGetThreadList (:1079)
func DbgGetThreadList(list THREADLIST *)  {
	// TODO: implement make json packet
}

// DbgSettingsUpdated (:1080)
func DbgSettingsUpdated()  {
	// TODO: implement make json packet
}

// DbgDisasmFastAt (:1081)
func DbgDisasmFastAt(addr duint, basicinfo BASIC_INSTRUCTION_INFO *)  {
	// TODO: implement make json packet
}

// DbgMenuEntryClicked (:1082)
func DbgMenuEntryClicked(hEntry int)  {
	// TODO: implement make json packet
}

// DbgFunctionGet (:1083)
func DbgFunctionGet(addr duint, start duint *, end duint *)  {
	// TODO: implement make json packet
}

// DbgFunctionOverlaps (:1084)
func DbgFunctionOverlaps(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgFunctionAdd (:1085)
func DbgFunctionAdd(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgFunctionDel (:1086)
func DbgFunctionDel(addr duint)  {
	// TODO: implement make json packet
}

// DbgArgumentGet (:1087)
func DbgArgumentGet(addr duint, start duint *, end duint *)  {
	// TODO: implement make json packet
}

// DbgArgumentOverlaps (:1088)
func DbgArgumentOverlaps(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgArgumentAdd (:1089)
func DbgArgumentAdd(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgArgumentDel (:1090)
func DbgArgumentDel(addr duint)  {
	// TODO: implement make json packet
}

// DbgLoopGet (:1091)
func DbgLoopGet(depth int, addr duint, start duint *, end duint *)  {
	// TODO: implement make json packet
}

// DbgLoopOverlaps (:1092)
func DbgLoopOverlaps(depth int, start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgLoopAdd (:1093)
func DbgLoopAdd(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgLoopDel (:1094)
func DbgLoopDel(depth int, addr duint)  {
	// TODO: implement make json packet
}

// DbgXrefAdd (:1095)
func DbgXrefAdd(addr duint, from duint)  {
	// TODO: implement make json packet
}

// DbgXrefDelAll (:1096)
func DbgXrefDelAll(addr duint)  {
	// TODO: implement make json packet
}

// DbgXrefGet (:1097)
func DbgXrefGet(addr duint, info XREF_INFO *)  {
	// TODO: implement make json packet
}

// DbgGetXrefCountAt (:1098)
func DbgGetXrefCountAt(addr duint)  {
	// TODO: implement make json packet
}

// DbgGetXrefTypeAt (:1099)
func DbgGetXrefTypeAt(addr duint)  {
	// TODO: implement make json packet
}

// DbgIsRunLocked (:1100)
func DbgIsRunLocked()  {
	// TODO: implement make json packet
}

// DbgIsBpDisabled (:1101)
func DbgIsBpDisabled(addr duint)  {
	// TODO: implement make json packet
}

// DbgSetAutoCommentAt (:1102)
func DbgSetAutoCommentAt(addr duint, text string)  {
	// TODO: implement make json packet
}

// DbgClearAutoCommentRange (:1103)
func DbgClearAutoCommentRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgSetAutoLabelAt (:1104)
func DbgSetAutoLabelAt(addr duint, text string)  {
	// TODO: implement make json packet
}

// DbgClearAutoLabelRange (:1105)
func DbgClearAutoLabelRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgSetAutoBookmarkAt (:1106)
func DbgSetAutoBookmarkAt(addr duint)  {
	// TODO: implement make json packet
}

// DbgClearAutoBookmarkRange (:1107)
func DbgClearAutoBookmarkRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgSetAutoFunctionAt (:1108)
func DbgSetAutoFunctionAt(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgClearAutoFunctionRange (:1109)
func DbgClearAutoFunctionRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgGetStringAt (:1110)
func DbgGetStringAt(addr duint, text int8 *)  {
	// TODO: implement make json packet
}

// DbgFunctions (:1111)
func DbgFunctions()  {
	// TODO: implement make json packet
}

// DbgWinEvent (:1112)
func DbgWinEvent(message int *, result int32 *)  {
	// TODO: implement make json packet
}

// DbgWinEventGlobal (:1113)
func DbgWinEventGlobal(message int *)  {
	// TODO: implement make json packet
}

// DbgIsRunning (:1114)
func DbgIsRunning()  {
	// TODO: implement make json packet
}

// DbgGetTimeWastedCounter (:1115)
func DbgGetTimeWastedCounter()  {
	// TODO: implement make json packet
}

// DbgGetArgTypeAt (:1116)
func DbgGetArgTypeAt(addr duint)  {
	// TODO: implement make json packet
}

// DbgGetEncodeTypeBuffer (:1117)
func DbgGetEncodeTypeBuffer(addr duint, size duint *)  {
	// TODO: implement make json packet
}

// DbgReleaseEncodeTypeBuffer (:1118)
func DbgReleaseEncodeTypeBuffer(buffer void *)  {
	// TODO: implement make json packet
}

// DbgGetEncodeTypeAt (:1119)
func DbgGetEncodeTypeAt(addr duint, size duint)  {
	// TODO: implement make json packet
}

// DbgGetEncodeSizeAt (:1120)
func DbgGetEncodeSizeAt(addr duint, codesize duint)  {
	// TODO: implement make json packet
}

// DbgSetEncodeType (:1121)
func DbgSetEncodeType(addr duint, size duint, type ENCODETYPE)  {
	// TODO: implement make json packet
}

// DbgDelEncodeTypeRange (:1122)
func DbgDelEncodeTypeRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// DbgDelEncodeTypeSegment (:1123)
func DbgDelEncodeTypeSegment(start duint)  {
	// TODO: implement make json packet
}

// DbgGetWatchList (:1124)
func DbgGetWatchList(list ListInfo *)  {
	// TODO: implement make json packet
}

// DbgSelChanged (:1125)
func DbgSelChanged(hWindow int, VA duint)  {
	// TODO: implement make json packet
}

// DbgGetProcessHandle (:1126)
func DbgGetProcessHandle()  {
	// TODO: implement make json packet
}

// DbgGetThreadHandle (:1127)
func DbgGetThreadHandle()  {
	// TODO: implement make json packet
}

// DbgGetProcessId (:1128)
func DbgGetProcessId()  {
	// TODO: implement make json packet
}

// DbgGetThreadId (:1129)
func DbgGetThreadId()  {
	// TODO: implement make json packet
}

// DbgGetPebAddress (:1130)
func DbgGetPebAddress(ProcessId int)  {
	// TODO: implement make json packet
}

// DbgGetTebAddress (:1131)
func DbgGetTebAddress(ThreadId int)  {
	// TODO: implement make json packet
}

// DbgAnalyzeFunction (:1132)
func DbgAnalyzeFunction(entry duint, graph BridgeCFGraphList *)  {
	// TODO: implement make json packet
}

// DbgEval (:1133)
func DbgEval(expression string, success bool *)  {
	// TODO: implement make json packet
}

// DbgGetSymbolInfo (:1134)
func DbgGetSymbolInfo(symbolptr const SYMBOLPTR *, info SYMBOLINFO *)  {
	// TODO: implement make json packet
}

// DbgGetDebugEngine (:1135)
func DbgGetDebugEngine()  {
	// TODO: implement make json packet
}

// DbgGetSymbolInfoAt (:1136)
func DbgGetSymbolInfoAt(addr duint, info SYMBOLINFO *)  {
	// TODO: implement make json packet
}

// DbgXrefAddMulti (:1137)
func DbgXrefAddMulti(edges const XREF_EDGE *, count duint)  {
	// TODO: implement make json packet
}

// DbgMenuPrepare (:1151)
func DbgMenuPrepare(hMenu GUIMENUTYPE)  {
	// TODO: implement make json packet
}

// GuiTranslateText (:1354)
func GuiTranslateText(Source string)  {
	// TODO: implement make json packet
}

// GuiDisasmAt (:1355)
func GuiDisasmAt(addr duint, cip duint)  {
	// TODO: implement make json packet
}

// GuiSetDebugState (:1356)
func GuiSetDebugState(state DBGSTATE)  {
	// TODO: implement make json packet
}

// GuiSetDebugStateFast (:1357)
func GuiSetDebugStateFast(state DBGSTATE)  {
	// TODO: implement make json packet
}

// GuiAddLogMessage (:1358)
func GuiAddLogMessage(msg string)  {
	// TODO: implement make json packet
}

// GuiAddLogMessageHtml (:1359)
func GuiAddLogMessageHtml(msg string)  {
	// TODO: implement make json packet
}

// GuiLogClear (:1360)
func GuiLogClear()  {
	// TODO: implement make json packet
}

// GuiLogSave (:1361)
func GuiLogSave(filename string)  {
	// TODO: implement make json packet
}

// GuiLogRedirect (:1362)
func GuiLogRedirect(filename string)  {
	// TODO: implement make json packet
}

// GuiLogRedirectStop (:1363)
func GuiLogRedirectStop()  {
	// TODO: implement make json packet
}

// GuiUpdateAllViews (:1364)
func GuiUpdateAllViews()  {
	// TODO: implement make json packet
}

// GuiUpdateRegisterView (:1365)
func GuiUpdateRegisterView()  {
	// TODO: implement make json packet
}

// GuiUpdateDisassemblyView (:1366)
func GuiUpdateDisassemblyView()  {
	// TODO: implement make json packet
}

// GuiUpdateBreakpointsView (:1367)
func GuiUpdateBreakpointsView()  {
	// TODO: implement make json packet
}

// GuiUpdateWindowTitle (:1368)
func GuiUpdateWindowTitle(filename string)  {
	// TODO: implement make json packet
}

// GuiGetWindowHandle (:1369)
func GuiGetWindowHandle()  {
	// TODO: implement make json packet
}

// GuiDumpAt (:1370)
func GuiDumpAt(va duint)  {
	// TODO: implement make json packet
}

// GuiScriptAdd (:1371)
func GuiScriptAdd(count int, lines string*)  {
	// TODO: implement make json packet
}

// GuiScriptClear (:1372)
func GuiScriptClear()  {
	// TODO: implement make json packet
}

// GuiScriptSetIp (:1373)
func GuiScriptSetIp(line int)  {
	// TODO: implement make json packet
}

// GuiScriptError (:1374)
func GuiScriptError(line int, message string)  {
	// TODO: implement make json packet
}

// GuiScriptSetTitle (:1375)
func GuiScriptSetTitle(title string)  {
	// TODO: implement make json packet
}

// GuiScriptSetInfoLine (:1376)
func GuiScriptSetInfoLine(line int, info string)  {
	// TODO: implement make json packet
}

// GuiScriptMessage (:1377)
func GuiScriptMessage(message string)  {
	// TODO: implement make json packet
}

// GuiScriptMsgyn (:1378)
func GuiScriptMsgyn(message string)  {
	// TODO: implement make json packet
}

// GuiScriptEnableHighlighting (:1379)
func GuiScriptEnableHighlighting(enable bool)  {
	// TODO: implement make json packet
}

// GuiSymbolLogAdd (:1380)
func GuiSymbolLogAdd(message string)  {
	// TODO: implement make json packet
}

// GuiSymbolLogClear (:1381)
func GuiSymbolLogClear()  {
	// TODO: implement make json packet
}

// GuiSymbolSetProgress (:1382)
func GuiSymbolSetProgress(percent int)  {
	// TODO: implement make json packet
}

// GuiSymbolUpdateModuleList (:1383)
func GuiSymbolUpdateModuleList(count int, modules SYMBOLMODULEINFO *)  {
	// TODO: implement make json packet
}

// GuiSymbolRefreshCurrent (:1384)
func GuiSymbolRefreshCurrent()  {
	// TODO: implement make json packet
}

// GuiReferenceAddColumn (:1385)
func GuiReferenceAddColumn(width int, title string)  {
	// TODO: implement make json packet
}

// GuiReferenceSetRowCount (:1386)
func GuiReferenceSetRowCount(count int)  {
	// TODO: implement make json packet
}

// GuiReferenceGetRowCount (:1387)
func GuiReferenceGetRowCount()  {
	// TODO: implement make json packet
}

// GuiReferenceSearchGetRowCount (:1388)
func GuiReferenceSearchGetRowCount()  {
	// TODO: implement make json packet
}

// GuiReferenceDeleteAllColumns (:1389)
func GuiReferenceDeleteAllColumns()  {
	// TODO: implement make json packet
}

// GuiReferenceInitialize (:1390)
func GuiReferenceInitialize(name string)  {
	// TODO: implement make json packet
}

// GuiReferenceSetCellContent (:1391)
func GuiReferenceSetCellContent(row int, col int, str string)  {
	// TODO: implement make json packet
}

// GuiReferenceGetCellContent (:1392)
func GuiReferenceGetCellContent(row int, col int)  {
	// TODO: implement make json packet
}

// GuiReferenceSearchGetCellContent (:1393)
func GuiReferenceSearchGetCellContent(row int, col int)  {
	// TODO: implement make json packet
}

// GuiReferenceReloadData (:1394)
func GuiReferenceReloadData()  {
	// TODO: implement make json packet
}

// GuiReferenceSetSingleSelection (:1395)
func GuiReferenceSetSingleSelection(index int, scroll bool)  {
	// TODO: implement make json packet
}

// GuiReferenceSetProgress (:1396)
func GuiReferenceSetProgress(progress int)  {
	// TODO: implement make json packet
}

// GuiReferenceSetCurrentTaskProgress (:1397)
func GuiReferenceSetCurrentTaskProgress(progress int, taskTitle string)  {
	// TODO: implement make json packet
}

// GuiReferenceSetSearchStartCol (:1398)
func GuiReferenceSetSearchStartCol(col int)  {
	// TODO: implement make json packet
}

// GuiStackDumpAt (:1399)
func GuiStackDumpAt(addr duint, csp duint)  {
	// TODO: implement make json packet
}

// GuiUpdateDumpView (:1400)
func GuiUpdateDumpView()  {
	// TODO: implement make json packet
}

// GuiUpdateWatchView (:1401)
func GuiUpdateWatchView()  {
	// TODO: implement make json packet
}

// GuiUpdateThreadView (:1402)
func GuiUpdateThreadView()  {
	// TODO: implement make json packet
}

// GuiUpdateMemoryView (:1403)
func GuiUpdateMemoryView()  {
	// TODO: implement make json packet
}

// GuiAddRecentFile (:1404)
func GuiAddRecentFile(file string)  {
	// TODO: implement make json packet
}

// GuiSetLastException (:1405)
func GuiSetLastException(exception uint)  {
	// TODO: implement make json packet
}

// GuiGetDisassembly (:1406)
func GuiGetDisassembly(addr duint, text int8 *)  {
	// TODO: implement make json packet
}

// GuiMenuAdd (:1407)
func GuiMenuAdd(hMenu int, title string)  {
	// TODO: implement make json packet
}

// GuiMenuAddEntry (:1408)
func GuiMenuAddEntry(hMenu int, title string)  {
	// TODO: implement make json packet
}

// GuiMenuAddSeparator (:1409)
func GuiMenuAddSeparator(hMenu int)  {
	// TODO: implement make json packet
}

// GuiMenuClear (:1410)
func GuiMenuClear(hMenu int)  {
	// TODO: implement make json packet
}

// GuiMenuRemove (:1411)
func GuiMenuRemove(hEntryMenu int)  {
	// TODO: implement make json packet
}

// GuiSelectionGet (:1412)
func GuiSelectionGet(hWindow GUISELECTIONTYPE, selection SELECTIONDATA *)  {
	// TODO: implement make json packet
}

// GuiSelectionSet (:1413)
func GuiSelectionSet(hWindow GUISELECTIONTYPE, selection const SELECTIONDATA *)  {
	// TODO: implement make json packet
}

// GuiGetLineWindow (:1414)
func GuiGetLineWindow(title string, text int8 *)  {
	// TODO: implement make json packet
}

// GuiAutoCompleteAddCmd (:1415)
func GuiAutoCompleteAddCmd(cmd string)  {
	// TODO: implement make json packet
}

// GuiAutoCompleteDelCmd (:1416)
func GuiAutoCompleteDelCmd(cmd string)  {
	// TODO: implement make json packet
}

// GuiAutoCompleteClearAll (:1417)
func GuiAutoCompleteClearAll()  {
	// TODO: implement make json packet
}

// GuiAddStatusBarMessage (:1418)
func GuiAddStatusBarMessage(msg string)  {
	// TODO: implement make json packet
}

// GuiUpdateSideBar (:1419)
func GuiUpdateSideBar()  {
	// TODO: implement make json packet
}

// GuiRepaintTableView (:1420)
func GuiRepaintTableView()  {
	// TODO: implement make json packet
}

// GuiUpdatePatches (:1421)
func GuiUpdatePatches()  {
	// TODO: implement make json packet
}

// GuiUpdateCallStack (:1422)
func GuiUpdateCallStack()  {
	// TODO: implement make json packet
}

// GuiUpdateSEHChain (:1423)
func GuiUpdateSEHChain()  {
	// TODO: implement make json packet
}

// GuiLoadSourceFileEx (:1424)
func GuiLoadSourceFileEx(path string, addr duint)  {
	// TODO: implement make json packet
}

// GuiMenuSetIcon (:1425)
func GuiMenuSetIcon(hMenu int, icon const ICONDATA *)  {
	// TODO: implement make json packet
}

// GuiMenuSetEntryIcon (:1426)
func GuiMenuSetEntryIcon(hEntry int, icon const ICONDATA *)  {
	// TODO: implement make json packet
}

// GuiMenuSetEntryChecked (:1427)
func GuiMenuSetEntryChecked(hEntry int, checked bool)  {
	// TODO: implement make json packet
}

// GuiMenuSetVisible (:1428)
func GuiMenuSetVisible(hMenu int, visible bool)  {
	// TODO: implement make json packet
}

// GuiMenuSetEntryVisible (:1429)
func GuiMenuSetEntryVisible(hEntry int, visible bool)  {
	// TODO: implement make json packet
}

// GuiMenuSetName (:1430)
func GuiMenuSetName(hMenu int, name string)  {
	// TODO: implement make json packet
}

// GuiMenuSetEntryName (:1431)
func GuiMenuSetEntryName(hEntry int, name string)  {
	// TODO: implement make json packet
}

// GuiMenuSetEntryHotkey (:1432)
func GuiMenuSetEntryHotkey(hEntry int, hack string)  {
	// TODO: implement make json packet
}

// GuiShowCpu (:1433)
func GuiShowCpu()  {
	// TODO: implement make json packet
}

// GuiShowThreads (:1434)
func GuiShowThreads()  {
	// TODO: implement make json packet
}

// GuiAddQWidgetTab (:1435)
func GuiAddQWidgetTab(qWidget void *)  {
	// TODO: implement make json packet
}

// GuiShowQWidgetTab (:1436)
func GuiShowQWidgetTab(qWidget void *)  {
	// TODO: implement make json packet
}

// GuiCloseQWidgetTab (:1437)
func GuiCloseQWidgetTab(qWidget void *)  {
	// TODO: implement make json packet
}

// GuiExecuteOnGuiThread (:1438)
func GuiExecuteOnGuiThread(cbGuiThread GUICALLBACK)  {
	// TODO: implement make json packet
}

// GuiUpdateTimeWastedCounter (:1439)
func GuiUpdateTimeWastedCounter()  {
	// TODO: implement make json packet
}

// GuiSetGlobalNotes (:1440)
func GuiSetGlobalNotes(text string)  {
	// TODO: implement make json packet
}

// GuiGetGlobalNotes (:1441)
func GuiGetGlobalNotes(text int8 **)  {
	// TODO: implement make json packet
}

// GuiSetDebuggeeNotes (:1442)
func GuiSetDebuggeeNotes(text string)  {
	// TODO: implement make json packet
}

// GuiGetDebuggeeNotes (:1443)
func GuiGetDebuggeeNotes(text int8 **)  {
	// TODO: implement make json packet
}

// GuiDumpAtN (:1444)
func GuiDumpAtN(va duint, index int)  {
	// TODO: implement make json packet
}

// GuiDisplayWarning (:1445)
func GuiDisplayWarning(title string, text string)  {
	// TODO: implement make json packet
}

// GuiRegisterScriptLanguage (:1446)
func GuiRegisterScriptLanguage(info SCRIPTTYPEINFO *)  {
	// TODO: implement make json packet
}

// GuiUnregisterScriptLanguage (:1447)
func GuiUnregisterScriptLanguage(id int)  {
	// TODO: implement make json packet
}

// GuiUpdateArgumentWidget (:1448)
func GuiUpdateArgumentWidget()  {
	// TODO: implement make json packet
}

// GuiFocusView (:1449)
func GuiFocusView(hWindow int)  {
	// TODO: implement make json packet
}

// GuiIsUpdateDisabled (:1450)
func GuiIsUpdateDisabled()  {
	// TODO: implement make json packet
}

// GuiUpdateEnable (:1451)
func GuiUpdateEnable(updateNow bool)  {
	// TODO: implement make json packet
}

// GuiUpdateDisable (:1452)
func GuiUpdateDisable()  {
	// TODO: implement make json packet
}

// GuiLoadGraph (:1453)
func GuiLoadGraph(graph BridgeCFGraphList *, addr duint)  {
	// TODO: implement make json packet
}

// GuiGraphAt (:1454)
func GuiGraphAt(addr duint)  {
	// TODO: implement make json packet
}

// GuiUpdateGraphView (:1455)
func GuiUpdateGraphView()  {
	// TODO: implement make json packet
}

// GuiDisableLog (:1456)
func GuiDisableLog()  {
	// TODO: implement make json packet
}

// GuiEnableLog (:1457)
func GuiEnableLog()  {
	// TODO: implement make json packet
}

// GuiIsLogEnabled (:1458)
func GuiIsLogEnabled()  {
	// TODO: implement make json packet
}

// GuiAddFavouriteTool (:1459)
func GuiAddFavouriteTool(name string, description string)  {
	// TODO: implement make json packet
}

// GuiAddFavouriteCommand (:1460)
func GuiAddFavouriteCommand(name string, shortcut string)  {
	// TODO: implement make json packet
}

// GuiSetFavouriteToolShortcut (:1461)
func GuiSetFavouriteToolShortcut(name string, shortcut string)  {
	// TODO: implement make json packet
}

// GuiFoldDisassembly (:1462)
func GuiFoldDisassembly(startAddress duint, length duint)  {
	// TODO: implement make json packet
}

// GuiSelectInMemoryMap (:1463)
func GuiSelectInMemoryMap(addr duint)  {
	// TODO: implement make json packet
}

// GuiGetActiveView (:1464)
func GuiGetActiveView(activeView ACTIVEVIEW *)  {
	// TODO: implement make json packet
}

// GuiAddInfoLine (:1465)
func GuiAddInfoLine(infoLine string)  {
	// TODO: implement make json packet
}

// GuiProcessEvents (:1466)
func GuiProcessEvents()  {
	// TODO: implement make json packet
}

// GuiTypeAddNode (:1467)
func GuiTypeAddNode(parent void *, type const TYPEDESCRIPTOR *)  {
	// TODO: implement make json packet
}

// GuiTypeClear (:1468)
func GuiTypeClear()  {
	// TODO: implement make json packet
}

// GuiUpdateTypeWidget (:1469)
func GuiUpdateTypeWidget()  {
	// TODO: implement make json packet
}

// GuiCloseApplication (:1470)
func GuiCloseApplication()  {
	// TODO: implement make json packet
}

// GuiFlushLog (:1471)
func GuiFlushLog()  {
	// TODO: implement make json packet
}

// GuiReferenceAddCommand (:1472)
func GuiReferenceAddCommand(title string, command string)  {
	// TODO: implement make json packet
}

// GuiUpdateTraceBrowser (:1473)
func GuiUpdateTraceBrowser()  {
	// TODO: implement make json packet
}

// GuiOpenTraceFile (:1474)
func GuiOpenTraceFile(fileName string)  {
	// TODO: implement make json packet
}

// GuiInvalidateSymbolSource (:1475)
func GuiInvalidateSymbolSource(base duint)  {
	// TODO: implement make json packet
}

// GuiExecuteOnGuiThreadEx (:1476)
func GuiExecuteOnGuiThreadEx(cbGuiThread GUICALLBACKEX, userdata void *)  {
	// TODO: implement make json packet
}

// GuiGetCurrentGraph (:1477)
func GuiGetCurrentGraph(graphList BridgeCFGraphList *)  {
	// TODO: implement make json packet
}

// GuiShowReferences (:1478)
func GuiShowReferences()  {
	// TODO: implement make json packet
}

// GuiSelectInSymbolsTab (:1479)
func GuiSelectInSymbolsTab(addr duint)  {
	// TODO: implement make json packet
}

// GuiGotoTrace (:1480)
func GuiGotoTrace(index duint)  {
	// TODO: implement make json packet
}

// GuiShowTrace (:1481)
func GuiShowTrace()  {
	// TODO: implement make json packet
}

// GuiGetMainThreadId (:1482)
func GuiGetMainThreadId()  {
	// TODO: implement make json packet
}

// json_object (:86)
func json_object()  {
	// TODO: implement make json packet
}

// json_array (:87)
func json_array()  {
	// TODO: implement make json packet
}

// json_string (:88)
func json_string(value string)  {
	// TODO: implement make json packet
}

// json_stringn (:89)
func json_stringn(value string, len uint)  {
	// TODO: implement make json packet
}

// json_string_nocheck (:90)
func json_string_nocheck(value string)  {
	// TODO: implement make json packet
}

// json_stringn_nocheck (:91)
func json_stringn_nocheck(value string, len uint)  {
	// TODO: implement make json packet
}

// json_integer (:92)
func json_integer(value int64)  {
	// TODO: implement make json packet
}

// json_real (:93)
func json_real(value float64)  {
	// TODO: implement make json packet
}

// json_true (:94)
func json_true()  {
	// TODO: implement make json packet
}

// json_false (:95)
func json_false()  {
	// TODO: implement make json packet
}

// json_null (:97)
func json_null()  {
	// TODO: implement make json packet
}

// json_incref (:100)
func json_incref(json json_t *)  {
	// TODO: implement make json packet
}

// json_delete (:108)
func json_delete(json json_t *)  {
	// TODO: implement make json packet
}

// json_decref (:111)
func json_decref(json json_t *)  {
	// TODO: implement make json packet
}

// json_decrefp (:119)
func json_decrefp(json json_t **)  {
	// TODO: implement make json packet
}

// json_object_seed (:149)
func json_object_seed(seed uint)  {
	// TODO: implement make json packet
}

// json_object_size (:150)
func json_object_size(object const json_t *)  {
	// TODO: implement make json packet
}

// json_object_get (:151)
func json_object_get(object const json_t *, key string)  {
	// TODO: implement make json packet
}

// json_object_set_new (:152)
func json_object_set_new(object json_t *, key string, value json_t *)  {
	// TODO: implement make json packet
}

// json_object_set_new_nocheck (:153)
func json_object_set_new_nocheck(object json_t *, key string, value json_t *)  {
	// TODO: implement make json packet
}

// json_object_del (:154)
func json_object_del(object json_t *, key string)  {
	// TODO: implement make json packet
}

// json_object_clear (:155)
func json_object_clear(object json_t *)  {
	// TODO: implement make json packet
}

// json_object_update (:156)
func json_object_update(object json_t *, other json_t *)  {
	// TODO: implement make json packet
}

// json_object_update_existing (:157)
func json_object_update_existing(object json_t *, other json_t *)  {
	// TODO: implement make json packet
}

// json_object_update_missing (:158)
func json_object_update_missing(object json_t *, other json_t *)  {
	// TODO: implement make json packet
}

// json_object_iter (:159)
func json_object_iter(object json_t *)  {
	// TODO: implement make json packet
}

// json_object_iter_at (:160)
func json_object_iter_at(object json_t *, key string)  {
	// TODO: implement make json packet
}

// json_object_key_to_iter (:161)
func json_object_key_to_iter(key string)  {
	// TODO: implement make json packet
}

// json_object_iter_next (:162)
func json_object_iter_next(object json_t *, iter void *)  {
	// TODO: implement make json packet
}

// json_object_iter_key (:163)
func json_object_iter_key(iter void *)  {
	// TODO: implement make json packet
}

// json_object_iter_value (:164)
func json_object_iter_value(iter void *)  {
	// TODO: implement make json packet
}

// json_object_iter_set_new (:165)
func json_object_iter_set_new(object json_t *, iter void *, value json_t *)  {
	// TODO: implement make json packet
}

// json_object_set (:185)
func json_object_set(object json_t *, key string, value json_t *)  {
	// TODO: implement make json packet
}

// json_object_set_nocheck (:191)
func json_object_set_nocheck(object json_t *, key string, value json_t *)  {
	// TODO: implement make json packet
}

// json_object_iter_set (:197)
func json_object_iter_set(object json_t *, iter void *, value json_t *)  {
	// TODO: implement make json packet
}

// json_array_size (:202)
func json_array_size(array const json_t *)  {
	// TODO: implement make json packet
}

// json_array_get (:203)
func json_array_get(array const json_t *, index uint)  {
	// TODO: implement make json packet
}

// json_array_set_new (:204)
func json_array_set_new(array json_t *, index uint, value json_t *)  {
	// TODO: implement make json packet
}

// json_array_append_new (:205)
func json_array_append_new(array json_t *, value json_t *)  {
	// TODO: implement make json packet
}

// json_array_insert_new (:206)
func json_array_insert_new(array json_t *, index uint, value json_t *)  {
	// TODO: implement make json packet
}

// json_array_remove (:207)
func json_array_remove(array json_t *, index uint)  {
	// TODO: implement make json packet
}

// json_array_clear (:208)
func json_array_clear(array json_t *)  {
	// TODO: implement make json packet
}

// json_array_extend (:209)
func json_array_extend(array json_t *, other json_t *)  {
	// TODO: implement make json packet
}

// json_array_set (:212)
func json_array_set(array json_t *, ind uint, value json_t *)  {
	// TODO: implement make json packet
}

// json_array_append (:218)
func json_array_append(array json_t *, value json_t *)  {
	// TODO: implement make json packet
}

// json_array_insert (:224)
func json_array_insert(array json_t *, ind uint, value json_t *)  {
	// TODO: implement make json packet
}

// json_string_value (:229)
func json_string_value(string const json_t *)  {
	// TODO: implement make json packet
}

// json_string_length (:230)
func json_string_length(string const json_t *)  {
	// TODO: implement make json packet
}

// json_integer_value (:231)
func json_integer_value(integer const json_t *)  {
	// TODO: implement make json packet
}

// json_real_value (:232)
func json_real_value(real const json_t *)  {
	// TODO: implement make json packet
}

// json_number_value (:233)
func json_number_value(json const json_t *)  {
	// TODO: implement make json packet
}

// json_string_set (:235)
func json_string_set(string json_t *, value string)  {
	// TODO: implement make json packet
}

// json_string_setn (:236)
func json_string_setn(string json_t *, value string, len uint)  {
	// TODO: implement make json packet
}

// json_string_set_nocheck (:237)
func json_string_set_nocheck(string json_t *, value string)  {
	// TODO: implement make json packet
}

// json_string_setn_nocheck (:238)
func json_string_setn_nocheck(string json_t *, value string, len uint)  {
	// TODO: implement make json packet
}

// json_integer_set (:239)
func json_integer_set(integer json_t *, value int64)  {
	// TODO: implement make json packet
}

// json_real_set (:240)
func json_real_set(real json_t *, value float64)  {
	// TODO: implement make json packet
}

// json_pack (:244)
func json_pack(fmt string)  {
	// TODO: implement make json packet
}

// json_pack_ex (:245)
func json_pack_ex(error json_error_t *, flags uint, fmt string)  {
	// TODO: implement make json packet
}

// json_vpack_ex (:246)
func json_vpack_ex(error json_error_t *, flags uint, fmt string, ap int)  {
	// TODO: implement make json packet
}

// json_unpack (:251)
func json_unpack(root json_t *, fmt string)  {
	// TODO: implement make json packet
}

// json_unpack_ex (:252)
func json_unpack_ex(root json_t *, error json_error_t *, flags uint, fmt string)  {
	// TODO: implement make json packet
}

// json_vunpack_ex (:253)
func json_vunpack_ex(root json_t *, error json_error_t *, flags uint, fmt string, ap int)  {
	// TODO: implement make json packet
}

// json_equal (:258)
func json_equal(value1 json_t *, value2 json_t *)  {
	// TODO: implement make json packet
}

// json_copy (:263)
func json_copy(value json_t *)  {
	// TODO: implement make json packet
}

// json_deep_copy (:264)
func json_deep_copy(value const json_t *)  {
	// TODO: implement make json packet
}

// json_loads (:277)
func json_loads(input string, flags uint, error json_error_t *)  {
	// TODO: implement make json packet
}

// json_loadb (:278)
func json_loadb(buffer string, buflen uint, flags uint, error json_error_t *)  {
	// TODO: implement make json packet
}

// json_loadf (:279)
func json_loadf(input int *, flags uint, error json_error_t *)  {
	// TODO: implement make json packet
}

// json_load_file (:280)
func json_load_file(path string, flags uint, error json_error_t *)  {
	// TODO: implement make json packet
}

// json_load_callback (:281)
func json_load_callback(callback json_load_callback_t, data void *, flags uint, error json_error_t *)  {
	// TODO: implement make json packet
}

// json_dumps (:298)
func json_dumps(json const json_t *, flags uint)  {
	// TODO: implement make json packet
}

// json_dumpf (:299)
func json_dumpf(json const json_t *, output int *, flags uint)  {
	// TODO: implement make json packet
}

// json_dump_file (:300)
func json_dump_file(json const json_t *, path string, flags uint)  {
	// TODO: implement make json packet
}

// json_dump_callback (:301)
func json_dump_callback(json const json_t *, callback json_dump_callback_t, data void *, flags uint)  {
	// TODO: implement make json packet
}

// json_set_alloc_funcs (:308)
func json_set_alloc_funcs(malloc_fn json_malloc_t, free_fn json_free_t)  {
	// TODO: implement make json packet
}

// json_get_alloc_funcs (:309)
func json_get_alloc_funcs(malloc_fn json_malloc_t *, free_fn json_free_t *)  {
	// TODO: implement make json packet
}

// _plugin_registercallback (:324)
func _plugin_registercallback(pluginHandle int, cbType CBTYPE, cbPlugin CBPLUGIN)  {
	// TODO: implement make json packet
}

// _plugin_unregistercallback (:325)
func _plugin_unregistercallback(pluginHandle int, cbType CBTYPE)  {
	// TODO: implement make json packet
}

// _plugin_registercommand (:326)
func _plugin_registercommand(pluginHandle int, command string, cbCommand CBPLUGINCOMMAND, debugonly bool)  {
	// TODO: implement make json packet
}

// _plugin_unregistercommand (:327)
func _plugin_unregistercommand(pluginHandle int, command string)  {
	// TODO: implement make json packet
}

// _plugin_logprintf (:328)
func _plugin_logprintf(format string)  {
	// TODO: implement make json packet
}

// _plugin_lograw_html (:329)
func _plugin_lograw_html(text string)  {
	// TODO: implement make json packet
}

// _plugin_logputs (:330)
func _plugin_logputs(text string)  {
	// TODO: implement make json packet
}

// _plugin_logprint (:331)
func _plugin_logprint(text string)  {
	// TODO: implement make json packet
}

// _plugin_debugpause (:332)
func _plugin_debugpause()  {
	// TODO: implement make json packet
}

// _plugin_debugskipexceptions (:333)
func _plugin_debugskipexceptions(skip bool)  {
	// TODO: implement make json packet
}

// _plugin_menuadd (:334)
func _plugin_menuadd(hMenu int, title string)  {
	// TODO: implement make json packet
}

// _plugin_menuaddentry (:335)
func _plugin_menuaddentry(hMenu int, hEntry int, title string)  {
	// TODO: implement make json packet
}

// _plugin_menuaddseparator (:336)
func _plugin_menuaddseparator(hMenu int)  {
	// TODO: implement make json packet
}

// _plugin_menuclear (:337)
func _plugin_menuclear(hMenu int)  {
	// TODO: implement make json packet
}

// _plugin_menuseticon (:338)
func _plugin_menuseticon(hMenu int, icon const ICONDATA *)  {
	// TODO: implement make json packet
}

// _plugin_menuentryseticon (:339)
func _plugin_menuentryseticon(pluginHandle int, hEntry int, icon const ICONDATA *)  {
	// TODO: implement make json packet
}

// _plugin_menuentrysetchecked (:340)
func _plugin_menuentrysetchecked(pluginHandle int, hEntry int, checked bool)  {
	// TODO: implement make json packet
}

// _plugin_menusetvisible (:341)
func _plugin_menusetvisible(pluginHandle int, hMenu int, visible bool)  {
	// TODO: implement make json packet
}

// _plugin_menuentrysetvisible (:342)
func _plugin_menuentrysetvisible(pluginHandle int, hEntry int, visible bool)  {
	// TODO: implement make json packet
}

// _plugin_menusetname (:343)
func _plugin_menusetname(pluginHandle int, hMenu int, name string)  {
	// TODO: implement make json packet
}

// _plugin_menuentrysetname (:344)
func _plugin_menuentrysetname(pluginHandle int, hEntry int, name string)  {
	// TODO: implement make json packet
}

// _plugin_menuentrysethotkey (:345)
func _plugin_menuentrysethotkey(pluginHandle int, hEntry int, hotkey string)  {
	// TODO: implement make json packet
}

// _plugin_menuremove (:346)
func _plugin_menuremove(hMenu int)  {
	// TODO: implement make json packet
}

// _plugin_menuentryremove (:347)
func _plugin_menuentryremove(pluginHandle int, hEntry int)  {
	// TODO: implement make json packet
}

// _plugin_startscript (:348)
func _plugin_startscript(cbScript CBPLUGINSCRIPT)  {
	// TODO: implement make json packet
}

// _plugin_waituntilpaused (:349)
func _plugin_waituntilpaused()  {
	// TODO: implement make json packet
}

// _plugin_registerexprfunction (:350)
func _plugin_registerexprfunction(pluginHandle int, name string, argc int, cbFunction CBPLUGINEXPRFUNCTION, userdata void *)  {
	// TODO: implement make json packet
}

// _plugin_registerexprfunctionex (:351)
func _plugin_registerexprfunctionex(pluginHandle int, name string, returnType ValueType, argTypes const ValueType *, argCount uint, cbFunction CBPLUGINEXPRFUNCTIONEX, userdata void *)  {
	// TODO: implement make json packet
}

// _plugin_unregisterexprfunction (:352)
func _plugin_unregisterexprfunction(pluginHandle int, name string)  {
	// TODO: implement make json packet
}

// _plugin_unload (:353)
func _plugin_unload(pluginName string)  {
	// TODO: implement make json packet
}

// _plugin_load (:354)
func _plugin_load(pluginName string)  {
	// TODO: implement make json packet
}

// _plugin_hash (:355)
func _plugin_hash(data const void *, size duint)  {
	// TODO: implement make json packet
}

// _plugin_registerformatfunction (:356)
func _plugin_registerformatfunction(pluginHandle int, type string, cbFunction CBPLUGINFORMATFUNCTION, userdata void *)  {
	// TODO: implement make json packet
}

// _plugin_unregisterformatfunction (:357)
func _plugin_unregisterformatfunction(pluginHandle int, type string)  {
	// TODO: implement make json packet
}

// Add (:20)
func Add(info const FunctionInfo *)  {
	// TODO: implement make json packet
}

// Get (:105)
func Get(reg RegisterEnum)  {
	// TODO: implement make json packet
}

// GetInfo (:24)
func GetInfo(addr duint, info LabelInfo *)  {
	// TODO: implement make json packet
}

// Overlaps (:23)
func Overlaps(start duint, end duint)  {
	// TODO: implement make json packet
}

// Delete (:25)
func Delete(addr duint)  {
	// TODO: implement make json packet
}

// DeleteRange (:26)
func DeleteRange(start duint, end duint)  {
	// TODO: implement make json packet
}

// Clear (:27)
func Clear()  {
	// TODO: implement make json packet
}

// GetList (:26)
func GetList(list ListInfo *)  {
	// TODO: implement make json packet
}

// Assemble (:10)
func Assemble(addr duint, dest byte *, size int *, instruction string)  {
	// TODO: implement make json packet
}

// AssembleEx (:11)
func AssembleEx(addr duint, dest byte *, size int *, instruction string, error int8 *)  {
	// TODO: implement make json packet
}

// AssembleMem (:12)
func AssembleMem(addr duint, instruction string)  {
	// TODO: implement make json packet
}

// AssembleMemEx (:13)
func AssembleMemEx(addr duint, instruction string, size int *, error int8 *, fillnop bool)  {
	// TODO: implement make json packet
}

// Set (:106)
func Set(reg RegisterEnum, value duint)  {
	// TODO: implement make json packet
}

// Wait (:17)
func Wait()  {
	// TODO: implement make json packet
}

// Run (:18)
func Run()  {
	// TODO: implement make json packet
}

// Pause (:19)
func Pause()  {
	// TODO: implement make json packet
}

// Stop (:20)
func Stop()  {
	// TODO: implement make json packet
}

// StepIn (:21)
func StepIn()  {
	// TODO: implement make json packet
}

// StepOver (:22)
func StepOver()  {
	// TODO: implement make json packet
}

// StepOut (:23)
func StepOut()  {
	// TODO: implement make json packet
}

// SetBreakpoint (:24)
func SetBreakpoint(address duint)  {
	// TODO: implement make json packet
}

// DeleteBreakpoint (:25)
func DeleteBreakpoint(address duint)  {
	// TODO: implement make json packet
}

// DisableBreakpoint (:26)
func DisableBreakpoint(address duint)  {
	// TODO: implement make json packet
}

// SetHardwareBreakpoint (:27)
func SetHardwareBreakpoint(address duint, type HardwareType)  {
	// TODO: implement make json packet
}

// DeleteHardwareBreakpoint (:28)
func DeleteHardwareBreakpoint(address duint)  {
	// TODO: implement make json packet
}

// GetZF (:26)
func GetZF()  {
	// TODO: implement make json packet
}

// SetZF (:27)
func SetZF(value bool)  {
	// TODO: implement make json packet
}

// GetOF (:28)
func GetOF()  {
	// TODO: implement make json packet
}

// SetOF (:29)
func SetOF(value bool)  {
	// TODO: implement make json packet
}

// GetCF (:30)
func GetCF()  {
	// TODO: implement make json packet
}

// SetCF (:31)
func SetCF(value bool)  {
	// TODO: implement make json packet
}

// GetPF (:32)
func GetPF()  {
	// TODO: implement make json packet
}

// SetPF (:33)
func SetPF(value bool)  {
	// TODO: implement make json packet
}

// GetSF (:34)
func GetSF()  {
	// TODO: implement make json packet
}

// SetSF (:35)
func SetSF(value bool)  {
	// TODO: implement make json packet
}

// GetTF (:36)
func GetTF()  {
	// TODO: implement make json packet
}

// SetTF (:37)
func SetTF(value bool)  {
	// TODO: implement make json packet
}

// GetAF (:38)
func GetAF()  {
	// TODO: implement make json packet
}

// SetAF (:39)
func SetAF(value bool)  {
	// TODO: implement make json packet
}

// GetDF (:40)
func GetDF()  {
	// TODO: implement make json packet
}

// SetDF (:41)
func SetDF(value bool)  {
	// TODO: implement make json packet
}

// GetIF (:42)
func GetIF()  {
	// TODO: implement make json packet
}

// SetIF (:43)
func SetIF(value bool)  {
	// TODO: implement make json packet
}

// SelectionGet (:62)
func SelectionGet(window Window, start duint *, end duint *)  {
	// TODO: implement make json packet
}

// SelectionSet (:63)
func SelectionSet(window Window, start duint, end duint)  {
	// TODO: implement make json packet
}

// SelectionGetStart (:64)
func SelectionGetStart(window Window)  {
	// TODO: implement make json packet
}

// SelectionGetEnd (:65)
func SelectionGetEnd(window Window)  {
	// TODO: implement make json packet
}

// Message (:66)
func Message(message string)  {
	// TODO: implement make json packet
}

// MessageYesNo (:67)
func MessageYesNo(message string)  {
	// TODO: implement make json packet
}

// InputLine (:68)
func InputLine(title string, text int8 *)  {
	// TODO: implement make json packet
}

// InputValue (:69)
func InputValue(title string, value duint *)  {
	// TODO: implement make json packet
}

// Refresh (:70)
func Refresh()  {
	// TODO: implement make json packet
}

// AddQWidgetTab (:71)
func AddQWidgetTab(qWidget void *)  {
	// TODO: implement make json packet
}

// ShowQWidgetTab (:72)
func ShowQWidgetTab(qWidget void *)  {
	// TODO: implement make json packet
}

// CloseQWidgetTab (:73)
func CloseQWidgetTab(qWidget void *)  {
	// TODO: implement make json packet
}

// FromString (:21)
func FromString(label string, addr duint *)  {
	// TODO: implement make json packet
}

// IsTemporary (:23)
func IsTemporary(addr duint)  {
	// TODO: implement make json packet
}

// Read (:10)
func Read(addr duint, data void *, size duint, sizeRead duint *)  {
	// TODO: implement make json packet
}

// Write (:12)
func Write(data byte *, datasize duint, pattern string)  {
	// TODO: implement make json packet
}

// IsValidPtr (:12)
func IsValidPtr(addr duint)  {
	// TODO: implement make json packet
}

// RemoteAlloc (:13)
func RemoteAlloc(addr duint, size duint)  {
	// TODO: implement make json packet
}

// RemoteFree (:14)
func RemoteFree(addr duint)  {
	// TODO: implement make json packet
}

// GetProtect (:15)
func GetProtect(addr duint, reserved bool, cache bool)  {
	// TODO: implement make json packet
}

// SetProtect (:16)
func SetProtect(addr duint, protect uint, size duint)  {
	// TODO: implement make json packet
}

// GetBase (:17)
func GetBase(addr duint, reserved bool, cache bool)  {
	// TODO: implement make json packet
}

// GetSize (:18)
func GetSize(addr duint, reserved bool, cache bool)  {
	// TODO: implement make json packet
}

// ReadByte (:20)
func ReadByte(addr duint)  {
	// TODO: implement make json packet
}

// WriteByte (:21)
func WriteByte(addr duint, data byte)  {
	// TODO: implement make json packet
}

// ReadWord (:22)
func ReadWord(addr duint)  {
	// TODO: implement make json packet
}

// WriteWord (:23)
func WriteWord(addr duint, data uint16)  {
	// TODO: implement make json packet
}

// ReadDword (:24)
func ReadDword(addr duint)  {
	// TODO: implement make json packet
}

// WriteDword (:25)
func WriteDword(addr duint, data uint)  {
	// TODO: implement make json packet
}

// ReadQword (:26)
func ReadQword(addr duint)  {
	// TODO: implement make json packet
}

// WriteQword (:27)
func WriteQword(addr duint, data uint64)  {
	// TODO: implement make json packet
}

// ReadPtr (:28)
func ReadPtr(addr duint)  {
	// TODO: implement make json packet
}

// WritePtr (:29)
func WritePtr(addr duint, data duint)  {
	// TODO: implement make json packet
}

// ParseExpression (:24)
func ParseExpression(expression string, value duint *)  {
	// TODO: implement make json packet
}

// RemoteGetProcAddress (:37)
func RemoteGetProcAddress(module string, api string)  {
	// TODO: implement make json packet
}

// ResolveLabel (:49)
func ResolveLabel(label string)  {
	// TODO: implement make json packet
}

// Alloc (:65)
func Alloc(size duint)  {
	// TODO: implement make json packet
}

// InfoFromAddr (:47)
func InfoFromAddr(addr duint, info ModuleInfo *)  {
	// TODO: implement make json packet
}

// InfoFromName (:48)
func InfoFromName(name string, info ModuleInfo *)  {
	// TODO: implement make json packet
}

// BaseFromAddr (:49)
func BaseFromAddr(addr duint)  {
	// TODO: implement make json packet
}

// BaseFromName (:50)
func BaseFromName(name string)  {
	// TODO: implement make json packet
}

// SizeFromAddr (:51)
func SizeFromAddr(addr duint)  {
	// TODO: implement make json packet
}

// SizeFromName (:52)
func SizeFromName(name string)  {
	// TODO: implement make json packet
}

// NameFromAddr (:53)
func NameFromAddr(addr duint, name int8 *)  {
	// TODO: implement make json packet
}

// PathFromAddr (:54)
func PathFromAddr(addr duint, path int8 *)  {
	// TODO: implement make json packet
}

// PathFromName (:55)
func PathFromName(name string, path int8 *)  {
	// TODO: implement make json packet
}

// EntryFromAddr (:56)
func EntryFromAddr(addr duint)  {
	// TODO: implement make json packet
}

// EntryFromName (:57)
func EntryFromName(name string)  {
	// TODO: implement make json packet
}

// SectionCountFromAddr (:58)
func SectionCountFromAddr(addr duint)  {
	// TODO: implement make json packet
}

// SectionCountFromName (:59)
func SectionCountFromName(name string)  {
	// TODO: implement make json packet
}

// SectionFromAddr (:60)
func SectionFromAddr(addr duint, number int, section ModuleSectionInfo *)  {
	// TODO: implement make json packet
}

// SectionFromName (:61)
func SectionFromName(name string, number int, section ModuleSectionInfo *)  {
	// TODO: implement make json packet
}

// SectionListFromAddr (:62)
func SectionListFromAddr(addr duint, list ListInfo *)  {
	// TODO: implement make json packet
}

// SectionListFromName (:63)
func SectionListFromName(name string, list ListInfo *)  {
	// TODO: implement make json packet
}

// GetMainModuleInfo (:64)
func GetMainModuleInfo(info ModuleInfo *)  {
	// TODO: implement make json packet
}

// GetMainModuleBase (:65)
func GetMainModuleBase()  {
	// TODO: implement make json packet
}

// GetMainModuleSize (:66)
func GetMainModuleSize()  {
	// TODO: implement make json packet
}

// GetMainModuleEntry (:67)
func GetMainModuleEntry()  {
	// TODO: implement make json packet
}

// GetMainModuleSectionCount (:68)
func GetMainModuleSectionCount()  {
	// TODO: implement make json packet
}

// GetMainModuleName (:69)
func GetMainModuleName(name int8 *)  {
	// TODO: implement make json packet
}

// GetMainModulePath (:70)
func GetMainModulePath(path int8 *)  {
	// TODO: implement make json packet
}

// GetMainModuleSectionList (:71)
func GetMainModuleSectionList(list ListInfo *)  {
	// TODO: implement make json packet
}

// GetExports (:73)
func GetExports(mod const ModuleInfo *, list ListInfo *)  {
	// TODO: implement make json packet
}

// GetImports (:74)
func GetImports(mod const ModuleInfo *, list ListInfo *)  {
	// TODO: implement make json packet
}

// Find (:10)
func Find(data byte *, datasize duint, pattern string)  {
	// TODO: implement make json packet
}

// FindMem (:11)
func FindMem(start duint, size duint, pattern string)  {
	// TODO: implement make json packet
}

// WriteMem (:13)
func WriteMem(start duint, size duint, pattern string)  {
	// TODO: implement make json packet
}

// SearchAndReplace (:14)
func SearchAndReplace(data byte *, datasize duint, searchpattern string, replacepattern string)  {
	// TODO: implement make json packet
}

// SearchAndReplaceMem (:15)
func SearchAndReplaceMem(start duint, size duint, searchpattern string, replacepattern string)  {
	// TODO: implement make json packet
}

// Size (:107)
func Size()  {
	// TODO: implement make json packet
}

// GetDR0 (:109)
func GetDR0()  {
	// TODO: implement make json packet
}

// SetDR0 (:110)
func SetDR0(value duint)  {
	// TODO: implement make json packet
}

// GetDR1 (:111)
func GetDR1()  {
	// TODO: implement make json packet
}

// SetDR1 (:112)
func SetDR1(value duint)  {
	// TODO: implement make json packet
}

// GetDR2 (:113)
func GetDR2()  {
	// TODO: implement make json packet
}

// SetDR2 (:114)
func SetDR2(value duint)  {
	// TODO: implement make json packet
}

// GetDR3 (:115)
func GetDR3()  {
	// TODO: implement make json packet
}

// SetDR3 (:116)
func SetDR3(value duint)  {
	// TODO: implement make json packet
}

// GetDR6 (:117)
func GetDR6()  {
	// TODO: implement make json packet
}

// SetDR6 (:118)
func SetDR6(value duint)  {
	// TODO: implement make json packet
}

// GetDR7 (:119)
func GetDR7()  {
	// TODO: implement make json packet
}

// SetDR7 (:120)
func SetDR7(value duint)  {
	// TODO: implement make json packet
}

// GetEAX (:122)
func GetEAX()  {
	// TODO: implement make json packet
}

// SetEAX (:123)
func SetEAX(value uint)  {
	// TODO: implement make json packet
}

// GetAX (:124)
func GetAX()  {
	// TODO: implement make json packet
}

// SetAX (:125)
func SetAX(value uint16)  {
	// TODO: implement make json packet
}

// GetAH (:126)
func GetAH()  {
	// TODO: implement make json packet
}

// SetAH (:127)
func SetAH(value byte)  {
	// TODO: implement make json packet
}

// GetAL (:128)
func GetAL()  {
	// TODO: implement make json packet
}

// SetAL (:129)
func SetAL(value byte)  {
	// TODO: implement make json packet
}

// GetEBX (:130)
func GetEBX()  {
	// TODO: implement make json packet
}

// SetEBX (:131)
func SetEBX(value uint)  {
	// TODO: implement make json packet
}

// GetBX (:132)
func GetBX()  {
	// TODO: implement make json packet
}

// SetBX (:133)
func SetBX(value uint16)  {
	// TODO: implement make json packet
}

// GetBH (:134)
func GetBH()  {
	// TODO: implement make json packet
}

// SetBH (:135)
func SetBH(value byte)  {
	// TODO: implement make json packet
}

// GetBL (:136)
func GetBL()  {
	// TODO: implement make json packet
}

// SetBL (:137)
func SetBL(value byte)  {
	// TODO: implement make json packet
}

// GetECX (:138)
func GetECX()  {
	// TODO: implement make json packet
}

// SetECX (:139)
func SetECX(value uint)  {
	// TODO: implement make json packet
}

// GetCX (:140)
func GetCX()  {
	// TODO: implement make json packet
}

// SetCX (:141)
func SetCX(value uint16)  {
	// TODO: implement make json packet
}

// GetCH (:142)
func GetCH()  {
	// TODO: implement make json packet
}

// SetCH (:143)
func SetCH(value byte)  {
	// TODO: implement make json packet
}

// GetCL (:144)
func GetCL()  {
	// TODO: implement make json packet
}

// SetCL (:145)
func SetCL(value byte)  {
	// TODO: implement make json packet
}

// GetEDX (:146)
func GetEDX()  {
	// TODO: implement make json packet
}

// SetEDX (:147)
func SetEDX(value uint)  {
	// TODO: implement make json packet
}

// GetDX (:148)
func GetDX()  {
	// TODO: implement make json packet
}

// SetDX (:149)
func SetDX(value uint16)  {
	// TODO: implement make json packet
}

// GetDH (:150)
func GetDH()  {
	// TODO: implement make json packet
}

// SetDH (:151)
func SetDH(value byte)  {
	// TODO: implement make json packet
}

// GetDL (:152)
func GetDL()  {
	// TODO: implement make json packet
}

// SetDL (:153)
func SetDL(value byte)  {
	// TODO: implement make json packet
}

// GetEDI (:154)
func GetEDI()  {
	// TODO: implement make json packet
}

// SetEDI (:155)
func SetEDI(value uint)  {
	// TODO: implement make json packet
}

// GetDI (:156)
func GetDI()  {
	// TODO: implement make json packet
}

// SetDI (:157)
func SetDI(value uint16)  {
	// TODO: implement make json packet
}

// GetESI (:158)
func GetESI()  {
	// TODO: implement make json packet
}

// SetESI (:159)
func SetESI(value uint)  {
	// TODO: implement make json packet
}

// GetSI (:160)
func GetSI()  {
	// TODO: implement make json packet
}

// SetSI (:161)
func SetSI(value uint16)  {
	// TODO: implement make json packet
}

// GetEBP (:162)
func GetEBP()  {
	// TODO: implement make json packet
}

// SetEBP (:163)
func SetEBP(value uint)  {
	// TODO: implement make json packet
}

// GetBP (:164)
func GetBP()  {
	// TODO: implement make json packet
}

// SetBP (:165)
func SetBP(value uint16)  {
	// TODO: implement make json packet
}

// GetESP (:166)
func GetESP()  {
	// TODO: implement make json packet
}

// SetESP (:167)
func SetESP(value uint)  {
	// TODO: implement make json packet
}

// GetSP (:168)
func GetSP()  {
	// TODO: implement make json packet
}

// SetSP (:169)
func SetSP(value uint16)  {
	// TODO: implement make json packet
}

// GetEIP (:170)
func GetEIP()  {
	// TODO: implement make json packet
}

// SetEIP (:171)
func SetEIP(value uint)  {
	// TODO: implement make json packet
}

// GetRAX (:174)
func GetRAX()  {
	// TODO: implement make json packet
}

// SetRAX (:175)
func SetRAX(value uint64)  {
	// TODO: implement make json packet
}

// GetRBX (:176)
func GetRBX()  {
	// TODO: implement make json packet
}

// SetRBX (:177)
func SetRBX(value uint64)  {
	// TODO: implement make json packet
}

// GetRCX (:178)
func GetRCX()  {
	// TODO: implement make json packet
}

// SetRCX (:179)
func SetRCX(value uint64)  {
	// TODO: implement make json packet
}

// GetRDX (:180)
func GetRDX()  {
	// TODO: implement make json packet
}

// SetRDX (:181)
func SetRDX(value uint64)  {
	// TODO: implement make json packet
}

// GetRSI (:182)
func GetRSI()  {
	// TODO: implement make json packet
}

// SetRSI (:183)
func SetRSI(value uint64)  {
	// TODO: implement make json packet
}

// GetSIL (:184)
func GetSIL()  {
	// TODO: implement make json packet
}

// SetSIL (:185)
func SetSIL(value byte)  {
	// TODO: implement make json packet
}

// GetRDI (:186)
func GetRDI()  {
	// TODO: implement make json packet
}

// SetRDI (:187)
func SetRDI(value uint64)  {
	// TODO: implement make json packet
}

// GetDIL (:188)
func GetDIL()  {
	// TODO: implement make json packet
}

// SetDIL (:189)
func SetDIL(value byte)  {
	// TODO: implement make json packet
}

// GetRBP (:190)
func GetRBP()  {
	// TODO: implement make json packet
}

// SetRBP (:191)
func SetRBP(value uint64)  {
	// TODO: implement make json packet
}

// GetBPL (:192)
func GetBPL()  {
	// TODO: implement make json packet
}

// SetBPL (:193)
func SetBPL(value byte)  {
	// TODO: implement make json packet
}

// GetRSP (:194)
func GetRSP()  {
	// TODO: implement make json packet
}

// SetRSP (:195)
func SetRSP(value uint64)  {
	// TODO: implement make json packet
}

// GetSPL (:196)
func GetSPL()  {
	// TODO: implement make json packet
}

// SetSPL (:197)
func SetSPL(value byte)  {
	// TODO: implement make json packet
}

// GetRIP (:198)
func GetRIP()  {
	// TODO: implement make json packet
}

// SetRIP (:199)
func SetRIP(value uint64)  {
	// TODO: implement make json packet
}

// GetR8 (:200)
func GetR8()  {
	// TODO: implement make json packet
}

// SetR8 (:201)
func SetR8(value uint64)  {
	// TODO: implement make json packet
}

// GetR8D (:202)
func GetR8D()  {
	// TODO: implement make json packet
}

// SetR8D (:203)
func SetR8D(value uint)  {
	// TODO: implement make json packet
}

// GetR8W (:204)
func GetR8W()  {
	// TODO: implement make json packet
}

// SetR8W (:205)
func SetR8W(value uint16)  {
	// TODO: implement make json packet
}

// GetR8B (:206)
func GetR8B()  {
	// TODO: implement make json packet
}

// SetR8B (:207)
func SetR8B(value byte)  {
	// TODO: implement make json packet
}

// GetR9 (:208)
func GetR9()  {
	// TODO: implement make json packet
}

// SetR9 (:209)
func SetR9(value uint64)  {
	// TODO: implement make json packet
}

// GetR9D (:210)
func GetR9D()  {
	// TODO: implement make json packet
}

// SetR9D (:211)
func SetR9D(value uint)  {
	// TODO: implement make json packet
}

// GetR9W (:212)
func GetR9W()  {
	// TODO: implement make json packet
}

// SetR9W (:213)
func SetR9W(value uint16)  {
	// TODO: implement make json packet
}

// GetR9B (:214)
func GetR9B()  {
	// TODO: implement make json packet
}

// SetR9B (:215)
func SetR9B(value byte)  {
	// TODO: implement make json packet
}

// GetR10 (:216)
func GetR10()  {
	// TODO: implement make json packet
}

// SetR10 (:217)
func SetR10(value uint64)  {
	// TODO: implement make json packet
}

// GetR10D (:218)
func GetR10D()  {
	// TODO: implement make json packet
}

// SetR10D (:219)
func SetR10D(value uint)  {
	// TODO: implement make json packet
}

// GetR10W (:220)
func GetR10W()  {
	// TODO: implement make json packet
}

// SetR10W (:221)
func SetR10W(value uint16)  {
	// TODO: implement make json packet
}

// GetR10B (:222)
func GetR10B()  {
	// TODO: implement make json packet
}

// SetR10B (:223)
func SetR10B(value byte)  {
	// TODO: implement make json packet
}

// GetR11 (:224)
func GetR11()  {
	// TODO: implement make json packet
}

// SetR11 (:225)
func SetR11(value uint64)  {
	// TODO: implement make json packet
}

// GetR11D (:226)
func GetR11D()  {
	// TODO: implement make json packet
}

// SetR11D (:227)
func SetR11D(value uint)  {
	// TODO: implement make json packet
}

// GetR11W (:228)
func GetR11W()  {
	// TODO: implement make json packet
}

// SetR11W (:229)
func SetR11W(value uint16)  {
	// TODO: implement make json packet
}

// GetR11B (:230)
func GetR11B()  {
	// TODO: implement make json packet
}

// SetR11B (:231)
func SetR11B(value byte)  {
	// TODO: implement make json packet
}

// GetR12 (:232)
func GetR12()  {
	// TODO: implement make json packet
}

// SetR12 (:233)
func SetR12(value uint64)  {
	// TODO: implement make json packet
}

// GetR12D (:234)
func GetR12D()  {
	// TODO: implement make json packet
}

// SetR12D (:235)
func SetR12D(value uint)  {
	// TODO: implement make json packet
}

// GetR12W (:236)
func GetR12W()  {
	// TODO: implement make json packet
}

// SetR12W (:237)
func SetR12W(value uint16)  {
	// TODO: implement make json packet
}

// GetR12B (:238)
func GetR12B()  {
	// TODO: implement make json packet
}

// SetR12B (:239)
func SetR12B(value byte)  {
	// TODO: implement make json packet
}

// GetR13 (:240)
func GetR13()  {
	// TODO: implement make json packet
}

// SetR13 (:241)
func SetR13(value uint64)  {
	// TODO: implement make json packet
}

// GetR13D (:242)
func GetR13D()  {
	// TODO: implement make json packet
}

// SetR13D (:243)
func SetR13D(value uint)  {
	// TODO: implement make json packet
}

// GetR13W (:244)
func GetR13W()  {
	// TODO: implement make json packet
}

// SetR13W (:245)
func SetR13W(value uint16)  {
	// TODO: implement make json packet
}

// GetR13B (:246)
func GetR13B()  {
	// TODO: implement make json packet
}

// SetR13B (:247)
func SetR13B(value byte)  {
	// TODO: implement make json packet
}

// GetR14 (:248)
func GetR14()  {
	// TODO: implement make json packet
}

// SetR14 (:249)
func SetR14(value uint64)  {
	// TODO: implement make json packet
}

// GetR14D (:250)
func GetR14D()  {
	// TODO: implement make json packet
}

// SetR14D (:251)
func SetR14D(value uint)  {
	// TODO: implement make json packet
}

// GetR14W (:252)
func GetR14W()  {
	// TODO: implement make json packet
}

// SetR14W (:253)
func SetR14W(value uint16)  {
	// TODO: implement make json packet
}

// GetR14B (:254)
func GetR14B()  {
	// TODO: implement make json packet
}

// SetR14B (:255)
func SetR14B(value byte)  {
	// TODO: implement make json packet
}

// GetR15 (:256)
func GetR15()  {
	// TODO: implement make json packet
}

// SetR15 (:257)
func SetR15(value uint64)  {
	// TODO: implement make json packet
}

// GetR15D (:258)
func GetR15D()  {
	// TODO: implement make json packet
}

// SetR15D (:259)
func SetR15D(value uint)  {
	// TODO: implement make json packet
}

// GetR15W (:260)
func GetR15W()  {
	// TODO: implement make json packet
}

// SetR15W (:261)
func SetR15W(value uint16)  {
	// TODO: implement make json packet
}

// GetR15B (:262)
func GetR15B()  {
	// TODO: implement make json packet
}

// SetR15B (:263)
func SetR15B(value byte)  {
	// TODO: implement make json packet
}

// GetCAX (:266)
func GetCAX()  {
	// TODO: implement make json packet
}

// SetCAX (:267)
func SetCAX(value duint)  {
	// TODO: implement make json packet
}

// GetCBX (:268)
func GetCBX()  {
	// TODO: implement make json packet
}

// SetCBX (:269)
func SetCBX(value duint)  {
	// TODO: implement make json packet
}

// GetCCX (:270)
func GetCCX()  {
	// TODO: implement make json packet
}

// SetCCX (:271)
func SetCCX(value duint)  {
	// TODO: implement make json packet
}

// GetCDX (:272)
func GetCDX()  {
	// TODO: implement make json packet
}

// SetCDX (:273)
func SetCDX(value duint)  {
	// TODO: implement make json packet
}

// GetCDI (:274)
func GetCDI()  {
	// TODO: implement make json packet
}

// SetCDI (:275)
func SetCDI(value duint)  {
	// TODO: implement make json packet
}

// GetCSI (:276)
func GetCSI()  {
	// TODO: implement make json packet
}

// SetCSI (:277)
func SetCSI(value duint)  {
	// TODO: implement make json packet
}

// GetCBP (:278)
func GetCBP()  {
	// TODO: implement make json packet
}

// SetCBP (:279)
func SetCBP(value duint)  {
	// TODO: implement make json packet
}

// GetCSP (:280)
func GetCSP()  {
	// TODO: implement make json packet
}

// SetCSP (:281)
func SetCSP(value duint)  {
	// TODO: implement make json packet
}

// GetCIP (:282)
func GetCIP()  {
	// TODO: implement make json packet
}

// SetCIP (:283)
func SetCIP(value duint)  {
	// TODO: implement make json packet
}

// GetCFLAGS (:284)
func GetCFLAGS()  {
	// TODO: implement make json packet
}

// SetCFLAGS (:285)
func SetCFLAGS(value duint)  {
	// TODO: implement make json packet
}

// Pop (:10)
func Pop()  {
	// TODO: implement make json packet
}

// Push (:11)
func Push(value duint)  {
	// TODO: implement make json packet
}

// Peek (:12)
func Peek(offset int)  {
	// TODO: implement make json packet
}

