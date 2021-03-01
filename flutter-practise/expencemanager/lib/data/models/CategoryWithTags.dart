import 'package:expencemanager/data/datasources/LocalDatabase.dart';

class CategoryWithTags {
  category mCategory;
  List<tag> tags;

  CategoryWithTags(this.mCategory, this.tags);

  List<int> checkedTagIndexes = [];
}