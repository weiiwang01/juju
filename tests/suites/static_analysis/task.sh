test_static_analysis() {
	if [ "$(skip 'test_static_analysis')" ]; then
		echo "==> TEST SKIPPED: skip static analysis"
		return
	fi

	set_verbosity

	test_copyright
	test_doc_go
	test_versions
	test_static_analysis_shell
	test_static_analysis_python
	test_schema

	# slow ones go last
	test_static_analysis_go
}
