<template>
  <div class="flex">
    <BackButton />
    <h1>ユーザー情報編集</h1>
  </div>
  <UserEditForm v-if="user" :user="user" />
</template>

<script setup lang="ts">
import BackButton from '@/components/parts/common/BackButton.vue';
import { useUserDetail } from '@/views/MyPage/funcs';
import UserEditForm from '@/views/MyPage/UserEdit/form/UserEditForm.vue';

const { user } = useUserDetail();
</script>

<style scoped>
h1 {
  font-size: 150%;
  font-weight: bold;
}
</style>
