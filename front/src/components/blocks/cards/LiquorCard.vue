<!--ランダムレコメンドのカード-->
<template>
  <div v-if="liquor" class="liquor-container">
    <router-link :to="{ name: 'LiquorDetail', params: { id: liquor.id } }">
      <div class="top-img-area">
        <img
          v-if="liquor.imageBase64"
          :src="`data:image/jpg;base64,${liquor.imageBase64}`"
          :alt="liquor.name"
        />
        <img v-else class="no-image" src="/no_image.svg" alt="no image" />
      </div>
    </router-link>
    <div class="bottom-content">
      <router-link
        :to="{ name: 'CategoryNarrowDown', params: { id: liquor.categoryId } }"
        ><p class="category-name">{{ liquor.categoryName }}</p></router-link
      >
      <router-link :to="{ name: 'LiquorDetail', params: { id: liquor.id } }"
        ><p class="title">{{ liquor.name }}</p></router-link
      >
    </div>
  </div>
</template>

<script setup lang="ts">
//propsのセット
import type { Liquor } from '@/graphQL/Index/random';

interface Props {
  liquor: Liquor;
}

const { liquor } = defineProps<Props>();
</script>

<style scoped>
div.liquor-container {
  display: grid;
  grid-template-rows: 1fr 80px; /* 上側を自動 (1fr)、下側を固定長 (100px) */
  height: 100%; /* コンテナ全体の高さを指定（親要素の高さに依存） */

  border: 1px solid #777;

  div.top-img-area {
    display: flex;
    align-items: center; /* ← 縦方向の中央揃え */
    justify-content: center; /* ← 横方向も中央にしたい場合（任意） */
    height: 100%; /* ← 必要に応じて高さを明示 */
    img {
      width: 100%;
      object-fit: contain; /* 縦横比を維持して収める */
      max-height: 100%; /* 親要素からはみ出さないように */
    }
    img.no-image {
      width: 80%;
      height: auto;
      object-fit: contain;
    }
  }
  div.bottom-content {
    p.category-name {
      font-size: 75%;
    }
    p.title {
      font-weight: bold;
    }
  }
}
</style>
