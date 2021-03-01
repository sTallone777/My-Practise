import 'package:expencemanager/data/datasources/LocalDatabase.dart';

class EntryWithCategoryAndWallet {

  EntryWithCategoryAndWallet(this.mEntry, this.mCategory, this.mTag, this.mWallet);

  entry mEntry;
  category mCategory;
  tag mTag;
  wallet mWallet;
}