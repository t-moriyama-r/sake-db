<template>
  <div>
    <HeaderArea @openSideBar="openSideBar" />
    <main class="flex">
      <!-- PC・タブレット用 -->
      <section id="sidebar" class="sidebar fixed hidden sm:block">
        <SideBar />
      </section>

      <!-- スマホ用 スライドサイドバー -->
      <transition name="slide">
        <div
          id="mobile-sidebar-wrapper"
          v-if="isOpenSideBar"
          class="fixed top-0 left-0 w-[180px] h-full bg-white shadow-lg z-50 sm:hidden"
        >
          <SideBar @closeSideBar="closeSideBar" />
        </div>
      </transition>
      <!-- 背景オーバーレイ -->
      <transition name="fade">
        <div
          v-show="isOpenSideBar"
          id="mobile-sidebar-wrapper-overlay"
          class="fixed inset-0 bg-black bg-opacity-50 z-40 sm:hidden"
          @click="closeSideBar"
        />
      </transition>

      <section
        class="main-container overflow-y-scroll ml-0 sm:ml-[180px] w-full sm:w-[calc(100%-180px)]"
      >
        <router-view></router-view>
      </section>
    </main>
    <FooterArea />
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';

import FooterArea from '@/components/layouts/main/FooterArea.vue';
import HeaderArea from '@/components/layouts/main/header/HeaderArea.vue';
import SideBar from '@/components/layouts/main/sideBar/SideBar.vue';

const isOpenSideBar = ref<boolean>(false);

const openSideBar = () => {
  isOpenSideBar.value = true;
};
const closeSideBar = () => {
  isOpenSideBar.value = false;
};
</script>

<style lang="scss" scoped>
$header-height: 30px; // TODO:実質的に手打ちなので、なんかもーちょっとうまい書き方にしたい
$main-height: calc(100vh - #{$header-height} - 30px);

main {
  section {
    height: #{$main-height};
  }
}

#sidebar {
  height: #{$main-height};
}

#mobile-sidebar-wrapper {
  margin-top: #{$header-height};
  height: calc(#{$main-height} + 30px); //フッターの高さを足す

  &.slide-enter-active,
  &.slide-leave-active {
    transition: transform 0.1s ease;
  }

  &.slide-enter-from,
  &.slide-leave-to {
    transform: translateX(-100%);
  }

  &.slide-enter-to,
  &.slide-leave-from {
    transform: translateX(0);
  }

  &.fade-enter-active,
  &.fade-leave-active {
    transition: opacity 0.1s ease;
  }

  &.fade-enter-from,
  &.fade-leave-to {
    opacity: 0;
  }

  &.fade-enter-to,
  &.fade-leave-from {
    opacity: 1;
  }
}
#mobile-sidebar-wrapper-overlay {
  height: 100vh;
}
</style>
