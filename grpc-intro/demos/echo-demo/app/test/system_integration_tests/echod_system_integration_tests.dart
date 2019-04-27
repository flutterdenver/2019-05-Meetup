import 'dart:convert';
import 'dart:io';

import 'package:flutter_test/flutter_test.dart';
import 'package:grpc/grpc.dart';
import 'package:echo_demo/api/echo_service.pbgrpc.dart';

void main() {
  Process process;
  ClientChannel channel;
  EchoServiceClient echoServiceClient;

  setUpAll(() async {
    var port = 9010;

    process = await Process.start(
      'docker',
      ['run', '-p', '$port:9000', 'flutterdenver/echod:latest'],
      environment: {
        'PORT': ':$port',
      },
    );
    process.stderr.transform(utf8.decoder).listen(print);
    process.stdout.transform(utf8.decoder).listen(print);

    channel = ClientChannel(
      '127.0.0.1',
      port: port,
      options: const ChannelOptions(
        credentials: const ChannelCredentials.insecure(),
        idleTimeout: Duration(seconds: 1),
      ),
    );

    echoServiceClient = EchoServiceClient(
      channel,
      options: CallOptions(),
    );

    // TODO: How do I "wait" for a client connection instead of sleeping?
    sleep(Duration(seconds: 1));
  });

  tearDownAll(() async {
    await channel.shutdown();
    expect(true, process.kill(ProcessSignal.sigint));
  });

  test('echoing a simple message', () async {
    final req = EchoMessage()..value = "Hello World!";
    final res = await echoServiceClient.echo(req);
    expect(res.value, "Hello World!");
  });

  test('StatusCode.aborted', () async {
    final req = EchoMessage()..value = "Respond with Aborted";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.aborted);
      expect(e.message, 'Planned error for {Respond with Aborted}; Code {10}');
    }
  });

  test('StatusCode.alreadyExists', () async {
    final req = EchoMessage()..value = "Respond with AlreadyExists";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.alreadyExists);
      expect(e.message,
          'Planned error for {Respond with AlreadyExists}; Code {6}');
    }
  });

  test('StatusCode.cancelled', () async {
    final req = EchoMessage()..value = "Respond with Cancelled";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.cancelled);
      expect(e.message, 'Planned error for {Respond with Cancelled}; Code {1}');
    }
  });

  test('StatusCode.dataLoss', () async {
    final req = EchoMessage()..value = "Respond with DataLoss";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.dataLoss);
      expect(e.message, 'Planned error for {Respond with DataLoss}; Code {15}');
    }
  });

  test('StatusCode.deadlineExceeded', () async {
    final req = EchoMessage()..value = "Respond with DeadlineExceeded";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.deadlineExceeded);
      expect(e.message,
          'Planned error for {Respond with DeadlineExceeded}; Code {4}');
    }
  });

  test('StatusCode.failedPrecondition', () async {
    final req = EchoMessage()..value = "Respond with FailedPrecondition";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.failedPrecondition);
      expect(e.message,
          'Planned error for {Respond with FailedPrecondition}; Code {9}');
    }
  });

  test('StatusCode.internal', () async {
    final req = EchoMessage()..value = "Respond with Internal";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.internal);
      expect(e.message, 'Planned error for {Respond with Internal}; Code {13}');
    }
  });

  test('StatusCode.invalidArgument', () async {
    final req = EchoMessage()..value = "Respond with InvalidArgument";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.invalidArgument);
      expect(e.message,
          'Planned error for {Respond with InvalidArgument}; Code {3}');
    }
  });

  test('StatusCode.notFound', () async {
    final req = EchoMessage()..value = "Respond with NotFound";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.notFound);
      expect(e.message, 'Planned error for {Respond with NotFound}; Code {5}');
    }
  });

  test('StatusCode.outOfRange', () async {
    final req = EchoMessage()..value = "Respond with OutOfRange";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.outOfRange);
      expect(
          e.message, 'Planned error for {Respond with OutOfRange}; Code {11}');
    }
  });

  test('StatusCode.permissionDenied', () async {
    final req = EchoMessage()..value = "Respond with PermissionDenied";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.permissionDenied);
      expect(e.message,
          'Planned error for {Respond with PermissionDenied}; Code {7}');
    }
  });

  test('StatusCode.resourceExhausted', () async {
    final req = EchoMessage()..value = "Respond with ResourceExhausted";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.resourceExhausted);
      expect(e.message,
          'Planned error for {Respond with ResourceExhausted}; Code {8}');
    }
  });

  test('StatusCode.unauthenticated', () async {
    final req = EchoMessage()..value = "Respond with Unauthenticated";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.unauthenticated);
      expect(e.message,
          'Planned error for {Respond with Unauthenticated}; Code {16}');
    }
  });

  test('StatusCode.unavailable', () async {
    final req = EchoMessage()..value = "Respond with Unavailable";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.unavailable);
      expect(
          e.message, 'Planned error for {Respond with Unavailable}; Code {14}');
    }
  });

  test('StatusCode.unimplemented', () async {
    final req = EchoMessage()..value = "Respond with Unimplemented";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.unimplemented);
      expect(e.message,
          'Planned error for {Respond with Unimplemented}; Code {12}');
    }
  });

  test('StatusCode.unknown', () async {
    final req = EchoMessage()..value = "Respond with Unknown";
    try {
      final res = await echoServiceClient.echo(req);
      expect(res, null);
    } on GrpcError catch (e) {
      expect(e.code, StatusCode.unknown);
      expect(e.message, 'Planned error for {Respond with Unknown}; Code {2}');
    }
  });
}
