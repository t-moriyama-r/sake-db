<template>
  <Menu @click="openSideBar">
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
  <CategoriesForSmart :isOpen="isOpenSideBar" @closeSideBar="closeSideBar" />
</template>

<script setup lang="ts">
import { Menu, MenuButton } from '@headlessui/vue';
import { ref } from 'vue';

import CategoriesForSmart from '@/components/layouts/main/header/categoriesForSmart/CategoriesForSmart.vue';
const isOpenSideBar = ref<boolean>(false);

const openSideBar = () => {
  isOpenSideBar.value = true;
};
const closeSideBar = () => {
  isOpenSideBar.value = false;
};
</script>

<style scoped></style>
