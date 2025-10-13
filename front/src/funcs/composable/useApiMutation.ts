import { useMutation } from '@tanstack/vue-query';
import axios, { type AxiosError, type AxiosResponse } from 'axios';
import { unref } from 'vue';

import { TOKEN_EXPIRE_MSG } from '@/funcs/composable/const';
import { debug, errorDebug } from '@/funcs/util/core/console';
import router from '@/router';
import { refreshToken, useUserStore } from '@/stores/userStore/userStore';
import type { APIType } from '@/type/api/APIType/APIType';

export function useApiMutation<
  Request extends object | null = object,
  Response extends object | null = object,
>(
  apiType: APIType<Request, Response>,
  options: { isAuth: boolean } = { isAuth: false },
) {
  const mutationFn = async (
    data: Request,
  ): Promise<AxiosResponse<Response>> => {
    debug('リクエストdata:', data);
    const headers = apiType.headers ?? { 'Content-Type': 'application/json' };

    // isAuthフラグがtrueの場合、JWTトークンを追加
    if (options.isAuth) {
      // ストアにアクセスしてアクセストークンを取得
      const { accessToken } = useUserStore();
      const token = accessToken.get();
      if (token) {
        headers['Authorization'] = `Bearer ${token}`;
      }
    }

    async function run(): Promise<AxiosResponse<Response>> {
      return axios({
        url: apiType.url,
        method: apiType.method,
        headers,
        data,
      }).then((response) => {
        debug('リクエスト成功 - レスポンス:', response.data);
        return response;
      });
    }

    try {
      return await run();
    } catch (err: unknown) {
      errorDebug('エラー返却：', err);
      const errResponse = (
        err as AxiosError<{ message: string; error_id: string }>
      ).response?.data;
      if (errResponse && errResponse.message == TOKEN_EXPIRE_MSG) {
        debug('リフレッシュトークン取得');
        await refreshToken(); //アクセストークン期限切れの場合、リフレッシュトークンを再取得

        //再度リクエスト
        return await run();
      } else if ((err as Error).message == 'unauthorized') {
        //認証エラーの場合はログインページにリダイレクト
        void router.push({ name: 'Login' });
      }
      throw new Error(errResponse?.message || '不明なエラー'); // 親にエラーを投げる
    }
  };

  return useMutation<AxiosResponse<Response>, unknown, Request>({
    mutationFn,
    ...unref(apiType.options), // optionsがMaybeRefDeepである可能性があるため、unrefで取り出す
  });
}
