<template>
  <div>
    <p class="flex items-center gap-2">
      <BackButton to="Index" />
      <span class="text-lg font-semibold">検索結果: "{{ keyword }}"</span>
    </p>
    <div v-if="loading" class="text-center py-8">検索中...</div>
    <div v-else-if="error" class="text-red-500 py-4">
      検索に失敗しました: {{ error.message }}
    </div>
    <div v-else-if="liquors && liquors.length > 0" id="search-results-area">
      <p class="text-sm text-gray-600 mb-4">{{ liquors.length }}件の結果</p>
      <CardContainer
        columns="repeat(auto-fill, minmax(100px, 1fr))"
        gap="0.5em"
        min="200px"
      >
        <LiquorCard
          v-for="liquor in liquors"
          :liquor="liquor"
          :key="liquor.id"
        />
      </CardContainer>
    </div>
    <div v-else class="text-center py-8 text-gray-500">
      <p>「{{ keyword }}」に一致するお酒が見つかりませんでした</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import LiquorCard from '@/components/blocks/cards/LiquorCard.vue';
import BackButton from '@/components/parts/common/BackButton.vue';
import CardContainer from '@/components/parts/common/CardContainer.vue';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import type { Liquor } from '@/graphQL/Index/random';
import {
  SEARCH_LIQUORS,
  type SearchLiquorsResponse,
} from '@/graphQL/Liquor/search';

const route = useRoute();
const { fetch } = useQuery<SearchLiquorsResponse>(SEARCH_LIQUORS);

const keyword = ref<string>('');
const liquors = ref<Liquor[] | null>(null);
const loading = ref<boolean>(false);
const error = ref<Error | null>(null);

// データフェッチ
const fetchData = async (searchKeyword: string): Promise<void> => {
  if (!searchKeyword || searchKeyword.trim() === '') {
    liquors.value = [];
    return;
  }

  loading.value = true;
  error.value = null;

  try {
    const response = await fetch({
      keyword: searchKeyword,
      limit: 100, // 最大100件まで表示
    });
    liquors.value = response.searchLiquors;
  } catch (e) {
    error.value = e as Error;
    liquors.value = [];
  } finally {
    loading.value = false;
  }
};

// `watch` を使ってルートパラメータの変更を監視
watch(
  () => route.query.keyword,
  (newKeyword) => {
    const searchKeyword = newKeyword as string;
    keyword.value = searchKeyword || '';
    if (searchKeyword) {
      fetchData(searchKeyword);
    }
  },
  { immediate: true },
);
</script>

<style scoped>
div#search-results-area {
  margin: auto;
  padding: 2em;
}
</style>
