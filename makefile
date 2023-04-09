api_output_dir = bin/medusa
api_entry_point = cmd/medusa/main.go

web_package = ./web

app-install:
	yarn --cwd ${web_package} install

app-build: app-install
	yarn --cwd ${web_package} build

app-dev:
	yarn --cwd ${web_package} dev


server-build:
	go build -v -o ${api_output_dir} ${api_entry_point}

server-build-pi:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC="zig cc -target arm-linux-musleabihf" CXX="zig c++ -target arm-linux-musleabihf"	go build -v -o ${api_output_dir}_pi ${api_entry_point}

server-run: app-build server-build
	${api_output_dir}

server-dev:
	DEBUG=1 go run ${api_entry_point}

server-test:
	go test -count=1 ./...