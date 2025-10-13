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
    v-if="liquors"
    :liquors="liquors"
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

const liquors = ref<Liquor[] | null>(null);
const categoryName = ref<string>('');
const categoryDescription = ref<string>('');
const truncatedDescription = computed(() =>
  truncateString({ str: categoryDescription.value ?? '', maxLength: 100 }),
);

// データフェッチ
const fetchData = async (id: number): Promise<void> => {
  sidebarStore.updateContent(id);
  const { listFromCategory: response } = await fetch({
    id,
  });
  liquors.value = response.liquors;
  categoryName.value = response.categoryName;
  categoryDescription.value = response.categoryDescription;
};

// `watch` を使ってルートパラメータの変更を監視
watch(
  () => route.params.id, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      return;
    }
    fetchData(Number(to));
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>

<style scoped></style>
