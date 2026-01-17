<template>
  <transition name="fade">
    <div
      v-show="isVisible"
      class="fixed inset-0 bg-black bg-opacity-50 h-screen"
      :style="{ zIndex: zIndex ?? Z_INDEX_DEFAULT }"
      @click="closeSidebar"
    />
  </transition>
</template>

<script setup lang="ts">
type Props = {
  isVisible: boolean; // transitionのv-if用
  zIndex?: number;
};
defineProps<Props>();

const Z_INDEX_DEFAULT = 40;

const emit = defineEmits<{
  (_: 'click'): void;
}>();
const closeSidebar = () => {
  emit('click');
};
</script>

<style scoped></style>
