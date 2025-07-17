<template>
  <div v-if="flavorMap" @click="emit('click')">
    <div class="flavor-map-container flex flex-col">
      <div class="axis-label y text-center order-[-1]">
        {{ flavorMap.yNames[0] }}
      </div>
      <div class="axis-label y text-center order-1">
        {{ flavorMap.yNames[1] }}
      </div>
      <div class="flex flex-row">
        <div class="axis-label x flex items-center order-[-1]">
          {{ flavorMap.xNames[1] }}
        </div>
        <div class="axis-label x flex items-center order-1">
          {{ flavorMap.xNames[0] }}
        </div>
        <!-- フレーバーマップのグリッド -->
        <div
          class="flavor-map-grid-container"
          :class="isSelectable ? 'selectable' : ''"
        >
          <div
            v-for="(_, yIndex) in 21"
            :key="10 - yIndex"
            class="grid-row flex"
          >
            <CellView
              v-for="(_, xIndex) in 21"
              :key="xIndex - 10"
              :cellData="
                flavorMap.mapData.find(
                  (data: FlavorCell) =>
                    data.x === xIndex - 10 && data.y === 10 - yIndex,
                )!
              "
              class="grid-cell"
              :isSelected="
                votedCoordinates?.x === xIndex - 10 &&
                votedCoordinates?.y === 10 - yIndex
              "
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  type Coordinates,
  type FlavorCell,
  type FlavorMap,
} from '@/graphQL/Liquor/flavorMap';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import { useUserStore } from '@/stores/userStore/userStore';
import CellView from '@/views/Discovery/Details/Liquor/flavorMap/CellMini.vue';

interface Props {
  liquor: Liquor;
  flavorMap: FlavorMap;
  votedCoordinates: Coordinates | null;
}

const { isLogin } = useUserStore();

const { flavorMap, votedCoordinates } = defineProps<Props>();
const emit = defineEmits(['click']); // 親に送るイベントを定義
const isSelectable = isLogin || votedCoordinates == null;
</script>

<style scoped lang="scss">
$map-size: 105px;

div.flavor-map-container {
  width: calc(#{$map-size} + 2em);
  height: calc(#{$map-size} + 2em);
  //margin-bottom: 0.6em; // 下属性部分

  div.axis-label {
    font-size: 65%;
    &.x {
      margin: auto;
      writing-mode: vertical-rl;
    }
  }
}

div.flavor-map-grid-container {
  background-color: #eee;
  &.selectable {
    cursor: pointer;
  }
  div.grid-row {
    height: calc(#{$map-size} / 21);
  }
}
</style>
