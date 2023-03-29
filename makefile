api_output_dir = bin/medusa
api_entry_point = cmd/medusa/main.go

web_package = ./web

api-build:
	go build -o ${api_output_dir} ${api_entry_point}

api-run: api-build
	${api_output_dir}

api-dev:
	go run ${api_entry_point}

api-test:
	go test -count=1 -v ./...


app-install:
	yarn --cwd ${web_package} install

app-build: app-install
	yarn --cwd ${web_package} build

app-export:
	cp -r ${web_package}/out/* ./website

app-run: app-build
	yarn --cwd ${web_package} start

app-dev:
	yarn --cwd ${web_package} dev