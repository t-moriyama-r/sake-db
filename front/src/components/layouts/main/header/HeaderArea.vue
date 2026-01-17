<template>
  <main class="flex items-center gap-1 sm:gap-2">
    <section class="block md:hidden">
      <CategoriesForSmartMenu />
    </section>
    <section class="flex flex-1 min-w-0">
      <KeywordSearch />
    </section>
    <section class="flex items-center shrink-0">
      <router-link :to="{ name: 'LiquorEdit' }"
        ><CommonButton :size="'small'" class="px-2 sm:px-6 md:px-3 lg:px-6"
          ><FontAwesomeIcon icon="fa-solid fa-plus" />投稿</CommonButton
        ></router-link
      >
    </section>
    <section class="block sm:hidden shrink-0 overflow-hidden">
      <LoginAndSignUpButton />
    </section>
    <section class="hidden sm:block shrink-0 overflow-hidden">
      <AccountMenu />
    </section>
  </main>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import KeywordSearch from '@/components/blocks/keywordSearch/KeywordSearch.vue';
import AccountMenu from '@/components/layouts/main/header/accountMenu/AccountMenu.vue';
import CategoriesForSmartMenu from '@/components/layouts/main/header/categoriesForSmart/CategoriesForSmartMenuButton.vue';
import LoginAndSignUpButton from '@/components/layouts/main/header/loginAndSignUp/LoginAndSignUpButton.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
</script>

<style scoped>
main {
  width: 100%;
}
</style>
