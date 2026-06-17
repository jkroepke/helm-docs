.PHONY: helm-docs
helm-docs:
	go build github.com/norwoodj/helm-docs/cmd/helm-docs

.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o web/wasm/helm-docs.wasm ./cmd/helm-docs-wasm
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" web/wasm/wasm_exec.js

.PHONY: wasm-serve
wasm-serve: wasm
	python3 -m http.server 8080 --directory web/wasm

.PHONY: install
install:
	go install github.com/norwoodj/helm-docs/cmd/helm-docs

.PHONY: generate-example-charts
generate-example-charts: helm-docs
	./helm-docs --chart-search-root=example-charts --template-files=./_templates.gotmpl --template-files=README.md.gotmpl --document-dependency-values

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: gosec
gosec:
	@which gosec > /dev/null || go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -exclude G304 ./...

.PHONY: lint
lint: fmt gosec

.PHONY: clean
clean:
	rm -f helm-docs
	rm -f web/wasm/helm-docs.wasm web/wasm/wasm_exec.js

.PHONY: dist
dist:
	goreleaser release --rm-dist --snapshot --skip=sign

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  helm-docs    - Build the helm-docs binary"
	@echo "  wasm         - Build the browser WebAssembly demo"
	@echo "  wasm-serve   - Build and serve the browser WebAssembly demo on localhost:8080"
	@echo "  install      - Install the helm-docs binary"
	@echo "  fmt          - Format Go code"
	@echo "  test         - Run tests"
	@echo "  gosec        - Run security scan with gosec"
	@echo "  lint         - Run all linters (fmt, gosec)"
	@echo "  clean        - Clean build artifacts"
	@echo "  dist         - Create distribution with goreleaser"
	@echo "  help         - Show this help message"
