<template>
  <div>
    <p class="flex">
      <BackButton to="Index" />
      <span>カテゴリで検索:</span
      ><span v-if="Number(route.params.id) < 10"> {{ categoryName }}</span>
      <router-link
        v-else
        :to="{ name: 'CategoryDetail', params: { id: route.params.id } }"
      >
        {{ categoryName }}</router-link
      >
    </p>
    <div class="max-sm:text-sm">
      {{ truncatedDescription }}
    </div>
  </div>
  <FromCategory
    :key="route.params.id as string"
    v-if="liquors && categoryId"
    :category-id="categoryId"
    :liquors="liquors"
    :has-more="hasMore"
    :is-loading="isLoading"
    @load-more="loadMore"
  />
</template>
<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import BackButton from '@/components/parts/common/BackButton.vue';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import truncateString from '@/funcs/util/transform/truncateString';
import type { Liquor } from '@/graphQL/Index/random';
import {
  LIQUOR_LIST_FROM_CATEGORY,
  type ListFromCategoryResponse,
} from '@/graphQL/Liquor/liquor';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import FromCategory from '@/views/Discovery/NarrowDowns/FromCategory.vue';

const route = useRoute(); // 現在のルートを取得
const sidebarStore = useSelectedCategoryStore();
const { fetch } = useQuery<ListFromCategoryResponse>(LIQUOR_LIST_FROM_CATEGORY);

const categoryId = ref<string | null>(null);
const liquors = ref<Liquor[]>([]);
const categoryName = ref<string>('');
const categoryDescription = ref<string>('');
const isLoading = ref<boolean>(false);
const hasMore = ref<boolean>(true);
const offset = ref<number>(0);
const limit = 20; // 1回あたりのフェッチ件数

const truncatedDescription = computed(() =>
  truncateString({ str: categoryDescription.value ?? '', maxLength: 100 }),
);

// データフェッチ（初回）
const fetchData = async (id: number): Promise<void> => {
  sidebarStore.updateContent(id);
  isLoading.value = true;
  offset.value = 0;
  hasMore.value = true;

  try {
    const { listFromCategory: response } = await fetch({
      id,
      limit,
      offset: 0,
    });
    liquors.value = response.liquors;
    categoryName.value = response.categoryName;
    categoryDescription.value = response.categoryDescription;
    hasMore.value = response.liquors.length >= limit;
    offset.value = response.liquors.length;
  } finally {
    isLoading.value = false;
  }
};

// 追加データフェッチ
const loadMore = async (): Promise<void> => {
  if (!categoryId.value || isLoading.value || !hasMore.value) {
    return;
  }

  isLoading.value = true;

  try {
    const { listFromCategory: response } = await fetch({
      id: Number(categoryId.value),
      limit,
      offset: offset.value,
    });

    if (response.liquors.length > 0) {
      liquors.value = [...liquors.value, ...response.liquors];
      offset.value += response.liquors.length;
      hasMore.value = response.liquors.length >= limit;
    } else {
      hasMore.value = false;
    }
  } finally {
    isLoading.value = false;
  }
};

// `watch` を使ってルートパラメータの変更を監視
watch(
  () => route.params.id,
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      categoryId.value = null;
      return;
    }
    categoryId.value = id;
    fetchData(Number(id));
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>

<style scoped></style>
