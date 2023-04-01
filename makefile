api_output_dir = bin/medusa
api_entry_point = cmd/medusa/main.go

web_package = ./web

app-install:
	yarn --cwd ${web_package} install

app-build: app-install
	yarn --cwd ${web_package} build

app-dev:
	yarn --cwd ${web_package} dev


server-build:app-build
	go build -o ${api_output_dir} ${api_entry_point}

server-run: server-build
	${api_output_dir}

server-dev:
	go run ${api_entry_point}

server-test:
	go test -count=1 -v ./...