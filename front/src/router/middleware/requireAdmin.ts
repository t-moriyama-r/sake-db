//管理者権限チェックミドルウェア
import type {
  NavigationGuardNext,
  RouteLocationNormalizedGeneric,
} from 'vue-router';

import { Roles } from '@/graphQL/Auth/types';
import { useUserStore } from '@/stores/userStore/userStore';

export function requireAdmin(
  to: RouteLocationNormalizedGeneric,
  next: NavigationGuardNext,
) {
  if (to.meta.requiresAdmin === undefined) {
    //管理者権限がそもそも不要なページの場合はそのまま遷移
    next();
    return;
  }

  const userStore = useUserStore();
  const roles = userStore.getRoles();

  // 管理者権限を持っているかチェック
  if (!roles.includes(Roles.Admin)) {
    console.error('管理者権限がありません！');
    // 管理者権限がない場合はトップページにリダイレクト
    next({ name: 'Index' });
    return;
  }

  console.log('管理者権限確認：OK');
  next();
}
