<template>
  <div id="layout" class="flex">
    <div id="form-container" class="flex-1 w-full md:min-w-[380px]">
      <div class="flex">
        <BackButton />
        <h1 v-if="route.params.id">カテゴリー編集</h1>
        <h1 v-else>カテゴリー作成</h1>
      </div>
      <div class="md:hidden">TODO: カテゴリーログボタン</div>
      <CategoryForm
        :initial-data="initialValues"
        :version-no="historyData?.now.versionNo ?? null"
        :readonly="historyData?.now.readonly ?? false"
      />
    </div>
    <div id="versions" v-if="historyData?.histories" class="hidden md:block">
      <CategoryLogs
        :logs="[historyData.now, ...historyData.histories]"
        @selectLog="reflectLog"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';

import BackButton from '@/components/parts/common/BackButton.vue';
import type { Category } from '@/graphQL/Category/categories';
import type { CategoryHistoryData } from '@/graphQL/Category/categoryLog';
import CategoryForm from '@/views/Edit/CategoryEdit/form/CategoryForm.vue';
import CategoryLogs from '@/views/Edit/CategoryEdit/form/CategoryLogs.vue';

// propsから受け取る初期値
const { historyData, initialParentId } = defineProps<{
  historyData: CategoryHistoryData | null;
  initialParentId?: number | null;
}>();

const route = useRoute(); // 現在のルートを取得

// 初期値を定義
const initialValues = ref<Category | null>(
  historyData?.now ??
    (initialParentId ? ({ parent: initialParentId } as Category) : null),
);

const reflectLog = (log: Category) => {
  if (historyData?.now.readonly) {
    //↓es-lintの除外設定してるつもりなんですけど、時間ないので一旦保留 es-lintとtypescript-eslintが意図せず共存してる？
    // eslint-disable-next-line no-unused-vars,@typescript-eslint/no-unused-vars
    const { parent: _, name: __, ...rest } = log;
    initialValues.value = {
      parent: historyData.now.parent,
      name: historyData.now.name,
      ...rest,
    };
  } else {
    initialValues.value = { ...log }; //過去のデータをそのまま初期値として代入する
  }
};
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
