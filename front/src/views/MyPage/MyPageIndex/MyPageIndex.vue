<template>
  <div class="my-page-container">
    <!-- ユーザープロフィールカード -->
    <div class="profile-card">
      <div class="profile-header">
        <div class="avatar-section">
          <RadiusImage
            v-if="props.user.imageBase64"
            :imageSrc="props.user.imageBase64"
            radius="60px"
            alt="プロフィール画像"
          />
          <div v-else class="default-avatar">
            <font-awesome-icon :icon="['fas', 'user']" class="user-icon" />
          </div>
        </div>
        <div class="user-info">
          <h1 class="user-name">{{ props.user.name }}</h1>
          <p v-if="props.user.email" class="user-email">
            {{ props.user.email }}
          </p>
          <p v-if="props.user.profile" class="user-profile">
            {{ props.user.profile }}
          </p>
        </div>
      </div>
      <div class="profile-actions">
        <router-link :to="{ name: 'MyPageEdit' }" class="edit-button-link">
          <CommonButton :color="ColorType.Primary" class="edit-button">
            <font-awesome-icon :icon="['fas', 'edit']" class="icon" />
            ユーザー情報を編集
          </CommonButton>
        </router-link>
      </div>
    </div>

    <!-- 評価済リストセクション -->
    <div class="rated-section">
      <div class="section-header">
        <h2 class="section-title">
          <font-awesome-icon :icon="['fas', 'star']" class="icon" />
          評価済リスト
        </h2>
      </div>
      <div class="section-content">
        <UserPosts v-if="evaluateList" :evaluates="evaluateList" />
        <div v-else class="loading-message">読み込み中...</div>
      </div>
    </div>

    <!-- ユーザーブックマークリストセクション -->
    <div class="bookmarks-section">
      <div class="section-header">
        <h2 class="section-title">
          <font-awesome-icon :icon="['fas', 'user-friends']" class="icon" />
          ユーザーブックマーク
        </h2>
      </div>
      <div class="section-content">
        <BookmarkList />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { onMounted, ref } from 'vue';

import CommonButton from '@/components/parts/common/CommonButton/CommonButton.vue';
import RadiusImage from '@/components/parts/common/RadiusImage.vue';
import useQuery from '@/funcs/composable/useQuery/useQuery';
import type { AuthUserFull } from '@/graphQL/Auth/types';
import {
  type EvaluateList,
  GET_USERDATA_FULL,
  type GetUserDetailResponse,
} from '@/graphQL/User/user';
import { ColorType } from '@/type/common/ColorType';
import BookmarkList from '@/views/MyPage/MyPageIndex/BookmarkList.vue';
import UserPosts from '@/views/User/UserPosts.vue';

interface Props {
  user: AuthUserFull;
}

const props = defineProps<Props>();

// 評価リストの取得
const evaluateList = ref<EvaluateList | null>(null);
const { fetch } = useQuery<GetUserDetailResponse>(GET_USERDATA_FULL, {
  isAuth: true,
});

onMounted(async () => {
  try {
    const response = await fetch({ id: props.user.id });
    evaluateList.value = response.getUserByIdDetail.evaluateList;
  } catch (error) {
    console.error('評価リストの取得に失敗しました:', error);
  }
});
</script>

<style scoped>
.my-page-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

/* プロフィールカード */
.profile-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  padding: 2rem;
  margin-bottom: 2rem;
}

.profile-header {
  display: flex;
  gap: 2rem;
  align-items: flex-start;
  margin-bottom: 1.5rem;
}

.avatar-section {
  flex-shrink: 0;
}

.default-avatar {
  width: 120px;
  height: 120px;
  border-radius: 60px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.user-icon {
  font-size: 3rem;
  color: white;
}

.user-info {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-size: 1.75rem;
  font-weight: 700;
  color: #1a202c;
  margin-bottom: 0.5rem;
}

.user-email {
  font-size: 0.95rem;
  color: #718096;
  margin-bottom: 0.75rem;
}

.user-profile {
  font-size: 1rem;
  color: #4a5568;
  line-height: 1.6;
  white-space: pre-wrap;
}

.profile-actions {
  display: flex;
  justify-content: flex-end;
}

.edit-button-link {
  text-decoration: none;
}

.edit-button .icon {
  margin-right: 0.5rem;
}

/* ブックマークセクション */
.bookmarks-section {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* 評価済リストセクション */
.rated-section {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
}

.section-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 1.5rem 2rem;
  border-radius: 12px 12px 0 0;
}

.section-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: white;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.section-title .icon {
  font-size: 1.25rem;
}

.section-content {
  padding: 2rem;
}

.loading-message {
  text-align: center;
  padding: 2rem;
  color: #718096;
  font-size: 1rem;
}

/* レスポンシブデザイン */
@media (max-width: 768px) {
  .my-page-container {
    padding: 1rem 0.5rem;
  }

  .profile-card {
    padding: 1.5rem;
  }

  .profile-header {
    flex-direction: column;
    align-items: center;
    text-align: center;
    gap: 1.5rem;
  }

  .user-info {
    width: 100%;
  }

  .profile-actions {
    justify-content: center;
  }

  .section-header {
    padding: 1rem 1.5rem;
  }

  .section-title {
    font-size: 1.25rem;
  }

  .section-content {
    padding: 1.5rem;
  }
}
</style>
