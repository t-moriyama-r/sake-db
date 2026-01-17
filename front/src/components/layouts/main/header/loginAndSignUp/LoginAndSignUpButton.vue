<template>
  <CommonButton :size="'small'" :color="ColorType.None" @click="openDialog">
    <slot>
      <FontAwesomeIcon
        icon="fa-solid fa-arrow-right-to-bracket"
        class="text-lg"
      />
    </slot>
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

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import CommonDialog from '@/components/parts/common/CommonDialog/CommonDialog.vue';
import { ColorType } from '@/type/common/ColorType';
import LoginForm from '@/views/Auth/Login/form/LoginForm.vue';

const isDialogOpen = ref<boolean>(false);
const openDialog = () => {
  isDialogOpen.value = true;
};
const closeDialog = () => {
  isDialogOpen.value = false;
};
</script>

<style scoped></style>
