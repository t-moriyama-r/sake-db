//memo:idはトークンから取るので、inputはRegisterと同値でかまわないが、ログイン判定を必要とするため呼び出すリゾルバが異なる
import { gql } from '@apollo/client/core';
import type { DocumentNode } from 'graphql/index';

import type { EvaluateList } from '@/graphQL/User/user';

export const Update: DocumentNode = gql`
  mutation updateUser($input: RegisterInput!) {
    updateUser(input: $input)
  }
`;

export interface GetMyEvaluateListResponse {
  readonly getUserByIdDetail: {
    readonly evaluateList: EvaluateList;
  };
}

// 自分の評価リストを取得するクエリ
// getUserByIdDetailを使って自分のIDで呼び出す
export const GET_MY_EVALUATE_LIST: DocumentNode = gql`
  query getMyEvaluateList($id: String!) {
    getUserByIdDetail(id: $id) {
      evaluateList {
        recentComments {
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate5Liquors {
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate4Liquors {
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate3Liquors {
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate2Liquors {
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        rate1Liquors {
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
        noRateLiquors {
          categoryId
          categoryName
          liquorId
          name
          imageBase64
          comment
          rate
          updatedAt
        }
      }
    }
  }
`;
