package sdk
// add (:31)
func add(a int, b int)  {
	// TODO: implement make json packet
}

// DevicePathToPathW (:11)
func DevicePathToPathW(szDevicePath const wint8_t *, szPath wint8_t *, nSizeInChars uint)  {
	// TODO: implement make json packet
}

// DevicePathToPathA (:12)
func DevicePathToPathA(szDevicePath string, szPath int8 *, nSizeInChars uint)  {
	// TODO: implement make json packet
}

// DevicePathFromFileHandleW (:13)
func DevicePathFromFileHandleW(hFile int, szDevicePath wint8_t *, nSizeInChars uint)  {
	// TODO: implement make json packet
}

// DevicePathFromFileHandleA (:14)
func DevicePathFromFileHandleA(hFile int, szDevicePath int8 *, nSizeInChars uint)  {
	// TODO: implement make json packet
}

// PathFromFileHandleW (:15)
func PathFromFileHandleW(hFile int, szPath wint8_t *, nSizeInChars uint)  {
	// TODO: implement make json packet
}

// PathFromFileHandleA (:16)
func PathFromFileHandleA(hFile int, szPath int8 *, nSizeInChars uint)  {
	// TODO: implement make json packet
}

// DumpProcess (:687)
func DumpProcess(hProcess int, ImageBase int, szDumpFileName string, EntryPoint int)  {
	// TODO: implement make json packet
}

// DumpProcessW (:688)
func DumpProcessW(hProcess int, ImageBase int, szDumpFileName const wint8_t *, EntryPoint int)  {
	// TODO: implement make json packet
}

// DumpProcessEx (:689)
func DumpProcessEx(ProcessId int, ImageBase int, szDumpFileName string, EntryPoint int)  {
	// TODO: implement make json packet
}

// DumpProcessExW (:690)
func DumpProcessExW(ProcessId int, ImageBase int, szDumpFileName const wint8_t *, EntryPoint int)  {
	// TODO: implement make json packet
}

// DumpMemory (:691)
func DumpMemory(hProcess int, MemoryStart int, MemorySize int, szDumpFileName string)  {
	// TODO: implement make json packet
}

// DumpMemoryW (:692)
func DumpMemoryW(hProcess int, MemoryStart int, MemorySize int, szDumpFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// DumpMemoryEx (:693)
func DumpMemoryEx(ProcessId int, MemoryStart int, MemorySize int, szDumpFileName string)  {
	// TODO: implement make json packet
}

// DumpMemoryExW (:694)
func DumpMemoryExW(ProcessId int, MemoryStart int, MemorySize int, szDumpFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// DumpRegions (:695)
func DumpRegions(hProcess int, szDumpFolder string, DumpAboveImageBaseOnly bool)  {
	// TODO: implement make json packet
}

// DumpRegionsW (:696)
func DumpRegionsW(hProcess int, szDumpFolder const wint8_t *, DumpAboveImageBaseOnly bool)  {
	// TODO: implement make json packet
}

// DumpRegionsEx (:697)
func DumpRegionsEx(ProcessId int, szDumpFolder string, DumpAboveImageBaseOnly bool)  {
	// TODO: implement make json packet
}

// DumpRegionsExW (:698)
func DumpRegionsExW(ProcessId int, szDumpFolder const wint8_t *, DumpAboveImageBaseOnly bool)  {
	// TODO: implement make json packet
}

// DumpModule (:699)
func DumpModule(hProcess int, ModuleBase int, szDumpFileName string)  {
	// TODO: implement make json packet
}

// DumpModuleW (:700)
func DumpModuleW(hProcess int, ModuleBase int, szDumpFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// DumpModuleEx (:701)
func DumpModuleEx(ProcessId int, ModuleBase int, szDumpFileName string)  {
	// TODO: implement make json packet
}

// DumpModuleExW (:702)
func DumpModuleExW(ProcessId int, ModuleBase int, szDumpFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// PastePEHeader (:703)
func PastePEHeader(hProcess int, ImageBase int, szDebuggedFileName string)  {
	// TODO: implement make json packet
}

// PastePEHeaderW (:704)
func PastePEHeaderW(hProcess int, ImageBase int, szDebuggedFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// ExtractSection (:705)
func ExtractSection(szFileName string, szDumpFileName string, SectionNumber int)  {
	// TODO: implement make json packet
}

// ExtractSectionW (:706)
func ExtractSectionW(szFileName const wint8_t *, szDumpFileName const wint8_t *, SectionNumber int)  {
	// TODO: implement make json packet
}

// ResortFileSections (:707)
func ResortFileSections(szFileName string)  {
	// TODO: implement make json packet
}

// ResortFileSectionsW (:708)
func ResortFileSectionsW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// FindOverlay (:709)
func FindOverlay(szFileName string, OverlayStart int, OverlaySize int)  {
	// TODO: implement make json packet
}

// FindOverlayW (:710)
func FindOverlayW(szFileName const wint8_t *, OverlayStart int, OverlaySize int)  {
	// TODO: implement make json packet
}

// ExtractOverlay (:711)
func ExtractOverlay(szFileName string, szExtactedFileName string)  {
	// TODO: implement make json packet
}

// ExtractOverlayW (:712)
func ExtractOverlayW(szFileName const wint8_t *, szExtactedFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// AddOverlay (:713)
func AddOverlay(szFileName string, szOverlayFileName string)  {
	// TODO: implement make json packet
}

// AddOverlayW (:714)
func AddOverlayW(szFileName const wint8_t *, szOverlayFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// CopyOverlay (:715)
func CopyOverlay(szInFileName string, szOutFileName string)  {
	// TODO: implement make json packet
}

// CopyOverlayW (:716)
func CopyOverlayW(szInFileName const wint8_t *, szOutFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// RemoveOverlay (:717)
func RemoveOverlay(szFileName string)  {
	// TODO: implement make json packet
}

// RemoveOverlayW (:718)
func RemoveOverlayW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// MakeAllSectionsRWE (:719)
func MakeAllSectionsRWE(szFileName string)  {
	// TODO: implement make json packet
}

// MakeAllSectionsRWEW (:720)
func MakeAllSectionsRWEW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// AddNewSectionEx (:721)
func AddNewSectionEx(szFileName string, szSectionName string, SectionSize int, SectionAttributes int, SectionContent int, ContentSize int)  {
	// TODO: implement make json packet
}

// AddNewSectionExW (:722)
func AddNewSectionExW(szFileName const wint8_t *, szSectionName string, SectionSize int, SectionAttributes int, SectionContent int, ContentSize int)  {
	// TODO: implement make json packet
}

// AddNewSection (:723)
func AddNewSection(szFileName string, szSectionName string, SectionSize int)  {
	// TODO: implement make json packet
}

// AddNewSectionW (:724)
func AddNewSectionW(szFileName const wint8_t *, szSectionName string, SectionSize int)  {
	// TODO: implement make json packet
}

// ResizeLastSection (:725)
func ResizeLastSection(szFileName string, NumberOfExpandBytes int, AlignResizeData bool)  {
	// TODO: implement make json packet
}

// ResizeLastSectionW (:726)
func ResizeLastSectionW(szFileName const wint8_t *, NumberOfExpandBytes int, AlignResizeData bool)  {
	// TODO: implement make json packet
}

// SetSharedOverlay (:727)
func SetSharedOverlay(szFileName string)  {
	// TODO: implement make json packet
}

// SetSharedOverlayW (:728)
func SetSharedOverlayW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// GetSharedOverlay (:729)
func GetSharedOverlay()  {
	// TODO: implement make json packet
}

// GetSharedOverlayW (:730)
func GetSharedOverlayW()  {
	// TODO: implement make json packet
}

// DeleteLastSection (:731)
func DeleteLastSection(szFileName string)  {
	// TODO: implement make json packet
}

// DeleteLastSectionW (:732)
func DeleteLastSectionW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// DeleteLastSectionEx (:733)
func DeleteLastSectionEx(szFileName string, NumberOfSections int)  {
	// TODO: implement make json packet
}

// DeleteLastSectionExW (:734)
func DeleteLastSectionExW(szFileName const wint8_t *, NumberOfSections int)  {
	// TODO: implement make json packet
}

// GetPE32DataFromMappedFile (:735)
func GetPE32DataFromMappedFile(FileMapVA int, WhichSection int, WhichData int)  {
	// TODO: implement make json packet
}

// GetPE32Data (:736)
func GetPE32Data(szFileName string, WhichSection int, WhichData int)  {
	// TODO: implement make json packet
}

// GetPE32DataW (:737)
func GetPE32DataW(szFileName const wint8_t *, WhichSection int, WhichData int)  {
	// TODO: implement make json packet
}

// GetPE32DataFromMappedFileEx (:738)
func GetPE32DataFromMappedFileEx(FileMapVA int, DataStorage int)  {
	// TODO: implement make json packet
}

// GetPE32DataEx (:739)
func GetPE32DataEx(szFileName string, DataStorage int)  {
	// TODO: implement make json packet
}

// GetPE32DataExW (:740)
func GetPE32DataExW(szFileName const wint8_t *, DataStorage int)  {
	// TODO: implement make json packet
}

// SetPE32DataForMappedFile (:741)
func SetPE32DataForMappedFile(FileMapVA int, WhichSection int, WhichData int, NewDataValue int)  {
	// TODO: implement make json packet
}

// SetPE32Data (:742)
func SetPE32Data(szFileName string, WhichSection int, WhichData int, NewDataValue int)  {
	// TODO: implement make json packet
}

// SetPE32DataW (:743)
func SetPE32DataW(szFileName const wint8_t *, WhichSection int, WhichData int, NewDataValue int)  {
	// TODO: implement make json packet
}

// SetPE32DataForMappedFileEx (:744)
func SetPE32DataForMappedFileEx(FileMapVA int, DataStorage int)  {
	// TODO: implement make json packet
}

// SetPE32DataEx (:745)
func SetPE32DataEx(szFileName string, DataStorage int)  {
	// TODO: implement make json packet
}

// SetPE32DataExW (:746)
func SetPE32DataExW(szFileName const wint8_t *, DataStorage int)  {
	// TODO: implement make json packet
}

// GetPE32SectionNumberFromVA (:747)
func GetPE32SectionNumberFromVA(FileMapVA int, AddressToConvert int)  {
	// TODO: implement make json packet
}

// ConvertVAtoFileOffset (:748)
func ConvertVAtoFileOffset(FileMapVA int, AddressToConvert int, ReturnType bool)  {
	// TODO: implement make json packet
}

// ConvertVAtoFileOffsetEx (:749)
func ConvertVAtoFileOffsetEx(FileMapVA int, FileSize int, ImageBase int, AddressToConvert int, AddressIsRVA bool, ReturnType bool)  {
	// TODO: implement make json packet
}

// ConvertFileOffsetToVA (:750)
func ConvertFileOffsetToVA(FileMapVA int, AddressToConvert int, ReturnType bool)  {
	// TODO: implement make json packet
}

// ConvertFileOffsetToVAEx (:751)
func ConvertFileOffsetToVAEx(FileMapVA int, FileSize int, ImageBase int, AddressToConvert int, ReturnType bool)  {
	// TODO: implement make json packet
}

// MemoryReadSafe (:752)
func MemoryReadSafe(hProcess int, lpBaseAddress int, lpBuffer int, nSize int, lpNumberOfBytesRead int *)  {
	// TODO: implement make json packet
}

// MemoryWriteSafe (:753)
func MemoryWriteSafe(hProcess int, lpBaseAddress int, lpBuffer int, nSize int, lpNumberOfBytesWritten int *)  {
	// TODO: implement make json packet
}

// FixHeaderCheckSum (:755)
func FixHeaderCheckSum(szFileName string)  {
	// TODO: implement make json packet
}

// FixHeaderCheckSumW (:756)
func FixHeaderCheckSumW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// RealignPE (:757)
func RealignPE(FileMapVA int, FileSize int, RealingMode int)  {
	// TODO: implement make json packet
}

// RealignPEEx (:758)
func RealignPEEx(szFileName string, RealingFileSize int, ForcedFileAlignment int)  {
	// TODO: implement make json packet
}

// RealignPEExW (:759)
func RealignPEExW(szFileName const wint8_t *, RealingFileSize int, ForcedFileAlignment int)  {
	// TODO: implement make json packet
}

// WipeSection (:760)
func WipeSection(szFileName string, WipeSectionNumber int, RemovePhysically bool)  {
	// TODO: implement make json packet
}

// WipeSectionW (:761)
func WipeSectionW(szFileName const wint8_t *, WipeSectionNumber int, RemovePhysically bool)  {
	// TODO: implement make json packet
}

// IsPE32FileValidEx (:762)
func IsPE32FileValidEx(szFileName string, CheckDepth int, FileStatusInfo int)  {
	// TODO: implement make json packet
}

// IsPE32FileValidExW (:763)
func IsPE32FileValidExW(szFileName const wint8_t *, CheckDepth int, FileStatusInfo int)  {
	// TODO: implement make json packet
}

// FixBrokenPE32FileEx (:764)
func FixBrokenPE32FileEx(szFileName string, FileStatusInfo int, FileFixInfo int)  {
	// TODO: implement make json packet
}

// FixBrokenPE32FileExW (:765)
func FixBrokenPE32FileExW(szFileName const wint8_t *, FileStatusInfo int, FileFixInfo int)  {
	// TODO: implement make json packet
}

// IsFileDLL (:766)
func IsFileDLL(szFileName string, FileMapVA int)  {
	// TODO: implement make json packet
}

// IsFileDLLW (:767)
func IsFileDLLW(szFileName const wint8_t *, FileMapVA int)  {
	// TODO: implement make json packet
}

// GetPEBLocation (:769)
func GetPEBLocation(hProcess int)  {
	// TODO: implement make json packet
}

// GetPEBLocation64 (:770)
func GetPEBLocation64(hProcess int)  {
	// TODO: implement make json packet
}

// GetTEBLocation (:771)
func GetTEBLocation(hThread int)  {
	// TODO: implement make json packet
}

// GetTEBLocation64 (:772)
func GetTEBLocation64(hThread int)  {
	// TODO: implement make json packet
}

// HideDebugger (:773)
func HideDebugger(hProcess int, PatchAPILevel int)  {
	// TODO: implement make json packet
}

// UnHideDebugger (:774)
func UnHideDebugger(hProcess int, PatchAPILevel int)  {
	// TODO: implement make json packet
}

// RelocaterCleanup (:776)
func RelocaterCleanup()  {
	// TODO: implement make json packet
}

// RelocaterInit (:777)
func RelocaterInit(MemorySize int, OldImageBase int, NewImageBase int)  {
	// TODO: implement make json packet
}

// RelocaterAddNewRelocation (:778)
func RelocaterAddNewRelocation(hProcess int, RelocateAddress int, RelocateState int)  {
	// TODO: implement make json packet
}

// RelocaterEstimatedSize (:779)
func RelocaterEstimatedSize()  {
	// TODO: implement make json packet
}

// RelocaterExportRelocation (:780)
func RelocaterExportRelocation(StorePlace int, StorePlaceRVA int, FileMapVA int)  {
	// TODO: implement make json packet
}

// RelocaterExportRelocationEx (:781)
func RelocaterExportRelocationEx(szFileName string, szSectionName string)  {
	// TODO: implement make json packet
}

// RelocaterExportRelocationExW (:782)
func RelocaterExportRelocationExW(szFileName const wint8_t *, szSectionName string)  {
	// TODO: implement make json packet
}

// RelocaterGrabRelocationTable (:783)
func RelocaterGrabRelocationTable(hProcess int, MemoryStart int, MemorySize int)  {
	// TODO: implement make json packet
}

// RelocaterGrabRelocationTableEx (:784)
func RelocaterGrabRelocationTableEx(hProcess int, MemoryStart int, MemorySize int, NtSizeOfImage int)  {
	// TODO: implement make json packet
}

// RelocaterMakeSnapshot (:785)
func RelocaterMakeSnapshot(hProcess int, szSaveFileName string, MemoryStart int, MemorySize int)  {
	// TODO: implement make json packet
}

// RelocaterMakeSnapshotW (:786)
func RelocaterMakeSnapshotW(hProcess int, szSaveFileName const wint8_t *, MemoryStart int, MemorySize int)  {
	// TODO: implement make json packet
}

// RelocaterCompareTwoSnapshots (:787)
func RelocaterCompareTwoSnapshots(hProcess int, LoadedImageBase int, NtSizeOfImage int, szDumpFile1 string, szDumpFile2 string, MemStart int)  {
	// TODO: implement make json packet
}

// RelocaterCompareTwoSnapshotsW (:788)
func RelocaterCompareTwoSnapshotsW(hProcess int, LoadedImageBase int, NtSizeOfImage int, szDumpFile1 const wint8_t *, szDumpFile2 const wint8_t *, MemStart int)  {
	// TODO: implement make json packet
}

// RelocaterChangeFileBase (:789)
func RelocaterChangeFileBase(szFileName string, NewImageBase int)  {
	// TODO: implement make json packet
}

// RelocaterChangeFileBaseW (:790)
func RelocaterChangeFileBaseW(szFileName const wint8_t *, NewImageBase int)  {
	// TODO: implement make json packet
}

// RelocaterRelocateMemoryBlock (:791)
func RelocaterRelocateMemoryBlock(FileMapVA int, MemoryLocation int, RelocateMemory void *, RelocateMemorySize int, CurrentLoadedBase int, RelocateBase int)  {
	// TODO: implement make json packet
}

// RelocaterWipeRelocationTable (:792)
func RelocaterWipeRelocationTable(szFileName string)  {
	// TODO: implement make json packet
}

// RelocaterWipeRelocationTableW (:793)
func RelocaterWipeRelocationTableW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// ResourcerLoadFileForResourceUse (:795)
func ResourcerLoadFileForResourceUse(szFileName string)  {
	// TODO: implement make json packet
}

// ResourcerLoadFileForResourceUseW (:796)
func ResourcerLoadFileForResourceUseW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// ResourcerFreeLoadedFile (:797)
func ResourcerFreeLoadedFile(LoadedFileBase int)  {
	// TODO: implement make json packet
}

// ResourcerExtractResourceFromFileEx (:798)
func ResourcerExtractResourceFromFileEx(hFile int, szResourceType string, szResourceName string, szExtractedFileName string)  {
	// TODO: implement make json packet
}

// ResourcerExtractResourceFromFile (:799)
func ResourcerExtractResourceFromFile(szFileName string, szResourceType string, szResourceName string, szExtractedFileName string)  {
	// TODO: implement make json packet
}

// ResourcerExtractResourceFromFileW (:800)
func ResourcerExtractResourceFromFileW(szFileName const wint8_t *, szResourceType int8 *, szResourceName string, szExtractedFileName string)  {
	// TODO: implement make json packet
}

// ResourcerFindResource (:801)
func ResourcerFindResource(szFileName string, szResourceType string, ResourceType int, szResourceName string, ResourceName int, ResourceLanguage int, pResourceData int, pResourceSize int)  {
	// TODO: implement make json packet
}

// ResourcerFindResourceW (:802)
func ResourcerFindResourceW(szFileName const wint8_t *, szResourceType const wint8_t *, ResourceType int, szResourceName const wint8_t *, ResourceName int, ResourceLanguage int, pResourceData int, pResourceSize int)  {
	// TODO: implement make json packet
}

// ResourcerFindResourceEx (:803)
func ResourcerFindResourceEx(FileMapVA int, FileSize int, szResourceType const wint8_t *, ResourceType int, szResourceName const wint8_t *, ResourceName int, ResourceLanguage int, pResourceData int, pResourceSize int)  {
	// TODO: implement make json packet
}

// ResourcerEnumerateResource (:804)
func ResourcerEnumerateResource(szFileName string, CallBack void *)  {
	// TODO: implement make json packet
}

// ResourcerEnumerateResourceW (:805)
func ResourcerEnumerateResourceW(szFileName const wint8_t *, CallBack void *)  {
	// TODO: implement make json packet
}

// ResourcerEnumerateResourceEx (:806)
func ResourcerEnumerateResourceEx(FileMapVA int, FileSize int, CallBack void *)  {
	// TODO: implement make json packet
}

// ThreaderImportRunningThreadData (:808)
func ThreaderImportRunningThreadData(ProcessId int)  {
	// TODO: implement make json packet
}

// ThreaderGetThreadInfo (:809)
func ThreaderGetThreadInfo(hThread int, ThreadId int)  {
	// TODO: implement make json packet
}

// ThreaderEnumThreadInfo (:810)
func ThreaderEnumThreadInfo(EnumCallBack void *)  {
	// TODO: implement make json packet
}

// ThreaderPauseThread (:811)
func ThreaderPauseThread(hThread int)  {
	// TODO: implement make json packet
}

// ThreaderResumeThread (:812)
func ThreaderResumeThread(hThread int)  {
	// TODO: implement make json packet
}

// ThreaderTerminateThread (:813)
func ThreaderTerminateThread(hThread int, ThreadExitCode int)  {
	// TODO: implement make json packet
}

// ThreaderPauseAllThreads (:814)
func ThreaderPauseAllThreads(LeaveMainRunning bool)  {
	// TODO: implement make json packet
}

// ThreaderResumeAllThreads (:815)
func ThreaderResumeAllThreads(LeaveMainPaused bool)  {
	// TODO: implement make json packet
}

// ThreaderPauseProcess (:816)
func ThreaderPauseProcess()  {
	// TODO: implement make json packet
}

// ThreaderResumeProcess (:817)
func ThreaderResumeProcess()  {
	// TODO: implement make json packet
}

// ThreaderCreateRemoteThread (:818)
func ThreaderCreateRemoteThread(ThreadStartAddress int, AutoCloseTheHandle bool, ThreadPassParameter int, ThreadId int)  {
	// TODO: implement make json packet
}

// ThreaderInjectAndExecuteCode (:819)
func ThreaderInjectAndExecuteCode(InjectCode int, StartDelta int, InjectSize int)  {
	// TODO: implement make json packet
}

// ThreaderCreateRemoteThreadEx (:820)
func ThreaderCreateRemoteThreadEx(hProcess int, ThreadStartAddress int, AutoCloseTheHandle bool, ThreadPassParameter int, ThreadId int)  {
	// TODO: implement make json packet
}

// ThreaderInjectAndExecuteCodeEx (:821)
func ThreaderInjectAndExecuteCodeEx(hProcess int, InjectCode int, StartDelta int, InjectSize int)  {
	// TODO: implement make json packet
}

// ThreaderSetCallBackForNextExitThreadEvent (:822)
func ThreaderSetCallBackForNextExitThreadEvent(exitThreadCallBack int)  {
	// TODO: implement make json packet
}

// ThreaderIsThreadStillRunning (:823)
func ThreaderIsThreadStillRunning(hThread int)  {
	// TODO: implement make json packet
}

// ThreaderIsThreadActive (:824)
func ThreaderIsThreadActive(hThread int)  {
	// TODO: implement make json packet
}

// ThreaderIsAnyThreadActive (:825)
func ThreaderIsAnyThreadActive()  {
	// TODO: implement make json packet
}

// ThreaderExecuteOnlyInjectedThreads (:826)
func ThreaderExecuteOnlyInjectedThreads()  {
	// TODO: implement make json packet
}

// ThreaderGetOpenHandleForThread (:827)
func ThreaderGetOpenHandleForThread(ThreadId int)  {
	// TODO: implement make json packet
}

// ThreaderIsExceptionInMainThread (:828)
func ThreaderIsExceptionInMainThread()  {
	// TODO: implement make json packet
}

// StaticDisassembleEx (:830)
func StaticDisassembleEx(DisassmStart int, DisassmAddress int)  {
	// TODO: implement make json packet
}

// StaticDisassemble (:831)
func StaticDisassemble(DisassmAddress int)  {
	// TODO: implement make json packet
}

// DisassembleEx (:832)
func DisassembleEx(hProcess int, DisassmAddress int, ReturnInstructionType bool)  {
	// TODO: implement make json packet
}

// Disassemble (:833)
func Disassemble(DisassmAddress int)  {
	// TODO: implement make json packet
}

// StaticLengthDisassemble (:834)
func StaticLengthDisassemble(DisassmAddress int)  {
	// TODO: implement make json packet
}

// LengthDisassembleEx (:835)
func LengthDisassembleEx(hProcess int, DisassmAddress int)  {
	// TODO: implement make json packet
}

// LengthDisassemble (:836)
func LengthDisassemble(DisassmAddress int)  {
	// TODO: implement make json packet
}

// InitDebug (:837)
func InitDebug(szFileName int8 *, szCommandLine int8 *, szCurrentFolder int8 *)  {
	// TODO: implement make json packet
}

// InitDebugW (:838)
func InitDebugW(szFileName const wint8_t *, szCommandLine const wint8_t *, szCurrentFolder const wint8_t *)  {
	// TODO: implement make json packet
}

// InitNativeDebug (:839)
func InitNativeDebug(szFileName int8 *, szCommandLine int8 *, szCurrentFolder int8 *)  {
	// TODO: implement make json packet
}

// InitNativeDebugW (:840)
func InitNativeDebugW(szFileName const wint8_t *, szCommandLine const wint8_t *, szCurrentFolder const wint8_t *)  {
	// TODO: implement make json packet
}

// InitDebugEx (:841)
func InitDebugEx(szFileName string, szCommandLine string, szCurrentFolder string, EntryCallBack int)  {
	// TODO: implement make json packet
}

// InitDebugExW (:842)
func InitDebugExW(szFileName const wint8_t *, szCommandLine const wint8_t *, szCurrentFolder const wint8_t *, EntryCallBack int)  {
	// TODO: implement make json packet
}

// InitDLLDebug (:843)
func InitDLLDebug(szFileName string, ReserveModuleBase bool, szCommandLine string, szCurrentFolder string, EntryCallBack int)  {
	// TODO: implement make json packet
}

// InitDLLDebugW (:844)
func InitDLLDebugW(szFileName const wint8_t *, ReserveModuleBase bool, szCommandLine const wint8_t *, szCurrentFolder const wint8_t *, EntryCallBack int)  {
	// TODO: implement make json packet
}

// StopDebug (:845)
func StopDebug()  {
	// TODO: implement make json packet
}

// SetBPXOptions (:846)
func SetBPXOptions(DefaultBreakPointType int32)  {
	// TODO: implement make json packet
}

// IsBPXEnabled (:847)
func IsBPXEnabled(bpxAddress int)  {
	// TODO: implement make json packet
}

// EnableBPX (:848)
func EnableBPX(bpxAddress int)  {
	// TODO: implement make json packet
}

// DisableBPX (:849)
func DisableBPX(bpxAddress int)  {
	// TODO: implement make json packet
}

// SetBPX (:850)
func SetBPX(bpxAddress int, bpxType int, bpxCallBack TITANCBSOFTBP)  {
	// TODO: implement make json packet
}

// DeleteBPX (:851)
func DeleteBPX(bpxAddress int)  {
	// TODO: implement make json packet
}

// SafeDeleteBPX (:852)
func SafeDeleteBPX(bpxAddress int)  {
	// TODO: implement make json packet
}

// SetAPIBreakPoint (:853)
func SetAPIBreakPoint(szDLLName string, szAPIName string, bpxType int, bpxPlace int, bpxCallBack int)  {
	// TODO: implement make json packet
}

// DeleteAPIBreakPoint (:854)
func DeleteAPIBreakPoint(szDLLName string, szAPIName string, bpxPlace int)  {
	// TODO: implement make json packet
}

// SafeDeleteAPIBreakPoint (:855)
func SafeDeleteAPIBreakPoint(szDLLName string, szAPIName string, bpxPlace int)  {
	// TODO: implement make json packet
}

// SetMemoryBPX (:856)
func SetMemoryBPX(MemoryStart int, SizeOfMemory int, bpxCallBack int)  {
	// TODO: implement make json packet
}

// SetMemoryBPXEx (:857)
func SetMemoryBPXEx(MemoryStart int, SizeOfMemory int, BreakPointType int, RestoreOnHit bool, bpxCallBack TITANCBMEMBP)  {
	// TODO: implement make json packet
}

// RemoveMemoryBPX (:858)
func RemoveMemoryBPX(MemoryStart int, SizeOfMemory int)  {
	// TODO: implement make json packet
}

// GetContextFPUDataEx (:859)
func GetContextFPUDataEx(hActiveThread int, FPUSaveArea void *)  {
	// TODO: implement make json packet
}

// Getx87FPURegisters (:860)
func Getx87FPURegisters(x87FPURegisters x87FPURegister_t *, titcontext TITAN_ENGINE_CONTEXT_t *)  {
	// TODO: implement make json packet
}

// GetMMXRegisters (:861)
func GetMMXRegisters(mmx int *, titcontext TITAN_ENGINE_CONTEXT_t *)  {
	// TODO: implement make json packet
}

// GetFullContextDataEx (:862)
func GetFullContextDataEx(hActiveThread int, titcontext TITAN_ENGINE_CONTEXT_t *)  {
	// TODO: implement make json packet
}

// SetFullContextDataEx (:863)
func SetFullContextDataEx(hActiveThread int, titcontext TITAN_ENGINE_CONTEXT_t *)  {
	// TODO: implement make json packet
}

// GetContextDataEx (:864)
func GetContextDataEx(hActiveThread int, IndexOfRegister int)  {
	// TODO: implement make json packet
}

// GetContextData (:865)
func GetContextData(IndexOfRegister int)  {
	// TODO: implement make json packet
}

// SetContextFPUDataEx (:866)
func SetContextFPUDataEx(hActiveThread int, FPUSaveArea void *)  {
	// TODO: implement make json packet
}

// SetContextDataEx (:867)
func SetContextDataEx(hActiveThread int, IndexOfRegister int, NewRegisterValue int)  {
	// TODO: implement make json packet
}

// SetContextData (:868)
func SetContextData(IndexOfRegister int, NewRegisterValue int)  {
	// TODO: implement make json packet
}

// GetAVXContext (:869)
func GetAVXContext(hActiveThread int, titcontext TITAN_ENGINE_CONTEXT_t *)  {
	// TODO: implement make json packet
}

// SetAVXContext (:870)
func SetAVXContext(hActiveThread int, titcontext TITAN_ENGINE_CONTEXT_t *)  {
	// TODO: implement make json packet
}

// ClearExceptionNumber (:871)
func ClearExceptionNumber()  {
	// TODO: implement make json packet
}

// CurrentExceptionNumber (:872)
func CurrentExceptionNumber()  {
	// TODO: implement make json packet
}

// MatchPatternEx (:873)
func MatchPatternEx(hProcess int, MemoryToCheck void *, SizeOfMemoryToCheck int, PatternToMatch void *, SizeOfPatternToMatch int, WildCard int)  {
	// TODO: implement make json packet
}

// MatchPattern (:874)
func MatchPattern(MemoryToCheck void *, SizeOfMemoryToCheck int, PatternToMatch void *, SizeOfPatternToMatch int, WildCard int)  {
	// TODO: implement make json packet
}

// FindEx (:875)
func FindEx(hProcess int, MemoryStart int, MemorySize int, SearchPattern int, PatternSize int, WildCard int)  {
	// TODO: implement make json packet
}

// Find (:10)
func Find(data byte *, datasize duint, pattern string)  {
	// TODO: implement make json packet
}

// FillEx (:877)
func FillEx(hProcess int, MemoryStart int, MemorySize int, FillByte int)  {
	// TODO: implement make json packet
}

// Fill (:878)
func Fill(MemoryStart int, MemorySize int, FillByte int)  {
	// TODO: implement make json packet
}

// PatchEx (:879)
func PatchEx(hProcess int, MemoryStart int, MemorySize int, ReplacePattern int, ReplaceSize int, AppendNOP bool, PrependNOP bool)  {
	// TODO: implement make json packet
}

// Patch (:880)
func Patch(MemoryStart int, MemorySize int, ReplacePattern int, ReplaceSize int, AppendNOP bool, PrependNOP bool)  {
	// TODO: implement make json packet
}

// ReplaceEx (:881)
func ReplaceEx(hProcess int, MemoryStart int, MemorySize int, SearchPattern int, PatternSize int, NumberOfRepetitions int, ReplacePattern int, ReplaceSize int, WildCard int)  {
	// TODO: implement make json packet
}

// Replace (:882)
func Replace(MemoryStart int, MemorySize int, SearchPattern int, PatternSize int, NumberOfRepetitions int, ReplacePattern int, ReplaceSize int, WildCard int)  {
	// TODO: implement make json packet
}

// GetDebugData (:883)
func GetDebugData()  {
	// TODO: implement make json packet
}

// GetTerminationData (:884)
func GetTerminationData()  {
	// TODO: implement make json packet
}

// GetExitCode (:885)
func GetExitCode()  {
	// TODO: implement make json packet
}

// GetDebuggedDLLBaseAddress (:886)
func GetDebuggedDLLBaseAddress()  {
	// TODO: implement make json packet
}

// GetDebuggedFileBaseAddress (:887)
func GetDebuggedFileBaseAddress()  {
	// TODO: implement make json packet
}

// GetRemoteString (:888)
func GetRemoteString(hProcess int, StringAddress int, StringStorage int, MaximumStringSize int)  {
	// TODO: implement make json packet
}

// GetFunctionParameter (:889)
func GetFunctionParameter(hProcess int, FunctionType int, ParameterNumber int, ParameterType int)  {
	// TODO: implement make json packet
}

// GetJumpDestinationEx (:890)
func GetJumpDestinationEx(hProcess int, InstructionAddress int, JustJumps bool)  {
	// TODO: implement make json packet
}

// GetJumpDestination (:891)
func GetJumpDestination(hProcess int, InstructionAddress int)  {
	// TODO: implement make json packet
}

// IsJumpGoingToExecuteEx (:892)
func IsJumpGoingToExecuteEx(hProcess int, hThread int, InstructionAddress int, RegFlags int)  {
	// TODO: implement make json packet
}

// IsJumpGoingToExecute (:893)
func IsJumpGoingToExecute()  {
	// TODO: implement make json packet
}

// SetCustomHandler (:894)
func SetCustomHandler(ExceptionId int, CallBack TITANCBCH)  {
	// TODO: implement make json packet
}

// ForceClose (:895)
func ForceClose()  {
	// TODO: implement make json packet
}

// StepInto (:896)
func StepInto(traceCallBack TITANCBSTEP)  {
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

// SingleStep (:899)
func SingleStep(StepCount int, StepCallBack TITANCBSTEP)  {
	// TODO: implement make json packet
}

// GetUnusedHardwareBreakPointRegister (:900)
func GetUnusedHardwareBreakPointRegister(RegisterIndex int)  {
	// TODO: implement make json packet
}

// SetHardwareBreakPointEx (:901)
func SetHardwareBreakPointEx(hActiveThread int, bpxAddress int, IndexOfRegister int, bpxType int, bpxSize int, bpxCallBack TITANCBHWBP, IndexOfSelectedRegister int)  {
	// TODO: implement make json packet
}

// SetHardwareBreakPoint (:902)
func SetHardwareBreakPoint(bpxAddress int, IndexOfRegister int, bpxType int, bpxSize int, bpxCallBack TITANCBHWBP)  {
	// TODO: implement make json packet
}

// DeleteHardwareBreakPoint (:903)
func DeleteHardwareBreakPoint(IndexOfRegister int)  {
	// TODO: implement make json packet
}

// RemoveAllBreakPoints (:904)
func RemoveAllBreakPoints(RemoveOption int)  {
	// TODO: implement make json packet
}

// TitanGetProcessInformation (:905)
func TitanGetProcessInformation()  {
	// TODO: implement make json packet
}

// TitanGetStartupInformation (:906)
func TitanGetStartupInformation()  {
	// TODO: implement make json packet
}

// DebugLoop (:907)
func DebugLoop()  {
	// TODO: implement make json packet
}

// SetDebugLoopTimeOut (:908)
func SetDebugLoopTimeOut(TimeOut int)  {
	// TODO: implement make json packet
}

// SetNextDbgContinueStatus (:909)
func SetNextDbgContinueStatus(SetDbgCode int)  {
	// TODO: implement make json packet
}

// AttachDebugger (:910)
func AttachDebugger(ProcessId int, KillOnExit bool, DebugInfo int, CallBack TITANCALLBACK)  {
	// TODO: implement make json packet
}

// DetachDebugger (:911)
func DetachDebugger(ProcessId int)  {
	// TODO: implement make json packet
}

// DetachDebuggerEx (:912)
func DetachDebuggerEx(ProcessId int)  {
	// TODO: implement make json packet
}

// DebugLoopEx (:913)
func DebugLoopEx(TimeOut int)  {
	// TODO: implement make json packet
}

// AutoDebugEx (:914)
func AutoDebugEx(szFileName string, ReserveModuleBase bool, szCommandLine string, szCurrentFolder string, TimeOut int, EntryCallBack TITANCBSOFTBP)  {
	// TODO: implement make json packet
}

// AutoDebugExW (:915)
func AutoDebugExW(szFileName const wint8_t *, ReserveModuleBase bool, szCommandLine const wint8_t *, szCurrentFolder const wint8_t *, TimeOut int, EntryCallBack TITANCBSOFTBP)  {
	// TODO: implement make json packet
}

// IsFileBeingDebugged (:916)
func IsFileBeingDebugged()  {
	// TODO: implement make json packet
}

// SetErrorModel (:917)
func SetErrorModel(DisplayErrorMessages bool)  {
	// TODO: implement make json packet
}

// FindOEPInit (:919)
func FindOEPInit()  {
	// TODO: implement make json packet
}

// FindOEPGenerically (:920)
func FindOEPGenerically(szFileName string, TraceInitCallBack int, CallBack int)  {
	// TODO: implement make json packet
}

// FindOEPGenericallyW (:921)
func FindOEPGenericallyW(szFileName const wint8_t *, TraceInitCallBack int, CallBack int)  {
	// TODO: implement make json packet
}

// ImporterAddNewDll (:923)
func ImporterAddNewDll(szDLLName string, FirstThunk int)  {
	// TODO: implement make json packet
}

// ImporterAddNewAPI (:924)
func ImporterAddNewAPI(szAPIName string, ThunkValue int)  {
	// TODO: implement make json packet
}

// ImporterAddNewOrdinalAPI (:925)
func ImporterAddNewOrdinalAPI(OrdinalNumber int, ThunkValue int)  {
	// TODO: implement make json packet
}

// ImporterGetAddedDllCount (:926)
func ImporterGetAddedDllCount()  {
	// TODO: implement make json packet
}

// ImporterGetAddedAPICount (:927)
func ImporterGetAddedAPICount()  {
	// TODO: implement make json packet
}

// ImporterExportIAT (:928)
func ImporterExportIAT(StorePlace int, FileMapVA int, hFileMap int)  {
	// TODO: implement make json packet
}

// ImporterEstimatedSize (:929)
func ImporterEstimatedSize()  {
	// TODO: implement make json packet
}

// ImporterExportIATEx (:930)
func ImporterExportIATEx(szDumpFileName string, szExportFileName string, szSectionName string)  {
	// TODO: implement make json packet
}

// ImporterExportIATExW (:931)
func ImporterExportIATExW(szDumpFileName const wint8_t *, szExportFileName const wint8_t *, szSectionName const wint8_t *)  {
	// TODO: implement make json packet
}

// ImporterFindAPIWriteLocation (:932)
func ImporterFindAPIWriteLocation(szAPIName string)  {
	// TODO: implement make json packet
}

// ImporterFindOrdinalAPIWriteLocation (:933)
func ImporterFindOrdinalAPIWriteLocation(OrdinalNumber int)  {
	// TODO: implement make json packet
}

// ImporterFindAPIByWriteLocation (:934)
func ImporterFindAPIByWriteLocation(APIWriteLocation int)  {
	// TODO: implement make json packet
}

// ImporterFindDLLByWriteLocation (:935)
func ImporterFindDLLByWriteLocation(APIWriteLocation int)  {
	// TODO: implement make json packet
}

// ImporterGetDLLName (:936)
func ImporterGetDLLName(APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetDLLNameW (:937)
func ImporterGetDLLNameW(APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetAPIName (:938)
func ImporterGetAPIName(APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetAPIOrdinalNumber (:939)
func ImporterGetAPIOrdinalNumber(APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetAPINameEx (:940)
func ImporterGetAPINameEx(APIAddress int, DLLBasesList int)  {
	// TODO: implement make json packet
}

// ImporterGetRemoteAPIAddress (:941)
func ImporterGetRemoteAPIAddress(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetRemoteAPIAddressEx (:942)
func ImporterGetRemoteAPIAddressEx(szDLLName string, szAPIName string)  {
	// TODO: implement make json packet
}

// ImporterGetLocalAPIAddress (:943)
func ImporterGetLocalAPIAddress(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetDLLNameFromDebugee (:944)
func ImporterGetDLLNameFromDebugee(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetDLLNameFromDebugeeW (:945)
func ImporterGetDLLNameFromDebugeeW(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetAPINameFromDebugee (:946)
func ImporterGetAPINameFromDebugee(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetAPIOrdinalNumberFromDebugee (:947)
func ImporterGetAPIOrdinalNumberFromDebugee(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetDLLIndexEx (:948)
func ImporterGetDLLIndexEx(APIAddress int, DLLBasesList int)  {
	// TODO: implement make json packet
}

// ImporterGetDLLIndex (:949)
func ImporterGetDLLIndex(hProcess int, APIAddress int, DLLBasesList int)  {
	// TODO: implement make json packet
}

// ImporterGetRemoteDLLBase (:950)
func ImporterGetRemoteDLLBase(hProcess int, LocalModuleBase int)  {
	// TODO: implement make json packet
}

// ImporterGetRemoteDLLBaseEx (:951)
func ImporterGetRemoteDLLBaseEx(hProcess int, szModuleName string)  {
	// TODO: implement make json packet
}

// ImporterGetRemoteDLLBaseExW (:952)
func ImporterGetRemoteDLLBaseExW(hProcess int, szModuleName const wint8_t *)  {
	// TODO: implement make json packet
}

// ImporterIsForwardedAPI (:953)
func ImporterIsForwardedAPI(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetForwardedAPIName (:954)
func ImporterGetForwardedAPIName(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetForwardedDLLName (:955)
func ImporterGetForwardedDLLName(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetForwardedDLLIndex (:956)
func ImporterGetForwardedDLLIndex(hProcess int, APIAddress int, DLLBasesList int)  {
	// TODO: implement make json packet
}

// ImporterGetForwardedAPIOrdinalNumber (:957)
func ImporterGetForwardedAPIOrdinalNumber(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetNearestAPIAddress (:958)
func ImporterGetNearestAPIAddress(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterGetNearestAPIName (:959)
func ImporterGetNearestAPIName(hProcess int, APIAddress int)  {
	// TODO: implement make json packet
}

// ImporterCopyOriginalIAT (:960)
func ImporterCopyOriginalIAT(szOriginalFile string, szDumpFile string)  {
	// TODO: implement make json packet
}

// ImporterCopyOriginalIATW (:961)
func ImporterCopyOriginalIATW(szOriginalFile const wint8_t *, szDumpFile const wint8_t *)  {
	// TODO: implement make json packet
}

// ImporterLoadImportTable (:962)
func ImporterLoadImportTable(szFileName string)  {
	// TODO: implement make json packet
}

// ImporterLoadImportTableW (:963)
func ImporterLoadImportTableW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// ImporterMoveOriginalIAT (:964)
func ImporterMoveOriginalIAT(szOriginalFile string, szDumpFile string, szSectionName string)  {
	// TODO: implement make json packet
}

// ImporterMoveOriginalIATW (:965)
func ImporterMoveOriginalIATW(szOriginalFile const wint8_t *, szDumpFile const wint8_t *, szSectionName string)  {
	// TODO: implement make json packet
}

// ImporterAutoSearchIAT (:966)
func ImporterAutoSearchIAT(ProcessId int, szFileName string, SearchStart int, pIATStart int, pIATSize int)  {
	// TODO: implement make json packet
}

// ImporterAutoSearchIATW (:967)
func ImporterAutoSearchIATW(ProcessIds int, szFileName const wint8_t *, SearchStart int, pIATStart int, pIATSize int)  {
	// TODO: implement make json packet
}

// ImporterAutoSearchIATEx (:968)
func ImporterAutoSearchIATEx(ProcessId int, ImageBase int, SearchStart int, pIATStart int, pIATSize int)  {
	// TODO: implement make json packet
}

// ImporterEnumAddedData (:969)
func ImporterEnumAddedData(EnumCallBack int)  {
	// TODO: implement make json packet
}

// ImporterAutoFixIATEx (:970)
func ImporterAutoFixIATEx(ProcessId int, szDumpedFile string, szSectionName string, DumpRunningProcess bool, RealignFile bool, EntryPointAddress int, ImageBase int, SearchStart int, TryAutoFix bool, FixEliminations bool, UnknownPointerFixCallback int)  {
	// TODO: implement make json packet
}

// ImporterAutoFixIATExW (:971)
func ImporterAutoFixIATExW(ProcessId int, szDumpedFile const wint8_t *, szSectionName const wint8_t *, DumpRunningProcess bool, RealignFile bool, EntryPointAddress int, ImageBase int, SearchStart int, TryAutoFix bool, FixEliminations bool, UnknownPointerFixCallback int)  {
	// TODO: implement make json packet
}

// ImporterAutoFixIAT (:972)
func ImporterAutoFixIAT(ProcessId int, szDumpedFile string, SearchStart int)  {
	// TODO: implement make json packet
}

// ImporterAutoFixIATW (:973)
func ImporterAutoFixIATW(ProcessId int, szDumpedFile const wint8_t *, SearchStart int)  {
	// TODO: implement make json packet
}

// ImporterDeleteAPI (:974)
func ImporterDeleteAPI(apiAddr int)  {
	// TODO: implement make json packet
}

// HooksSafeTransitionEx (:976)
func HooksSafeTransitionEx(HookAddressArray int, NumberOfHooks int, TransitionStart bool)  {
	// TODO: implement make json packet
}

// HooksSafeTransition (:977)
func HooksSafeTransition(HookAddress int, TransitionStart bool)  {
	// TODO: implement make json packet
}

// HooksIsAddressRedirected (:978)
func HooksIsAddressRedirected(HookAddress int)  {
	// TODO: implement make json packet
}

// HooksGetTrampolineAddress (:979)
func HooksGetTrampolineAddress(HookAddress int)  {
	// TODO: implement make json packet
}

// HooksGetHookEntryDetails (:980)
func HooksGetHookEntryDetails(HookAddress int)  {
	// TODO: implement make json packet
}

// HooksInsertNewRedirection (:981)
func HooksInsertNewRedirection(HookAddress int, RedirectTo int, HookType int)  {
	// TODO: implement make json packet
}

// HooksInsertNewIATRedirectionEx (:982)
func HooksInsertNewIATRedirectionEx(FileMapVA int, LoadedModuleBase int, szHookFunction string, RedirectTo int)  {
	// TODO: implement make json packet
}

// HooksInsertNewIATRedirection (:983)
func HooksInsertNewIATRedirection(szModuleName string, szHookFunction string, RedirectTo int)  {
	// TODO: implement make json packet
}

// HooksRemoveRedirection (:984)
func HooksRemoveRedirection(HookAddress int, RemoveAll bool)  {
	// TODO: implement make json packet
}

// HooksRemoveRedirectionsForModule (:985)
func HooksRemoveRedirectionsForModule(ModuleBase int)  {
	// TODO: implement make json packet
}

// HooksRemoveIATRedirection (:986)
func HooksRemoveIATRedirection(szModuleName string, szHookFunction string, RemoveAll bool)  {
	// TODO: implement make json packet
}

// HooksDisableRedirection (:987)
func HooksDisableRedirection(HookAddress int, DisableAll bool)  {
	// TODO: implement make json packet
}

// HooksDisableRedirectionsForModule (:988)
func HooksDisableRedirectionsForModule(ModuleBase int)  {
	// TODO: implement make json packet
}

// HooksDisableIATRedirection (:989)
func HooksDisableIATRedirection(szModuleName string, szHookFunction string, DisableAll bool)  {
	// TODO: implement make json packet
}

// HooksEnableRedirection (:990)
func HooksEnableRedirection(HookAddress int, EnableAll bool)  {
	// TODO: implement make json packet
}

// HooksEnableRedirectionsForModule (:991)
func HooksEnableRedirectionsForModule(ModuleBase int)  {
	// TODO: implement make json packet
}

// HooksEnableIATRedirection (:992)
func HooksEnableIATRedirection(szModuleName string, szHookFunction string, EnableAll bool)  {
	// TODO: implement make json packet
}

// HooksScanModuleMemory (:993)
func HooksScanModuleMemory(ModuleBase int, CallBack int)  {
	// TODO: implement make json packet
}

// HooksScanEntireProcessMemory (:994)
func HooksScanEntireProcessMemory(CallBack int)  {
	// TODO: implement make json packet
}

// HooksScanEntireProcessMemoryEx (:995)
func HooksScanEntireProcessMemoryEx()  {
	// TODO: implement make json packet
}

// TracerInit (:997)
func TracerInit()  {
	// TODO: implement make json packet
}

// TracerLevel1 (:998)
func TracerLevel1(hProcess int, AddressToTrace int)  {
	// TODO: implement make json packet
}

// HashTracerLevel1 (:999)
func HashTracerLevel1(hProcess int, AddressToTrace int, InputNumberOfInstructions int)  {
	// TODO: implement make json packet
}

// TracerDetectRedirection (:1000)
func TracerDetectRedirection(hProcess int, AddressToTrace int)  {
	// TODO: implement make json packet
}

// TracerFixKnownRedirection (:1001)
func TracerFixKnownRedirection(hProcess int, AddressToTrace int, RedirectionId int)  {
	// TODO: implement make json packet
}

// TracerFixRedirectionViaModule (:1002)
func TracerFixRedirectionViaModule(hModuleHandle int, hProcess int, AddressToTrace int, IdParameter int)  {
	// TODO: implement make json packet
}

// TracerFixRedirectionViaImpRecPlugin (:1003)
func TracerFixRedirectionViaImpRecPlugin(hProcess int, szPluginName string, AddressToTrace int)  {
	// TODO: implement make json packet
}

// ExporterCleanup (:1005)
func ExporterCleanup()  {
	// TODO: implement make json packet
}

// ExporterSetImageBase (:1006)
func ExporterSetImageBase(ImageBase int)  {
	// TODO: implement make json packet
}

// ExporterInit (:1007)
func ExporterInit(MemorySize int, ImageBase int, ExportOrdinalBase int, szExportModuleName string)  {
	// TODO: implement make json packet
}

// ExporterAddNewExport (:1008)
func ExporterAddNewExport(szExportName string, ExportRelativeAddress int)  {
	// TODO: implement make json packet
}

// ExporterAddNewOrdinalExport (:1009)
func ExporterAddNewOrdinalExport(OrdinalNumber int, ExportRelativeAddress int)  {
	// TODO: implement make json packet
}

// ExporterGetAddedExportCount (:1010)
func ExporterGetAddedExportCount()  {
	// TODO: implement make json packet
}

// ExporterEstimatedSize (:1011)
func ExporterEstimatedSize()  {
	// TODO: implement make json packet
}

// ExporterBuildExportTable (:1012)
func ExporterBuildExportTable(StorePlace int, FileMapVA int)  {
	// TODO: implement make json packet
}

// ExporterBuildExportTableEx (:1013)
func ExporterBuildExportTableEx(szExportFileName string, szSectionName string)  {
	// TODO: implement make json packet
}

// ExporterBuildExportTableExW (:1014)
func ExporterBuildExportTableExW(szExportFileName const wint8_t *, szSectionName string)  {
	// TODO: implement make json packet
}

// ExporterLoadExportTable (:1015)
func ExporterLoadExportTable(szFileName string)  {
	// TODO: implement make json packet
}

// ExporterLoadExportTableW (:1016)
func ExporterLoadExportTableW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// LibrarianSetBreakPoint (:1018)
func LibrarianSetBreakPoint(szLibraryName string, bpxType int, SingleShoot bool, bpxCallBack int)  {
	// TODO: implement make json packet
}

// LibrarianRemoveBreakPoint (:1019)
func LibrarianRemoveBreakPoint(szLibraryName string, bpxType int)  {
	// TODO: implement make json packet
}

// LibrarianGetLibraryInfo (:1020)
func LibrarianGetLibraryInfo(szLibraryName string)  {
	// TODO: implement make json packet
}

// LibrarianGetLibraryInfoW (:1021)
func LibrarianGetLibraryInfoW(szLibraryName const wint8_t *)  {
	// TODO: implement make json packet
}

// LibrarianGetLibraryInfoEx (:1022)
func LibrarianGetLibraryInfoEx(BaseOfDll void *)  {
	// TODO: implement make json packet
}

// LibrarianGetLibraryInfoExW (:1023)
func LibrarianGetLibraryInfoExW(BaseOfDll void *)  {
	// TODO: implement make json packet
}

// LibrarianEnumLibraryInfo (:1024)
func LibrarianEnumLibraryInfo(EnumCallBack void *)  {
	// TODO: implement make json packet
}

// LibrarianEnumLibraryInfoW (:1025)
func LibrarianEnumLibraryInfoW(EnumCallBack void *)  {
	// TODO: implement make json packet
}

// GetActiveProcessId (:1027)
func GetActiveProcessId(szImageName string)  {
	// TODO: implement make json packet
}

// GetActiveProcessIdW (:1028)
func GetActiveProcessIdW(szImageName const wint8_t *)  {
	// TODO: implement make json packet
}

// EnumProcessesWithLibrary (:1029)
func EnumProcessesWithLibrary(szLibraryName string, EnumFunction void *)  {
	// TODO: implement make json packet
}

// TitanOpenProcess (:1030)
func TitanOpenProcess(dwDesiredAccess int, bInheritHandle bool, dwProcessId int)  {
	// TODO: implement make json packet
}

// TitanOpenThread (:1031)
func TitanOpenThread(dwDesiredAccess int, bInheritHandle bool, dwThreadId int)  {
	// TODO: implement make json packet
}

// TLSBreakOnCallBack (:1033)
func TLSBreakOnCallBack(ArrayOfCallBacks int, NumberOfCallBacks int, bpxCallBack int)  {
	// TODO: implement make json packet
}

// TLSGrabCallBackData (:1034)
func TLSGrabCallBackData(szFileName string, ArrayOfCallBacks int, NumberOfCallBacks int)  {
	// TODO: implement make json packet
}

// TLSGrabCallBackDataW (:1035)
func TLSGrabCallBackDataW(szFileName const wint8_t *, ArrayOfCallBacks int, NumberOfCallBacks int)  {
	// TODO: implement make json packet
}

// TLSBreakOnCallBackEx (:1036)
func TLSBreakOnCallBackEx(szFileName string, bpxCallBack int)  {
	// TODO: implement make json packet
}

// TLSBreakOnCallBackExW (:1037)
func TLSBreakOnCallBackExW(szFileName const wint8_t *, bpxCallBack int)  {
	// TODO: implement make json packet
}

// TLSRemoveCallback (:1038)
func TLSRemoveCallback(szFileName string)  {
	// TODO: implement make json packet
}

// TLSRemoveCallbackW (:1039)
func TLSRemoveCallbackW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// TLSRemoveTable (:1040)
func TLSRemoveTable(szFileName string)  {
	// TODO: implement make json packet
}

// TLSRemoveTableW (:1041)
func TLSRemoveTableW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// TLSBackupData (:1042)
func TLSBackupData(szFileName string)  {
	// TODO: implement make json packet
}

// TLSBackupDataW (:1043)
func TLSBackupDataW(szFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// TLSRestoreData (:1044)
func TLSRestoreData()  {
	// TODO: implement make json packet
}

// TLSBuildNewTable (:1045)
func TLSBuildNewTable(FileMapVA int, StorePlace int, StorePlaceRVA int, ArrayOfCallBacks int, NumberOfCallBacks int)  {
	// TODO: implement make json packet
}

// TLSBuildNewTableEx (:1046)
func TLSBuildNewTableEx(szFileName string, szSectionName string, ArrayOfCallBacks int, NumberOfCallBacks int)  {
	// TODO: implement make json packet
}

// TLSBuildNewTableExW (:1047)
func TLSBuildNewTableExW(szFileName const wint8_t *, szSectionName string, ArrayOfCallBacks int, NumberOfCallBacks int)  {
	// TODO: implement make json packet
}

// TranslateNativeName (:1049)
func TranslateNativeName(szNativeName string)  {
	// TODO: implement make json packet
}

// TranslateNativeNameW (:1050)
func TranslateNativeNameW(szNativeName const wint8_t *)  {
	// TODO: implement make json packet
}

// HandlerGetActiveHandleCount (:1052)
func HandlerGetActiveHandleCount(ProcessId int)  {
	// TODO: implement make json packet
}

// HandlerIsHandleOpen (:1053)
func HandlerIsHandleOpen(ProcessId int, hHandle int)  {
	// TODO: implement make json packet
}

// HandlerGetHandleName (:1054)
func HandlerGetHandleName(hProcess int, ProcessId int, hHandle int, TranslateName bool)  {
	// TODO: implement make json packet
}

// HandlerGetHandleNameW (:1055)
func HandlerGetHandleNameW(hProcess int, ProcessId int, hHandle int, TranslateName bool)  {
	// TODO: implement make json packet
}

// HandlerEnumerateOpenHandles (:1056)
func HandlerEnumerateOpenHandles(ProcessId int, HandleBuffer int, MaxHandleCount int)  {
	// TODO: implement make json packet
}

// HandlerGetHandleDetails (:1057)
func HandlerGetHandleDetails(hProcess int, ProcessId int, hHandle int, InformationReturn int)  {
	// TODO: implement make json packet
}

// HandlerCloseRemoteHandle (:1058)
func HandlerCloseRemoteHandle(hProcess int, hHandle int)  {
	// TODO: implement make json packet
}

// HandlerEnumerateLockHandles (:1059)
func HandlerEnumerateLockHandles(szFileOrFolderName int8 *, NameIsFolder bool, NameIsTranslated bool, HandleDataBuffer int, MaxHandleCount int)  {
	// TODO: implement make json packet
}

// HandlerEnumerateLockHandlesW (:1060)
func HandlerEnumerateLockHandlesW(szFileOrFolderName const wint8_t *, NameIsFolder bool, NameIsTranslated bool, HandleDataBuffer int, MaxHandleCount int)  {
	// TODO: implement make json packet
}

// HandlerCloseAllLockHandles (:1061)
func HandlerCloseAllLockHandles(szFileOrFolderName string, NameIsFolder bool, NameIsTranslated bool)  {
	// TODO: implement make json packet
}

// HandlerCloseAllLockHandlesW (:1062)
func HandlerCloseAllLockHandlesW(szFileOrFolderName const wint8_t *, NameIsFolder bool, NameIsTranslated bool)  {
	// TODO: implement make json packet
}

// HandlerIsFileLocked (:1063)
func HandlerIsFileLocked(szFileOrFolderName string, NameIsFolder bool, NameIsTranslated bool)  {
	// TODO: implement make json packet
}

// HandlerIsFileLockedW (:1064)
func HandlerIsFileLockedW(szFileOrFolderName const wint8_t *, NameIsFolder bool, NameIsTranslated bool)  {
	// TODO: implement make json packet
}

// HandlerEnumerateOpenMutexes (:1066)
func HandlerEnumerateOpenMutexes(hProcess int, ProcessId int, HandleBuffer int, MaxHandleCount int)  {
	// TODO: implement make json packet
}

// HandlerGetOpenMutexHandle (:1067)
func HandlerGetOpenMutexHandle(hProcess int, ProcessId int, szMutexString string)  {
	// TODO: implement make json packet
}

// HandlerGetOpenMutexHandleW (:1068)
func HandlerGetOpenMutexHandleW(hProcess int, ProcessId int, szMutexString const wint8_t *)  {
	// TODO: implement make json packet
}

// HandlerGetProcessIdWhichCreatedMutex (:1069)
func HandlerGetProcessIdWhichCreatedMutex(szMutexString string)  {
	// TODO: implement make json packet
}

// HandlerGetProcessIdWhichCreatedMutexW (:1070)
func HandlerGetProcessIdWhichCreatedMutexW(szMutexString const wint8_t *)  {
	// TODO: implement make json packet
}

// RemoteLoadLibrary (:1072)
func RemoteLoadLibrary(hProcess int, szLibraryFile string, WaitForThreadExit bool)  {
	// TODO: implement make json packet
}

// RemoteLoadLibraryW (:1073)
func RemoteLoadLibraryW(hProcess int, szLibraryFile const wint8_t *, WaitForThreadExit bool)  {
	// TODO: implement make json packet
}

// RemoteFreeLibrary (:1074)
func RemoteFreeLibrary(hProcess int, hModule int, szLibraryFile string, WaitForThreadExit bool)  {
	// TODO: implement make json packet
}

// RemoteFreeLibraryW (:1075)
func RemoteFreeLibraryW(hProcess int, hModule int, szLibraryFile const wint8_t *, WaitForThreadExit bool)  {
	// TODO: implement make json packet
}

// RemoteExitProcess (:1076)
func RemoteExitProcess(hProcess int, ExitCode int)  {
	// TODO: implement make json packet
}

// StaticFileLoad (:1078)
func StaticFileLoad(szFileName string, DesiredAccess int, SimulateLoad bool, FileHandle int, LoadedSize int, FileMap int, FileMapVA int)  {
	// TODO: implement make json packet
}

// StaticFileLoadW (:1079)
func StaticFileLoadW(szFileName const wint8_t *, DesiredAccess int, SimulateLoad bool, FileHandle int, LoadedSize int, FileMap int, FileMapVA int)  {
	// TODO: implement make json packet
}

// StaticFileUnload (:1080)
func StaticFileUnload(szFileName string, CommitChanges bool, FileHandle int, LoadedSize int, FileMap int, FileMapVA int)  {
	// TODO: implement make json packet
}

// StaticFileUnloadW (:1081)
func StaticFileUnloadW(szFileName const wint8_t *, CommitChanges bool, FileHandle int, LoadedSize int, FileMap int, FileMapVA int)  {
	// TODO: implement make json packet
}

// StaticFileOpen (:1082)
func StaticFileOpen(szFileName string, DesiredAccess int, FileHandle int, FileSizeLow int, FileSizeHigh int)  {
	// TODO: implement make json packet
}

// StaticFileOpenW (:1083)
func StaticFileOpenW(szFileName const wint8_t *, DesiredAccess int, FileHandle int, FileSizeLow int, FileSizeHigh int)  {
	// TODO: implement make json packet
}

// StaticFileGetContent (:1084)
func StaticFileGetContent(FileHandle int, FilePositionLow int, FilePositionHigh int, Buffer void *, Size int)  {
	// TODO: implement make json packet
}

// StaticFileClose (:1085)
func StaticFileClose(FileHandle int)  {
	// TODO: implement make json packet
}

// StaticMemoryDecrypt (:1086)
func StaticMemoryDecrypt(MemoryStart int, MemorySize int, DecryptionType int, DecryptionKeySize int, DecryptionKey int)  {
	// TODO: implement make json packet
}

// StaticMemoryDecryptEx (:1087)
func StaticMemoryDecryptEx(MemoryStart int, MemorySize int, DecryptionKeySize int, DecryptionCallBack void *)  {
	// TODO: implement make json packet
}

// StaticMemoryDecryptSpecial (:1088)
func StaticMemoryDecryptSpecial(MemoryStart int, MemorySize int, DecryptionKeySize int, SpecDecryptionType int, DecryptionCallBack void *)  {
	// TODO: implement make json packet
}

// StaticSectionDecrypt (:1089)
func StaticSectionDecrypt(FileMapVA int, SectionNumber int, SimulateLoad bool, DecryptionType int, DecryptionKeySize int, DecryptionKey int)  {
	// TODO: implement make json packet
}

// StaticMemoryDecompress (:1090)
func StaticMemoryDecompress(Source void *, SourceSize int, Destination void *, DestinationSize int, Algorithm int)  {
	// TODO: implement make json packet
}

// StaticRawMemoryCopy (:1091)
func StaticRawMemoryCopy(hFile int, FileMapVA int, VitualAddressToCopy int, Size int, AddressIsRVA bool, szDumpFileName string)  {
	// TODO: implement make json packet
}

// StaticRawMemoryCopyW (:1092)
func StaticRawMemoryCopyW(hFile int, FileMapVA int, VitualAddressToCopy int, Size int, AddressIsRVA bool, szDumpFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// StaticRawMemoryCopyEx (:1093)
func StaticRawMemoryCopyEx(hFile int, RawAddressToCopy int, Size int, szDumpFileName string)  {
	// TODO: implement make json packet
}

// StaticRawMemoryCopyExW (:1094)
func StaticRawMemoryCopyExW(hFile int, RawAddressToCopy int, Size int, szDumpFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// StaticRawMemoryCopyEx64 (:1095)
func StaticRawMemoryCopyEx64(hFile int, RawAddressToCopy int, Size int, szDumpFileName string)  {
	// TODO: implement make json packet
}

// StaticRawMemoryCopyEx64W (:1096)
func StaticRawMemoryCopyEx64W(hFile int, RawAddressToCopy int, Size int, szDumpFileName const wint8_t *)  {
	// TODO: implement make json packet
}

// StaticHashMemory (:1097)
func StaticHashMemory(MemoryToHash void *, SizeOfMemory int, HashDigest void *, OutputString bool, Algorithm int)  {
	// TODO: implement make json packet
}

// StaticHashFileW (:1098)
func StaticHashFileW(szFileName const wint8_t *, HashDigest int8 *, OutputString bool, Algorithm int)  {
	// TODO: implement make json packet
}

// StaticHashFile (:1099)
func StaticHashFile(szFileName string, HashDigest int8 *, OutputString bool, Algorithm int)  {
	// TODO: implement make json packet
}

// EngineUnpackerInitialize (:1101)
func EngineUnpackerInitialize(szFileName string, szUnpackedFileName string, DoLogData bool, DoRealignFile bool, DoMoveOverlay bool, EntryCallBack void *)  {
	// TODO: implement make json packet
}

// EngineUnpackerInitializeW (:1102)
func EngineUnpackerInitializeW(szFileName const wint8_t *, szUnpackedFileName const wint8_t *, DoLogData bool, DoRealignFile bool, DoMoveOverlay bool, EntryCallBack void *)  {
	// TODO: implement make json packet
}

// EngineUnpackerSetBreakCondition (:1103)
func EngineUnpackerSetBreakCondition(SearchStart void *, SearchSize int, SearchPattern void *, PatternSize int, PatternDelta int, BreakType int, SingleBreak bool, Parameter1 int, Parameter2 int)  {
	// TODO: implement make json packet
}

// EngineUnpackerSetEntryPointAddress (:1104)
func EngineUnpackerSetEntryPointAddress(UnpackedEntryPointAddress int)  {
	// TODO: implement make json packet
}

// EngineUnpackerFinalizeUnpacking (:1105)
func EngineUnpackerFinalizeUnpacking()  {
	// TODO: implement make json packet
}

// SetEngineVariable (:1107)
func SetEngineVariable(VariableId int, VariableSet bool)  {
	// TODO: implement make json packet
}

// EngineCreateMissingDependencies (:1108)
func EngineCreateMissingDependencies(szFileName string, szOutputFolder string, LogCreatedFiles bool)  {
	// TODO: implement make json packet
}

// EngineCreateMissingDependenciesW (:1109)
func EngineCreateMissingDependenciesW(szFileName const wint8_t *, szOutputFolder const wint8_t *, LogCreatedFiles bool)  {
	// TODO: implement make json packet
}

// EngineFakeMissingDependencies (:1110)
func EngineFakeMissingDependencies(hProcess int)  {
	// TODO: implement make json packet
}

// EngineDeleteCreatedDependencies (:1111)
func EngineDeleteCreatedDependencies()  {
	// TODO: implement make json packet
}

// EngineCreateUnpackerWindow (:1112)
func EngineCreateUnpackerWindow(WindowUnpackerTitle string, WindowUnpackerLongTitle string, WindowUnpackerName string, WindowUnpackerAuthor string, StartUnpackingCallBack void *)  {
	// TODO: implement make json packet
}

// EngineAddUnpackerWindowLogMessage (:1113)
func EngineAddUnpackerWindowLogMessage(szLogMessage string)  {
	// TODO: implement make json packet
}

// EngineCheckStructAlignment (:1114)
func EngineCheckStructAlignment(StructureType int, StructureSize int)  {
	// TODO: implement make json packet
}

// ExtensionManagerIsPluginLoaded (:1116)
func ExtensionManagerIsPluginLoaded(szPluginName string)  {
	// TODO: implement make json packet
}

// ExtensionManagerIsPluginEnabled (:1117)
func ExtensionManagerIsPluginEnabled(szPluginName string)  {
	// TODO: implement make json packet
}

// ExtensionManagerDisableAllPlugins (:1118)
func ExtensionManagerDisableAllPlugins()  {
	// TODO: implement make json packet
}

// ExtensionManagerDisablePlugin (:1119)
func ExtensionManagerDisablePlugin(szPluginName string)  {
	// TODO: implement make json packet
}

// ExtensionManagerEnableAllPlugins (:1120)
func ExtensionManagerEnableAllPlugins()  {
	// TODO: implement make json packet
}

// ExtensionManagerEnablePlugin (:1121)
func ExtensionManagerEnablePlugin(szPluginName string)  {
	// TODO: implement make json packet
}

// ExtensionManagerUnloadAllPlugins (:1122)
func ExtensionManagerUnloadAllPlugins()  {
	// TODO: implement make json packet
}

// ExtensionManagerUnloadPlugin (:1123)
func ExtensionManagerUnloadPlugin(szPluginName string)  {
	// TODO: implement make json packet
}

// ExtensionManagerGetPluginInfo (:1124)
func ExtensionManagerGetPluginInfo(szPluginName string)  {
	// TODO: implement make json packet
}

// XEDParseAssemble (:47)
func XEDParseAssemble(XEDParse XEDPARSE *)  {
	// TODO: implement make json packet
}

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

// json_string (:28)
func json_string(str const int &)  {
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

// SymFindDebugInfoFile (:152)
func SymFindDebugInfoFile(hProcess int, FileName int, DebugFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// SymFindDebugInfoFileW (:169)
func SymFindDebugInfoFileW(hProcess int, FileName int, DebugFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// FindDebugInfoFile (:179)
func FindDebugInfoFile(FileName int, SymbolPath int, DebugFilePath int)  {
	// TODO: implement make json packet
}

// FindDebugInfoFileEx (:187)
func FindDebugInfoFileEx(FileName int, SymbolPath int, DebugFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// FindDebugInfoFileExW (:197)
func FindDebugInfoFileExW(FileName int, SymbolPath int, DebugFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// DECLSPEC_IMPORT (:0)
func DECLSPEC_IMPORT( int *)  {
	// TODO: implement make json packet
}

// SymFindFileInPath (:213)
func SymFindFileInPath(hprocess int, SearchPath int, FileName int, id int, two int, three int, flags int, FoundFile int, callback int, context int)  {
	// TODO: implement make json packet
}

// SymFindFileInPathW (:234)
func SymFindFileInPathW(hprocess int, SearchPath int, FileName int, id int, two int, three int, flags int, FoundFile int, callback int, context int)  {
	// TODO: implement make json packet
}

// SymFindExecutableImage (:256)
func SymFindExecutableImage(hProcess int, FileName int, ImageFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// SymFindExecutableImageW (:273)
func SymFindExecutableImageW(hProcess int, FileName int, ImageFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// FindExecutableImage (:283)
func FindExecutableImage(FileName int, SymbolPath int, ImageFilePath int)  {
	// TODO: implement make json packet
}

// FindExecutableImageEx (:291)
func FindExecutableImageEx(FileName int, SymbolPath int, ImageFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// FindExecutableImageExW (:301)
func FindExecutableImageExW(FileName int, SymbolPath int, ImageFilePath int, Callback int, CallerData int)  {
	// TODO: implement make json packet
}

// ImageNtHeader (:311)
func ImageNtHeader(Base int)  {
	// TODO: implement make json packet
}

// ImageDirectoryEntryToDataEx (:317)
func ImageDirectoryEntryToDataEx(Base int, MappedAsImage int, DirectoryEntry int, Size int, FoundHeader int *)  {
	// TODO: implement make json packet
}

// ImageDirectoryEntryToData (:327)
func ImageDirectoryEntryToData(Base int, MappedAsImage int, DirectoryEntry int, Size int)  {
	// TODO: implement make json packet
}

// ImageRvaToSection (:336)
func ImageRvaToSection(NtHeaders int, Base int, Rva int)  {
	// TODO: implement make json packet
}

// ImageRvaToVa (:344)
func ImageRvaToVa(NtHeaders int, Base int, Rva int, PIMAGE_SECTION_HEADER int)  {
	// TODO: implement make json packet
}

// SearchTreeForFile (:421)
func SearchTreeForFile(RootPath int, InputPathName int, OutputPathBuffer int)  {
	// TODO: implement make json packet
}

// SearchTreeForFileW (:429)
func SearchTreeForFileW(RootPath int, InputPathName int, OutputPathBuffer int)  {
	// TODO: implement make json packet
}

// EnumDirTree (:443)
func EnumDirTree(hProcess int, RootPath int, InputPathName int, OutputPathBuffer int, cb int, data int)  {
	// TODO: implement make json packet
}

// EnumDirTreeW (:460)
func EnumDirTreeW(hProcess int, RootPath int, InputPathName int, OutputPathBuffer int, cb int, data int)  {
	// TODO: implement make json packet
}

// MakeSureDirectoryPathExists (:471)
func MakeSureDirectoryPathExists(DirPath int)  {
	// TODO: implement make json packet
}

// UnDecorateSymbolName (:501)
func UnDecorateSymbolName(name int, outputString int, maxStringLength int, flags int)  {
	// TODO: implement make json packet
}

// UnDecorateSymbolNameW (:511)
func UnDecorateSymbolNameW(name int, outputString int, maxStringLength int, flags int)  {
	// TODO: implement make json packet
}

// Far (:774)
func Far( int *)  {
	// TODO: implement make json packet
}

// Virtual (:2429)
func Virtual( int *)  {
	// TODO: implement make json packet
}

// StackWalk64 (:835)
func StackWalk64(MachineType int, hProcess int, hThread int, StackFrame LPSTACKFRAME64, ContextRecord int, ReadMemoryRoutine PREAD_PROCESS_MEMORY_ROUTINE64, FunctionTableAccessRoutine PFUNCTION_TABLE_ACCESS_ROUTINE64, GetModuleBaseRoutine PGET_MODULE_BASE_ROUTINE64, TranslateAddress PTRANSLATE_ADDRESS_ROUTINE64)  {
	// TODO: implement make json packet
}

// ImagehlpApiVersionEx (:925)
func ImagehlpApiVersionEx(AppVersion LPAPI_VERSION)  {
	// TODO: implement make json packet
}

// GetTimestampForLoadedLibrary (:931)
func GetTimestampForLoadedLibrary(Module int)  {
	// TODO: implement make json packet
}

// PdbUnmatched (:1290)
func PdbUnmatched( int *)  {
	// TODO: implement make json packet
}

// DbgUnmatched (:1291)
func DbgUnmatched( int *)  {
	// TODO: implement make json packet
}

// LineNumbers (:1292)
func LineNumbers( int *)  {
	// TODO: implement make json packet
}

// GlobalSymbols (:1293)
func GlobalSymbols( int *)  {
	// TODO: implement make json packet
}

// TypeInfo (:1294)
func TypeInfo( int *)  {
	// TODO: implement make json packet
}

// SourceIndexed (:1296)
func SourceIndexed( int *)  {
	// TODO: implement make json packet
}

// Publics (:1297)
func Publics( int *)  {
	// TODO: implement make json packet
}

// Key (:1856)
func Key( CALLBACK *)  {
	// TODO: implement make json packet
}

// buf (:1417)
func buf( CALLBACK *)  {
	// TODO: implement make json packet
}

// object (:1448)
func object( CALLBACK *)  {
	// TODO: implement make json packet
}

// SymSetParentWindow (:1519)
func SymSetParentWindow(hwnd int)  {
	// TODO: implement make json packet
}

// SymSetHomeDirectory (:1525)
func SymSetHomeDirectory(hProcess int, dir int)  {
	// TODO: implement make json packet
}

// SymSetHomeDirectoryW (:1532)
func SymSetHomeDirectoryW(hProcess int, dir int)  {
	// TODO: implement make json packet
}

// SymGetHomeDirectory (:1539)
func SymGetHomeDirectory(type int, dir int, size uint)  {
	// TODO: implement make json packet
}

// SymGetHomeDirectoryW (:1547)
func SymGetHomeDirectoryW(type int, dir int, size uint)  {
	// TODO: implement make json packet
}

// SymGetOmaps (:1569)
func SymGetOmaps(hProcess int, BaseOfDll int, OmapTo POMAP *, cOmapTo int, OmapFrom POMAP *, cOmapFrom int)  {
	// TODO: implement make json packet
}

// SymSetOptions (:1613)
func SymSetOptions(SymOptions int)  {
	// TODO: implement make json packet
}

// SymCleanup (:1625)
func SymCleanup(hProcess int)  {
	// TODO: implement make json packet
}

// SymMatchString (:1631)
func SymMatchString(string int, expression int, fCase bool *)  {
	// TODO: implement make json packet
}

// SymMatchStringA (:1639)
func SymMatchStringA(string int, expression int, fCase bool *)  {
	// TODO: implement make json packet
}

// SymMatchStringW (:1647)
func SymMatchStringW(string int, expression int, fCase bool *)  {
	// TODO: implement make json packet
}

// SymEnumSourceFiles (:1664)
func SymEnumSourceFiles(hProcess int, ModBase int, Mask int, cbSrcFiles int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumSourceFilesW (:1680)
func SymEnumSourceFilesW(hProcess int, ModBase int, Mask int, cbSrcFiles int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumerateModules64 (:1690)
func SymEnumerateModules64(hProcess int, EnumModulesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumerateModulesW64 (:1698)
func SymEnumerateModulesW64(hProcess int, EnumModulesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// EnumerateLoadedModulesEx (:1718)
func EnumerateLoadedModulesEx(hProcess int, EnumLoadedModulesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// EnumerateLoadedModulesExW (:1726)
func EnumerateLoadedModulesExW(hProcess int, EnumLoadedModulesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// EnumerateLoadedModules64 (:1734)
func EnumerateLoadedModules64(hProcess int, EnumLoadedModulesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// EnumerateLoadedModulesW64 (:1742)
func EnumerateLoadedModulesW64(hProcess int, EnumLoadedModulesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymFunctionTableAccess64 (:1762)
func SymFunctionTableAccess64(hProcess int, AddrBase int)  {
	// TODO: implement make json packet
}

// SymGetUnwindInfo (:1780)
func SymGetUnwindInfo(hProcess int, Address int, Buffer uintptr *, Size int)  {
	// TODO: implement make json packet
}

// SymGetModuleInfo64 (:1789)
func SymGetModuleInfo64(hProcess int, qwAddr int, ModuleInfo PIMAGEHLP_MODULE64)  {
	// TODO: implement make json packet
}

// SymGetModuleInfoW64 (:1797)
func SymGetModuleInfoW64(hProcess int, qwAddr int, ModuleInfo PIMAGEHLP_MODULEW64)  {
	// TODO: implement make json packet
}

// SymGetModuleBase64 (:1826)
func SymGetModuleBase64(hProcess int, qwAddr int)  {
	// TODO: implement make json packet
}

// SymEnumLines (:1872)
func SymEnumLines(hProcess int, Base int, Obj int, File int, EnumLinesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumLinesW (:1889)
func SymEnumLinesW(hProcess int, Base int, Obj int, File int, EnumLinesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymGetLineFromAddr64 (:1900)
func SymGetLineFromAddr64(hProcess int, qwAddr int, pdwDisplacement int, Line64 PIMAGEHLP_LINE64)  {
	// TODO: implement make json packet
}

// SymGetLineFromAddrW64 (:1909)
func SymGetLineFromAddrW64(hProcess int, dwAddr int, pdwDisplacement int, Line PIMAGEHLP_LINEW64)  {
	// TODO: implement make json packet
}

// SymEnumSourceLines (:1918)
func SymEnumSourceLines(hProcess int, Base int, Obj int, File int, Line int, Flags int, EnumLinesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumSourceLinesW (:1931)
func SymEnumSourceLinesW(hProcess int, Base int, Obj int, File int, Line int, Flags int, EnumLinesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymGetLineFromName64 (:1974)
func SymGetLineFromName64(hProcess int, ModuleName int, FileName int, dwLineNumber int, plDisplacement int, Line PIMAGEHLP_LINE64)  {
	// TODO: implement make json packet
}

// SymGetLineFromNameW64 (:1985)
func SymGetLineFromNameW64(hProcess int, ModuleName int, FileName int, dwLineNumber int, plDisplacement int, Line PIMAGEHLP_LINEW64)  {
	// TODO: implement make json packet
}

// SymGetLineNext64 (:2011)
func SymGetLineNext64(hProcess int, Line PIMAGEHLP_LINE64)  {
	// TODO: implement make json packet
}

// SymGetLineNextW64 (:2018)
func SymGetLineNextW64(hProcess int, Line PIMAGEHLP_LINEW64)  {
	// TODO: implement make json packet
}

// SymGetLinePrev64 (:2043)
func SymGetLinePrev64(hProcess int, Line PIMAGEHLP_LINE64)  {
	// TODO: implement make json packet
}

// SymGetLinePrevW64 (:2050)
func SymGetLinePrevW64(hProcess int, Line PIMAGEHLP_LINEW64)  {
	// TODO: implement make json packet
}

// SymGetFileLineOffsets64 (:2075)
func SymGetFileLineOffsets64(hProcess int, ModuleName int, FileName int, Buffer int, BufferLines int)  {
	// TODO: implement make json packet
}

// SymMatchFileName (:2085)
func SymMatchFileName(FileName int, Match int, FileNameStop int *, MatchStop int *)  {
	// TODO: implement make json packet
}

// SymMatchFileNameW (:2094)
func SymMatchFileNameW(FileName int, Match int, FileNameStop int *, MatchStop int *)  {
	// TODO: implement make json packet
}

// SymGetSourceFile (:2103)
func SymGetSourceFile(hProcess int, Base int, Params int, FileSpec int, FilePath int, Size int)  {
	// TODO: implement make json packet
}

// SymGetSourceFileW (:2114)
func SymGetSourceFileW(hProcess int, Base int, Params int, FileSpec int, FilePath int, Size int)  {
	// TODO: implement make json packet
}

// SymGetSourceFileToken (:2125)
func SymGetSourceFileToken(hProcess int, Base int, FileSpec int, Token uintptr *, Size int *)  {
	// TODO: implement make json packet
}

// SymGetSourceFileTokenW (:2135)
func SymGetSourceFileTokenW(hProcess int, Base int, FileSpec int, Token uintptr *, Size int *)  {
	// TODO: implement make json packet
}

// SymGetSourceFileFromToken (:2145)
func SymGetSourceFileFromToken(hProcess int, Token uintptr *, Params int, FilePath int, Size int)  {
	// TODO: implement make json packet
}

// SymGetSourceFileFromTokenW (:2155)
func SymGetSourceFileFromTokenW(hProcess int, Token uintptr *, Params int, FilePath int, Size int)  {
	// TODO: implement make json packet
}

// SymGetSourceVarFromToken (:2165)
func SymGetSourceVarFromToken(hProcess int, Token uintptr *, Params int, VarName int, Value int, Size int)  {
	// TODO: implement make json packet
}

// SymGetSourceVarFromTokenW (:2176)
func SymGetSourceVarFromTokenW(hProcess int, Token uintptr *, Params int, VarName int, Value int, Size int)  {
	// TODO: implement make json packet
}

// SymEnumSourceFileTokens (:2189)
func SymEnumSourceFileTokens(hProcess int, Base int, Callback int)  {
	// TODO: implement make json packet
}

// SymInitialize (:2197)
func SymInitialize(hProcess int, UserSearchPath int, fInvadeProcess bool *)  {
	// TODO: implement make json packet
}

// SymInitializeW (:2205)
func SymInitializeW(hProcess int, UserSearchPath int, fInvadeProcess bool *)  {
	// TODO: implement make json packet
}

// SymGetSearchPath (:2213)
func SymGetSearchPath(hProcess int, SearchPath int, SearchPathLength int)  {
	// TODO: implement make json packet
}

// SymGetSearchPathW (:2221)
func SymGetSearchPathW(hProcess int, SearchPath int, SearchPathLength int)  {
	// TODO: implement make json packet
}

// SymSetSearchPath (:2229)
func SymSetSearchPath(hProcess int, SearchPath int)  {
	// TODO: implement make json packet
}

// SymSetSearchPathW (:2236)
func SymSetSearchPathW(hProcess int, SearchPath int)  {
	// TODO: implement make json packet
}

// SymLoadModuleEx (:2247)
func SymLoadModuleEx(hProcess int, hFile int, ImageName int, ModuleName int, BaseOfDll int, DllSize int, Data PMODLOAD_DATA, Flags int)  {
	// TODO: implement make json packet
}

// SymLoadModuleExW (:2260)
func SymLoadModuleExW(hProcess int, hFile int, ImageName int, ModuleName int, BaseOfDll int, DllSize int, Data PMODLOAD_DATA, Flags int)  {
	// TODO: implement make json packet
}

// SymUnloadModule64 (:2273)
func SymUnloadModule64(hProcess int, BaseOfDll int)  {
	// TODO: implement make json packet
}

// SymUnDName64 (:2291)
func SymUnDName64(sym PIMAGEHLP_SYMBOL64, UnDecName int, UnDecNameLength int)  {
	// TODO: implement make json packet
}

// SymRegisterCallback64 (:2311)
func SymRegisterCallback64(hProcess int, CallbackFunction int, UserContext int)  {
	// TODO: implement make json packet
}

// SymRegisterCallbackW64 (:2319)
func SymRegisterCallbackW64(hProcess int, CallbackFunction int, UserContext int)  {
	// TODO: implement make json packet
}

// SymRegisterFunctionEntryCallback64 (:2327)
func SymRegisterFunctionEntryCallback64(hProcess int, CallbackFunction int, UserContext int)  {
	// TODO: implement make json packet
}

// SymSetContext (:2438)
func SymSetContext(hProcess int, StackFrame PIMAGEHLP_STACK_FRAME, Context int)  {
	// TODO: implement make json packet
}

// SymSetScopeFromAddr (:2446)
func SymSetScopeFromAddr(hProcess int, Address int)  {
	// TODO: implement make json packet
}

// SymSetScopeFromIndex (:2453)
func SymSetScopeFromIndex(hProcess int, BaseOfDll int, Index int)  {
	// TODO: implement make json packet
}

// SymEnumProcesses (:2467)
func SymEnumProcesses(EnumProcessesCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymFromAddr (:2474)
func SymFromAddr(hProcess int, Address int, Displacement int, Symbol PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymFromAddrW (:2483)
func SymFromAddrW(hProcess int, Address int, Displacement int, Symbol PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymFromToken (:2492)
func SymFromToken(hProcess int, Base int, Token int, Symbol PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymFromTokenW (:2501)
func SymFromTokenW(hProcess int, Base int, Token int, Symbol PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymNext (:2510)
func SymNext(hProcess int, si PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymNextW (:2517)
func SymNextW(hProcess int, siw PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymPrev (:2524)
func SymPrev(hProcess int, si PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymPrevW (:2531)
func SymPrevW(hProcess int, siw PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymFromName (:2544)
func SymFromName(hProcess int, Name int, Symbol PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymFromNameW (:2552)
func SymFromNameW(hProcess int, Name int, Symbol PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymEnumSymbols (:2567)
func SymEnumSymbols(hProcess int, BaseOfDll int, Mask int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumSymbolsW (:2584)
func SymEnumSymbolsW(hProcess int, BaseOfDll int, Mask int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumSymbolsForAddr (:2594)
func SymEnumSymbolsForAddr(hProcess int, Address int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumSymbolsForAddrW (:2603)
func SymEnumSymbolsForAddrW(hProcess int, Address int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymSearch (:2617)
func SymSearch(hProcess int, BaseOfDll int, Index int, SymTag int, Mask int, Address int, EnumSymbolsCallback int, UserContext uintptr *, Options int)  {
	// TODO: implement make json packet
}

// SymSearchW (:2631)
func SymSearchW(hProcess int, BaseOfDll int, Index int, SymTag int, Mask int, Address int, EnumSymbolsCallback int, UserContext uintptr *, Options int)  {
	// TODO: implement make json packet
}

// SymGetScope (:2645)
func SymGetScope(hProcess int, BaseOfDll int, Index int, Symbol PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymGetScopeW (:2654)
func SymGetScopeW(hProcess int, BaseOfDll int, Index int, Symbol PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymFromIndex (:2663)
func SymFromIndex(hProcess int, BaseOfDll int, Index int, Symbol PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymFromIndexW (:2672)
func SymFromIndexW(hProcess int, BaseOfDll int, Index int, Symbol PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymGetTypeInfo (:2726)
func SymGetTypeInfo(hProcess int, ModBase int, TypeId int, GetType IMAGEHLP_SYMBOL_TYPE_INFO, pInfo uintptr *)  {
	// TODO: implement make json packet
}

// SymGetTypeInfoEx (:2761)
func SymGetTypeInfoEx(hProcess int, ModBase int, Params PIMAGEHLP_GET_TYPE_INFO_PARAMS)  {
	// TODO: implement make json packet
}

// SymEnumTypes (:2769)
func SymEnumTypes(hProcess int, BaseOfDll int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumTypesW (:2778)
func SymEnumTypesW(hProcess int, BaseOfDll int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumTypesByName (:2787)
func SymEnumTypesByName(hProcess int, BaseOfDll int, mask int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumTypesByNameW (:2797)
func SymEnumTypesByNameW(hProcess int, BaseOfDll int, mask int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymGetTypeFromName (:2807)
func SymGetTypeFromName(hProcess int, BaseOfDll int, Name int, Symbol PSYMBOL_INFO)  {
	// TODO: implement make json packet
}

// SymGetTypeFromNameW (:2816)
func SymGetTypeFromNameW(hProcess int, BaseOfDll int, Name int, Symbol PSYMBOL_INFOW)  {
	// TODO: implement make json packet
}

// SymAddSymbol (:2825)
func SymAddSymbol(hProcess int, BaseOfDll int, Name int, Address int, Size int, Flags int)  {
	// TODO: implement make json packet
}

// SymAddSymbolW (:2836)
func SymAddSymbolW(hProcess int, BaseOfDll int, Name int, Address int, Size int, Flags int)  {
	// TODO: implement make json packet
}

// SymDeleteSymbol (:2847)
func SymDeleteSymbol(hProcess int, BaseOfDll int, Name int, Address int, Flags int)  {
	// TODO: implement make json packet
}

// SymDeleteSymbolW (:2857)
func SymDeleteSymbolW(hProcess int, BaseOfDll int, Name int, Address int, Flags int)  {
	// TODO: implement make json packet
}

// SymRefreshModuleList (:2867)
func SymRefreshModuleList(hProcess int)  {
	// TODO: implement make json packet
}

// SymAddSourceStream (:2873)
func SymAddSourceStream(hProcess int, Base int, StreamFile int, Buffer int, Size uint)  {
	// TODO: implement make json packet
}

// SymAddSourceStreamA (:2885)
func SymAddSourceStreamA(hProcess int, Base int, StreamFile int, Buffer int, Size uint)  {
	// TODO: implement make json packet
}

// SymAddSourceStreamW (:2897)
func SymAddSourceStreamW(hProcess int, Base int, FileSpec int, Buffer int, Size uint)  {
	// TODO: implement make json packet
}

// SymSrvIsStoreW (:2907)
func SymSrvIsStoreW(hProcess int, path int)  {
	// TODO: implement make json packet
}

// SymSrvIsStore (:2914)
func SymSrvIsStore(hProcess int, path int)  {
	// TODO: implement make json packet
}

// SymSrvDeltaName (:2921)
func SymSrvDeltaName(hProcess int, SymPath int, Type int, File1 int, File2 int)  {
	// TODO: implement make json packet
}

// SymSrvDeltaNameW (:2931)
func SymSrvDeltaNameW(hProcess int, SymPath int, Type int, File1 int, File2 int)  {
	// TODO: implement make json packet
}

// SymSrvGetSupplement (:2941)
func SymSrvGetSupplement(hProcess int, SymPath int, Node int, File int)  {
	// TODO: implement make json packet
}

// SymSrvGetSupplementW (:2950)
func SymSrvGetSupplementW(hProcess int, SymPath int, Node int, File int)  {
	// TODO: implement make json packet
}

// SymSrvGetFileIndexes (:2959)
func SymSrvGetFileIndexes(File int, Id int *, Val1 int, Val2 int, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvGetFileIndexesW (:2969)
func SymSrvGetFileIndexesW(File int, Id int *, Val1 int, Val2 int, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvGetFileIndexStringW (:2979)
func SymSrvGetFileIndexStringW(hProcess int, SrvPath int, File int, Index int, Size uint, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvGetFileIndexString (:2990)
func SymSrvGetFileIndexString(hProcess int, SrvPath int, File int, Index int, Size uint, Flags int)  {
	// TODO: implement make json packet
}

// stripped (:3017)
func stripped( int *)  {
	// TODO: implement make json packet
}

// SymSrvGetFileIndexInfo (:3029)
func SymSrvGetFileIndexInfo(File int, Info PSYMSRV_INDEX_INFO, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvGetFileIndexInfoW (:3037)
func SymSrvGetFileIndexInfoW(File int, Info PSYMSRV_INDEX_INFOW, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvStoreSupplement (:3045)
func SymSrvStoreSupplement(hProcess int, SrvPath int, Node int, File int, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvStoreSupplementW (:3055)
func SymSrvStoreSupplementW(hProcess int, SymPath int, Node int, File int, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvStoreFile (:3065)
func SymSrvStoreFile(hProcess int, SrvPath int, File int, Flags int)  {
	// TODO: implement make json packet
}

// SymSrvStoreFileW (:3074)
func SymSrvStoreFileW(hProcess int, SrvPath int, File int, Flags int)  {
	// TODO: implement make json packet
}

// SymGetSymbolFile (:3094)
func SymGetSymbolFile(hProcess int, SymPath int, ImageFile int, Type int, SymbolFile int, cSymbolFile uint, DbgFile int, cDbgFile uint)  {
	// TODO: implement make json packet
}

// SymGetSymbolFileW (:3107)
func SymGetSymbolFileW(hProcess int, SymPath int, ImageFile int, Type int, SymbolFile int, cSymbolFile uint, DbgFile int, cDbgFile uint)  {
	// TODO: implement make json packet
}

// WINAPI (:4660)
func WINAPI( int *)  {
	// TODO: implement make json packet
}

// SymGetSymFromAddr64 (:3152)
func SymGetSymFromAddr64(hProcess int, qwAddr int, pdwDisplacement int, Symbol PIMAGEHLP_SYMBOL64)  {
	// TODO: implement make json packet
}

// SymGetSymFromName64 (:3181)
func SymGetSymFromName64(hProcess int, Name int, Symbol PIMAGEHLP_SYMBOL64)  {
	// TODO: implement make json packet
}

// FindFileInPath (:3388)
func FindFileInPath(hprocess int, SearchPath int, FileName int, id uintptr *, two uint32 *, three uint32 *, flags uint32 *, FilePath int)  {
	// TODO: implement make json packet
}

// FindFileInSearchPath (:3405)
func FindFileInSearchPath(hprocess int, SearchPath int, FileName int, one uint32 *, two uint32 *, three uint32 *, FilePath int)  {
	// TODO: implement make json packet
}

// SymEnumSym (:3418)
func SymEnumSym(hProcess int, BaseOfDll int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumerateSymbols64 (:3428)
func SymEnumerateSymbols64(hProcess int, BaseOfDll int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymEnumerateSymbolsW64 (:3438)
func SymEnumerateSymbolsW64(hProcess int, BaseOfDll int, EnumSymbolsCallback int, UserContext uintptr *)  {
	// TODO: implement make json packet
}

// SymLoadModule64 (:3475)
func SymLoadModule64(hProcess int, hFile int, ImageName int, ModuleName int, BaseOfDll int, SizeOfDll uint32 *)  {
	// TODO: implement make json packet
}

// SymGetSymNext64 (:3501)
func SymGetSymNext64(hProcess int, Symbol PIMAGEHLP_SYMBOL64)  {
	// TODO: implement make json packet
}

// SymGetSymNextW64 (:3508)
func SymGetSymNextW64(hProcess int, Symbol PIMAGEHLP_SYMBOLW64)  {
	// TODO: implement make json packet
}

// SymGetSymPrev64 (:3534)
func SymGetSymPrev64(hProcess int, Symbol PIMAGEHLP_SYMBOL64)  {
	// TODO: implement make json packet
}

// SymGetSymPrevW64 (:3541)
func SymGetSymPrevW64(hProcess int, Symbol PIMAGEHLP_SYMBOLW64)  {
	// TODO: implement make json packet
}

// Rva (:3612)
func Rva( WINAPI *)  {
	// TODO: implement make json packet
}

// StreamDirectoryRva (:3645)
func StreamDirectoryRva( WINAPI *)  {
	// TODO: implement make json packet
}

// CSDVersionRva (:3831)
func CSDVersionRva( WINAPI *)  {
	// TODO: implement make json packet
}

// ModuleNameRva (:4126)
func ModuleNameRva( WINAPI *)  {
	// TODO: implement make json packet
}

// ThreadId (:4003)
func ThreadId( WINAPI *)  {
	// TODO: implement make json packet
}

// ClientPointers (:4006)
func ClientPointers( int *)  {
	// TODO: implement make json packet
}

// NextInfoRva (:4032)
func NextInfoRva( WINAPI *)  {
	// TODO: implement make json packet
}

// TypeNameRva (:4052)
func TypeNameRva( WINAPI *)  {
	// TODO: implement make json packet
}

// ObjectNameRva (:4053)
func ObjectNameRva( WINAPI *)  {
	// TODO: implement make json packet
}

// ObjectInfoRva (:4058)
func ObjectInfoRva( WINAPI *)  {
	// TODO: implement make json packet
}

// Buffer (:4418)
func Buffer( CALLBACK *)  {
	// TODO: implement make json packet
}

// CvRecord (:4389)
func CvRecord( CALLBACK *)  {
	// TODO: implement make json packet
}

// MiscRecord (:4391)
func MiscRecord( CALLBACK *)  {
	// TODO: implement make json packet
}

// CheckCancel (:4465)
func CheckCancel( int *)  {
	// TODO: implement make json packet
}

// Cancel (:4466)
func Cancel( int *)  {
	// TODO: implement make json packet
}

// Continue (:4472)
func Continue( int *)  {
	// TODO: implement make json packet
}

// CallbackParam (:4615)
func CallbackParam( CALLBACK *)  {
	// TODO: implement make json packet
}

// json_hex (:8)
func json_hex(value uint64)  {
	// TODO: implement make json packet
}

// json_hex_value (:16)
func json_hex_value(hex const json_t *)  {
	// TODO: implement make json packet
}

// LZ4_compress (:63)
func LZ4_compress(source string, dest int8 *, inputSize int)  {
	// TODO: implement make json packet
}

// LZ4_decompress_safe (:64)
func LZ4_decompress_safe(source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compressBound (:101)
func LZ4_compressBound(isize int)  {
	// TODO: implement make json packet
}

// LZ4_compress_limitedOutput (:115)
func LZ4_compress_limitedOutput(source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_decompress_fast (:128)
func LZ4_decompress_fast(source string, dest int8 *, originalSize int)  {
	// TODO: implement make json packet
}

// LZ4_decompress_safe_partial (:143)
func LZ4_decompress_safe_partial(source string, dest int8 *, inputSize int, targetOutputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_sizeofState (:157)
func LZ4_sizeofState()  {
	// TODO: implement make json packet
}

// LZ4_compress_withState (:158)
func LZ4_compress_withState(state void *, source string, dest int8 *, inputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compress_limitedOutput_withState (:159)
func LZ4_compress_limitedOutput_withState(state void *, source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_create (:165)
func LZ4_create(inputBuffer string)  {
	// TODO: implement make json packet
}

// LZ4_compress_continue (:166)
func LZ4_compress_continue(LZ4_Data void *, source string, dest int8 *, inputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compress_limitedOutput_continue (:167)
func LZ4_compress_limitedOutput_continue(LZ4_Data void *, source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_slideInputBuffer (:168)
func LZ4_slideInputBuffer(LZ4_Data void *)  {
	// TODO: implement make json packet
}

// LZ4_free (:169)
func LZ4_free(LZ4_Data void *)  {
	// TODO: implement make json packet
}

// LZ4_sizeofStreamState (:201)
func LZ4_sizeofStreamState()  {
	// TODO: implement make json packet
}

// LZ4_resetStreamState (:202)
func LZ4_resetStreamState(state void *, inputBuffer string)  {
	// TODO: implement make json packet
}

// LZ4_decompress_safe_withPrefix64k (:226)
func LZ4_decompress_safe_withPrefix64k(source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_decompress_fast_withPrefix64k (:227)
func LZ4_decompress_fast_withPrefix64k(source string, dest int8 *, outputSize int)  {
	// TODO: implement make json packet
}

// LZ4_uncompress (:244)
func LZ4_uncompress(source string, dest int8 *, outputSize int)  {
	// TODO: implement make json packet
}

// LZ4_uncompress_unknownOutputSize (:245)
func LZ4_uncompress_unknownOutputSize(source string, dest int8 *, isize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compress_file (:19)
func LZ4_compress_file(input_filename string, output_filename string)  {
	// TODO: implement make json packet
}

// LZ4_compress_fileW (:20)
func LZ4_compress_fileW(input_filename const wint8_t *, output_filename const wint8_t *)  {
	// TODO: implement make json packet
}

// LZ4_decompress_file (:21)
func LZ4_decompress_file(input_filename string, output_filename string)  {
	// TODO: implement make json packet
}

// LZ4_decompress_fileW (:22)
func LZ4_decompress_fileW(input_filename const wint8_t *, output_filename const wint8_t *)  {
	// TODO: implement make json packet
}

// LZ4_compressHC (:43)
func LZ4_compressHC(source string, dest int8 *, inputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC_limitedOutput (:53)
func LZ4_compressHC_limitedOutput(source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC2 (:67)
func LZ4_compressHC2(source string, dest int8 *, inputSize int, compressionLevel int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC2_limitedOutput (:68)
func LZ4_compressHC2_limitedOutput(source string, dest int8 *, inputSize int, maxOutputSize int, compressionLevel int)  {
	// TODO: implement make json packet
}

// LZ4_sizeofStateHC (:85)
func LZ4_sizeofStateHC()  {
	// TODO: implement make json packet
}

// LZ4_compressHC_withStateHC (:86)
func LZ4_compressHC_withStateHC(state void *, source string, dest int8 *, inputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC_limitedOutput_withStateHC (:87)
func LZ4_compressHC_limitedOutput_withStateHC(state void *, source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC2_withStateHC (:89)
func LZ4_compressHC2_withStateHC(state void *, source string, dest int8 *, inputSize int, compressionLevel int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC2_limitedOutput_withStateHC (:90)
func LZ4_compressHC2_limitedOutput_withStateHC(state void *, source string, dest int8 *, inputSize int, maxOutputSize int, compressionLevel int)  {
	// TODO: implement make json packet
}

// LZ4_createHC (:108)
func LZ4_createHC(inputBuffer string)  {
	// TODO: implement make json packet
}

// LZ4_compressHC_continue (:109)
func LZ4_compressHC_continue(LZ4HC_Data void *, source string, dest int8 *, inputSize int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC_limitedOutput_continue (:110)
func LZ4_compressHC_limitedOutput_continue(LZ4HC_Data void *, source string, dest int8 *, inputSize int, maxOutputSize int)  {
	// TODO: implement make json packet
}

// LZ4_slideInputBufferHC (:111)
func LZ4_slideInputBufferHC(LZ4HC_Data void *)  {
	// TODO: implement make json packet
}

// LZ4_freeHC (:112)
func LZ4_freeHC(LZ4HC_Data void *)  {
	// TODO: implement make json packet
}

// LZ4_compressHC2_continue (:114)
func LZ4_compressHC2_continue(LZ4HC_Data void *, source string, dest int8 *, inputSize int, compressionLevel int)  {
	// TODO: implement make json packet
}

// LZ4_compressHC2_limitedOutput_continue (:115)
func LZ4_compressHC2_limitedOutput_continue(LZ4HC_Data void *, source string, dest int8 *, inputSize int, maxOutputSize int, compressionLevel int)  {
	// TODO: implement make json packet
}

// LZ4_sizeofStreamStateHC (:146)
func LZ4_sizeofStreamStateHC()  {
	// TODO: implement make json packet
}

// LZ4_resetStreamStateHC (:147)
func LZ4_resetStreamStateHC(state void *, inputBuffer string)  {
	// TODO: implement make json packet
}

