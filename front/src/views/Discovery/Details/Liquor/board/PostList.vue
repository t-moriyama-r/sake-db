<template>
  <div class="bg-white rounded-lg shadow-md p-6">
    <h3 class="text-xl font-semibold text-gray-900 mb-4">みんなの感想</h3>
    <div class="space-y-4">
      <div
        v-for="(post, index) in props.posts"
        :key="index"
        class="border border-gray-200 rounded-lg p-4 hover:shadow-md transition-shadow"
      >
        <div class="flex items-start gap-4">
          <!-- ユーザー情報 -->
          <div class="flex-shrink-0">
            <div
              v-if="post.userImageBase64"
              class="w-10 h-10 rounded-full overflow-hidden bg-gray-200"
            >
              <img
                :src="`data:image/jpeg;base64,${post.userImageBase64}`"
                :alt="post.userName || '名無し'"
                class="w-full h-full object-cover"
              />
            </div>
            <div
              v-else
              class="w-10 h-10 rounded-full bg-gray-200 flex items-center justify-center"
            >
              <span class="text-gray-600 text-sm font-medium">
                {{ getUserInitial(post.userName) }}
              </span>
            </div>
          </div>

          <!-- コンテンツ -->
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-2">
              <router-link
                v-if="post.userId"
                :to="{ name: 'UserPage', params: { id: post.userId } }"
                class="font-semibold text-blue-600 hover:text-blue-800"
              >
                {{ post.userName }}
              </router-link>
              <span v-else class="font-semibold text-gray-600">名無し</span>

              <!-- 評価 -->
              <div
                class="flex items-center gap-1 px-2 py-1 rounded-full text-sm"
                :class="post.rate ? 'bg-yellow-100' : 'bg-gray-100'"
              >
                <span v-if="post.rate" class="flex items-center">
                  <span
                    v-for="i in 5"
                    :key="i"
                    class="text-yellow-500"
                    :class="i <= post.rate ? 'opacity-100' : 'opacity-30'"
                  >
                    ★
                  </span>
                </span>
                <span v-else class="text-gray-500">未評価</span>
              </div>
            </div>

            <!-- 投稿テキスト -->
            <p class="text-gray-700 leading-relaxed">{{ post.text }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import type { Post } from '@/graphQL/Liquor/board';

interface Props {
  posts: Post[];
}

const props = defineProps<Props>();

// ユーザー名のイニシャルを取得する
const getUserInitial = (userName: string | null | undefined): string => {
  if (!userName) return '?';
  return userName.charAt(0).toUpperCase();
};
</script>

<style scoped></style>
