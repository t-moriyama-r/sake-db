<template>
  <div class="grid place-items-center min-h-[200px]">
    <div>
      <p>このカテゴリに登録されたお酒がありません。</p>
    </div>
    <div>
      <CommonButton @click="toPostPage"> お酒を登録 </CommonButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';

interface Props {
  categoryId: string;
}
const props = defineProps<Props>();

const router = useRouter();
const toPostPage = () => {
  console.log('CategoryId:', props.categoryId);
  router.push({
    name: 'LiquorPostByCategory',
    params: { categoryId: props.categoryId },
  });
};
</script>

<style scoped></style>
