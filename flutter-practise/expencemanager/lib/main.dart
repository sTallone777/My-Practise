import 'package:expencemanager/data/datasources/LocalDatabase.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

void main() => runApp(MyApp());

class ExpenseManagerApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return Provider<LocalDatabase>(
      create: (context) => LocalDatabase(),
      child: MaterialApp(
        title: 'Expense Manager',
        theme: ThemeData.light(),
        onGenerateRoute: ,
        initialRoute: ,
      ),
      dispose: (context, db) => db.close(),
    );
  }
}