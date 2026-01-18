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
      <router-link
        :to="{ name: 'Register' }"
        @click="closeDialog"
        class="mt-3 flex w-full items-center justify-center gap-2 rounded-md bg-blue-500 px-4 py-2 font-medium text-white transition hover:bg-blue-600"
      >
        <FontAwesomeIcon icon="fa-solid fa-person-circle-plus" />
        <span>新規登録</span>
      </router-link>
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
