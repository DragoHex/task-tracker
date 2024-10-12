BINARY_NAME=tt

.PHONY: bin
bin:
	go build -o ${BINARY_NAME}

.PHONY: run
run:
	go run main.go

# .SILENT:
.PHONY: clean
clean:
	go clean
	@rm ${BINARY_NAME} 2> /dev/null

$(V).SILENT:
