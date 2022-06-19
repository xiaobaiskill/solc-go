
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vat
vet:
	go vet ./...

.PHONY: build-mac-arm64
build-mac-arm64: fmt vet
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 SDKROOT=$(xcrun --sdk macosx --show-sdk-path) \
	go build -o build/solc .