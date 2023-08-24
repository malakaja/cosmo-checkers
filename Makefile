gen-protoc-ts:
	mkdir -p ./client/src/types/generated/
	ls proto/checkers | xargs -I {} protoc \
		--plugin="./scripts/node_modules/.bin/protoc-gen-ts_proto" \
		--ts_proto_out="./client/src/types/generated" \
		--proto_path="./proto" \
		--proto_path="$(HOME)/protoc/protobuf/src" \
		--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
		checkers/{}




build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/checkersd-linux-amd64 ./cmd/checkersd/main.go
	GOOS=linux GOARCH=arm64 go build -o ./build/checkersd-linux-arm64 ./cmd/checkersd/main.go

do-checksum-linux:
	cd build && sha256sum \
		checkersd-linux-amd64 checkersd-linux-arm64 \
		> checkers-checksum-linux

build-linux-with-checksum: build-linux do-checksum-linux



build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/checkersd-darwin-amd64 ./cmd/checkersd/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/checkersd-darwin-arm64 ./cmd/checkersd/main.go

build-all: build-linux build-darwin

do-checksum-darwin:
	cd build && sha256sum \
		checkersd-darwin-amd64 checkersd-darwin-arm64 \
		> checkers-checksum-darwin

build-darwin-with-checksum: build-darwin do-checksum-darwin

build-with-checksum: build-linux-with-checksum build-darwin-with-checksum