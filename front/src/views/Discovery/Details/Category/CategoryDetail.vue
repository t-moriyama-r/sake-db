<template>
  カテゴリー詳細
  <div>
    <p>{{ category.name }}</p>
  </div>
  <img v-if="category.imageUrl" :src="category.imageUrl" alt="画像" />
  <div>
    <span v-for="(line, index) in descriptionLines" :key="index">
      {{ line }}<br v-if="index !== descriptionLines.length - 1" />
    </span>
  </div>
  <CategoryChildren
    v-if="category.children"
    :category-list="category.children"
  />
  <router-link :to="{ name: 'CategoryEdit', params: { id: category.id } }">
    <CommonButton>編集する</CommonButton></router-link
  >
</template>

<script setup lang="ts">
import { computed } from 'vue';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import type { Category } from '@/graphQL/Category/categories';
import CategoryChildren from '@/views/Discovery/Details/Category/CategoryChildren.vue';

interface Props {
  category: Category;
}
const { category } = defineProps<Props>();

const descriptionLines = computed<string[]>(() => {
  return category.description.split('\n');
});
</script>

<style scoped></style>
