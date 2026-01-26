<!--関連銘柄の表示コンポーネント-->
<template>
  <div v-if="relatedLiquors.length > 0" class="related-liquors-container">
    <h2>似ているお酒</h2>
    <div class="liquor-grid">
      <LiquorCard
        v-for="liquor in relatedLiquors"
        :key="liquor.id"
        :liquor="liquor"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

import LiquorCard from '@/components/blocks/cards/LiquorCard.vue';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import type { Liquor } from '@/graphQL/Index/random';
import {
  RELATED_LIQUORS_GET,
  type RelatedLiquorsResponse,
} from '@/graphQL/Liquor/relatedLiquors';

interface Props {
  liquorId: string;
}

const props = defineProps<Props>();

const relatedLiquors = ref<Liquor[]>([]);
const { fetch } = useQuery<RelatedLiquorsResponse>(RELATED_LIQUORS_GET);

// データフェッチ
const fetchData = async (id: string): Promise<void> => {
  try {
    const { relatedLiquors: response } = await fetch({
      id,
    });
    relatedLiquors.value = response;
  } catch (error) {
    // エラーが発生した場合は空配列を設定（関連銘柄が表示されないだけで、他の機能には影響しない）
    console.error('Failed to fetch related liquors:', error);
    relatedLiquors.value = [];
  }
};

// liquorIdの変更を監視
watch(
  () => props.liquorId,
  (newId) => {
    if (newId) {
      fetchData(newId);
    }
  },
  { immediate: true },
);
</script>

<style scoped>
.related-liquors-container {
  margin-top: 3rem;
  padding-top: 2rem;
  border-top: 2px solid #e0e0e0;
}

h2 {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
  color: #333;
}

.liquor-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 1.5rem;
}

@media (max-width: 768px) {
  .liquor-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 1rem;
  }
}
</style>
