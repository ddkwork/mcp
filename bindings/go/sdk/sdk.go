package sdk

import (
	"encoding/hex"
	"github.com/ddkwork/golibrary/std/stream"
	"strconv"
)

type Debugger struct{}

func (Debugger) VmxSupportDetection() bool {
	return request[bool]("VmxSupportDetection", nil)
}
func (Debugger) CpuReadVendorString() string {
	return request[string]("CpuReadVendorString", nil)
}
func (Debugger) LoadVmmModule() string {
	return request[string]("LoadVmmModule", nil)
}
func (Debugger) UnloadVmm() string {
	return request[string]("UnloadVmm", nil)
}
func (Debugger) InstallVmmDriver() string {
	return request[string]("InstallVmmDriver", nil)
}
func (Debugger) UninstallVmmDriver() string {
	return request[string]("UninstallVmmDriver", nil)
}
func (Debugger) StopVmmDriver() string {
	return request[string]("StopVmmDriver", nil)
}
func (Debugger) RunCommand(command string) string {
	return request[string]("RunCommand", map[string]string{"command": command})
}
func (Debugger) TestCommandParserShowTokens(command string) string {
	return request[string]("TestCommandParserShowTokens", map[string]string{"command": command})
}
func (Debugger) ShowSignature() string {
	return request[string]("ShowSignature", nil)
}
func (Debugger) ContinuePreviousCommand() bool {
	return request[bool]("ContinuePreviousCommand", nil)
}
func (Debugger) CheckMultilineCommand(current_command string, reset bool) bool {
	return request[bool]("CheckMultilineCommand", map[string]string{
		"current_command": current_command,
		"reset":           strconv.FormatBool(reset),
	})
}
func (Debugger) ConnectLocalDebugger() {
	request[void]("ConnectLocalDebugger", nil)
}
func (Debugger) ConnectRemoteDebugger(ip string, port string) string {
	return request[string]("ConnectRemoteDebugger", map[string]string{
		"ip":   ip,
		"port": port,
	})
}
func (Debugger) Continue() {
	request[void]("Continue", nil)
}
func (Debugger) Pause() {
	request[void]("Pause", nil)
}
func (Debugger) SetBreakPoint(address uint64, pid uint32, tid uint32, core_numer uint32) {
	request[void]("SetBreakPoint", map[string]string{
		"address":    strconv.FormatUint(address, 10),
		"pid":        strconv.FormatUint(uint64(pid), 10),
		"tid":        strconv.FormatUint(uint64(tid), 10),
		"core_numer": strconv.FormatUint(uint64(core_numer), 10),
	})
}
func (Debugger) SetCustomDriverPath(driver_file_path string, driver_name string) bool {
	return request[bool]("SetCustomDriverPath", map[string]string{
		"driver_file_path": driver_file_path,
		"driver_name":      driver_name,
	})
}
func (Debugger) UseDefaultDriverPath() {
	request[void]("UseDefaultDriverPath", nil)
}
func (Debugger) ReadMemory(target_address uint64, memory_type DEBUGGER_READ_MEMORY_TYPE, reading_Type DEBUGGER_READ_READING_TYPE, pid uint32, size uint32, get_address_mode bool, address_mode DEBUGGER_READ_MEMORY_ADDRESS_MODE, target_buffer_to_store []byte, return_length uint32) bool {
	return request[bool]("ReadMemory", map[string]string{
		"target_address":         strconv.FormatUint(target_address, 10),
		"memory_type":            strconv.FormatUint(uint64(memory_type), 10),
		"reading_Type":           strconv.FormatUint(uint64(reading_Type), 10),
		"pid":                    strconv.FormatUint(uint64(pid), 10),
		"size":                   strconv.FormatUint(uint64(size), 10),
		"get_address_mode":       strconv.FormatBool(get_address_mode),
		"address_mode":           strconv.FormatUint(uint64(address_mode), 10),
		"target_buffer_to_store": hex.EncodeToString(target_buffer_to_store),
		"return_length":          strconv.FormatUint(uint64(return_length), 10),
	})
}
func (Debugger) ShowMemoryOrDisassemble(style DEBUGGER_SHOW_MEMORY_STYLE, address uint64, memory_type DEBUGGER_READ_MEMORY_TYPE, reading_type DEBUGGER_READ_READING_TYPE, pid uint32, size uint32, dt_details DEBUGGER_DT_COMMAND_OPTIONS) {
	request[void]("ShowMemoryOrDisassemble", map[string]string{
		"style":        strconv.FormatUint(uint64(style), 10),
		"address":      strconv.FormatUint(address, 10),
		"memory_type":  strconv.FormatUint(uint64(memory_type), 10),
		"reading_type": strconv.FormatUint(uint64(reading_type), 10),
		"pid":          strconv.FormatUint(uint64(pid), 10),
		"size":         strconv.FormatUint(uint64(size), 10),
		"dt_details":   string(stream.MarshalJSON(dt_details)),
	})
}
func (Debugger) ReadAllRegisters(guest_registers GUEST_REGS, extra_registers GUEST_EXTRA_REGISTERS) bool {
	return request[bool]("ReadAllRegisters", map[string]string{
		"guest_registers": string(stream.MarshalJSON(guest_registers)),
		"extra_registers": string(stream.MarshalJSON(extra_registers)),
	})
}
func (Debugger) ReadTargetRegister(register_id REGS_ENUM, target_register uint64) bool {
	return request[bool]("ReadTargetRegister", map[string]string{
		"register_id":     strconv.FormatUint(uint64(register_id), 10),
		"target_register": strconv.FormatUint(uint64(target_register), 10),
	})
}
func (Debugger) WriteTargetRegister(register_id REGS_ENUM, value uint64) bool {
	return request[bool]("WriteTargetRegister", map[string]string{
		"register_id": strconv.FormatUint(uint64(register_id), 10),
		"value":       strconv.FormatUint(value, 10),
	})
}
func (Debugger) RegisterShowAll() bool {
	return request[bool]("RegisterShowAll", nil)
}
func (Debugger) RegisterShowTargetRegister(register_id REGS_ENUM) bool {
	return request[bool]("RegisterShowTargetRegister", map[string]string{"register_id": strconv.FormatUint(uint64(register_id), 10)})
}
func (Debugger) WriteMemory(destination_address uint64, memory_type DEBUGGER_EDIT_MEMORY_TYPE, process_id uint32, source_address uint64, number_of_bytes uint32) bool {
	return request[bool]("WriteMemory", map[string]string{
		"destination_address": strconv.FormatUint(uint64(destination_address), 10),
		"memory_type":         strconv.FormatUint(uint64(memory_type), 10),
		"process_id":          strconv.FormatUint(uint64(process_id), 10),
		"source_address":      strconv.FormatUint(uint64(source_address), 10),
		"number_of_bytes":     strconv.FormatUint(uint64(number_of_bytes), 10),
	})
}
func (Debugger) DebuggerGetKernelBase() uint64 {
	return request[uint64]("DebuggerGetKernelBase", nil)
}
func (Debugger) DebugRemoteDeviceUsingComPort(port_name string, baudrate uint32, pause_after_connection bool) bool {
	return request[bool]("DebugRemoteDeviceUsingComPort", map[string]string{
		"port_name":              port_name,
		"baudrate":               strconv.FormatUint(uint64(baudrate), 10),
		"pause_after_connection": strconv.FormatBool(pause_after_connection),
	})
}
func (Debugger) DebugRemoteDeviceUsingNamedPipe(named_pipe string, pause_after_connection bool) bool {
	return request[bool]("DebugRemoteDeviceUsingNamedPipe", map[string]string{
		"named_pipe":             named_pipe,
		"pause_after_connection": strconv.FormatBool(pause_after_connection),
	})
}
func (Debugger) DebugCloseRemoteDebugger() bool {
	return request[bool]("DebugCloseRemoteDebugger", nil)
}
func (Debugger) DebugCurrentDeviceUsingComPort(port_name string, baudrate uint32) bool {
	return request[bool]("DebugCurrentDeviceUsingComPort", map[string]string{
		"port_name": port_name,
		"baudrate":  strconv.FormatUint(uint64(baudrate), 10),
	})
}
func (Debugger) StartProcess(path string) bool {
	return request[bool]("StartProcess", map[string]string{"path": path})
}
func (Debugger) StartProcessWithArgs(path string, arguments string) bool {
	return request[bool]("StartProcessWithArgs", map[string]string{
		"path":      path,
		"arguments": arguments,
	})
}
func (Debugger) AssembleGetLength(assembly_code string, start_address uint64, length uint32) bool {
	return request[bool]("AssembleGetLength", map[string]string{
		"assembly_code": assembly_code,
		"start_address": strconv.FormatUint(start_address, 10),
		"length":        strconv.FormatUint(uint64(length), 10),
	})
}
func (Debugger) Assemble(assembly_code string, start_address uint64, buffer_to_store_assembled_data uint64, buffer_size uint32) bool {
	return request[bool]("Assemble", map[string]string{
		"assembly_code":                  assembly_code,
		"start_address":                  strconv.FormatUint(start_address, 10),
		"buffer_to_store_assembled_data": strconv.FormatUint(uint64(buffer_to_store_assembled_data), 10),
		"buffer_size":                    strconv.FormatUint(uint64(buffer_size), 10),
	})
}
func (Debugger) SetupPathForFileName(filename string, file_location string, buffer_len uint32, check_file_existence bool) bool {
	return request[bool]("SetupPathForFileName", map[string]string{
		"filename":             filename,
		"file_location":        file_location,
		"buffer_len":           strconv.FormatUint(uint64(buffer_len), 10),
		"check_file_existence": strconv.FormatBool(check_file_existence),
	})
}
func (Debugger) SteppingInstrumentationStepIn() bool {
	return request[bool]("SteppingInstrumentationStepIn", nil)
}
func (Debugger) SteppingRegularStepIn() bool {
	return request[bool]("SteppingRegularStepIn", nil)
}
func (Debugger) SteppingStepOver() bool {
	return request[bool]("SteppingStepOver", nil)
}
func (Debugger) SteppingInstrumentationStepInForTracking() bool {
	return request[bool]("SteppingInstrumentationStepInForTracking", nil)
}
func (Debugger) SteppingStepOverForGu(last_instruction bool) bool {
	return request[bool]("SteppingStepOverForGu", map[string]string{"last_instruction": strconv.FormatBool(last_instruction)})
}
func (Debugger) GetLocalApic(local_apic PLAPIC_PAGE, is_using_x2apic bool) bool {
	return request[bool]("GetLocalApic", map[string]string{
		"local_apic":      string(stream.MarshalJSON(local_apic)),
		"is_using_x2apic": strconv.FormatBool(is_using_x2apic),
	})
}
func (Debugger) GetIoApic(io_apic IO_APIC_ENTRY_PACKETS) bool {
	return request[bool]("GetIoApic", map[string]string{"io_apic": string(stream.MarshalJSON(io_apic))})
}
func (Debugger) GetIdtEntry(idt_packet INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS) bool {
	return request[bool]("GetIdtEntry", map[string]string{"idt_packet": string(stream.MarshalJSON(idt_packet))})
}
func (Debugger) HwdbgScriptRunScript(script string, instance_filepath_to_read string, hardware_script_file_path_to_save string, initial_bram_buffer_size uint32) bool {
	return request[bool]("HwdbgScriptRunScript", map[string]string{
		"script":                            script,
		"instance_filepath_to_read":         instance_filepath_to_read,
		"hardware_script_file_path_to_save": hardware_script_file_path_to_save,
		"initial_bram_buffer_size":          strconv.FormatUint(uint64(initial_bram_buffer_size), 10),
	})
}
func (Debugger) ScriptEngineWrapperTestParserForHwdbg(Expr string) {
	request[void]("ScriptEngineWrapperTestParserForHwdbg", map[string]string{"Expr": Expr})
}
func (Debugger) EnableTransparentMode(ProcessId uint32, ProcessName string, IsProcessId bool) bool {
	return request[bool]("EnableTransparentMode", map[string]string{
		"ProcessId":   strconv.FormatUint(uint64(ProcessId), 10),
		"ProcessName": ProcessName,
		"IsProcessId": strconv.FormatBool(IsProcessId),
	})
}
func (Debugger) DisableTransparentMode() bool {
	return request[bool]("DisableTransparentMode", nil)
}
