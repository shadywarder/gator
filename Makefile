TARGET ?= gator

.PHONY: build
build:
	@echo "go build for the target $(TARGET) is in progress!"
	@mkdir -p ./bin
	@go build -o ./bin/$(TARGET) ./cmd/$(TARGET)

.PHONY: run
run: build
	@echo "running $(TARGET)"
	@./bin/$(TARGET) $(cmd) $(user) $(if $(name), "$(name)") $(url) $(time) $(limit)

.PHONY: tidy
tidy:
	@go mod tidy -v
	@go fmt ./...

.PHONY: clean
clear:
	@rm -rf ./bin