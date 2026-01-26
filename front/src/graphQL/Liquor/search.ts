import { type DocumentNode, gql } from '@apollo/client/core';

import { type Liquor as CardLiquor } from '../Index/random';

export interface SearchLiquorsResponse {
  readonly searchLiquors: CardLiquor[];
}

export const SEARCH_LIQUORS: DocumentNode = gql`
  query searchLiquors($keyword: String!, $limit: Int) {
    searchLiquors(keyword: $keyword, limit: $limit) {
      id
      name
      categoryId
      categoryName
      description
      imageBase64
      updatedAt
    }
  }
`;
