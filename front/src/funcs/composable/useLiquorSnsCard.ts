import { useHead } from '@vueuse/head';
import { computed, type Ref } from 'vue';

import type { Liquor } from '@/graphQL/Liquor/liquor';

/**
 * お酒ページのSNSカード（Open Graph/Twitter Card）メタタグを設定するコンポーザブル
 * @param liquor - お酒データのリアクティブ参照
 */
export const useLiquorSnsCard = (liquor: Ref<Liquor | null>) => {
  // ページタイトル
  const pageTitle = computed(() =>
    liquor.value ? `${liquor.value.name} - Sake DB` : 'Sake DB',
  );

  // ページ説明文（200文字まで）
  const pageDescription = computed(() => {
    if (!liquor.value) return '日本酒データベース';
    const desc = liquor.value.description || liquor.value.name;
    return desc.length > 200 ? desc.substring(0, 200) + '...' : desc;
  });

  // ページ画像（SNSカード用にimageUrlのみ使用）
  const pageImage = computed(() => {
    // SNSカードにはimageUrlのみを使用（base64はサポートされない可能性がある）
    return liquor.value?.imageUrl || '';
  });

  // ページURL（正規化）
  const pageUrl = computed(() => {
    if (typeof window !== 'undefined' && liquor.value) {
      // クエリパラメータを除いた正規のURLを使用
      const origin = window.location.origin;
      return `${origin}/liquor/${liquor.value.id}`;
    }
    return '';
  });

  // メタタグの設定
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

  return {
    pageTitle,
    pageDescription,
    pageImage,
    pageUrl,
  };
};
