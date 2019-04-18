///
//  Generated code. Do not modify.
//  source: echo_service.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async' as $async;

import 'package:grpc/grpc.dart';

import 'echo_service.pb.dart';
export 'echo_service.pb.dart';

class EchoServiceClient extends Client {
  static final _$echo = new ClientMethod<EchoMessage, EchoMessage>(
      '/echo.EchoService/Echo',
      (EchoMessage value) => value.writeToBuffer(),
      (List<int> value) => new EchoMessage.fromBuffer(value));

  EchoServiceClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseFuture<EchoMessage> echo(EchoMessage request, {CallOptions options}) {
    final call = $createCall(_$echo, new $async.Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }
}

abstract class EchoServiceBase extends Service {
  String get $name => 'echo.EchoService';

  EchoServiceBase() {
    $addMethod(new ServiceMethod<EchoMessage, EchoMessage>(
        'Echo',
        echo_Pre,
        false,
        false,
        (List<int> value) => new EchoMessage.fromBuffer(value),
        (EchoMessage value) => value.writeToBuffer()));
  }

  $async.Future<EchoMessage> echo_Pre(
      ServiceCall call, $async.Future request) async {
    return echo(call, await request);
  }

  $async.Future<EchoMessage> echo(ServiceCall call, EchoMessage request);
}
