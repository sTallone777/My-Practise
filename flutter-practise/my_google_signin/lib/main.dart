import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:firebase_auth/firebase_auth.dart' as auth;
import 'package:my_google_signin/widget/home.dart';
import 'package:my_google_signin/widget/user.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  MyApp({Key key}) : super(key: key);

  @override
  _MyAppState createState() {
    return _MyAppState();
  }
}

class _MyAppState extends State<MyApp> {
  String _username = "";
  Widget currentPage;
  GoogleSignIn googleSignIn;
  Widget userPage;

  @override
  void initState() {
    super.initState();
    userPage = Home(
      onSiginin: (){
        _signin();
        print('Sign');
      },
      onLogout: _logout,
      showLoading: false,
    );
  }

  Future<auth.User> _signin() async{
    setState(() {
      userPage = Home(
        onSiginin: null,
        onLogout: _logout,
      );
    });

    auth.FirebaseAuth _auth = auth.FirebaseAuth.instance;
    try{
      googleSignIn = GoogleSignIn();
      GoogleSignInAccount googleSignInAccount = await googleSignIn.signIn();
      final GoogleSignInAuthentication gauth = await googleSignInAccount.authentication;
      final auth.AuthCredential credential = auth.GoogleAuthProvider.credential(
        accessToken: gauth.accessToken,
        idToken: gauth.idToken,
      );
      final auth.UserCredential authRes = await _auth.signInWithCredential(credential);
      final auth.User user = authRes.user;

      setState(() {
        _username = user.displayName;
        userPage = User(
          onLogout: _logout,
          user: user,
        );
      });

      return user;
    }catch(e){
      print(e.toString());
    }
    return null;
  }

  void _logout() async{
    await googleSignIn.signOut();
    setState(() {
      userPage = Home(
        onSiginin: (){
          _signin();
          print('Sign');
        },
        onLogout: _logout,
        showLoading: false,
      );
    });

    print('Logged Out');
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return MaterialApp(
      home: userPage,
    );
  }
}