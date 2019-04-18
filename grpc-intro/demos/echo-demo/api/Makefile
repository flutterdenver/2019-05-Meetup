.DEFAULT_GOAL := generate

generate: generate-dart generate-go

generate-dart:
	@mkdir -p ../app/lib/api
	@protoc --dart_out=grpc:../app/lib/api *.proto

generate-go:
	@mkdir -p ../backend/api
	@protoc --go_out=plugins=grpc:../backend/api *.proto
