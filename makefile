pre-commit:
	pip install pre-commit
	@pre-commit --version
	@pre-commit install
	@pre-commit autoupdate

.PHONY: pre-commit
