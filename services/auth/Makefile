
.PHONY: generate_mock
generate_mock:
	mockery -all -dir ./token

.PHONY: clean_mock
clean_mock:
	rm -rf mocks

.PHONY: test
test:
	go test -v -cover ./...

