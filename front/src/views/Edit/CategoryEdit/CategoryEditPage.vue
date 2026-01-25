<template>
  <CategoryEdit
    v-if="!isLoading"
    :history-data="category"
    :initial-parent-id="initialParentId"
  />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import { isEmpty } from '@/funcs/util/isEmpty';
import {
  type CategoryHistoryData,
  GET_LOGS_FOR_ROLLBACK,
  type HistoryResponse,
} from '@/graphQL/Category/categoryLog';
import CategoryEdit from '@/views/Edit/CategoryEdit/form/CategoryEdit.vue';

const isLoading = ref<boolean>(true);
const category = ref<CategoryHistoryData | null>(null); //フィールドにあるカテゴリ情報
const initialParentId = ref<number | null>(null); // 初期親カテゴリを指定して新規投稿する場合

const route = useRoute(); // 現在のルートを取得
const { fetch } = useQuery<HistoryResponse>(GET_LOGS_FOR_ROLLBACK);

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  const id: string = route.params.id as string; // ルートパラメータからidを取得
  const parentCategoryId: string = route.params.parentCategoryId as string;

  // 親カテゴリIDが指定されている場合は初期値として設定
  initialParentId.value = !isEmpty(parentCategoryId)
    ? Number(parentCategoryId)
    : null;

  if (isEmpty(id)) {
    isLoading.value = false;
    return;
  }
  await fetch(
    {
      id: Number(id),
    },
    {
      fetchPolicy: 'no-cache',
    },
  )
    .then((response) => {
      category.value = response.histories;
    })
    .finally(() => {
      isLoading.value = false;
    });
});
</script>

<style scoped></style>
