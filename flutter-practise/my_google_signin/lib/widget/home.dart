import 'package:flutter/material.dart';

class Home extends StatelessWidget {
  Home({Key key,
  @required this.onSiginin,
  @required this.onLogout,
  @required this.showLoading}) : super(key: key);

  final VoidCallback onSiginin;
  final VoidCallback onLogout;
  bool showLoading = false;

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Scaffold(
      appBar: AppBar(title: Text('Sign In')),
      body: Container(
        padding: const EdgeInsets.all(20.0),
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              showLoading
              ? CircularProgressIndicator()
                  : RaisedButton(
                onPressed: this.onSiginin,
                child: Text('Sign In'),
                color: Colors.lightBlueAccent,
              )
            ],
          ),
        ),
      ),
    );
  }
}
