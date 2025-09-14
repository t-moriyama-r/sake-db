<template>
  <main class="flex bg-gray-50">
    <!-- PC・タブレット用 -->
    <section id="sidebar" class="sidebar fixed hidden md:block">
      <SideBar />
    </section>

    <!-- スマホ用 スライドサイドバー -->
    <transition name="slide">
      <div
        id="mobile-sidebar-wrapper"
        v-if="isOpenSideBar"
        class="fixed top-0 left-0 w-[180px] h-full bg-white shadow-lg z-50 md:hidden"
      >
        <SideBar @closeSideBar="closeSideBar" />
      </div>
    </transition>
    <!-- 背景オーバーレイ -->
    <transition name="fade">
      <div
        v-show="isOpenSideBar"
        id="mobile-sidebar-wrapper-overlay"
        class="fixed inset-0 bg-black bg-opacity-50 z-40 md:hidden"
        @click="closeSideBar"
      />
    </transition>

    <section
      id="main-container"
      class="ml-0 md:ml-[180px] w-full md:w-[calc(100%-180px)] overflow-y-auto h-screen"
    >
      <header
        ref="headerRef"
        class="sticky top-0 bg-white/30 backdrop-blur-sm w-full"
      >
        <HeaderArea @openSideBar="openSideBar" />
      </header>
      <div id="page-area">
        <router-view></router-view>
      </div>

      <div id="footer" class="w-full">
        <FooterArea />
      </div>
    </section>
  </main>
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
$header-height: 40px; // TODO:実質的に手打ちなので、なんかもーちょっとうまい書き方にしたい
$footer-height: 30px; // TODO:実質的に手打ちなので、なんかもーちょっとうまい書き方にしたい
$main-height: calc(100vh - #{$header-height} - #{$footer-height});

main {
  #main-container {
    min-height: 100vh;

    header {
      height: $header-height;
    }

    #page-area {
      padding-top: $header-height;
      min-height: calc(100vh - #{$footer-height});
    }
  }
}

#sidebar {
  height: 100vh;
}

#mobile-sidebar-wrapper {
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

#footer {
  height: #{$footer-height};
}
</style>
