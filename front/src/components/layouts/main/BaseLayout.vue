<template>
  <div>
    <HeaderArea />
    <main class="flex">
      <section class="side-bar fixed">
        <SideBar />
      </section>
      <section class="main-container overflow-y-scroll">
        <router-view></router-view>
      </section>
    </main>
    <FooterArea />
  </div>
</template>
<script setup lang="ts">
import FooterArea from '@/components/layouts/main/FooterArea.vue';
import HeaderArea from '@/components/layouts/main/header/HeaderArea.vue';
import SideBar from '@/components/layouts/main/sideBar/SideBar.vue';
</script>

<style lang="scss" scoped>
$header-height: 50px;

div {
  height: calc(100vh - #{$header-height});

  main {
    section {
      height: calc(100vh - #{$header-height});
      &.main-container {
        margin-left: 180px;
        width: calc(100% - 180px);
      }
    }
  }
}
</style>
