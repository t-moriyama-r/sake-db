<template>
  <span class="arrow mx-1" @click="onBack">
    <FontAwesomeIcon icon="fa-solid fa-arrow-left" />
  </span>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { type RouteLocationAsRelativeGeneric, useRouter } from 'vue-router';

type RouteParams = Omit<RouteLocationAsRelativeGeneric, 'name'>;

interface Props {
  to?: string;
  option?: RouteParams;
}
const props = defineProps<Props>();
const router = useRouter();

const onBack = () => {
  if (props.to) {
    // name + 残りの params, query, hash を安全に再構成
    router.push({
      name: props.to,
      ...props.option,
    });
    return;
  }
  if (window.history.length > 1) {
    // 履歴があれば戻る
    router.back();
    return;
  }
  // 履歴がなければ Index
  router.push({ name: 'Index' });
};
</script>

<style scoped>
span.arrow {
  cursor: pointer;
  display: inline-flex; /* アイコンをフレックスコンテナに */
  align-items: center; /* 垂直方向中央揃え */
}
</style>
