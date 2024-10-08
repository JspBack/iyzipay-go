run:
	go test -v ./test -vet=all

pre-commit:
	pip install pre-commit
	@pre-commit --version
	@pre-commit install
	@pre-commit autoupdate

.PHONY: run pre-commit
