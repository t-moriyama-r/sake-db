<template>
  <main class="flex">
    <section class="block md:hidden">
      <Menu @click="handleOpenSidebarClick">
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
    </section>
    <section class="flex flex-1">
      <KeywordSearch />
    </section>
    <section>
      <MainMenu />
    </section>
  </main>
</template>

<script setup lang="ts">
import { Menu, MenuButton } from '@headlessui/vue';

import KeywordSearch from '@/components/blocks/keywordSearch/KeywordSearch.vue';
import MainMenu from '@/components/layouts/main/header/menu/MainMenu.vue';

const emit = defineEmits<{
  (_: 'openSideBar'): void;
}>();
const handleOpenSidebarClick = () => {
  emit('openSideBar');
};
</script>

<style scoped>
main {
  width: 100%;
}
</style>
