<template>
  <CommonButton
    :size="'small'"
    :color="ColorType.None"
    class="px-1 sm:px-6 md:px-3 lg:px-6"
    @click="openDialog"
  >
    <FontAwesomeIcon icon="fa-solid fa-arrow-right-to-bracket" />
    <span class="hidden sm:inline">ログイン / 新規登録</span>
  </CommonButton>
  <CommonDialog
    v-model="isDialogOpen"
    title="ログイン"
    :isUnUseDefaultButtons="true"
  >
    <div class="text-left">
      <LoginForm />
      <div class="mt-4 text-center">
        <router-link :to="{ name: 'Register' }" @click="closeDialog">
          <CommonButton :size="'small'">新規登録はこちら</CommonButton>
        </router-link>
      </div>
    </div>
  </CommonDialog>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { ref } from 'vue';

import LoginForm from '@/components/blocks/auth/LoginForm/LoginForm.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import CommonDialog from '@/components/parts/common/CommonDialog/CommonDialog.vue';
import { ColorType } from '@/type/common/ColorType';

const isDialogOpen = ref<boolean>(false);
const openDialog = () => {
  isDialogOpen.value = true;
};
const closeDialog = () => {
  isDialogOpen.value = false;
};
</script>

<style scoped></style>
