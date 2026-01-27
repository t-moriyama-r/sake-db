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

import LoginForm from '@/components/blocks/auth/LoginForm/LoginForm.vue';
import BackButton from '@/components/parts/common/BackButton.vue';
import { useUserStore } from '@/stores/userStore/userStore';

const { isLogin } = useUserStore();
const router = useRouter();

//すでにログイン済の場合、マイページへリダイレクト
onMounted(() => {
  if (isLogin) {
    router.push({ name: 'MyPageIndex' });
  }
});
</script>

<style scoped>
h1 {
  font-size: 150%;
  font-weight: bold;
}
</style>
