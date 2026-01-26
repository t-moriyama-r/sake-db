<template>
  <div>
    <p class="flex items-center gap-2">
      <BackButton to="Index" />
      <span>タグで検索: {{ tag }}</span>
    </p>
  </div>
  <FromTag
    :key="tag"
    v-if="liquors !== null"
    :tag="tag"
    :liquors="liquors"
    :has-more="hasMore"
    :is-loading="isLoading"
    @load-more="loadMore"
  />
  <div v-else class="loading-initial">
    <p>読み込み中...</p>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import BackButton from '@/components/parts/common/BackButton.vue';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import type { Liquor } from '@/graphQL/Index/random';
import {
  SearchLiquorsByTag,
  type SearchLiquorsByTagResponse,
} from '@/graphQL/Liquor/tags';
import FromTag from '@/views/Discovery/TagSearch/FromTag.vue';

const route = useRoute();
const { fetch } = useQuery<SearchLiquorsByTagResponse>(SearchLiquorsByTag);

const tag = ref<string>('');
const liquors = ref<Liquor[] | null>(null);
const isLoading = ref<boolean>(false);
const hasMore = ref<boolean>(true);
const offset = ref<number>(0);
const limit = 20; // 1回あたりのフェッチ件数

// データフェッチ（初回）
const fetchData = async (searchTag: string): Promise<void> => {
  isLoading.value = true;
  offset.value = 0;
  hasMore.value = true;

  try {
    const { searchLiquorsByTag: response } = await fetch({
      tag: searchTag,
      limit,
      offset: 0,
    });
    liquors.value = response;
    hasMore.value = response.length >= limit;
    offset.value = response.length;
  } catch (error) {
    console.error('タグ検索中にエラーが発生しました:', error);
    liquors.value = [];
    hasMore.value = false;
  } finally {
    isLoading.value = false;
  }
};

// 追加データフェッチ
const loadMore = async (): Promise<void> => {
  if (!tag.value || isLoading.value || !hasMore.value || !liquors.value) {
    return;
  }

  isLoading.value = true;

  try {
    const { searchLiquorsByTag: response } = await fetch({
      tag: tag.value,
      limit,
      offset: offset.value,
    });

    if (response.length > 0) {
      liquors.value = [...liquors.value, ...response];
      offset.value += response.length;
      hasMore.value = response.length >= limit;
    } else {
      hasMore.value = false;
    }
  } catch (error) {
    console.error('追加データのフェッチ中にエラーが発生しました:', error);
    hasMore.value = false;
  } finally {
    isLoading.value = false;
  }
};

// `watch` を使ってルートパラメータの変更を監視
watch(
  () => route.params.tag,
  (to) => {
    const tagParam = to as string;
    if (!tagParam) {
      tag.value = '';
      liquors.value = null;
      return;
    }
    tag.value = tagParam;
    fetchData(tagParam);
  },
  { immediate: true },
);
</script>

<style scoped>
.loading-initial {
  text-align: center;
  padding: 2em;
  color: #666;
}
</style>
