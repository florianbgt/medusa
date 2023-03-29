api_output_dir = bin/medusa
api_entry_point = cmd/medusa/main.go
api_test_package = ./...

web_package = ./web

api-compile:
	go build -o ${api_output_dir} ${api_entry_point}

api-run:
	go build -o ${api_output_dir} ${api_entry_point}
	${api_output_dir}

api-dev:
	go run ${api_entry_point}

api-test:
	go test -count=1 -v ${api_test_package}

app-dev:
	yarn --cwd ${web_package} dev