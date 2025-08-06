<!--トップページ-->
<template>
  <IndexMain />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';

import { useSelectedCategoryStore } from '@/stores/sidebar';
import IndexMain from '@/views/Index/IndexMain.vue';

const sidebarStore = useSelectedCategoryStore();
onMounted(() => {
  sidebarStore.updateContent(null);
});
</script>

<style scoped></style>
