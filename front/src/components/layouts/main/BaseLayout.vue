<template>
  <main class="flex bg-gray-50">
    <!-- PC・タブレット用 -->
    <section id="sidebar" class="sidebar fixed hidden md:block">
      <SideBar />
    </section>

    <section
      id="main-container"
      class="ml-0 md:ml-[180px] w-full md:w-[calc(100%-180px)] overflow-y-auto h-screen"
    >
      <header
        ref="headerRef"
        class="sticky top-0 bg-white/30 backdrop-blur-sm w-full"
      >
        <HeaderArea />
      </header>
      <div id="page-area">
        <router-view />
      </div>

      <div id="footer" class="w-full">
        <FooterArea />
      </div>
    </section>
  </main>
</template>

<script setup lang="ts">
import FooterArea from '@/components/layouts/main/FooterArea.vue';
import HeaderArea from '@/components/layouts/main/header/HeaderArea.vue';
import SideBar from '@/components/layouts/main/sideBar/SideBar.vue';
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
      min-height: #{$main-height};
    }
  }
}

#sidebar {
  height: 100vh;
}

#footer {
  height: #{$footer-height};
}
</style>
