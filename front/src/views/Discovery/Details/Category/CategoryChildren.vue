<!--子カテゴリのリスト-->
<template>
  <p class="title">子カテゴリ一覧</p>
  <ul v-for="category in categoryList" :key="category.id">
    <router-link :to="{ name: 'CategoryDetail', params: { id: category.id } }"
      ><li>{{ category.name }}</li></router-link
    >
  </ul>
</template>
<script setup lang="ts">
import type { Category } from '@/graphQL/Category/categories';

interface Props {
  categoryList: Category[];
}

const { categoryList } = defineProps<Props>();
</script>

<style scoped>
p.title {
  font-size: 125%;
  font-weight: bold;
}
</style>
