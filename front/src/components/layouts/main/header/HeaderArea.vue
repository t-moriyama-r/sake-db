<template>
  <header class="sticky">
    <main class="flex">
      <section class="block sm:hidden">
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
              ></path>
            </svg>
          </MenuButton>
        </Menu>
      </section>
      <section class="flex flex-1">
        <router-link to="/"><p class="block">酒データベース(α)</p></router-link>
        <router-link :to="{ name: 'LiquorEdit' }">投稿する</router-link>
      </section>
      <section v-if="getRoles().includes(Roles.Admin)">
        <router-link :to="{ name: 'Admin' }">管理者ページ</router-link>
      </section>
      <section>
        <AccountInfo />
      </section>
      <section>
        <MainMenu />
      </section>
    </main>
  </header>
</template>

<script setup lang="ts">
import { Menu, MenuButton } from '@headlessui/vue';

import AccountInfo from '@/components/layouts/main/header/AccountInfo.vue';
import MainMenu from '@/components/layouts/main/header/menu/MainMenu.vue';
import { Roles } from '@/graphQL/Auth/types';
import { useUserStore } from '@/stores/userStore/userStore';

const { getRoles } = useUserStore();

const emit = defineEmits<{
  (_: 'openSideBar'): void;
}>();
const handleOpenSidebarClick = () => {
  emit('openSideBar');
};
</script>

<style scoped>
header {
  width: 100%;
  background-color: black;
  color: white;
}
</style>
