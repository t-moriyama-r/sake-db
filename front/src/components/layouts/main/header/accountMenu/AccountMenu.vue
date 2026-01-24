<template>
  <div ref="menuRef">
    <!-- トリガー: アカウント情報 -->
    <button
      ref="triggerRef"
      @click="toggleMenu"
      class="flex items-center gap-1 px-2 py-1 rounded hover:bg-gray-100 transition-colors cursor-pointer"
    >
      <AccountInfo />
      <FontAwesomeIcon
        icon="fa-solid fa-chevron-down"
        class="text-xs transition-transform"
        :class="{ 'rotate-180': isOpen }"
      />
    </button>

    <!-- ドロップダウンメニュー (Teleport で body に出す) -->
    <Teleport to="body">
      <Transition name="dropdown">
        <div
          v-if="isOpen"
          ref="dropdownRef"
          class="fixed bg-white border border-gray-200 rounded-lg shadow-lg min-w-40 z-50 overflow-hidden"
          :style="dropdownStyle"
        >
          <ul class="py-1">
            <!-- マイページ -->
            <li>
              <router-link
                :to="{ name: 'MyPageIndex' }"
                class="flex items-center gap-2 px-4 py-2 hover:bg-gray-100 transition-colors"
                @click="closeMenu"
              >
                <FontAwesomeIcon icon="fa-solid fa-user" class="w-4" />
                <span>マイページ</span>
              </router-link>
            </li>

            <!-- 管理者ページ (管理者のみ) -->
            <li v-if="userStore.getRoles().includes(Roles.Admin)">
              <router-link
                :to="{ name: 'Admin' }"
                class="flex items-center gap-2 px-4 py-2 hover:bg-gray-100 transition-colors"
                @click="closeMenu"
              >
                <FontAwesomeIcon icon="fa-solid fa-gear" class="w-4" />
                <span>管理者ページ</span>
              </router-link>
            </li>

            <li class="border-t border-gray-200 my-1"></li>

            <!-- ログアウト -->
            <li>
              <button
                @click="handleLogout"
                class="flex items-center gap-2 px-4 py-2 w-full text-left hover:bg-gray-100 transition-colors text-red-600"
              >
                <FontAwesomeIcon
                  icon="fa-solid fa-right-from-bracket"
                  class="w-4"
                />
                <span>ログアウト</span>
              </button>
            </li>
          </ul>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { computed, onMounted, onUnmounted, ref } from 'vue';
import { useRouter } from 'vue-router';

import AccountInfo from '@/components/layouts/main/header/accountMenu/AccountInfo.vue';
import { useToast } from '@/funcs/composable/useToast';
import { Roles } from '@/graphQL/Auth/types';
import { useUserStore } from '@/stores/userStore/userStore';

const userStore = useUserStore();
const toast = useToast();
const router = useRouter();

const isOpen = ref(false);
const menuRef = ref<HTMLElement | null>(null);
const triggerRef = ref<HTMLElement | null>(null);
const dropdownRef = ref<HTMLElement | null>(null);
const triggerRect = ref<DOMRect | null>(null);

const dropdownStyle = computed(() => {
  if (!triggerRect.value) return {};
  return {
    top: `${triggerRect.value.bottom + 4}px`,
    right: `${window.innerWidth - triggerRect.value.right}px`,
  };
});

function updatePosition() {
  if (triggerRef.value) {
    triggerRect.value = triggerRef.value.getBoundingClientRect();
  }
}

async function toggleMenu() {
  if (!isOpen.value) {
    updatePosition();
  }
  isOpen.value = !isOpen.value;
}

function closeMenu() {
  isOpen.value = false;
}

async function handleLogout() {
  closeMenu();
  try {
    await userStore.logout();
    toast.showToast({
      message: 'ログアウトしました',
    });
    void router.push({ name: 'Index' });
  } catch (error) {
    toast.showToast({
      message: 'ログアウトに失敗しました',
    });
  }
}

// 外側クリックでメニューを閉じる
function handleClickOutside(event: MouseEvent) {
  const target = event.target as Node;
  const isInsideTrigger = menuRef.value?.contains(target);
  const isInsideDropdown = dropdownRef.value?.contains(target);

  if (!isInsideTrigger && !isInsideDropdown) {
    closeMenu();
  }
}

// スクロール・リサイズ時にメニューを閉じる
function handleScrollOrResize() {
  if (isOpen.value) {
    closeMenu();
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
  window.addEventListener('scroll', handleScrollOrResize, true);
  window.addEventListener('resize', handleScrollOrResize);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
  window.removeEventListener('scroll', handleScrollOrResize, true);
  window.removeEventListener('resize', handleScrollOrResize);
});
</script>

<style>
.dropdown-enter-active,
.dropdown-leave-active {
  transition:
    opacity 0.15s ease,
    transform 0.15s ease;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}
</style>
