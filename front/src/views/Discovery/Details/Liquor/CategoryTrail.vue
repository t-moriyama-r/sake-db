<!--カテゴリ一覧のパンくずリスト-->
<template>
  <nav class="flex items-center text-sm text-gray-600" aria-label="Breadcrumb">
    <div
      v-for="(category, index) in categoryTrails"
      :key="category.id"
      class="inline-flex items-center"
    >
      <router-link
        :to="{ name: 'CategoryNarrowDown', params: { id: category.id } }"
        class="hover:text-blue-600 transition-colors"
      >
        {{ category.name }}
      </router-link>
      <span
        v-if="index !== categoryTrails.length - 1"
        class="mx-2 text-gray-400"
      >
        &gt;
      </span>
    </div>
  </nav>
</template>

<script setup lang="ts">
interface Props {
  categoryTrails: {
    id: number;
    name: string;
  }[];
}
const { categoryTrails } = defineProps<Props>();
</script>

<style scoped></style>
