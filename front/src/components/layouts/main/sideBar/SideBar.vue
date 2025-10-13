<!--サイドバー-->
<template>
  <div class="container flex flex-col bg-gray-100">
    <aside class="flex-1">
      <div class="flex">
        <div class="flex-1">
          <router-link
            class="inline-block"
            :to="{ name: 'Index' }"
            @click="emitCloseSidebar"
            ><FontAwesomeIcon
              icon="fa-solid fa-wine-bottle"
            />酒データベース(α)</router-link
          >
        </div>
        <div class="block md:hidden" @click="emitCloseSidebar">X</div>
      </div>
      <section class="menu">
        <div :class="{ 'font-bold': isNonSelectedCategory }">
          <router-link class="inline-block" :to="{ name: 'Index' }"
            >すべてのお酒</router-link
          >
        </div>
        <CategoryParent
          v-for="category in categoryList"
          :key="category.id"
          :category="category"
          :display-ids="filteredCategoryIdList"
        />
      </section>
    </aside>
    <aside class="new-post">
      <router-link :to="{ name: 'CategoryEdit' }" @click="emitCloseSidebar"
        >+新規カテゴリ追加</router-link
      >
    </aside>
  </div>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { computed, type ComputedRef, onMounted, ref, watch } from 'vue';

import CategoryParent from '@/components/layouts/main/sideBar/CategoryParent.vue';
import { getDisplayCategoryIds } from '@/components/layouts/main/sideBar/func/sideBarFunc';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import {
  type Categories,
  type Category,
  GET_QUERY,
} from '@/graphQL/Category/categories';
import { useSelectedCategoryStore } from '@/stores/sidebar';

const sidebarStore = useSelectedCategoryStore();

const { fetch } = useQuery<Categories>(GET_QUERY);

const categoryList = ref<Category[] | null>();

const emit = defineEmits<{
  (_: 'closeSideBar'): void;
}>();
const emitCloseSidebar = () => {
  emit('closeSideBar');
};

async function fetchData() {
  const { categories: response } = await fetch(null, {
    fetchPolicy: 'network-only',
  });
  categoryList.value = [...response];
  sidebarStore.setReloadFlgFalse();
}

// 読み込み時に情報をAPIから取得
onMounted(async () => {
  void fetchData();
});

watch(
  () => sidebarStore.isReloadFlg,
  () => {
    if (sidebarStore.isReloadFlg) {
      void fetchData();
    }
  },
);

const isNonSelectedCategory = computed(() => sidebarStore.content === null);

// sidebarStore.contentに基づいてカテゴリをフィルタリングする
const filteredCategoryIdList: ComputedRef<number[]> = computed(() => {
  if (!categoryList.value) return []; //そもそも存在していなければ処理終了
  if (!sidebarStore.content) return categoryList.value.map((c) => c.id); // contentがない場合は全ての大カテゴリを返す
  return getDisplayCategoryIds(categoryList.value, sidebarStore.content);
});
</script>

<style scoped>
div.container {
  width: 180px;
  height: 100%;
  border: 1px solid #ccc;
}

section.menu {
  padding-left: 5px;
}
</style>
