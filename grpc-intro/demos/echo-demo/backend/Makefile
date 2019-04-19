.DEFAULT_GOAL = build

build-echoc:
	@go install ./cmd/echoc

build-echod:
	@docker build --tag flutterdenver/echod .

build: build-echoc build-echod

run-echoc: build-echoc
	echoc
	echoc -r

run-echod: build-echod
	@docker run -p 9000:9000 flutterdenver/echod:latest

cover:
	@go test -coverprofile=${TMPDIR}/coverprofile ./...
	@go tool cover -html=${TMPDIR}/coverprofile

test:
	@go test -cover ./...

watch:
	@reflex -d none -c reflex.conf
