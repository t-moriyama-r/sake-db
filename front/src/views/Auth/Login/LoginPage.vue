<template>
  <div class="px-5">
    <div class="flex">
      <BackButton />
      <h1>ログイン</h1>
    </div>
    <LoginForm />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';

import BackButton from '@/components/parts/common/BackButton.vue';
import { useUserStore } from '@/stores/userStore/userStore';
import LoginForm from '@/views/Auth/Login/form/LoginForm.vue';

const { isLogin } = useUserStore();
const router = useRouter();

//すでにログイン済の場合、ユーザーページへリダイレクト
onMounted(() => {
  if (isLogin) {
    router.push({ name: 'UserIndex' });
  }
});
</script>

<style scoped>
h1 {
  font-size: 150%;
  font-weight: bold;
}
</style>
