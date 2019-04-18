.DEFAULT_GOAL := generate

generate: \
	generate-dart \
	generate-go

generate-dart:
	@mkdir -p ../app/lib/api
	@protoc --dart_out=grpc:../app/lib/api echo_service.proto

generate-go:
	@protoc --go_out=plugins=grpc:../backend/api echo_service.proto
