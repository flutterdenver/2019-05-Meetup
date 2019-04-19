
build:
	@docker build --tag flutterdenver/echod .

run:
	@docker run -p 9000:9000 flutterdenver/echod:latest

cover:
	@go test -coverprofile=${TMPDIR}/coverprofile ./...
	@go tool cover -html=${TMPDIR}/coverprofile

test:
	@go test -cover ./...

watch:
	@reflex -d none -c reflex.conf
