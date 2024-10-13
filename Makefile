BINARY_NAME=tt

.PHONY: bin
bin:
	go build -o artifacts/${BINARY_NAME} cmd/tt/main.go

.PHONY: run
run:
	go run cmd/tt/main.go

# .SILENT:
.PHONY: clean
clean:
	go clean
	@rm -r artifacts 2> /dev/null

$(V).SILENT:
