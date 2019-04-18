// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var echo_service_pb = require('./echo_service_pb.js');

function serialize_echo_EchoMessage(arg) {
  if (!(arg instanceof echo_service_pb.EchoMessage)) {
    throw new Error('Expected argument of type echo.EchoMessage');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_echo_EchoMessage(buffer_arg) {
  return echo_service_pb.EchoMessage.deserializeBinary(new Uint8Array(buffer_arg));
}


var EchoServiceService = exports.EchoServiceService = {
  echo: {
    path: '/echo.EchoService/Echo',
    requestStream: false,
    responseStream: false,
    requestType: echo_service_pb.EchoMessage,
    responseType: echo_service_pb.EchoMessage,
    requestSerialize: serialize_echo_EchoMessage,
    requestDeserialize: deserialize_echo_EchoMessage,
    responseSerialize: serialize_echo_EchoMessage,
    responseDeserialize: deserialize_echo_EchoMessage,
  },
};

exports.EchoServiceClient = grpc.makeGenericClientConstructor(EchoServiceService);
