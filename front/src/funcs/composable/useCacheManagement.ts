/**
 * Apollo Clientのキャッシュ管理ユーティリティ
 */
import client from '@/apolloClient';
import type { Categories, Category } from '@/graphQL/Category/categories';
import { GET_QUERY } from '@/graphQL/Category/categories';

/**
 * 特定のカテゴリとその親カテゴリに関連するお酒のキャッシュをクリア
 * お酒の作成、更新、削除後に呼び出して、サーバーから最新データを取得できるようにする
 * @param categoryId - 変更されたお酒のカテゴリID
 */
export function clearLiquorCache(categoryId?: number) {
  if (categoryId !== undefined) {
    // キャッシュからカテゴリツリーを読み込む
    try {
      const cachedData = client.cache.readQuery<Categories>({
        query: GET_QUERY,
      });

      if (cachedData?.categories) {
        // すべての親カテゴリID（対象カテゴリ自身を含む）を取得
        const categoryIds = findCategoryPathIds(
          cachedData.categories,
          categoryId,
        );

        // パス内の各カテゴリのキャッシュをクリア
        categoryIds.forEach((catId) => {
          client.cache.evict({
            id: 'ROOT_QUERY',
            fieldName: 'listFromCategory',
            args: { categoryId: catId },
          });
        });
      } else {
        // カテゴリツリーがキャッシュにない場合、フォールバックとしてすべてクリア
        client.cache.evict({
          id: 'ROOT_QUERY',
          fieldName: 'listFromCategory',
        });
      }
    } catch (error) {
      // キャッシュの読み込みに失敗した場合、フォールバックとしてすべてクリア
      console.warn('Failed to read category tree from cache:', error);
      client.cache.evict({
        id: 'ROOT_QUERY',
        fieldName: 'listFromCategory',
      });
    }
  } else {
    // カテゴリIDが未指定の場合、すべてクリア（後方互換性のため）
    client.cache.evict({
      id: 'ROOT_QUERY',
      fieldName: 'listFromCategory',
    });
  }

  // ランダムおすすめリストは常にクリア
  client.cache.evict({
    id: 'ROOT_QUERY',
    fieldName: 'randomRecommendList',
  });

  // ガベージコレクションを実行して孤立したキャッシュエントリをクリーンアップ
  client.cache.gc();
}

/**
 * 特定のお酒の詳細キャッシュをクリア
 * @param id - キャッシュからクリアするお酒のID
 */
export function clearLiquorDetailCache(id: string) {
  client.cache.evict({
    id: 'ROOT_QUERY',
    fieldName: 'liquor',
    args: { id },
  });

  client.cache.gc();
}

/**
 * タグ検索のキャッシュをクリア
 * タグの登録・削除後に呼び出して、検索結果に最新の情報が反映されるようにする
 * @param tag - クリアするタグのテキスト（オプション、指定しない場合は全てのタグ検索キャッシュをクリア）
 */
export function clearTagSearchCache(tag?: string) {
  if (tag) {
    // 特定のタグ検索キャッシュをクリア
    client.cache.evict({
      id: 'ROOT_QUERY',
      fieldName: 'searchLiquorsByTag',
      args: { tag },
    });
  } else {
    // すべてのタグ検索キャッシュをクリア
    client.cache.evict({
      id: 'ROOT_QUERY',
      fieldName: 'searchLiquorsByTag',
    });
  }

  client.cache.gc();
}

/**
 * 指定されたカテゴリIDから親カテゴリまでのIDパスを取得
 * @param categories - カテゴリツリー
 * @param targetId - 検索対象のカテゴリID
 * @param path - 蓄積されたパス（再帰処理で内部的に使用）
 * @returns ルートから対象カテゴリまでのカテゴリID配列（対象自身を含む）
 */
function findCategoryPathIds(
  categories: Category[],
  targetId: number,
  path: number[] = [],
): number[] {
  for (const category of categories) {
    if (category.id === targetId) {
      return [...path, category.id];
    }
    if (category.children) {
      const result = findCategoryPathIds(category.children, targetId, [
        ...path,
        category.id,
      ]);
      if (result.length > 0) {
        return result;
      }
    }
  }
  return [];
}
