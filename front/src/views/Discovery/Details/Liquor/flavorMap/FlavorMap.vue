<template>
  <div v-if="flavorMap" class="flavor-map-container">
    <FlavorMapMini
      :liquor="liquor"
      :flavorMap="flavorMap"
      :votedCoordinates="votedCoordinates"
      @click="openPostDialog"
    />
    <div class="description">
      <p>
        投票数：{{ flavorMap.guestFullAmount + flavorMap.userFullAmount }}人
      </p>
      <p>登録済ユーザ：{{ flavorMap.userFullAmount }}人</p>
      <p>未登録ユーザ：{{ flavorMap.guestFullAmount }}人</p>
    </div>
  </div>
  <CommonDialog v-model="isOpenPostDialog" isCentered>
    <PostDialog
      v-if="flavorMap"
      :liquor="liquor"
      :flavorMap="flavorMap"
      :voted-coordinates="votedCoordinates"
      @onSubmit="onFetch"
    />
  </CommonDialog>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

import CommonDialog from '@/components/parts/common/CommonDialog/CommonDialog.vue';
import { useMutation, useQuery } from '@/funcs/composable/useQuery/useQuery';
import {
  type Coordinates,
  type FlavorMap,
  type FlavorMapResponse,
  GetFlavorMap,
  GetVoted,
  type VotedResponse,
} from '@/graphQL/Liquor/flavorMap';
import type { Liquor } from '@/graphQL/Liquor/liquor';
import { guestFlavorMapStore } from '@/stores/guestFlavorMapStore';
import { useUserStore } from '@/stores/userStore/userStore';
import FlavorMapMini from '@/views/Discovery/Details/Liquor/flavorMap/FlavorMapMini.vue';
import PostDialog from '@/views/Discovery/Details/Liquor/flavorMap/postDialog/PostDialog.vue';

interface Props {
  liquor: Liquor;
}

const { isLogin } = useUserStore();
const { getById } = guestFlavorMapStore();
const { execute: fetch } = useMutation<FlavorMapResponse>(GetFlavorMap, {
  isAuth: true,
});
const { fetch: votedFetch } = useQuery<VotedResponse>(GetVoted, {
  isAuth: true,
});

const { liquor } = defineProps<Props>();
const flavorMap = ref<FlavorMap | null>(null);
const votedCoordinates = ref<Coordinates | null>(null);
const isOpenPostDialog = ref<boolean>(false);

onMounted(() => {
  void onFetch();
});

function openPostDialog() {
  isOpenPostDialog.value = true;
}

async function onFetch() {
  void setVoted(); //再読込ごとに投票済みかどうかを確認
  const response: FlavorMapResponse = await fetch(
    { liquorId: liquor.id },
    {
      fetchPolicy: 'network-only',
    },
  );
  flavorMap.value = response.getFlavorMap;
}

async function setVoted(): Promise<void> {
  if (isLogin) {
    const response: VotedResponse = await votedFetch(
      { liquorId: liquor.id },
      {
        fetchPolicy: 'network-only',
      },
    );
    votedCoordinates.value = response.getVoted;
    return;
  }
  //ゲストユーザーはローカルストレージから取得
  votedCoordinates.value = getById(liquor.id);
}
</script>

<style scoped lang="scss">
div.flavor-map-container {
  display: block;
  float: right;
  margin: 0 0 1em 1em; // 左下に余白をつけて回り込みを自然に

  div.description {
    text-align: right;
    font-size: 65%;
    margin-right: 1em;
  }
}
</style>
