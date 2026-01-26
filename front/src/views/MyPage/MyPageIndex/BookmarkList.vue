<template>
  <div class="bookmark-list-container">
    <div v-if="bookmarks && bookmarks.length > 0" class="bookmarks-grid">
      <div v-for="user in bookmarks" :key="user.userId" class="bookmark-card">
        <div class="user-content">
          <div class="user-avatar">
            <RadiusImage
              v-if="user.imageBase64"
              :imageSrc="user.imageBase64"
              radius="30px"
              alt="ユーザーアイコン"
            />
            <div
              v-else
              class="default-user-avatar"
              role="img"
              aria-label="デフォルトユーザーアイコン"
            >
              <font-awesome-icon :icon="['fas', 'user']" class="user-icon" />
            </div>
          </div>
          <div class="user-details">
            <router-link
              :to="{ name: 'UserPage', params: { id: user.userId } }"
              class="user-name-link"
            >
              {{ user.name }}
            </router-link>
            <p class="bookmark-date">
              <font-awesome-icon
                :icon="['fas', 'calendar-plus']"
                class="icon"
              />
              {{ format(date(user.createdAt), 'yyyy/MM/dd') }}
            </p>
          </div>
        </div>
        <div class="bookmark-actions">
          <BookMarkLogics :target-id="user.userId" v-slot="{ remove }">
            <CommonButton
              :color="ColorType.Danger"
              size="small"
              @click="deleteUser(user.userId, remove)"
            >
              <font-awesome-icon :icon="['fas', 'trash']" class="icon" />
              削除
            </CommonButton>
          </BookMarkLogics>
        </div>
      </div>
    </div>
    <div v-else class="empty-state">
      <font-awesome-icon :icon="['fas', 'bookmark']" class="empty-icon" />
      <p class="empty-message">ブックマークしているユーザーはいません</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { format } from 'date-fns';
import { onMounted, ref } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import BookMarkLogics from '@/components/slots/BookMarkLogics.vue';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import date from '@/funcs/util/date';
import {
  type Bookmark,
  type GetBookmarkListResponse,
  LIST,
} from '@/graphQL/Bookmark/bookmark';
import { ColorType } from '@/type/common/ColorType';

const { fetch } = useQuery<GetBookmarkListResponse>(LIST, {
  isAuth: true,
});
const bookmarks = ref<Bookmark[] | null>(null);

onMounted(() => {
  void reFetch({
    isUseCache: false, //誰かをブックマークしてからリストに戻っても反映されないのでキャッシュを使わないことにした
  });
});

const reFetch = async ({
  isUseCache = false,
}: {
  isUseCache: boolean;
}): Promise<void> => {
  const response = await fetch(undefined, {
    fetchPolicy: isUseCache ? 'cache-first' : 'network-only',
  });
  bookmarks.value = response.getBookMarkList ?? [];
};

//削除ボタンの動作
const deleteUser = async (userId: string, removeFn: () => Promise<void>) => {
  await removeFn();
  //再度取得する
  void reFetch({
    isUseCache: false,
  });
};
</script>

<style scoped>
.bookmark-list-container {
  min-height: 200px;
}

.bookmarks-grid {
  display: grid;
  gap: 1rem;
}

.bookmark-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  transition: all 0.3s ease;
  background: #ffffff;
}

.bookmark-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translateY(-2px);
}

.user-content {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 1;
}

.user-avatar {
  flex-shrink: 0;
}

.default-user-avatar {
  width: 60px;
  height: 60px;
  border-radius: 30px;
  background: linear-gradient(135deg, #a8b3cf 0%, #8a9ab8 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.default-user-avatar .user-icon {
  font-size: 1.5rem;
  color: white;
}

.user-details {
  flex: 1;
  min-width: 0;
}

.user-name-link {
  font-size: 1.125rem;
  font-weight: 600;
  color: #2d3748;
  text-decoration: none;
  display: block;
  margin-bottom: 0.25rem;
  transition: color 0.2s;
}

.user-name-link:hover {
  color: #667eea;
}

.bookmark-date {
  font-size: 0.875rem;
  color: #718096;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 0;
}

.bookmark-date .icon {
  font-size: 0.75rem;
}

.bookmark-actions {
  flex-shrink: 0;
}

.bookmark-actions .icon {
  margin-right: 0.25rem;
}

/* 空の状態 */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 3rem 1rem;
  color: #a0aec0;
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  opacity: 0.5;
}

.empty-message {
  font-size: 1.125rem;
  margin: 0;
}

/* レスポンシブデザイン */
@media (max-width: 640px) {
  .bookmark-card {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }

  .user-content {
    width: 100%;
  }

  .bookmark-actions {
    width: 100%;
  }

  .bookmark-actions button {
    width: 100%;
  }
}
</style>
