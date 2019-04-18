
build:
	@docker build --tag echo-backend .

run: build
	@docker run echo-backend:latest

cover:
	@go test -coverprofile=${TMPDIR}/coverprofile ./...
	@go tool cover -html=${TMPDIR}/coverprofile

test:
	@go test -cover ./...

watch:
	@reflex -d none -c reflex.conf
