
.PHONY: generate_mock
generate_mocks:
	mockery -all -keeptree

.PHONY: clean_mock
clean_mocks:
	rm -rf mocks

.PHONY: test
test:
	go test -v -cover ./...

