<template>
  <div id="layout" class="flex">
    <div id="form-container" class="flex-1 w-full md:min-w-[380px]">
      <div class="flex">
        <BackButton />
        <h1 v-if="route.params.id">お酒ページ編集</h1>
        <h1 v-else>お酒ページ作成</h1>
      </div>

      <div class="md:hidden">
        <LiquorLogsButton />
      </div>
      <LiquorForm
        :initial-data="initialValues"
        :initial-category-id="initialCategoryId"
        :version-no="historyData?.now.versionNo ?? null"
      />
    </div>
    <div id="versions" v-if="historyData?.histories" class="hidden md:block">
      <LiquorLogs
        :logs="[historyData.now, ...historyData.histories]"
        @selectLog="reflectLog"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import BackButton from '@/components/parts/common/BackButton.vue';
import { assertNonNullable } from '@/funcs/util/core/assertNonNullable';
import { isEmpty } from '@/funcs/util/isEmpty';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import type { LiquorHistoryData } from '@/graphQL/Liquor/liquorLog';
import LiquorForm from '@/views/Edit/LiquorEdit/form/LiquorForm.vue';
import LiquorLogs from '@/views/Edit/LiquorEdit/form/LiquorLogs/LiquorLogs.vue';
import LiquorLogsButton from '@/views/Edit/LiquorEdit/form/LiquorLogs/LiquorLogsButton/LiquorLogsButton.vue';

const props = defineProps<{
  historyData: LiquorHistoryData | null;
}>();

const route = useRoute(); // 現在のルートを取得

// 初期値を定義
const initialValues = ref<Liquor | null>(props.historyData?.now ?? null);
const initialCategoryId = ref<number | null>(null); // 初期カテゴリを指定して新規投稿する場合

onMounted(async () => {
  const categoryId: string = route.params.categoryId as string;
  initialCategoryId.value = !isEmpty(categoryId) ? Number(categoryId) : null;
});

//バージョン履歴をコピーする時に呼ばれる関数
const reflectLog = (log: Liquor) => {
  assertNonNullable(props.historyData);
  initialValues.value = { ...log, id: props.historyData.now.id }; //過去のデータをそのまま初期値として代入する(IDだけは現在の値で上書き。ここが呼び出された時点で初期値は存在しないとおかしいのでアサーション)
};

watch(
  () => props.historyData,
  (liquor) => {
    initialValues.value = liquor?.now ?? null;
  },
);
</script>

<style scoped>
div#layout {
  margin: 0 1rem;
}

h1 {
  font-size: 150%;
  font-weight: bold;
}

div#form-container {
  margin: 0 1rem;
}

div#versions {
  margin: 0 1rem;
}
</style>
