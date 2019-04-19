
build:
	@docker build --tag echod .

run: build
	@docker run -p 9000:9000 echod:latest

cover:
	@go test -coverprofile=${TMPDIR}/coverprofile ./...
	@go tool cover -html=${TMPDIR}/coverprofile

test:
	@go test -cover ./...

watch:
	@reflex -d none -c reflex.conf
