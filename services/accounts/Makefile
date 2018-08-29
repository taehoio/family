
.PHONY: generate_mock
generate_mocks:
	mockery -dir=./pkg/crypt -name=Crypt
	mockery -dir=./internal/repo -name=AccountsRepo

.PHONY: clean_mock
clean_mocks:
	find . -name "mock_*.go" -type f -delete
	rm -rf mocks

.PHONY: test
test:
	go test -v -cover ./...
