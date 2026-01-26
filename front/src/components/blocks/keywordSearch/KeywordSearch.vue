<template>
  <div class="flex-1 px-0 sm:px-6 max-w-[480px]">
    <Form
      @submit="onSubmit"
      :initial-values="{ keyword: '' }"
      :validation-schema="validationSchema"
      ><div class="flex">
        <FormField name="keyword" classes="w-full" :showErrors="'hidden'" />
        <SubmitButton :size="'small'" class="ml-1 px-2"
          ><FontAwesomeIcon icon="fa-solid fa-magnifying-glass"
        /></SubmitButton>
      </div>
    </Form>
    <!-- ダイアログをボタン表示する場合-->
    <!--    <div class="flex xs:hidden items-center justify-center">-->
    <!--      <CommonButton :size="'small'" class="px-2" @click="onDialogOpen"-->
    <!--        ><FontAwesomeIcon icon="fa-solid fa-magnifying-glass"-->
    <!--      /></CommonButton>-->
    <!--    </div>-->
    <!--    <KeywordSearchDialog v-model="isDialogOpen" :onSubmit="onSubmit" />-->
  </div>
</template>

<script setup lang="ts">
/**
 * キーワード検索コンポーネント
 */

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { Form, type SubmissionHandler } from 'vee-validate';
import { useRouter } from 'vue-router';

//import { ref } from 'vue';
import {
  type FormValues,
  validationSchema,
} from '@/components/blocks/keywordSearch/formSettings';
//import KeywordSearchDialog from '@/components/blocks/keywordSearch/KeywordSearchDialog.vue';
//import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';

const router = useRouter();

// const isDialogOpen = ref(false);
// const onDialogOpen = () => {
//   isDialogOpen.value = true;
// };
// const onDialogClose = () => {
//   isDialogOpen.value = false;
// };

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: FormValues) => {
  // 検索結果ページに遷移
  await router.push({
    name: 'SearchResults',
    query: { keyword: values.keyword },
  });
  //onDialogClose();
};
</script>

<style scoped></style>
