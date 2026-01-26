<template>
  <div v-if="userStore.user" class="flex items-center gap-1">
    <RadiusImage
      v-if="userStore.user.imageBase64"
      :imageSrc="userStore.user.imageBase64"
      radius="10px"
    />
    <span>
      {{ truncateString({ str: userStore.user.name, maxLength: 10 }) }}
    </span>
  </div>
</template>

<script setup lang="ts">
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import truncateString from '@/funcs/util/transform/truncateString';
import { useUserStore } from '@/stores/userStore/userStore';

const userStore = useUserStore();
</script>

<style scoped></style>
