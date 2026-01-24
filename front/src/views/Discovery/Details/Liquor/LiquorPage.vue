<!--酒の情報ページ-->
<template>
  <LiquorDetail v-if="liquor" :liquor="liquor" />
</template>

<script setup lang="ts">
import { useHead } from '@vueuse/head';
import { computed, ref, watch } from 'vue';
import { useRoute } from 'vue-router';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import {
  type Liquor,
  LIQUOR_DETAIL_GET,
  type LiquorResponse,
} from '@/graphQL/Liquor/liquor';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import LiquorDetail from '@/views/Discovery/Details/Liquor/LiquorDetail.vue';

const isLoading = ref<boolean>(true);
const liquor = ref<Liquor | null>(null);

const route = useRoute(); // 現在のルートを取得
const sidebarStore = useSelectedCategoryStore();
const { fetch } = useQuery<LiquorResponse<Liquor>>(LIQUOR_DETAIL_GET);

const isNoCache: boolean = window.history.state?.noCache ?? false; //TODO:何故か常にtrueになってる...？

// メタタグの設定
const pageTitle = computed(() =>
  liquor.value ? `${liquor.value.name} - Sake DB` : 'Sake DB',
);

const pageDescription = computed(() => {
  if (!liquor.value) return '日本酒データベース';
  const desc = liquor.value.description || liquor.value.name;
  return desc.length > 200 ? desc.substring(0, 200) + '...' : desc;
});

const pageImage = computed(() => {
  // SNSカードにはimageUrlのみを使用（base64はサポートされない可能性がある）
  return liquor.value?.imageUrl || '';
});

const pageUrl = computed(() => {
  if (typeof window !== 'undefined' && liquor.value) {
    // クエリパラメータを除いた正規のURLを使用
    const origin = window.location.origin;
    return `${origin}/liquor/${liquor.value.id}`;
  }
  return '';
});

useHead({
  title: pageTitle,
  meta: [
    { name: 'description', content: pageDescription },
    // Open Graph tags
    { property: 'og:title', content: pageTitle },
    { property: 'og:description', content: pageDescription },
    { property: 'og:image', content: pageImage },
    { property: 'og:url', content: pageUrl },
    { property: 'og:type', content: 'website' },
    { property: 'og:site_name', content: 'Sake DB' },
    // Twitter Card tags
    { name: 'twitter:card', content: 'summary_large_image' },
    { name: 'twitter:title', content: pageTitle },
    { name: 'twitter:description', content: pageDescription },
    { name: 'twitter:image', content: pageImage },
  ],
});

// データフェッチ
const fetchData = async (id: string): Promise<void> => {
  const { liquor: response } = await fetch(
    {
      id,
    },
    {
      fetchPolicy: isNoCache ? 'network-only' : undefined, //更新直後だとキャッシュが残っているため、キャッシュを無効化
    },
  );
  liquor.value = response;
  sidebarStore.updateContent(response.categoryId);
  isLoading.value = false;
};

watch(
  () => route.params.id, // ルートのパスやクエリ、パラメータなどを監視
  (to) => {
    // ルートが変更された際に実行される処理
    const id = to as string; // ルートパラメータからidを取得
    if (!id) {
      isLoading.value = false;
      return;
    }
    fetchData(id);
  },
  { immediate: true }, // 初回レンダリング時に実行される
);
</script>
