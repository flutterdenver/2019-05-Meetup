import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Center(
          child: RaisedButton(
            child: Text('Ahoy!'),
            onPressed: () {
              print('Ahoy!');
            },
          ),
        ),
      ),
    );
  }
}
