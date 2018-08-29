
.PHONY: generate_mock
generate_mocks:
	mockery -dir=./pkg/token -name=Token

.PHONY: clean_mock
clean_mocks:
	find . -name "mock_*.go" -type f -delete
	rm -rf mocks

.PHONY: test
test:
	go test -v -cover ./...

