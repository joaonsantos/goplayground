.PHONY: run build serve

run: build serve

build:
	@cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ./static/
	GOOS=js GOARCH=wasm go build -o ./static/main.wasm ./cmd/wasm/main.go

serve:
	go run ./cmd/serve/main.go

