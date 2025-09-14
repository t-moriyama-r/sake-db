<template>
  <CommonDialog v-model="isOpen">
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
  </CommonDialog>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { Form, type SubmissionHandler } from 'vee-validate';

import { validationSchema } from '@/components/blocks/keywordSearch/formSettings';
import CommonDialog from '@/components/parts/common/CommonDialog/CommonDialog.vue';
import FormField from '@/components/parts/forms/core/FormField.vue';
import SubmitButton from '@/components/parts/forms/core/SubmitButton.vue';

const isOpen = defineModel<boolean>();
const { onSubmit } = defineProps<{
  onSubmit: SubmissionHandler;
}>();
</script>

<style scoped></style>
