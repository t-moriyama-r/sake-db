<template>
  <Form
    @submit="onSubmit"
    :validation-schema="{
      [EMAIL_NAME]: string().required().email(),
    }"
  >
    <FormField name="email" label="メールアドレス" type="email" />
    <p class="mb-4 text-sm text-gray-500">
      登録済みのメールアドレスを入力してください。<br />
      パスワードリセット用のリンクをお送りします。
    </p>
    <div class="flex justify-center">
      <SubmitButton>送信</SubmitButton>
    </div>
  </Form>
  <div v-if="showBackToLogin" class="mt-4 text-center">
    <button
      type="button"
      class="text-sm text-gray-500 hover:underline"
      @click="emit('backToLogin')"
    >
      ログインに戻る
    </button>
  </div>
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import { string } from 'yup';

import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import { PASSWORD_RESET } from '@/graphQL/Auth/auth';
import { ToastType } from '@/plugins/toast';

defineProps<{
  showBackToLogin?: boolean;
}>();

const emit = defineEmits<{
  backToLogin: [];
}>();

const EMAIL_NAME = 'email';

const toast = useToast();

const { execute } = useMutation<{
  passwordReset: boolean;
}>(PASSWORD_RESET, {
  isUseSpinner: true,
});

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: {
  [EMAIL_NAME]: string;
}) => {
  await execute({
    email: values[EMAIL_NAME],
  }).then(() => {
    toast.showToast({
      message: 'パスワードリセットメールを送信しました。',
      type: ToastType.Success,
    });
  });
};
</script>

<style scoped></style>
