<template>
  <CommonTag
    text="+追加"
    isHiddenHash
    @click="showModal"
    :classes="['addTag']"
  />
  <CommonDialog
    title="タグを追加"
    v-model="isDialogOpen"
    is-un-use-default-buttons
    v-slot="{ close }"
  >
    <!--submitの警告は、vee-validateのGenericObjectの型が広すぎることによるので無視してOK-->
    <Form
      :initial-values="defaultValues(props.liquorId)"
      :validation-schema="validationSchema"
      @submit="
        async (values: FormValues) => {
          await onSubmit(values);
          close();
        }
      "
    >
      <FormField :name="PostTagKeys.LiquorId" type="hidden" />
      <div class="tag-input-container">
        <FormField
          :name="PostTagKeys.Tag"
          :showErrors="'show'"
          placeholder="タグを入力してください（最大20文字）"
          label="タグ名"
        />
      </div>
      <div class="button-container">
        <SubmitButton size="small" class="submit-button">登録</SubmitButton>
        <CommonButton size="small" @click="close" class="cancel-button"
          >キャンセル</CommonButton
        >
      </div>
    </Form>
  </CommonDialog>
</template>

<script setup lang="ts">
import { Form } from 'vee-validate';
import { ref } from 'vue';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import CommonDialog from '@/components/parts/common/CommonDialog/CommonDialog.vue';
import CommonTag from '@/components/parts/common/CommonTag/CommonTag.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import {
  PostTag,
  PostTagKeys,
  type PostTagResponse,
} from '@/graphQL/Liquor/tags';

import { defaultValues, type FormValues, validationSchema } from './form';

const props = defineProps<{
  liquorId: string;
}>();

const { execute } = useMutation<PostTagResponse>(PostTag, {
  isAuth: true,
});
const toast = useToast();

const emit = defineEmits(['submitted']);

const isDialogOpen = ref<boolean>(false);

async function onSubmit(values: FormValues) {
  const response = await execute({ input: values });
  toast.showToast({
    message: 'タグの登録に成功しました',
  });
  emit('submitted', response.postTag);
}

function showModal() {
  isDialogOpen.value = true;
}
</script>

<style scoped>
.addTag {
  color: #333;
  border: 1px solid #ccc;

  &:hover {
    background-color: #f8d7da;
    color: #721c24;
  }
}

.tag-input-container {
  margin: 1.5rem 0;
  text-align: left;
}

.button-container {
  display: flex;
  gap: 0.75rem;
  justify-content: center;
  margin-top: 1.5rem;
  padding-top: 1rem;
  border-top: 1px solid #e5e7eb;
}

.submit-button {
  background-color: #3b82f6;
  color: white;
  font-weight: 500;
  padding: 0.5rem 1.5rem;
  transition: all 0.2s;

  &:hover:not(:disabled) {
    background-color: #2563eb;
    transform: translateY(-1px);
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  }

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
}

.cancel-button {
  background-color: #f3f4f6;
  color: #374151;
  font-weight: 500;
  padding: 0.5rem 1.5rem;
  transition: all 0.2s;

  &:hover {
    background-color: #e5e7eb;
  }
}
</style>
