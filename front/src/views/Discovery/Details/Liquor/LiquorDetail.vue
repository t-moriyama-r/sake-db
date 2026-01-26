<!--酒詳細ページ-->
<template>
  <div v-if="liquor">
    <FlavorMap :liquor="liquor" />
    <div class="flex">
      <BackButton />
      <p class="title">{{ liquor.name }}</p>
    </div>
    <CategoryTrail :category-trails="liquor.categoryTrail" />
    <img
      v-if="liquor.imageUrl"
      :src="liquor.imageUrl"
      class="image"
      alt="画像"
    />
    <div>
      <span v-for="(line, index) in descriptionLines" :key="index">
        {{ line }}<br v-if="index !== descriptionLines.length - 1" />
      </span>
    </div>
    <div class="max-w-[515px]">
      <iframe
        v-if="liquor.youtube"
        class="w-full aspect-video"
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
    <router-link :to="{ name: 'LiquorEdit', params: { id: liquor.id } }">
      <CommonButton>編集する</CommonButton></router-link
    >
    <LiquorTags :liquor-id="liquor.id" />
    <AffiliateContainer :name="liquor.name" />
    <LiquorBoard :liquorId="liquor.id" />
    <RelatedLiquors :liquor-id="liquor.id" />
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
import RelatedLiquors from '@/views/Discovery/Details/Liquor/RelatedLiquors.vue';
import LiquorTags from '@/views/Discovery/Details/Liquor/tag/LiquorTagArea.vue';

interface Props {
  liquor: Liquor;
}

const { liquor } = defineProps<Props>();

const descriptionLines = computed<string[]>(() => {
  return liquor.description.split('\n');
});

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

<style scoped>
p.title {
  font-size: 150%;
  font-weight: bold;
}

img.image {
  max-height: 300px;
  max-width: 500px;
}
</style>
