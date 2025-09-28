<template>
  <router-link :to="{ name: 'MyPageIndex' }">
    <div v-if="userStore.user" class="flex container">
      <RadiusImage
        v-if="userStore.user.imageBase64"
        :imageSrc="userStore.user.imageBase64"
        radius="10px"
      />
      <div>
        {{ truncateString({ str: userStore.user.name, maxLength: 10 }) }}
      </div>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import truncateString from '@/funcs/util/transform/truncateString';
import { useUserStore } from '@/stores/userStore/userStore';

const userStore = useUserStore();
</script>

<style scoped></style>
