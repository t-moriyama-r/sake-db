import type { Category } from '@/graphQL/Category/categories';

/**
 * カテゴリをソートする関数。「その他」を常に最後に配置する。
 * @param categories - ソート対象のカテゴリ配列
 * @returns ソート済みのカテゴリ配列
 */
export function sortCategoriesWithOtherLast(
  categories: Category[],
): Category[] {
  const sortedCategories = [...categories];
  return sortedCategories.sort((a, b) => {
    const aIsOther = a.name === 'その他';
    const bIsOther = b.name === 'その他';

    // 「その他」は常に最後
    if (aIsOther) return 1;
    if (bIsOther) return -1;

    // それ以外は元の順序を維持（IDで並べる）
    return a.id - b.id;
  });
}
