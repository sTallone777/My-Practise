import 'package:expencemanager/data/datasources/AccountBookDao.dart';
import 'package:expencemanager/data/datasources/CategoryDao.dart';
import 'package:expencemanager/data/datasources/EntryDao.dart';
import 'package:expencemanager/data/datasources/WalletDao.dart';
import 'package:expencemanager/data/models/AccountBook.dart';
import 'package:expencemanager/data/models/Category.dart';
import 'package:expencemanager/data/models/Entry.dart';
import 'package:expencemanager/data/models/Tag.dart';
import 'package:expencemanager/data/models/Wallet.dart';
import 'package:moor_flutter/moor_flutter.dart';

part 'LocalDatabase.g.dart';

@UseMoor(
    tables: [AccountBook, Category, Entry, Tag, Wallet],
    daos: [AccountBookDao, CategoryDao, EntryDao, WalletDao])
class LocalDatabase extends _$LocalDatabase {
  LocalDatabase()
      : super(FlutterQueryExecutor.inDatabaseFolder(
            path: 'db.sqlite', logStatements: true));

  @override
  int get schemaVersion => 1;
}