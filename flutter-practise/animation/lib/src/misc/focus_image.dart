import 'package:flutter/material.dart';

class FocusImageDemo extends StatelessWidget{
  static String routeName = '/misc/focus_image';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text('Focus Image')),
      body: Grid(),
    );
  }
}

class Grid extends StatelessWidget{
  @override
  Widget build(BuildContext context){
    return Scaffold(
      body: GridView.builder(
        padding: EdgeInsets.all(100.0),
        itemCount: 20,
        gridDelegate:
          SliverGridDelegateWithFixedCrossAxisCount(crossAxisCount: 4),
        itemBuilder:(context, index) {
          return (index >= 10)
              ? SmallCard('assets/eat_cape_town_sm.jpg')
              : SmallCard('assets/eat_new_orleans_sm.jpg');
        },
      ),
    );
  }
}

Route _createRoute(BuildContext parentContext, String image){
  return PageRouteBuilder<void>(
    pageBuilder: (context, animation, secondaryAnimation){
      return _SecondPage(image);
    },
    transitionsBuilder: (context, animation, secondaryAnimation, child) {
      var rectAnimation = _createTween(parentContext)
          .chain(CurveTween(curve: Curves.ease))
          .animate(animation);

      return Stack(
        children: [
          PositionedTransition(rect: rectAnimation, child: child),
        ],
      );
    }
  );
}

Tween<RelativeRect> _createTween(BuildContext context){
  var windowSize = MediaQuery.of(context).size;
  var box = context.findRenderObject() as RenderBox;
  var rect = box.localToGlobal(Offset.zero) & box.size;
  var relativeRect = RelativeRect.fromSize(rect, windowSize);
  return RelativeRectTween(
    begin: relativeRect,
    end: RelativeRect.fill,
  );
}

class SmallCard extends StatelessWidget{
  final String imageAssetName;
  SmallCard(this.imageAssetName);
  @override
  Widget build(BuildContext context){
    return Card(
      child: Material(
        child: InkWell(
          onTap: (){
            var nav = Navigator.of(context);
            nav.push<void>(_createRoute(context, imageAssetName));
          },
          child: Image.asset(
            imageAssetName,
            fit: BoxFit.cover,
          ),
        ),
      ),
    );
  }
}

class _SecondPage extends StatelessWidget{
  final String imageAssetName;
  _SecondPage(this.imageAssetName);
  @override
  Widget build(BuildContext context){
    return Scaffold(
      backgroundColor: Colors.black,
      body: Center(
        child: Material(
          child: InkWell(
            onTap: () => Navigator.of(context).pop(),
            child: AspectRatio(
              aspectRatio: 1,
              child: Image.asset(
                imageAssetName,
                fit: BoxFit.cover,
              ),
            ),
          ),
        ),
      ),
    );
  }
}
