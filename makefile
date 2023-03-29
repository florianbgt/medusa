default: run

output_dir = bin/medusa
entry_point = cmd/medusa/main.go
testpackage = ./...

compile:
	go build -o ${output_dir} ${entry_point}

run:
	go build -o ${output_dir} ${entry_point}
	bin/medusa

dev:
	go run ${entry_point}

test:
	go test -count=1 -v ${testpackage}