init:
	@git config core.hooksPath .githooks
	@chmod +x .githooks/*
	@echo "Initialized git hooks."

tidy:
	@go mod tidy
	@go mod vendor

lint:
	@sh -c "$(PWD)/scripts/linter.sh"

test_unit:
	@TEST_MODE=unit go test ./...
