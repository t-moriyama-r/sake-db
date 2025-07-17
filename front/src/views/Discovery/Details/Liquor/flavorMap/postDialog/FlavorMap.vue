<template>
  <div v-if="flavorMap">
    <div class="flavor-map-container flex flex-col">
      <div class="text-center order-[-1]">{{ flavorMap.yNames[0] }}</div>
      <div class="text-center order-1">{{ flavorMap.yNames[1] }}</div>
      <div class="flex flex-row">
        <div class="flex items-center order-[-1]">
          {{ flavorMap.xNames[1] }}
        </div>
        <div class="flex items-center order-1">{{ flavorMap.xNames[0] }}</div>
        <div class="relative">
          <!-- 上の矢印 -->
          <div class="absolute top-0 left-1/2 transform -translate-x-1/2">
            <div class="arrow-up"></div>
          </div>

          <!-- 左の矢印 -->
          <div class="absolute top-1/2 left-0 transform -translate-y-1/2">
            <div class="arrow-left"></div>
          </div>

          <!-- 右の矢印 -->
          <div class="absolute top-1/2 right-0 transform -translate-y-1/2">
            <div class="arrow-right"></div>
          </div>

          <!-- 下の矢印 -->
          <div class="absolute bottom-0 left-1/2 transform -translate-x-1/2">
            <div class="arrow-down"></div>
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
              <FlavorCell
                v-for="(_, xIndex) in 21"
                :key="xIndex - 10"
                :cellData="
                  flavorMap.mapData.find(
                    (data: FlavorCellType) =>
                      data.x === xIndex - 10 && data.y === 10 - yIndex,
                  )!
                "
                class="grid-cell"
                :isSelected="
                  votedCoordinates?.x === xIndex - 10 &&
                  votedCoordinates?.y === 10 - yIndex
                "
                :isSelectable="isSelectable"
                @click="
                  () => {
                    if (isLogin || votedCoordinates == null) {
                      //なんか到達不能って出てるけど、そんなことない......
                      emit('clickCell', { x: xIndex - 10, y: 10 - yIndex });
                    }
                  }
                "
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  type Coordinates,
  type FlavorCell as FlavorCellType,
  type FlavorMap,
} from '@/graphQL/Liquor/flavorMap';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import { useUserStore } from '@/stores/userStore/userStore';
import FlavorCell from '@/views/Discovery/Details/Liquor/flavorMap/postDialog/FlavorCell.vue';

interface Props {
  liquor: Liquor;
  flavorMap: FlavorMap;
  votedCoordinates: Coordinates | null;
}

const { flavorMap, votedCoordinates } = defineProps<Props>();
const emit = defineEmits(['clickCell']); // 親に送るイベントを定義

const { isLogin } = useUserStore();
const isSelectable = isLogin || votedCoordinates == null;
</script>

<style scoped lang="scss">
$map-size: 420px;

div.flavor-map-container {
  width: calc(#{$map-size} + 2em);
  height: calc(#{$map-size} + 2em);
  margin-bottom: 1em; // 下属性部分
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

// 矢印のスタイル
.arrow-up,
.arrow-down,
.arrow-left,
.arrow-right {
  width: 0;
  height: 0;
  border-style: solid;
}

// 矢印の具体的なスタイル設定
.arrow-up {
  border-width: 0 10px 10px 10px;
  border-color: transparent transparent #333 transparent;
}

.arrow-down {
  border-width: 10px 10px 0 10px;
  border-color: #333 transparent transparent transparent;
}

.arrow-left {
  border-width: 10px 10px 10px 0;
  border-color: transparent #333 transparent transparent;
}

.arrow-right {
  border-width: 10px 0 10px 10px;
  border-color: transparent transparent transparent #333;
}
</style>
