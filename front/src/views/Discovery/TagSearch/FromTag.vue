<template>
  <div id="tag-search-area" v-if="liquors.length > 0">
    <CardContainer gap="0.5em" min="200px">
      <LiquorCard v-for="liquor in liquors" :liquor="liquor" :key="liquor.id" />
    </CardContainer>
  </div>
  <EmptyTag v-else :tag="tag" />
</template>

<script setup lang="ts">
import LiquorCard from '@/components/blocks/cards/LiquorCard.vue';
import CardContainer from '@/components/parts/common/CardContainer.vue';
import type { Liquor } from '@/graphQL/Index/random';
import EmptyTag from '@/views/Discovery/TagSearch/EmptyTag.vue';

interface Props {
  tag: string;
  liquors: Liquor[];
}

defineProps<Props>();
</script>

<style scoped>
div#tag-search-area {
  margin: auto;
  padding: 2em;
}
</style>
