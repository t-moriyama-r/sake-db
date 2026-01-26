/**
 * 登録情報編集フォーム
 */
import * as yup from 'yup';
import { string } from 'yup';

import type { AuthUserFull } from '@/graphQL/Auth/auth';
import yupLocaleJP from '@/lib/yup/yupLocaleJa';

import {
  FormKeys as RegisterFormKeys,
  type FormValues as RegisterFormValues,
  validationSchema as RegisterValidationSchema,
} from '../../../Auth/Register/form/RegisterForm';

yup.setLocale(yupLocaleJP);

export const FormKeys = {
  ...RegisterFormKeys,
  IMAGE: 'image',
  PROFILE: 'profile',
} as const;

export interface FormValues extends RegisterFormValues {
  [FormKeys.IMAGE]: File | null;
  [FormKeys.PROFILE]: string;
}

export function generateInitialValues(user: AuthUserFull): FormValues {
  return {
    [FormKeys.NAME]: user.name,
    [FormKeys.MAIL]: user.email,
    [FormKeys.PASSWORD]: '', //パスワードは入力が必須ではない
    [FormKeys.IMAGE]: null,
    [FormKeys.PROFILE]: user.profile,
  };
}
export const validationSchema = {
  ...RegisterValidationSchema,
  [FormKeys.PASSWORD]: string()
    .transform((value) => (value === '' ? null : value))
    .nullable()
    .test('min-length', '7文字以上で入力してください', (value) => {
      // transformで空文字列はnullに変換済みなので、null/undefinedのみチェック
      if (value === null || value === undefined) {
        return true; // 空の場合はOK
      }
      return value.length >= 7;
    }),
};
