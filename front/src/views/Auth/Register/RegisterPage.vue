<template>
  <div class="px-5">
    <div class="flex">
      <BackButton />
      <h1>新規登録</h1>
    </div>
    <RegisterForm />
  </div>
</template>

<script setup lang="ts">
import BackButton from '@/components/parts/common/BackButton.vue';
import RegisterForm from '@/views/Auth/Register/form/RegisterForm.vue';
</script>

<style scoped>
h1 {
  font-size: 150%;
  font-weight: bold;
}
</style>
