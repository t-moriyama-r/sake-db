<template>
  <Form
    @submit="onSubmit"
    :initial-values="initialValues"
    :validation-schema="validationSchema"
    :validate-on-mount="false"
    validate-on-input
    v-slot="{ meta }"
  >
    <FormField
      :name="FormKeys.MAIL"
      label="メールアドレス"
      type="email"
      :showErrors="'hidden'"
    />
    <FormField
      :name="FormKeys.PASSWORD"
      label="パスワード"
      type="password"
      :showErrors="'hidden'"
    />
    <div class="mt-2 flex justify-center">
      <SubmitButton :isDisabled="!meta.valid">ログイン</SubmitButton>
    </div>
  </Form>
  <div class="mt-2 text-center">
    <button
      v-if="isModal"
      type="button"
      class="text-sm text-gray-500 hover:underline"
      @click="emit('goToPasswordReset')"
    >
      パスワードをお忘れですか？
    </button>
    <router-link
      v-else
      :to="{ name: 'PasswordReset' }"
      class="text-sm text-gray-500 hover:underline"
    >
      パスワードをお忘れですか？
    </router-link>
  </div>
  <XLogin />
</template>

<script setup lang="ts">
import { Form, type SubmissionHandler } from 'vee-validate';
import { useRouter } from 'vue-router';

import XLogin from '@/components/blocks/auth/XLogin.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import { LOGIN } from '@/graphQL/Auth/auth';
import type { LoginMutation } from '@/graphQL/auto-generated';
import { getAuthPayloadForUI } from '@/stores/userStore/type';
import { useUserStore } from '@/stores/userStore/userStore';

import {
  FormKeys,
  type FormValues,
  initialValues,
  validationSchema,
} from './LoginForm';

defineProps<{
  isModal?: boolean;
}>();

const emit = defineEmits<{
  goToPasswordReset: [];
}>();

const router = useRouter();
const userStore = useUserStore();
const toast = useToast();
const { execute } = useMutation<LoginMutation>(LOGIN);

// extends GenericObjectは型が広すぎるのでキャストして対応する
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
//@ts-expect-error
const onSubmit: SubmissionHandler = async (values: FormValues) => {
  await execute({
    input: {
      email: values[FormKeys.MAIL],
      password: values[FormKeys.PASSWORD],
    },
  })
    .then((res) => {
      //トークンをセットし、トップへリンク
      userStore.setUserData(getAuthPayloadForUI(res.login)); //ストアの情報を更新する
      router.push({ name: 'Index' });
    })
    .catch((err) => {
      toast.errorToast(err.message);
    });
};
</script>

<style scoped></style>
