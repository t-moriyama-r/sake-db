<template>
  <div class="mt-4 flex items-center gap-3">
    <div class="h-px flex-1 bg-gray-300"></div>
    <span class="text-sm text-gray-500">または</span>
    <div class="h-px flex-1 bg-gray-300"></div>
  </div>
  <button
    @click="loginWithX"
    class="mt-4 flex w-full items-center justify-center gap-2 rounded-md bg-black px-4 py-2 font-medium text-white transition hover:bg-neutral-800"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      fill="currentColor"
      class="h-5 w-5"
    >
      <path
        d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z"
      />
    </svg>
    <span>Xでログイン</span>
  </button>
</template>

<script setup lang="ts">
import axios from 'axios';

import createURL from '@/funcs/util/core/createURL';

async function loginWithX() {
  try {
    // バックエンドからX認証用のURLを取得
    const response = await axios.get(createURL('x/login'));
    const { redirectUrl } = response.data;

    // Xの認証ページへリダイレクト
    window.location.href = redirectUrl;
  } catch (error) {
    console.error('Failed to get login URL:', error);
  }
}
</script>

<style scoped></style>
