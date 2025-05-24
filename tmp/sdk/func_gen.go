package sdk
// ts_parser_new (:207)
func ts_parser_new()  {
	// TODO: implement make json packet
}

// ts_parser_delete (:212)
func ts_parser_delete(self TSParser *)  {
	// TODO: implement make json packet
}

// ts_parser_language (:217)
func ts_parser_language(self const TSParser *)  {
	// TODO: implement make json packet
}

// ts_parser_set_language (:229)
func ts_parser_set_language(self TSParser *, language const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_parser_set_included_ranges (:253)
func ts_parser_set_included_ranges(self TSParser *, ranges const TSRange *, count int)  {
	// TODO: implement make json packet
}

// ts_parser_included_ranges (:266)
func ts_parser_included_ranges(self const TSParser *, count int *)  {
	// TODO: implement make json packet
}

// ts_parser_parse (:316)
func ts_parser_parse(self TSParser *, old_tree const TSTree *, input TSInput)  {
	// TODO: implement make json packet
}

// ts_parser_parse_with_options (:329)
func ts_parser_parse_with_options(self TSParser *, old_tree const TSTree *, input TSInput, parse_options TSParseOptions)  {
	// TODO: implement make json packet
}

// ts_parser_parse_string (:342)
func ts_parser_parse_string(self TSParser *, old_tree const TSTree *, string string, length int)  {
	// TODO: implement make json packet
}

// ts_parser_parse_string_encoding (:355)
func ts_parser_parse_string_encoding(self TSParser *, old_tree const TSTree *, string string, length int, encoding TSInputEncoding)  {
	// TODO: implement make json packet
}

// ts_parser_reset (:372)
func ts_parser_reset(self TSParser *)  {
	// TODO: implement make json packet
}

// ts_parser_set_timeout_micros (:383)
func ts_parser_set_timeout_micros(self TSParser *, timeout_micros int)  {
	// TODO: implement make json packet
}

// ts_parser_timeout_micros (:390)
func ts_parser_timeout_micros(self const TSParser *)  {
	// TODO: implement make json packet
}

// ts_parser_set_cancellation_flag (:401)
func ts_parser_set_cancellation_flag(self TSParser *, flag const uint *)  {
	// TODO: implement make json packet
}

// ts_parser_cancellation_flag (:408)
func ts_parser_cancellation_flag(self const TSParser *)  {
	// TODO: implement make json packet
}

// ts_parser_set_logger (:417)
func ts_parser_set_logger(self TSParser *, logger TSLogger)  {
	// TODO: implement make json packet
}

// ts_parser_logger (:422)
func ts_parser_logger(self const TSParser *)  {
	// TODO: implement make json packet
}

// ts_parser_print_dot_graphs (:430)
func ts_parser_print_dot_graphs(self TSParser *, fd int)  {
	// TODO: implement make json packet
}

// ts_tree_copy (:442)
func ts_tree_copy(self const TSTree *)  {
	// TODO: implement make json packet
}

// ts_tree_delete (:447)
func ts_tree_delete(self TSTree *)  {
	// TODO: implement make json packet
}

// ts_tree_root_node (:452)
func ts_tree_root_node(self const TSTree *)  {
	// TODO: implement make json packet
}

// ts_tree_root_node_with_offset (:458)
func ts_tree_root_node_with_offset(self const TSTree *, offset_bytes int, offset_extent TSPoint)  {
	// TODO: implement make json packet
}

// ts_tree_language (:467)
func ts_tree_language(self const TSTree *)  {
	// TODO: implement make json packet
}

// ts_tree_included_ranges (:474)
func ts_tree_included_ranges(self const TSTree *, length int *)  {
	// TODO: implement make json packet
}

// ts_tree_edit (:483)
func ts_tree_edit(self TSTree *, edit const TSInputEdit *)  {
	// TODO: implement make json packet
}

// ts_tree_get_changed_ranges (:506)
func ts_tree_get_changed_ranges(old_tree const TSTree *, new_tree const TSTree *, length int *)  {
	// TODO: implement make json packet
}

// ts_tree_print_dot_graph (:515)
func ts_tree_print_dot_graph(self const TSTree *, file_descriptor int)  {
	// TODO: implement make json packet
}

// ts_node_type (:524)
func ts_node_type(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_symbol (:529)
func ts_node_symbol(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_language (:534)
func ts_node_language(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_grammar_type (:540)
func ts_node_grammar_type(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_grammar_symbol (:547)
func ts_node_grammar_symbol(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_start_byte (:552)
func ts_node_start_byte(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_start_point (:557)
func ts_node_start_point(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_end_byte (:562)
func ts_node_end_byte(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_end_point (:567)
func ts_node_end_point(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_string (:575)
func ts_node_string(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_is_null (:582)
func ts_node_is_null(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_is_named (:589)
func ts_node_is_named(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_is_missing (:595)
func ts_node_is_missing(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_is_extra (:601)
func ts_node_is_extra(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_has_changes (:606)
func ts_node_has_changes(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_has_error (:611)
func ts_node_has_error(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_is_error (:616)
func ts_node_is_error(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_parse_state (:621)
func ts_node_parse_state(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_next_parse_state (:626)
func ts_node_next_parse_state(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_parent (:633)
func ts_node_parent(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_child_with_descendant (:640)
func ts_node_child_with_descendant(self TSNode, descendant TSNode)  {
	// TODO: implement make json packet
}

// ts_node_child (:646)
func ts_node_child(self TSNode, child_index int)  {
	// TODO: implement make json packet
}

// ts_node_field_name_for_child (:652)
func ts_node_field_name_for_child(self TSNode, child_index int)  {
	// TODO: implement make json packet
}

// ts_node_field_name_for_named_child (:658)
func ts_node_field_name_for_named_child(self TSNode, named_child_index int)  {
	// TODO: implement make json packet
}

// ts_node_child_count (:663)
func ts_node_child_count(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_named_child (:670)
func ts_node_named_child(self TSNode, child_index int)  {
	// TODO: implement make json packet
}

// ts_node_named_child_count (:677)
func ts_node_named_child_count(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_child_by_field_name (:682)
func ts_node_child_by_field_name(self TSNode, name string, name_length int)  {
	// TODO: implement make json packet
}

// ts_node_child_by_field_id (:694)
func ts_node_child_by_field_id(self TSNode, field_id TSFieldId)  {
	// TODO: implement make json packet
}

// ts_node_next_sibling (:699)
func ts_node_next_sibling(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_prev_sibling (:700)
func ts_node_prev_sibling(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_next_named_sibling (:705)
func ts_node_next_named_sibling(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_prev_named_sibling (:706)
func ts_node_prev_named_sibling(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_first_child_for_byte (:711)
func ts_node_first_child_for_byte(self TSNode, byte int)  {
	// TODO: implement make json packet
}

// ts_node_first_named_child_for_byte (:716)
func ts_node_first_named_child_for_byte(self TSNode, byte int)  {
	// TODO: implement make json packet
}

// ts_node_descendant_count (:721)
func ts_node_descendant_count(self TSNode)  {
	// TODO: implement make json packet
}

// ts_node_descendant_for_byte_range (:727)
func ts_node_descendant_for_byte_range(self TSNode, start int, end int)  {
	// TODO: implement make json packet
}

// ts_node_descendant_for_point_range (:728)
func ts_node_descendant_for_point_range(self TSNode, start TSPoint, end TSPoint)  {
	// TODO: implement make json packet
}

// ts_node_named_descendant_for_byte_range (:734)
func ts_node_named_descendant_for_byte_range(self TSNode, start int, end int)  {
	// TODO: implement make json packet
}

// ts_node_named_descendant_for_point_range (:735)
func ts_node_named_descendant_for_point_range(self TSNode, start TSPoint, end TSPoint)  {
	// TODO: implement make json packet
}

// ts_node_edit (:746)
func ts_node_edit(self TSNode *, edit const TSInputEdit *)  {
	// TODO: implement make json packet
}

// ts_node_eq (:751)
func ts_node_eq(self TSNode, other TSNode)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_new (:767)
func ts_tree_cursor_new(node TSNode)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_delete (:772)
func ts_tree_cursor_delete(self TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_reset (:778)
func ts_tree_cursor_reset(self TSTreeCursor *, node TSNode)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_reset_to (:786)
func ts_tree_cursor_reset_to(dst TSTreeCursor *, src const TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_current_node (:791)
func ts_tree_cursor_current_node(self const TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_current_field_name (:799)
func ts_tree_cursor_current_field_name(self const TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_current_field_id (:807)
func ts_tree_cursor_current_field_id(self const TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_parent (:818)
func ts_tree_cursor_goto_parent(self TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_next_sibling (:829)
func ts_tree_cursor_goto_next_sibling(self TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_previous_sibling (:844)
func ts_tree_cursor_goto_previous_sibling(self TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_first_child (:852)
func ts_tree_cursor_goto_first_child(self TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_last_child (:864)
func ts_tree_cursor_goto_last_child(self TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_descendant (:871)
func ts_tree_cursor_goto_descendant(self TSTreeCursor *, goal_descendant_index int)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_current_descendant_index (:877)
func ts_tree_cursor_current_descendant_index(self const TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_current_depth (:883)
func ts_tree_cursor_current_depth(self const TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_first_child_for_byte (:892)
func ts_tree_cursor_goto_first_child_for_byte(self TSTreeCursor *, goal_byte int)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_goto_first_child_for_point (:893)
func ts_tree_cursor_goto_first_child_for_point(self TSTreeCursor *, goal_point TSPoint)  {
	// TODO: implement make json packet
}

// ts_tree_cursor_copy (:895)
func ts_tree_cursor_copy(cursor const TSTreeCursor *)  {
	// TODO: implement make json packet
}

// ts_query_new (:912)
func ts_query_new(language const TSLanguage *, source string, source_len int, error_offset int *, error_type TSQueryError *)  {
	// TODO: implement make json packet
}

// ts_query_delete (:923)
func ts_query_delete(self TSQuery *)  {
	// TODO: implement make json packet
}

// ts_query_pattern_count (:928)
func ts_query_pattern_count(self const TSQuery *)  {
	// TODO: implement make json packet
}

// ts_query_capture_count (:929)
func ts_query_capture_count(self const TSQuery *)  {
	// TODO: implement make json packet
}

// ts_query_string_count (:930)
func ts_query_string_count(self const TSQuery *)  {
	// TODO: implement make json packet
}

// ts_query_start_byte_for_pattern (:938)
func ts_query_start_byte_for_pattern(self const TSQuery *, pattern_index int)  {
	// TODO: implement make json packet
}

// ts_query_end_byte_for_pattern (:946)
func ts_query_end_byte_for_pattern(self const TSQuery *, pattern_index int)  {
	// TODO: implement make json packet
}

// ts_query_predicates_for_pattern (:964)
func ts_query_predicates_for_pattern(self const TSQuery *, pattern_index int, step_count int *)  {
	// TODO: implement make json packet
}

// ts_query_is_pattern_rooted (:973)
func ts_query_is_pattern_rooted(self const TSQuery *, pattern_index int)  {
	// TODO: implement make json packet
}

// ts_query_is_pattern_non_local (:983)
func ts_query_is_pattern_non_local(self const TSQuery *, pattern_index int)  {
	// TODO: implement make json packet
}

// ts_query_is_pattern_guaranteed_at_step (:989)
func ts_query_is_pattern_guaranteed_at_step(self const TSQuery *, byte_offset int)  {
	// TODO: implement make json packet
}

// ts_query_capture_name_for_id (:996)
func ts_query_capture_name_for_id(self const TSQuery *, index int, length int *)  {
	// TODO: implement make json packet
}

// ts_query_capture_quantifier_for_id (:1006)
func ts_query_capture_quantifier_for_id(self const TSQuery *, pattern_index int, capture_index int)  {
	// TODO: implement make json packet
}

// ts_query_string_value_for_id (:1012)
func ts_query_string_value_for_id(self const TSQuery *, index int, length int *)  {
	// TODO: implement make json packet
}

// ts_query_disable_capture (:1025)
func ts_query_disable_capture(self TSQuery *, name string, length int)  {
	// TODO: implement make json packet
}

// ts_query_disable_pattern (:1033)
func ts_query_disable_pattern(self TSQuery *, pattern_index int)  {
	// TODO: implement make json packet
}

// ts_query_cursor_new (:1057)
func ts_query_cursor_new()  {
	// TODO: implement make json packet
}

// ts_query_cursor_delete (:1062)
func ts_query_cursor_delete(self TSQueryCursor *)  {
	// TODO: implement make json packet
}

// ts_query_cursor_exec (:1067)
func ts_query_cursor_exec(self TSQueryCursor *, query const TSQuery *, node TSNode)  {
	// TODO: implement make json packet
}

// ts_query_cursor_exec_with_options (:1072)
func ts_query_cursor_exec_with_options(self TSQueryCursor *, query const TSQuery *, node TSNode, query_options const TSQueryCursorOptions *)  {
	// TODO: implement make json packet
}

// ts_query_cursor_did_exceed_match_limit (:1090)
func ts_query_cursor_did_exceed_match_limit(self const TSQueryCursor *)  {
	// TODO: implement make json packet
}

// ts_query_cursor_match_limit (:1091)
func ts_query_cursor_match_limit(self const TSQueryCursor *)  {
	// TODO: implement make json packet
}

// ts_query_cursor_set_match_limit (:1092)
func ts_query_cursor_set_match_limit(self TSQueryCursor *, limit int)  {
	// TODO: implement make json packet
}

// ts_query_cursor_set_timeout_micros (:1103)
func ts_query_cursor_set_timeout_micros(self TSQueryCursor *, timeout_micros int)  {
	// TODO: implement make json packet
}

// ts_query_cursor_timeout_micros (:1112)
func ts_query_cursor_timeout_micros(self const TSQueryCursor *)  {
	// TODO: implement make json packet
}

// ts_query_cursor_set_byte_range (:1129)
func ts_query_cursor_set_byte_range(self TSQueryCursor *, start_byte int, end_byte int)  {
	// TODO: implement make json packet
}

// ts_query_cursor_set_point_range (:1146)
func ts_query_cursor_set_point_range(self TSQueryCursor *, start_point TSPoint, end_point TSPoint)  {
	// TODO: implement make json packet
}

// ts_query_cursor_next_match (:1154)
func ts_query_cursor_next_match(self TSQueryCursor *, match TSQueryMatch *)  {
	// TODO: implement make json packet
}

// ts_query_cursor_remove_match (:1155)
func ts_query_cursor_remove_match(self TSQueryCursor *, match_id int)  {
	// TODO: implement make json packet
}

// ts_query_cursor_next_capture (:1163)
func ts_query_cursor_next_capture(self TSQueryCursor *, match TSQueryMatch *, capture_index int *)  {
	// TODO: implement make json packet
}

// ts_query_cursor_set_max_start_depth (:1183)
func ts_query_cursor_set_max_start_depth(self TSQueryCursor *, max_start_depth int)  {
	// TODO: implement make json packet
}

// ts_language_copy (:1192)
func ts_language_copy(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_delete (:1198)
func ts_language_delete(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_symbol_count (:1203)
func ts_language_symbol_count(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_state_count (:1208)
func ts_language_state_count(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_symbol_for_name (:1213)
func ts_language_symbol_for_name(self const TSLanguage *, string string, length int, is_named bool)  {
	// TODO: implement make json packet
}

// ts_language_field_count (:1223)
func ts_language_field_count(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_field_name_for_id (:1228)
func ts_language_field_name_for_id(self const TSLanguage *, id TSFieldId)  {
	// TODO: implement make json packet
}

// ts_language_field_id_for_name (:1233)
func ts_language_field_id_for_name(self const TSLanguage *, name string, name_length int)  {
	// TODO: implement make json packet
}

// ts_language_supertypes (:1238)
func ts_language_supertypes(self const TSLanguage *, length int *)  {
	// TODO: implement make json packet
}

// ts_language_subtypes (:1245)
func ts_language_subtypes(self const TSLanguage *, supertype TSSymbol, length int *)  {
	// TODO: implement make json packet
}

// ts_language_symbol_name (:1254)
func ts_language_symbol_name(self const TSLanguage *, symbol TSSymbol)  {
	// TODO: implement make json packet
}

// ts_language_symbol_type (:1262)
func ts_language_symbol_type(self const TSLanguage *, symbol TSSymbol)  {
	// TODO: implement make json packet
}

// ts_language_version (:1273)
func ts_language_version(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_abi_version (:1282)
func ts_language_abi_version(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_metadata (:1291)
func ts_language_metadata(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_language_next_state (:1298)
func ts_language_next_state(self const TSLanguage *, state TSStateId, symbol TSSymbol)  {
	// TODO: implement make json packet
}

// ts_language_name (:1303)
func ts_language_name(self const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_new (:1324)
func ts_lookahead_iterator_new(self const TSLanguage *, state TSStateId)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_delete (:1329)
func ts_lookahead_iterator_delete(self TSLookaheadIterator *)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_reset_state (:1337)
func ts_lookahead_iterator_reset_state(self TSLookaheadIterator *, state TSStateId)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_reset (:1345)
func ts_lookahead_iterator_reset(self TSLookaheadIterator *, language const TSLanguage *, state TSStateId)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_language (:1350)
func ts_lookahead_iterator_language(self const TSLookaheadIterator *)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_next (:1357)
func ts_lookahead_iterator_next(self TSLookaheadIterator *)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_current_symbol (:1362)
func ts_lookahead_iterator_current_symbol(self const TSLookaheadIterator *)  {
	// TODO: implement make json packet
}

// ts_lookahead_iterator_current_symbol_name (:1368)
func ts_lookahead_iterator_current_symbol_name(self const TSLookaheadIterator *)  {
	// TODO: implement make json packet
}

// ts_wasm_store_new (:1393)
func ts_wasm_store_new(engine TSWasmEngine *, error TSWasmError *)  {
	// TODO: implement make json packet
}

// ts_wasm_store_delete (:1401)
func ts_wasm_store_delete( TSWasmStore *)  {
	// TODO: implement make json packet
}

// ts_wasm_store_load_language (:1410)
func ts_wasm_store_load_language( TSWasmStore *, name string, wasm string, wasm_len int, error TSWasmError *)  {
	// TODO: implement make json packet
}

// ts_wasm_store_language_count (:1421)
func ts_wasm_store_language_count( const TSWasmStore *)  {
	// TODO: implement make json packet
}

// ts_language_is_wasm (:1427)
func ts_language_is_wasm( const TSLanguage *)  {
	// TODO: implement make json packet
}

// ts_parser_set_wasm_store (:1433)
func ts_parser_set_wasm_store( TSParser *,  TSWasmStore *)  {
	// TODO: implement make json packet
}

// ts_parser_take_wasm_store (:1439)
func ts_parser_take_wasm_store( TSParser *)  {
	// TODO: implement make json packet
}

// ts_set_allocator (:1461)
func ts_set_allocator(new_malloc void *(*)(uint), new_calloc void *(*)(uint, uint), new_realloc void *(*)(void *, uint), new_free void (*)(void *))  {
	// TODO: implement make json packet
}

