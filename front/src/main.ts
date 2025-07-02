import './styles/tailwind.css';
import './styles/main.css';

import { DefaultApolloClient } from '@vue/apollo-composable';
import { createApp, h, provide } from 'vue';

import client from '@/apolloClient';
import { errorDebug } from '@/funcs/util/core/console';
import { registerPlugins } from '@/plugins';
import { getGlobalToast } from '@/plugins/toast';

import App from './App.vue';

const app = createApp({
  setup() {
    provide(DefaultApolloClient, client);
  },
  render: () => h(App),
});

// プラグインの登録
registerPlugins(app);

app.mount('#app');

app.config.errorHandler = (err, instance) => {
  const toast = getGlobalToast();
  if (!toast) return;
  if (instance) {
    const componentName =
      instance.$options.name ||
      instance.$options.__name ||
      '匿名コンポーネント';
    errorDebug('エラー発生コンポーネント:', componentName);
  }
  if (
    err &&
    typeof err === 'object' &&
    'message' in err &&
    typeof err.message === 'string'
  ) {
    toast.errorToast(err.message);
  } else {
    toast.errorToast('不明なエラー：' + err);
  }
};
