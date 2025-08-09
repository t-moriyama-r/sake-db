<template>
  <Form
    @submit="onSubmit"
    :initial-values="{ keyword: '' }"
    :validation-schema="validationSchema"
    ><div class="flex">
      <FormField name="keyword" :showErrors="'hidden'" />
      <SubmitButton :size="'small'" class="ml-1 px-2"
        ><FontAwesomeIcon icon="fa-solid fa-magnifying-glass"
      /></SubmitButton>
    </div>
  </Form>
</template>

<script setup lang="ts">
/**
 * TODO: ロジックは後で実装
 */

import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { Form } from 'vee-validate';
import { string } from 'yup';

import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';

const validationSchema = {
  keyword: string().required(),
};

const onSubmit = () => {
  //TODO
};
</script>

<style scoped></style>
