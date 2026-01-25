<template>
  <div class="flex">
    <BackButton
      to="CategoryNarrowDown"
      :option="{ params: { id: category.id } }"
    />
    <p class="title">{{ category.name }}</p>
  </div>
  <img
    v-if="category.imageUrl"
    :src="category.imageUrl"
    class="image"
    alt="画像"
  />
  <div>
    <span v-for="(line, index) in descriptionLines" :key="index">
      {{ line }}<br v-if="index !== descriptionLines.length - 1" />
    </span>
  </div>
  <AffiliateContainer :name="category.name" />
  <div class="mt-2 mb-2">
    <CategoryChildren
      v-if="category.children"
      :category-list="sortedChildren"
    />
  </div>
  <div class="p-2">
    <router-link :to="{ name: 'CategoryEdit', params: { id: category.id } }">
      <CommonButton>編集する</CommonButton></router-link
    >
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import AffiliateContainer from '@/components/blocks/common/amazon/AffiliateContainer.vue';
import BackButton from '@/components/parts/common/BackButton.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import { sortCategoriesWithOtherLast } from '@/funcs/util/sortCategories';
import type { Category } from '@/graphQL/Category/categories';
import CategoryChildren from '@/views/Discovery/Details/Category/CategoryChildren.vue';

interface Props {
  category: Category;
}
const { category } = defineProps<Props>();

const descriptionLines = computed<string[]>(() => {
  if (category.description.length === 0) return ['説明なし'];
  return category.description.split('\n');
});

const sortedChildren = computed(() => {
  if (!category.children) return [];
  return sortCategoriesWithOtherLast(category.children);
});
</script>

<style scoped>
p.title {
  font-size: 150%;
  font-weight: bold;
}
img.image {
  max-height: 300px;
  max-width: 500px;
}
</style>
