all: sqlc-gen-unsafe-paths sqlc-gen-unsafe-paths.wasm

sqlc-gen-unsafe-paths: main.go go.mod go.sum
	go build .

sqlc-gen-unsafe-paths.wasm: main.go go.mod go.sum
	GOOS=wasip1 GOARCH=wasm go build -o sqlc-gen-unsafe-paths.wasm main.go

sha256: sqlc-gen-unsafe-paths.wasm
	openssl sha256 sqlc-gen-unsafe-paths.wasm