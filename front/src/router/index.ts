import { createRouter, createWebHistory } from 'vue-router';

import MainRouter from '@/router/main';
import { authenticate } from '@/router/middleware/authenticate';
import { requireAdmin } from '@/router/middleware/requireAdmin';
import { useSelectedCategoryStore } from '@/stores/sidebar';
import MetaComponent from '@/views/MetaInfo.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'MetaView',
      component: MetaComponent, // 親ルート用の空コンポーネント
      children: [MainRouter],
    },
  ],
});

// カテゴリコンテキストを保持すべきルート名のリスト
const CATEGORY_CONTEXT_ROUTES = [
  'CategoryNarrowDown',
  'CategoryDetail',
  'LiquorDetail',
  'LiquorEdit',
  'LiquorPostByCategory',
  'CategoryEdit',
  'CategoryPostByParent',
];

//認証ミドルウェア
router.beforeEach((to, from, next) => {
  // カテゴリコンテキストを保持すべきルート以外では、カテゴリ選択をクリア
  if (to.name && !CATEGORY_CONTEXT_ROUTES.includes(to.name as string)) {
    const categoryStore = useSelectedCategoryStore();
    categoryStore.updateContent(null);
  }

  // 認証チェックを実行（認証が必要なページで未ログインの場合はログインページにリダイレクト）
  authenticate(to, (authNext) => {
    if (authNext === undefined || typeof authNext === 'boolean') {
      // 認証チェックが通過した場合、管理者権限チェックを実行
      requireAdmin(to, next);
    } else {
      // 認証チェックで別のページにリダイレクトされる場合はそのまま実行
      next(authNext);
    }
  });
});

export default router;
