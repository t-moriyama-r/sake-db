<template>
  <Menu @click="expandSidebar">
    <MenuButton>
      <svg
        class="w-6 h-6"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h16"
        />
      </svg>
    </MenuButton>
  </Menu>
  <CommonOverlay :isVisible="isExpanded" @click="collapseSidebar" />
</template>

<script setup lang="ts">
import { Menu, MenuButton } from '@headlessui/vue';
import { ref } from 'vue';

import CommonOverlay from '@/components/blocks/common/CommonOverlay.vue';

const isExpanded = ref<boolean>(false);
const expandSidebar = () => {
  isExpanded.value = true;
};
const collapseSidebar = () => {
  isExpanded.value = false;
};
</script>

<style scoped></style>
