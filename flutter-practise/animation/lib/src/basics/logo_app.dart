import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class LogoApp extends StatefulWidget{
  static String routeName = '/basic/logo_app.dart';
  _LogoAppState createState () => new _LogoAppState();
}

class _LogoAppState extends State<LogoApp> with SingleTickerProviderStateMixin{
  Animation<double> animation;
  AnimationController controller;

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    controller = new AnimationController(
      duration: const Duration(milliseconds: 2000), vsync: this
    );
    animation = new Tween(begin: 0.0, end: 300.0).animate(controller)
    ..addListener(() {
      setState(() {

      });
    });
    controller.forward();
  }

  Widget build(BuildContext context){
    return new Center(
      child: new Container(
        margin: new EdgeInsets.symmetric(vertical: 10.0),
        height: animation.value,
        width: animation.value,
        child: new FlutterLogo(),
      ),
    );
  }

  @override
  void dispose() {
    controller.dispose();
    // TODO: implement dispose
    super.dispose();
  }
}