<template>
  <div>
    <p class="flex items-center gap-2">
      <BackButton to="Index" />
      <span>タグで検索: {{ tag }}</span>
    </p>
  </div>
  <FromTag :key="tag" v-if="liquors" :tag="tag" :liquors="liquors" />
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

// データフェッチ
const fetchData = async (searchTag: string): Promise<void> => {
  try {
    const { searchLiquorsByTag: response } = await fetch({
      tag: searchTag,
    });
    liquors.value = response;
  } catch (error) {
    console.error('タグ検索中にエラーが発生しました:', error);
    liquors.value = [];
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

<style scoped></style>
