<template>
  <div id="tag-search-area" v-if="liquors.length > 0">
    <CardContainer gap="0.5em" min="200px">
      <LiquorCard v-for="liquor in liquors" :liquor="liquor" :key="liquor.id" />
    </CardContainer>
    <div ref="loadMoreTrigger" class="load-more-trigger"></div>
    <div v-if="isLoading" class="loading-indicator">
      <p>読み込み中...</p>
    </div>
  </div>
  <EmptyTag v-else-if="!isLoading" :tag="tag" />
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue';

import LiquorCard from '@/components/blocks/cards/LiquorCard.vue';
import CardContainer from '@/components/parts/common/CardContainer.vue';
import type { Liquor } from '@/graphQL/Index/random';
import EmptyTag from '@/views/Discovery/TagSearch/EmptyTag.vue';

interface Props {
  tag: string;
  liquors: Liquor[];
  hasMore: boolean;
  isLoading: boolean;
}

const props = defineProps<Props>();

const emit = defineEmits<{
  (_e: 'load-more'): void;
}>();

const loadMoreTrigger = ref<HTMLElement | null>(null);
let observer: IntersectionObserver | null = null;

onMounted(() => {
  // Intersection Observerを設定
  observer = new IntersectionObserver(
    (entries) => {
      const target = entries[0];
      if (target.isIntersecting && props.hasMore && !props.isLoading) {
        emit('load-more');
      }
    },
    {
      rootMargin: '100px', // 100px手前で発火
    },
  );

  if (loadMoreTrigger.value) {
    observer.observe(loadMoreTrigger.value);
  }
});

onUnmounted(() => {
  if (observer) {
    observer.disconnect();
  }
});
</script>

<style scoped>
div#tag-search-area {
  margin: auto;
  padding: 2em;
}

.load-more-trigger {
  height: 1px;
  width: 100%;
}

.loading-indicator {
  text-align: center;
  padding: 2em;
  color: #666;
}
</style>
