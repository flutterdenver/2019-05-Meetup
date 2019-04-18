.DEFAULT_GOAL := generate

generate: \
	generate-dart \
	generate-go \
	generate-node

generate-dart:
	@# README https://grpc.io/docs/quickstart/dart.html#prerequisites
	@mkdir -p gen/pb-dart
	@protoc --dart_out=grpc:gen/pb-dart echo_service.proto

generate-go:
	@docker run -v `pwd`:/defs namely/protoc-all -f echo_service.proto -l go

generate-node:
	@docker run -v `pwd`:/defs namely/protoc-all -f echo_service.proto -l node
