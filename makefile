app-install:
	yarn --cwd ./web install

app-build: app-install
	yarn --cwd ./web build

app-dev:
	yarn --cwd ./web dev


server-build:
	go build -v -o bin/medusa cmd/medusa/main.go

server-build-pi:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC="zig cc -target arm-linux-musleabihf" CXX="zig c++ -target arm-linux-musleabihf"	go build -v -o ${api_output_dir}_pi ${api_entry_point}

server-run: app-build server-build
	bin/medusa

server-dev:
	DEBUG=1 go run cmd/medusa/main.go

server-test:
	go test -count=1 ./...