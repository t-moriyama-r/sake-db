<template>
  <main class="flex">
    <section class="block md:hidden">
      <CategoriesForSmartMenu />
    </section>
    <section class="flex flex-1">
      <KeywordSearch />
    </section>
    <section class="flex items-center justify-center">
      <router-link :to="{ name: 'LiquorEdit' }"
        ><CommonButton :size="'small'" class="px-6 md:px-3 lg:px-6"
          ><FontAwesomeIcon icon="fa-solid fa-plus" />投稿</CommonButton
        ></router-link
      >
    </section>
    <section class="block sm:hidden">
      <AccountMenuForSmart />
    </section>
    <section class="hidden sm:block">
      <MainMenu />
    </section>
  </main>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import KeywordSearch from '@/components/blocks/keywordSearch/KeywordSearch.vue';
import MainMenu from '@/components/layouts/main/header/accountMenu/AccountMenu.vue';
import AccountMenuForSmart from '@/components/layouts/main/header/accountMenuForSmart/AccountMenuForSmartButton.vue';
import CategoriesForSmartMenu from '@/components/layouts/main/header/categoriesForSmart/CategoriesForSmartMenuButton.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
</script>

<style scoped>
main {
  width: 100%;
}
</style>
