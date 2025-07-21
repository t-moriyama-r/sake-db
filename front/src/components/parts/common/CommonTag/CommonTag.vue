<template>
  <span
    class="tag transition-colors duration-300 ease-in-out"
    :class="[
      baseColorClass,
      hoverColorClass,
      { showHash: !props.isHiddenHash },
      ...(props.classes ?? []),
    ]"
    @click="emit('click')"
    >{{ props.text }}
    <span
      v-if="props.isClose"
      class="close inline-block"
      @click.stop="emit('close')"
    >
      <span
        class="text-gray-500 transition-colors duration-300 ease-in-out"
      ></span></span
  ></span>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import { colorClassMap } from '@/components/parts/common/CommonTag/colors';
import type { TagProps } from '@/components/parts/common/CommonTag/type';

const props = defineProps<TagProps>();

const emit = defineEmits(['click', 'close']);

const baseColorClass = computed(
  () => colorClassMap[props.color ?? 'default'].base,
);
const hoverColorClass = computed(
  () => colorClassMap[props.color ?? 'default'].hover,
);
</script>

<style scoped>
span.tag {
  display: inline-block;
  margin: 0 0.1em 0.6em 0;
  padding: 0.6em;
  line-height: 1;
  text-decoration: none;
  background-color: #fff;
  border-width: 1px;
  border-style: solid;
  border-radius: 2em;
  cursor: pointer;

  &.showHash:before {
    content: '#';
  }
}

span.close {
  span {
    display: inline-block;
    vertical-align: middle;
    color: #333;
    line-height: 1;
    width: 1em;
    height: 0.1em;
    background: currentColor;
    border-radius: 0.1em;
    position: relative;
    transform: rotate(45deg);

    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background: inherit;
      border-radius: inherit;
      transform: rotate(90deg);
    }
  }

  &:hover span {
    color: #ff0000; /* ホバー時に色を赤に変更 */
  }
}
</style>
