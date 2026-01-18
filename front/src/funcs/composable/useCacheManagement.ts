/**
 * Apollo Client cache management utilities
 */
import client from '@/apolloClient';

/**
 * Clear liquor-related cache entries
 * This should be called after creating, updating, or deleting liquor entries
 * to ensure fresh data is fetched from the server
 */
export function clearLiquorCache() {
  // Evict all liquor list queries from cache
  client.cache.evict({
    id: 'ROOT_QUERY',
    fieldName: 'listFromCategory',
  });

  client.cache.evict({
    id: 'ROOT_QUERY',
    fieldName: 'randomRecommendList',
  });

  // Run garbage collection to clean up orphaned cache entries
  client.cache.gc();
}

/**
 * Clear specific liquor detail from cache
 * @param id - The liquor ID to clear from cache
 */
export function clearLiquorDetailCache(id: string) {
  client.cache.evict({
    id: 'ROOT_QUERY',
    fieldName: 'liquor',
    args: { id },
  });

  client.cache.gc();
}
