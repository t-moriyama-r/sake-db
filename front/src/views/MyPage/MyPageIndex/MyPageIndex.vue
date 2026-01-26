<template>
  名前：{{ props.user.name }}
  <router-link :to="{ name: 'MyPageEdit' }">ユーザー情報編集</router-link>
  <router-link v-if="isAdmin" :to="{ name: 'Admin' }">管理画面</router-link>
  <div>
    <BookmarkList />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import type { AuthUserFull } from '@/graphQL/Auth/auth';
import { Roles } from '@/graphQL/Auth/types';
import BookmarkList from '@/views/MyPage/MyPageIndex/BookmarkList.vue';

interface Props {
  user: AuthUserFull;
}

const props = defineProps<Props>();

const isAdmin = computed(() => {
  return props.user.roles.includes(Roles.Admin);
});
</script>

<style scoped></style>
