<template>
  <div class="px-5">
    <div class="flex">
      <BackButton />
      <h1>パスワードリセット</h1>
    </div>
    <PasswordResetForm />
  </div>
</template>

<script setup lang="ts">
import PasswordResetForm from '@/components/blocks/auth/PasswordResetForm/PasswordResetForm.vue';
import BackButton from '@/components/parts/common/BackButton.vue';
</script>

<style scoped>
h1 {
  font-size: 150%;
  font-weight: bold;
}
</style>
