import 'package:animation/src/basics/01_animated_container.dart';
import 'package:animation/src/misc/animated_positioned.dart';
import 'package:animation/src/misc/combo_animate.dart';
import 'package:animation/src/misc/focus_image.dart';
import 'package:animation/src/misc/parallel_animate.dart';
import 'package:flutter/material.dart';

import 'package:animation/src/basics/logo_app.dart';
import 'package:animation/src/misc/curve_page.dart';

void main() {
  runApp(MyApp());
}

class Demo {
  final String name;
  final String route;
  final WidgetBuilder builder;

  const Demo({this.name, this.route, this.builder});
}

final basicDemos = [
  Demo(
    name: 'AnimatedContainer',
    route: AnimatedContainerDemo.routeName,
    builder: (context) => AnimatedContainerDemo(),
  ),
  Demo(
    name: 'LogoApp',
    route: LogoApp.routeName,
    builder: (context) => LogoApp(),
  ),
];

final miscDemos = [
  Demo(
    name: 'Focus Image',
    route: FocusImageDemo.routeName,
    builder: (context) => FocusImageDemo(),
  ),
  Demo(
      name: 'AnimatedPositioned',
      route: AnimatedPositionedDemo.routeName,
      builder: (context) => AnimatedPositionedDemo()
  ),
  Demo(
      name: 'CurvePage',
      route: CurvePage.routeName,
      builder: (context) => CurvePage()
  ),
  Demo(
      name: 'Parallel',
      route: ParallelAnimate.routeName,
      builder: (context) => ParallelAnimate()
  ),
  Demo(
      name: 'ComboAnimate',
      route: StaggerRoute.routeName,
      builder: (context) => StaggerRoute()
  ),
];

final basicDemoRoutes = Map.fromEntries(basicDemos.map((d) => MapEntry(d.route, d.builder)));
final miscDemoRoutes = Map.fromEntries(miscDemos.map((d) => MapEntry(d.route, d.builder)));

final allRoutes = <String, WidgetBuilder>{
  ...basicDemoRoutes,
  ...miscDemoRoutes,
};

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Animation Samples',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        // This makes the visual density adapt to the platform that you run
        // the app on. For desktop platforms, the controls will be smaller and
        // closer together (more dense) than on mobile platforms.
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      routes: allRoutes,
      home: MyHomePage(),
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final headerStyle = Theme
        .of(context)
        .textTheme
        .headline6;
    return Scaffold(
      appBar: AppBar(
        title: Text('Animation Samples'),
      ),
      body: ListView(
          children: [
            ListTile(title: Text('Basics', style: headerStyle)),
            ...basicDemos.map((d) => DemoTile(d)),
            ListTile(title: Text('Misc', style: headerStyle)),
            ...miscDemos.map((d) => DemoTile(d)),
          ]
      ),
    );
  }
}

class DemoTile extends StatelessWidget{
  final Demo demo;
  DemoTile(this.demo);
  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text(demo.name),
      onTap: () {
        Navigator.pushNamed(context, demo.route);
      }
    );
  }
}
