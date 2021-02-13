import 'package:flutter/material.dart';
import 'package:firebase_auth/firebase_auth.dart' as auth;
import 'package:flutter/cupertino.dart';

class User extends StatelessWidget {
  User({Key key, @required this.onLogout, @required this.user}) : super(key: key);

  VoidCallback onLogout;
  String username;
  auth.User user;

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Scaffold(
      appBar: AppBar(
        title: Text('Welcome'),
        actions: [
          IconButton(
            icon: Icon(Icons.exit_to_app),
            onPressed: this.onLogout,
          ),
        ],
      ),
      body: Container(
        padding: const EdgeInsets.all(20.0),
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Image.network(user.photoURL),
              Text(
                user.displayName,
                textScaleFactor: 1.5,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
