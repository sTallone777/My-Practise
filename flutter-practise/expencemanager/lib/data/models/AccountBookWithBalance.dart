import 'package:expencemanager/data/datasources/LocalDatabase.dart';

class AccountBookWithBalance{
  account_book book;
  double income;
  double expense;

  AccountBookWithBalance(this.book, this.income, this.expense);
}