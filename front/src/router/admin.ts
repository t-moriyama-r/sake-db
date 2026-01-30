import AdminIndex from '@/views/Admin/AdminIndex.vue';

/**
 * 管理関連画面のルーター
 */
const AdminRouter = {
  path: 'admin',
  meta: { requiresAuth: true, requiresAdmin: true },
  children: [
    {
      path: 'index',
      name: 'Admin',
      component: AdminIndex,
    },
  ],
};

export default AdminRouter;
