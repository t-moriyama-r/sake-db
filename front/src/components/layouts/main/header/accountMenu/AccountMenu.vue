<template>
  <div class="flex pr-2 gap-2">
    <div v-if="!userStore.isLogin">
      <LoginAndSignUpButton class="px-1 sm:px-6 md:px-3 lg:px-6">
        <FontAwesomeIcon icon="fa-solid fa-arrow-right-to-bracket" />ログイン /
        新規登録
      </LoginAndSignUpButton>
    </div>
    <div class="flex" v-else>
      <AccountInfo />
      <router-link :to="{ name: 'Index' }" @click="logout"
        ><CommonButton :size="'small'" class="sm:px-6 md:px-3 lg:px-6"
          >ログアウト</CommonButton
        ></router-link
      >
    </div>
    <div v-if="userStore.getRoles().includes(Roles.Admin)">
      <router-link :to="{ name: 'Admin' }">管理者ページ</router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import AccountInfo from '@/components/layouts/main/header/accountMenu/AccountInfo.vue';
import LoginAndSignUpButton from '@/components/layouts/main/header/loginAndSignUp/LoginAndSignUpButton.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import { useToast } from '@/funcs/composable/useToast';
import { Roles } from '@/graphQL/Auth/types';
import { useUserStore } from '@/stores/userStore/userStore';

const userStore = useUserStore();
const toast = useToast();

function logout() {
  userStore.logout();
  toast.showToast({
    message: 'ログアウトしました',
  });
}
</script>

<style scoped></style>
