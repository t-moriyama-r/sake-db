<template>
  <div v-if="flavorMap">
    <FlavorMap
      :liquor="liquor"
      :flavorMap="flavorMap"
      :votedCoordinates="votedCoordinates"
      @clickCell="onClickCell"
    />
    <div>
      <p>
        投票数：{{ flavorMap.guestFullAmount + flavorMap.userFullAmount }}人
      </p>
      <p>登録済ユーザ：{{ flavorMap.userFullAmount }}人</p>
      <p>未登録ユーザ：{{ flavorMap.guestFullAmount }}人</p>
    </div>
  </div>
  <YesNoDialog
    v-if="savedCoordinates && flavorMap"
    v-model="isShowPostDialog"
    :on-yes="onSubmit"
  >
    <div v-if="savedCoordinates.x >= 0">
      {{ flavorMap.xNames[0] }}: {{ savedCoordinates.x }}
    </div>
    <div v-else>{{ flavorMap.xNames[1] }}: {{ savedCoordinates.x * -1 }}</div>
    <div v-if="savedCoordinates.y >= 0">
      {{ flavorMap.yNames[0] }}: {{ savedCoordinates.y }}
    </div>
    <div v-else>{{ flavorMap.yNames[1] }}: {{ savedCoordinates.y * -1 }}</div>
    <div>
      登録しますか？
      <div v-if="!isLogin" class="text-red-800">
        ※(注意)未ログイン状態の場合、一度投票したら変えられません。
      </div>
    </div>
  </YesNoDialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';

import YesNoDialog from '@/components/parts/common/CommonDialog/Variations/YesNoDialog.vue';
import { useMutation } from '@/funcs/composable/useQuery/useQuery';
import { useToast } from '@/funcs/composable/useToast';
import {
  type Coordinates,
  type FlavorMap as FlavorMapType,
  type FlavorMapResponse,
  PostFlavorMap,
} from '@/graphQL/Liquor/flavorMap';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import { guestFlavorMapStore } from '@/stores/guestFlavorMapStore';
import { useUserStore } from '@/stores/userStore/userStore';
import FlavorMap from '@/views/Discovery/Details/Liquor/flavorMap/postDialog/FlavorMap.vue';

interface Props {
  liquor: Liquor;
  flavorMap: FlavorMapType;
  votedCoordinates: Coordinates | null;
}

const { isLogin } = useUserStore();
const { addItem } = guestFlavorMapStore();
const { execute: post } = useMutation<FlavorMapResponse>(PostFlavorMap, {
  isAuth: true,
});
const toast = useToast();

const { liquor } = defineProps<Props>();
const emit = defineEmits(['onSubmit']); // 親に送るイベントを定義

const isShowPostDialog = ref<boolean>(false);
const savedCoordinates = ref<Coordinates | null>(null);

function onClickCell(coordinates: Coordinates) {
  savedCoordinates.value = coordinates;
  isShowPostDialog.value = true;
}

async function onSubmit() {
  await post({
    input: {
      liquorId: liquor.id,
      x: savedCoordinates.value!.x,
      y: savedCoordinates.value!.y,
    },
  });
  toast.showToast({
    message: '投票が完了しました',
  });
  if (!isLogin) {
    addItem({
      liquorId: liquor.id,
      x: savedCoordinates.value!.x,
      y: savedCoordinates.value!.y,
    });
  }
  emit('onSubmit');
}
</script>

<style scoped lang="scss">
div.flavor-map-container {
  width: calc(420px + 2em);
  height: calc(420px + 2em);
}

div.flavor-map-grid-container {
  background-color: #eee;
  div.grid-row {
    height: calc(420px / 21);
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
