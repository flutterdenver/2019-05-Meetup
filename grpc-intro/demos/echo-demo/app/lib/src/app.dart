import 'package:echo_demo/src/screens/home_screen.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';

class App extends StatefulWidget {
  final String host;
  final int port;

  App({
    this.host = '127.0.0.1',
    this.port = 9000,
  });

  @override
  _AppState createState() => _AppState(
        host: this.host,
        port: this.port,
      );
}

class _AppState extends State<App> {
  ClientChannel _channel;

  _AppState({
    String host,
    int port,
  }) {
    print('App created');
    _channel = ClientChannel(
      host,
      port: port,
      options: const ChannelOptions(
        credentials: const ChannelCredentials.insecure(),
        idleTimeout: const Duration(seconds: 1),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Echo Demo',
      home: HomeScreen(),
    );
  }

  @override
  @mustCallSuper
  void dispose() {
    _channel.shutdown();
    super.dispose();
    print('App disposed');
  }
}
