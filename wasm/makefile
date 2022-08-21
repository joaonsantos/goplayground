.PHONY: run build serve

run: build serve

build:
	GOOS=js GOARCH=wasm go build -o ./static/main.wasm ./cmd/wasm/main.go

serve:
	go run ./cmd/serve/main.go

