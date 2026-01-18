<template>
  <transition name="slide">
    <div
      id="mobile-sidebar-wrapper"
      v-if="isOpen"
      class="fixed top-0 left-0 w-[180px] bg-white shadow-lg z-50 md:hidden"
    >
      <SideBar @closeSideBar="closeSidebar" />
    </div>
  </transition>
  <CommonOverlay :isVisible="isOpen" :zIndex="45" @click="closeSidebar" />
</template>

<script setup lang="ts">
import CommonOverlay from '@/components/blocks/common/CommonOverlay.vue';
import SideBar from '@/components/layouts/main/sideBar/SideBar.vue';

type Props = {
  isOpen: boolean; // transitionのv-if用
};
defineProps<Props>();

const emit = defineEmits<{
  (_: 'closeSideBar'): void;
}>();
const closeSidebar = () => {
  emit('closeSideBar');
};
</script>

<style scoped>
#mobile-sidebar-wrapper {
  height: 100vh;
  &.slide-enter-active,
  &.slide-leave-active {
    transition: transform 0.1s ease;
  }

  &.slide-enter-from,
  &.slide-leave-to {
    transform: translateX(-100%);
  }

  &.slide-enter-to,
  &.slide-leave-from {
    transform: translateX(0);
  }
}
</style>
