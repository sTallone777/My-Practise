import 'package:expencemanager/data/datasources/LocalDatabase.dart';
import 'package:expencemanager/data/models/Category.dart';
import 'package:expencemanager/data/models/CategoryWithTag.dart';
import 'package:expencemanager/data/models/Entry.dart';
import 'package:expencemanager/data/models/Tag.dart';
import 'package:expencemanager/data/models/Wallet.dart';
import 'package:moor_flutter/moor_flutter.dart';

part 'CategoryDao.g.dart';

@UseDao(tables: [Entry, Tag, Category, Wallet])
class CategoryDao extends DatabaseAccessor<LocalDatabase>
    with _$CategoryDaoMixin {
  LocalDatabase _database;

  CategoryDao(LocalDatabase attachedDatabase) : super(attachedDatabase) {
    this._database = attachedDatabase;
  }

  Future<List<category>> getCategories(bool isIncome, int bookId) {
    return (select(_database.category)
      ..where((tbl) {
        return tbl.isIncome.equals(isIncome) & tbl.bookId.equals(bookId);
      }))
        .get();
  }
  Stream<List<CategoryWithTag>> getCategoriesWithTags(
      bool isIncome, int bookId) {
    final query = select(_database.tag).join([
      leftOuterJoin(_database.category,
          _database.category.id.equalsExp(_database.tag.categoryId))
    ])
      ..where(_database.tag.bookId.equals(bookId) &
      _database.category.isIncome.equals(isIncome));
    return query.get().asStream().map((event) {
      return event.map((e) {
        return CategoryWithTag(
            e.readTable(this.category as TableInfo<$CategoryTable, category>),
            e.readTable(tag as TableInfo<$TagTable, tag>));
      }).toList();
    });
  }

  Future<List<CategoryWithTag>> getAllCategoriesWithTags(int bookId) {
    final query = select(_database.tag).join([
      leftOuterJoin(_database.category,
          _database.category.id.equalsExp(_database.tag.categoryId))
    ])
      ..where(_database.tag.bookId.equals(bookId))
      ..orderBy([OrderingTerm.asc(_database.tag.categoryId)]);
    return query.map((event) {
      return CategoryWithTag(
          event.readTable(this.category), event.readTable(this.tag));
    }).get();
  }

  Stream<List<tag>> getTagsForACategory(int categoryId) {
    return (select(_database.tag)
      ..where((tbl) => tbl.categoryId.equals(categoryId)))
        .get()
        .asStream();
  }

  Future<int> insertTag(tag tag) {
    return into(_database.tag).insert(tag);
  }

  Stream<void> insertTags(List<tag> tag) {
    return batch((batch) {
      batch.insertAll(_database.tag, tag);
    }).asStream();
  }

  Future<int> insertCategory(category category) {
    return into(_database.category).insert(category);
  }

  Stream<List<int>> insertCategories(List<category> categories) {
    return batch((batch) {
      batch.insertAll(_database.category, categories);
    }).asStream();
  }

  Future<int> deleteCategory(category category) {
    return delete(_database.category).delete(category);
  }

  Future<int> deleteTag(tag tag) {
    return delete(_database.tag).delete(tag);
  }

  Future<category> findCategory(String name, bool isIncome, int bookId) {
    return (select(_database.category)
      ..limit(1)
      ..where((tbl) =>
      tbl.name.equals(name) &
      tbl.isIncome.equals(isIncome) &
      tbl.bookId.equals(bookId)))
        .getSingle();
  }

  Stream<category> getCategory(int categoryId) {
    return (select(_database.category)
      ..where((tbl) => tbl.id.equals(categoryId)))
        .getSingle()
        .asStream();
  }

  Stream<tag> getATag(int tagId) {
    return (select(_database.tag)..where((tbl) => tbl.id.equals(tagId)))
        .getSingle()
        .asStream();
  }
}