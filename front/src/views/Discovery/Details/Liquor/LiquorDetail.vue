<!--酒詳細ページ-->
<template>
  <div v-if="liquor" class="max-w-5xl mx-auto px-4 py-6 space-y-6">
    <!-- フレーバーマップ -->
    <FlavorMap :liquor="liquor" />

    <!-- メインコンテンツカード -->
    <div class="bg-white rounded-lg shadow-md overflow-hidden">
      <!-- ヘッダー部分 -->
      <div class="p-6 border-b border-gray-200">
        <div class="flex items-start gap-4 mb-3">
          <BackButton />
          <h1 class="text-3xl font-bold text-gray-900 flex-1">
            {{ liquor.name }}
          </h1>
        </div>
        <CategoryTrail :category-trails="liquor.categoryTrail" />
      </div>

      <!-- 画像とYouTube埋め込み -->
      <div class="p-6 space-y-6">
        <div
          v-if="liquor.imageUrl || liquor.youtube"
          class="flex flex-col md:flex-row gap-6"
        >
          <img
            v-if="liquor.imageUrl"
            :src="liquor.imageUrl"
            class="w-full md:w-1/2 h-auto object-contain rounded-lg shadow-sm max-h-96"
            alt="画像"
          />
          <div
            v-if="liquor.youtube"
            class="w-full md:w-1/2 aspect-video rounded-lg overflow-hidden shadow-sm"
          >
            <iframe
              class="w-full h-full"
              :src="embedUrl ?? undefined"
              title="YouTube video player"
              allow="
                accelerometer;
                clipboard-write;
                encrypted-media;
                gyroscope;
                picture-in-picture;
                web-share;
              "
              allowfullscreen
            />
          </div>
        </div>

        <!-- 説明文 -->
        <div>
          <p class="text-gray-700 leading-relaxed whitespace-pre-line">
            {{ liquor.description }}
          </p>
        </div>

        <!-- 編集ボタン -->
        <div class="flex justify-end pt-4">
          <router-link :to="{ name: 'LiquorEdit', params: { id: liquor.id } }">
            <CommonButton>編集する</CommonButton>
          </router-link>
        </div>
      </div>
    </div>

    <!-- タグ -->
    <LiquorTags :liquor-id="liquor.id" />

    <!-- アフィリエイト -->
    <AffiliateContainer :name="liquor.name" />

    <!-- 掲示板 -->
    <LiquorBoard :liquorId="liquor.id" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import AffiliateContainer from '@/components/blocks/common/amazon/AffiliateContainer.vue';
import BackButton from '@/components/parts/common/BackButton.vue';
import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import LiquorBoard from '@/views/Discovery/Details/Liquor/board/LiquorBoard.vue';
import CategoryTrail from '@/views/Discovery/Details/Liquor/CategoryTrail.vue';
import FlavorMap from '@/views/Discovery/Details/Liquor/flavorMap/FlavorMap.vue';
import LiquorTags from '@/views/Discovery/Details/Liquor/tag/LiquorTagArea.vue';

interface Props {
  liquor: Liquor;
}

const { liquor } = defineProps<Props>();

// YouTubeのURLをembed形式に変換するcomputedプロパティ
const embedUrl = computed<string | null>(() => {
  if (!liquor.youtube) return null;
  const videoIdMatch = liquor.youtube.match(
    /(?:youtube\.com\/watch\?v=|youtu\.be\/)([a-zA-Z0-9_-]{11})/,
  );
  return videoIdMatch
    ? `https://www.youtube.com/embed/${videoIdMatch[1]}`
    : null;
});
</script>

<style scoped></style>
