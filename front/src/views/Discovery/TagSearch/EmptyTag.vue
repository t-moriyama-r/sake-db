<template>
  <div class="px-8 mt-4">
    <p>「{{ tag }}」タグの日本酒が見つかりませんでした。</p>
  </div>
</template>

<script setup lang="ts">
interface Props {
  tag: string;
}

defineProps<Props>();
</script>

<style scoped></style>
